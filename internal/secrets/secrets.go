package secrets

import (
	"encoding/json"
	"io/ioutil"
)

// Secrets - Application secrets
type Secrets struct {
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPass     string `json:"db_pass"`
	DBDatabase string `json:"db_database"`
	JWTSecret  string `json:"jwt_secret"`
}

var LoadedSecrets Secrets

// LoadSecrets - first attempts to load secrets from secrets.json file
// if no secrets.json file is found, will attempt to load from aws
// secrets manager
func LoadSecrets() (Secrets, error) {
	secrets := Secrets{}
	configJSON, err := ioutil.ReadFile("configs/secrets.json")
	if err != nil {
		return secrets, err
	}
	err = json.Unmarshal([]byte(configJSON), &secrets)
	if err != nil {
		return secrets, err
	}
	LoadedSecrets = secrets
	return secrets, nil
}
