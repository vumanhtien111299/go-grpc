package config

type Config struct {
	MongoConfig MongoConfig `mapstructure:"mongodb"`
	App         App         `mapstructure:"app"`
}

type App struct {
	Port int `mapstructure:"port"`
}

type MongoConfig struct {
	DbName   string `mapstructure:"DB_NAME"`
	Address  string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}
