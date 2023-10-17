package viper

import (
	"blinkable/pkg/errs"

	"github.com/spf13/viper"
)

type Cfg struct {
	Viper *viper.Viper
}

func Init(configName string) Cfg {
	cfg := Cfg{
		Viper: viper.New(),
	}
	v := cfg.Viper
	v.SetConfigType("yml")
	v.SetConfigName(configName)
	v.AddConfigPath("../../configs/")

	if err := v.ReadInConfig(); err != nil {
		errs.HandleErrWithPanic("读取配置文件失败", err)
	}

	return cfg
}
