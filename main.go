package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dangquyitt/go-github-action/handler"
	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"PORT"`
}

func main() {
	// load .env file
	config, err := LoadConfig(".")

	if err != nil {
		log.Fatal(err)
	}

	// New serve mux
	mux := http.NewServeMux()

	// Register API handler
	mux.HandleFunc("GET /", handler.HelloHandler)
	mux.HandleFunc("GET /items", handler.ItemsHandler)
	mux.HandleFunc("GET /randomuser", handler.GetRandomUser)

	// ListenAndServe
	fmt.Printf("Server listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux))
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
