package configuration

import "log"

type DBConfig struct {
	Address  string
	Name     string
	User     string
	Password string
}

var dbConfig DBConfig

var serverPort string

func init() {

	dbConfig = DBConfig{
		Address:  getEnv("DB_ADDRESS", ""),
		Name:     getEnv("DB_NAME", ""),
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
	}

	serverPort = getEnv("SERVER_PORT", "3000")

	log.Print("Constants loaded")

}

func GetDBConfig() *DBConfig {
	return &dbConfig
}

func GetServerPort() *string {
	return &serverPort
}
