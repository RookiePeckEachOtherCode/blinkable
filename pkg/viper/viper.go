package viper

import (
	"log"

	"github.com/spf13/viper"
)

type Cfg struct {
	Viper *viper.Viper
}

func Load(configName string) Cfg {
	cfg := Cfg{
		Viper: viper.New(),
	}
	v := cfg.Viper
	v.SetConfigType("yml")
	v.SetConfigName(configName)
	v.AddConfigPath("./configs/")
	v.AddConfigPath("../configs/")
	v.AddConfigPath("../../configs/")
	v.AddConfigPath("../../../configs/")

	if err := v.ReadInConfig(); err != nil {
		log.Panicf("读取配置文件失败: %s", err)
	}

	return cfg
}
