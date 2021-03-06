package jobs

import (
	"log"
	"time"

	"github.com/arschles/kubeapp/api/rc"
	"github.com/deis/workflow-manager/config"
	"github.com/deis/workflow-manager/data"
	apiclient "github.com/deis/workflow-manager/pkg/swagger/client"
	"github.com/deis/workflow-manager/pkg/swagger/client/operations"
)

// Periodic is an interface for managing periodic job invocation
type Periodic interface {
	// will have a Do method to begin execution
	Do() error
}

// SendVersions fulfills the Periodic interface
type sendVersions struct {
	secretGetterCreator data.KubeSecretGetterCreator
	rcLister            rc.Lister
	apiClient           *apiclient.WorkflowManager
	availableVersions   data.AvailableVersions
}

// NewSendVersionsPeriodic creates a new SendVersions using sgc and rcl as the the secret getter / creator and replication controller lister implementations (respectively)
func NewSendVersionsPeriodic(
	apiClient *apiclient.WorkflowManager,
	sgc data.KubeSecretGetterCreator,
	rcl rc.Lister,
	availableVersions data.AvailableVersions,
) Periodic {
	return &sendVersions{
		secretGetterCreator: sgc,
		rcLister:            rcl,
		apiClient:           apiClient,
		availableVersions:   availableVersions,
	}
}

// Do method of SendVersions
func (s sendVersions) Do() error {
	if config.Spec.CheckVersions {
		err := sendVersionsImpl(s.apiClient, s.secretGetterCreator, s.rcLister, s.availableVersions)
		if err != nil {
			return err
		}
	}
	return nil
}

type getLatestVersionData struct {
	vsns                  data.AvailableVersions
	installedData         data.InstalledData
	clusterID             data.ClusterID
	availableComponentVsn data.AvailableComponentVersion
	sgc                   data.KubeSecretGetterCreator
}

// NewGetLatestVersionDataPeriodic creates a new periodic implementation that gets latest version data. It uses sgc and rcl as the secret getter/creator and replication controller lister implementations (respectively)
func NewGetLatestVersionDataPeriodic(
	sgc data.KubeSecretGetterCreator,
	rcl rc.Lister,
	installedData data.InstalledData,
	clusterID data.ClusterID,
	availVsn data.AvailableVersions,
	availCompVsn data.AvailableComponentVersion,
) Periodic {

	return &getLatestVersionData{
		vsns:                  availVsn,
		installedData:         installedData,
		clusterID:             clusterID,
		availableComponentVsn: availCompVsn,
		sgc: sgc,
	}
}

// Do method of GetLatestVersionData
func (u *getLatestVersionData) Do() error {
	cluster, err := data.GetCluster(u.installedData, u.clusterID, u.availableComponentVsn, u.sgc)
	if err != nil {
		return err
	}
	if _, err := u.vsns.Refresh(cluster); err != nil {
		return err
	}
	return nil
}

// DoPeriodic is a function for running jobs at a fixed interval
func DoPeriodic(p []Periodic, interval time.Duration) chan struct{} {
	ch := make(chan struct{})
	// schedule later job runs at a regular, periodic interval
	ticker := time.NewTicker(interval)
	go func() {
		// run the period jobs once at invocation time
		runJobs(p)
		for {
			select {
			case <-ticker.C:
				runJobs(p)
			case <-ch:
				ticker.Stop()
				return
			}
		}
	}()
	return ch
}

// runJobs is a helper function to run a list of jobs
func runJobs(p []Periodic) {
	for _, job := range p {
		err := job.Do()
		if err != nil {
			log.Printf("periodic job ran and returned error (%s)", err)
		}
	}
}

//  sendVersions sends cluster version data
func sendVersionsImpl(
	apiClient *apiclient.WorkflowManager,
	secretGetterCreator data.KubeSecretGetterCreator,
	rcLister rc.Lister,
	availableVersions data.AvailableVersions,
) error {
	cluster, err := data.GetCluster(
		data.NewInstalledDeisData(rcLister),
		data.NewClusterIDFromPersistentStorage(secretGetterCreator),
		data.NewLatestReleasedComponent(secretGetterCreator, rcLister, availableVersions),
		secretGetterCreator,
	)
	if err != nil {
		log.Println("error getting installed components data")
		return err
	}

	_, err = apiClient.Operations.CreateClusterDetails(&operations.CreateClusterDetailsParams{Body: &cluster})
	if err != nil {
		log.Println("error sending diagnostic data")
		return err
	}
	return nil
}
