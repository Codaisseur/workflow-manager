package config

import "github.com/kelseyhightower/envconfig"

// Specification config struct
type Specification struct {
	Port           string `default:"8080" envconfig:"PORT"`
	Polling        int    `default:"43200" envconfig:"POLL_INTERVAL_SEC"` // 43200 seconds = 12 hours
	VersionsAPIURL string `envconfig:"VERSIONS_API_URL" default:"https://versions-staging.deis.com"`
	DoctorAPIURL   string `envconfig:"DOCTOR_API_URL" default:"https://doctor-staging.deis.com"`
	APIVersion     string `envconfig:"API_VERSION" default:"v2"`
	CheckVersions  bool   `default:"true" envconfig:"CHECK_VERSIONS"`
	DeisNamespace  string `default:"deis" envconfig:"DEIS_NAMESPACE"`
}

// Spec is an exportable variable that contains workflow manager config data
var Spec Specification

func init() {
	envconfig.Process("workflow_manager", &Spec)
}
