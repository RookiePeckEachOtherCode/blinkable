package config

import "blinkable/server/kitex_gen/user/userservice"

var (
	GlobalServerConfig = &ServerConfig{}
	GlobalNacosConfig  = &NacosConfig{}

	GlobalUserClient userservice.Client
)
