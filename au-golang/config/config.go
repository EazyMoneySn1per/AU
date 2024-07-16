package config

type Server struct {
	//JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	// gorm
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
	Interview Interview `mapstructure:"interview" json:"interview" yaml:"interview"`
}
