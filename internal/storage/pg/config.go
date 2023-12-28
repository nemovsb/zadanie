package pg

type Config struct {
	Username string
	Host     string
	Port     string
	Password string
	DBName   string
}

func NewConfig(username string, host string, port string, password string, dbname string) *Config {
	return &Config{
		Username: username,
		Host:     host,
		Port:     port,
		Password: password,
		DBName:   dbname,
	}
}
