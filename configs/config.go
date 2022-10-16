package configs

type SqlConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

type Configuration struct {
	Port        string
	LogFilePath string
	JWTSecret   []byte
	SqlConfig   SqlConfig
}

var config Configuration

func init() {
	config = Configuration{
		Port: "8080",
		SqlConfig: SqlConfig{
			Host:     "postgres",
			User:     "postgres",
			Password: "postgres",
			Port:     "5432",
			Database: "companies",
		},
		LogFilePath: "./log",
		JWTSecret:   []byte("companies_jwt_token"),
	}
}

func GetConfig() Configuration {
	return config
}
