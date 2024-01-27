package config

type ServerConfig struct {
	Name      string      `mapstructure:"name"  json:"name"`
	Host      string      `mapstructure:"host"  json:"host"`
	Port      string      `mapstructure:"port"  json:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	OtelInfo  OtelConfig  `mapstructure:"otel"  json:"otel"`
	RedisInfo RedisConfig `mapstructure:"redis" json:"redis"`
	JWTInfo   JWTConfig   `mapstructure:"jwt" json:"jwt"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Name     string `mapstructure:"name" json:"name"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      uint64 `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataid" json:"data_id"`
	Group     string `mapstructure:"group" json:"group"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
