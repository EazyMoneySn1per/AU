package global

import (
	"go.uber.org/zap"

	"au-go/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB *gorm.DB
	//GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
)
