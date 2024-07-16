package config

type MpKey struct {
	AppSecret string
	AppId     string
}

type MiniProgramConfig struct {
	Module []Module `mapstructure:"module" json:"module" yaml:"module"`
}

type Module struct {
	Name      string `mapstructure:"name" json:"name" yaml:"name"`
	ImagePath string `mapstructure:"imagePath" json:"imagePath" yaml:"imagePath"`
	Click     string `mapstructure:"click" json:"click" yaml:"click"`
}
