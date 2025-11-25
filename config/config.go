package config

import (
	"os"
	"user-service/common/utils"

	"github.com/sirupsen/logrus"
	_ "github.com/spf13/viper/remote"
)

var Config AppConfig

type AppConfig struct {
	Port                  int            `json:"port"`
	AppName               string         `json:"appName"`
	AppEnv                string         `json:"appEnv"`
	SignatureKey          string         `json:"signatureKey"`
	Database              DatabaseConfig `json:"database"`
	RateLimiterMaxRequest float64        `json:"rateLimiterRequest"`
	RateLimiterTimeSecond float64        `json:"rateLimiterTimeSecond"`
	JWTSecretKey          string         `json:"jwtSecretKey"`
	JWTExpirationTime     float64        `json:"jwtExpirationTime"`
}

type DatabaseConfig struct {
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	Name                   string `json:"name"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	MaxOpenConnections     int    `json:"maxOpenConnections"`
	MaxLifetimeConnections int    `json:"maxLifetimeconnections"`
	MaxIdleConnections     int    `json:"maxIdleConnections"`
	MaxIdleTime            int    `json:"maxIdleTime"`
	MaxConnection          int    `json:"maxConnection"`
}

func Init() {
	err := utils.BindFromJSON(&Config, "config", ".")
	if err != nil {
		logrus.Infof("failed to bind config!")
		err = utils.BindFromConsulKV(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))

		if err != nil {
			logrus.Infof("failed to bind config from consul!")
			panic(err)
		}
	}

}
