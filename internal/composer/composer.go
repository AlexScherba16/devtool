package composer

import (
	"devtool/config"
	"devtool/internal/api"
	"github.com/spf13/viper"
	"log"
	"time"
)

type composer struct {
	server *api.Server
}

func NewApplication() (composer, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return composer{}, err
	}

	handler, err := api.NewHandler()
	if err != nil {
		return composer{}, err
	}

	return composer{
		server: api.NewServer(
			&config.ServerConfig{
				Address:      viper.GetString("server.address"),
				Port:         viper.GetString("server.port"),
				WriteTimeout: viper.GetDuration("server.write_timeout_s") * time.Second,
				ReadTimeout:  viper.GetDuration("server.read_timeout_s") * time.Second,
			},
			handler.Routes()),
	}, nil
}

func (c *composer) Start() {
	c.server.Run()
}

func (c *composer) Shutdown() {
	if err := c.server.Stop(); err != nil {
		log.Fatalf("Error occurred on server shutting down: %s", err.Error())
	}
}
