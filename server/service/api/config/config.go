package config

type ServerConfig struct {
	Name       string          `mapstructure:"name"  json:"name"`
	Host       string          `mapstructure:"host"  json:"host"`
	Port       string          `mapstructure:"port"  json:"port"`
	JWTInfo    JWTConfig       `mapstructure:"jwt" json:"jwt"`
	OtelInfo   OtelConfig      `mapstructure:"otel"  json:"otel"`
	UserServer RPCServerConfig `mapstructure:"user_server" json:"user_server"`
	MinioInfo  MinioConfig     `mapstructure:"minio" json:"minio"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
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

type RPCServerConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type MinioConfig struct {
	EndPoint        string   `mapstructure:"endpoint" json:"endpoint"`
	AccessKeyID     string   `mapstructure:"access_key_id" json:"access_key_id"`
	SecretAccessKey string   `mapstructure:"secret_access_key" json:"secret_access_key"`
	Bucket          []string `mapstructure:"bucket" json:"bucket"`
}
