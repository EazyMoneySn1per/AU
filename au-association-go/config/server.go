package config

type Server struct {
	Mysql             Mysql             `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System            System            `mapstructure:"system" json:"system" yaml:"system"`
	MpKey             MpKey             `mapstructure:"mpKey" json:"myKey" yaml:"mpKey"`
	Interview         Interview         `mapstructure:"interview" json:"interview" yaml:"interview"`
	Local             Local             `mapstructure:"local" json:"local" yaml:"local"`
	StoragePath       StoragePath       `mapstructure:"storagePath" json:"storagePath" yaml:"storagePath"`
	MiniProgramConfig MiniProgramConfig `mapstructure:"miniProgramConfig" json:"miniProgramConfig" yaml:"miniProgramConfig"`
}
