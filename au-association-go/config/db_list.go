package config

type GeneralDB struct {
	Path         string `json:"path" yaml:"path"`
	Port         string `json:"port" yaml:"port"`
	Config       string `json:"config" yaml:"config"`
	Dbname       string `json:"dbname" yaml:"dbname"`
	Username     string `json:"username" yaml:"username"`
	Password     string `json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
}
