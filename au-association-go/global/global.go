package global

import (
	"au-golang/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Server
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
)
