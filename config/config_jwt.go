package config

type Jwt struct {
	Secret  string `json:"secret" yaml:"secret"`
	Expires int    `json:"expires" yaml:"expires"`
	Issuer  string `json:"issuer" yaml:"issuer"`
}
