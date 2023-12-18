package config

import (
	"fmt"
	"reflect"
	"time"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" envDocs:"Service name used for resource tagging" envDefault:"observability-manager"`
	ClientName  string `env:"CLIENT_NAME" envDocs:"Publisher name used for resource tagging" envDefault:"accelbyte"`
	ProjectName string `env:"PROJECT_NAME" envDocs:"Project name used for resource tagging" envDefault:"justice"`

	Port   string `env:"SERVER_PORT" envDocs:"The port which the service will listen to" envDefault:"8080"`
	Realm  string `env:"REALM_NAME" envDocs:"The realm of the deployment"  envDefault:"dev"`
	Region string `env:"AWS_REGION" envDocs:"AWS region" envDefault:"us-west-2"`

	BasePath string `env:"SERVER_BASE_PATH" envDocs:"The base path of service" envDefault:"/template"`
	LogLevel string `env:"LOG_LEVEL" envDocs:"Log level" envDefault:"info"`

	PostgresHost            string        `env:"POSTGRES_HOST" envDocs:"The host for connecting to postgresql"`
	PostgresPort            string        `env:"POSTGRES_PORT" envDocs:"The port for connecting to postgresql"`
	PostgresPassword        string        `env:"POSTGRES_PASSWORD" envDocs:"The password for connecting to postgresql"`
	PostgresUser            string        `env:"POSTGRES_USER" envDocs:"The user for connecting to postgresql"`
	PostgresDatabase        string        `env:"POSTGRES_DB" envDocs:"Database name"`
	PostgresTimeout         string        `env:"POSTGRES_TIMEOUT" envDocs:"Connection Timeout postgresql (in Second)" envDefault:"5"`
	PostgresSSLMode         string        `env:"POSTGRES_SSL_MODE" envDocs:"SSL Mode"`
	PostgresDBReplicaHost   []string      `env:"POSTGRES_REPLICA_HOST" envDocs:"Postgres Replica Host" envDefault:""`
	PostgresMaxOpenConn     int           `env:"POSTGRES_MAX_OPEN_CONNECTIONS" envDocs:"Postgres maximum number of concurrent connection to DB" envDefault:"30"`
	PostgresMaxIdleConn     int           `env:"POSTGRES_MAX_IDLE_CONNECTIONS" envDocs:"Max buffered postgres connection" envDefault:"10"`
	PostgresConnMaxLifetime time.Duration `env:"POSTGRES_CONNECTION_MAX_LIFETIME" envDocs:"Max lifetime of idle postgres connection in seconds" envDefault:"3600s"`
}

func (config Config) HelpDocs() []string {
	reflectEnvVar := reflect.TypeOf(config)
	doc := make([]string, 1+reflectEnvVar.NumField())
	doc[0] = "Environment variables config:"
	for i := 0; i < reflectEnvVar.NumField(); i++ {
		field := reflectEnvVar.Field(i)
		envName := field.Tag.Get("env")
		envDafault := field.Tag.Get("envDefault")
		envDocs := field.Tag.Get("envDocs")
		doc[i+1] = fmt.Sprintf("  %v\t %v (default: %v)", envName, envDocs, envDafault)
	}

	return doc
}
