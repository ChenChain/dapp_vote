package conf

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var cfg *Config

func init() {
	c, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg = c
}

func Cfg() *Config {
	return cfg
}

type Config struct {
	URL     string
	Account Account
	Vote    Contract
}

type Account struct {
	Account string
	PriKey  string
}

type Contract struct {
	Name    string
	Address string
}

func loadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	config := &Config{
		URL: os.Getenv("URL"),
		Account: Account{
			Account: os.Getenv("ACCOUNT"),
			PriKey:  os.Getenv("PRIKEY"),
		},
		Vote: Contract{
			Name:    "VOTE_ADDRESS",
			Address: os.Getenv("VOTE_ADDRESS"),
		},
	}
	log.Println("config", config)
	return config, nil
}
