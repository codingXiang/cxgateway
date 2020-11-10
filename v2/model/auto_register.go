package model

type ServiceRegister struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AutoRegistrationInfo struct {
	Name         string            `json:"name" yaml:"name"`
	Url          string            `json:"url" yaml:"url"`
	Destinations map[string]string `json:"destinations" yaml:"destinations"`
}
