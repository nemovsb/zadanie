package config

type Config struct {
	Server  `mapstructure:"server"`
	Storage `mapstructure:"storage"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Storage struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"storage_pswrd"`
	DBName   string `mapstructure:"db_name"`
}
