package config

type Configuration struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}
