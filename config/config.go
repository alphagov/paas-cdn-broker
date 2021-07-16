package config

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Settings struct {
	Port                 string `envconfig:"port" default:"3000"`
	BrokerUsername       string `envconfig:"broker_username" required:"true"`
	BrokerPassword       string `envconfig:"broker_password" required:"true"`
	DatabaseUrl          string `envconfig:"database_url" required:"true"`
	CloudFrontPrefix     string `envconfig:"cloudfront_prefix" default:""`
	AwsAccessKeyId       string `envconfig:"aws_access_key_id" required:"true"`
	AwsSecretAccessKey   string `envconfig:"aws_secret_access_key" required:"true"`
	AwsDefaultRegion     string `envconfig:"aws_default_region" required:"true"`
	ServerSideEncryption string `envconfig:"server_side_encryption"`
	APIAddress           string `envconfig:"api_address" required:"true"`
	ClientID             string `envconfig:"client_id" required:"true"`
	ClientSecret         string `envconfig:"client_secret" required:"true"`
	DefaultOrigin        string `envconfig:"default_origin" required:"true"`
	DefaultDefaultTTL    int64  `envconfig:"default_default_ttl" default:"0"`
	Schedule             string `envconfig:"schedule" default:"0 0 * * * *"`
}

func NewSettings() (Settings, error) {
	var settings Settings
	err := envconfig.Process("cdn", &settings)
	if err != nil {
		return Settings{}, err
	}
	return settings, nil
}

func Connect(settings Settings) (*gorm.DB, error) {
	return gorm.Open("postgres", settings.DatabaseUrl)
}
