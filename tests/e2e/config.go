package e2e

import (
	"encoding/json"
	"fmt"
	"os"
)

const ConfigEnv = "E2E_CONFIG"

type Config struct {
	ApiURL              string `json:"api_url"`
	ApiKey              string `json:"api_key"`
	AuthJWTTokenURL     string `json:"jwt_token_url"`
	AuthJWTClientID     string `json:"jwt_client_id"`
	AuthJWTClientSecret string `json:"jwt_client_secret"`
}

func NewConfig() (*Config, error) {
	cfgStr := os.Getenv(ConfigEnv)
	if cfgStr == "" {
		return nil, fmt.Errorf("expected test data at environment variable '%s'", ConfigEnv)
	}

	cfg := &Config{}
	if err := json.Unmarshal([]byte(cfgStr), cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
