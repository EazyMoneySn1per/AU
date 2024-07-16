package config

type Interview struct {
	OpenTime string `mapstructure:"open-time" json:"open-time" yaml:"open-time"`
	EndTime  string `mapstructure:"end-time" json:"end-time" yaml:"end-time"`
}
