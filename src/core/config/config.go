package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

var (
	cfg  ConfigImpl
	once sync.Once
)

type configBuilder struct {
	dotenvFile           string
	errorOnMissingDotenv bool
}

func ConfigBuilder() configBuilder {
	return configBuilder{}
}

func (b configBuilder) WithDotenvFile(file string) configBuilder {
	b.dotenvFile = file
	return b
}

func (b configBuilder) PanicOnMissingDotenv(status bool) configBuilder {
	b.errorOnMissingDotenv = status
	return b
}

// Create a new instance of ConfigImpl layering up config:
//   1. dotenv file
//   2. environment variables
// with the latter overriding the former
// It will look for a file in the same directory called .env but you can override this with your own
// file by specifying the environment variable ENV_FILE=/path/to/dotenv file
func (b configBuilder) Build() ConfigImpl {
	cfg = NewConfig()

	dotenvFile := ".env"
	if b.dotenvFile != "" {
		dotenvFile = b.dotenvFile
	}

	dotenvFeeder := feeder.DotEnv{Path: dotenvFile}
	envFeeder := feeder.Env{}

	err := config.New().AddStruct(&cfg).AddFeeder(dotenvFeeder).Feed()
	if err != nil {
		if strings.Contains(err.Error(), "no such file") && b.errorOnMissingDotenv {
			log.Fatalf("error loading config from dotenv file %s: %s", dotenvFile, err.Error())
		}
	}
	err = config.New().AddStruct(&cfg).AddFeeder(envFeeder).Feed()
	if err != nil {
		log.Fatalf("error loding config from environemnt: %s", err.Error())
	}
	return cfg
}

// get, and initialise if not already done, the app config
func Config() ConfigImpl {
	once.Do(func() {
		envFile := ".env"
		if os.Getenv("ENV_FILE") != "" {
			envFile = os.Getenv("ENV_FILE")
		}
		cfg = ConfigBuilder().WithDotenvFile(envFile).Build()
	})
	return cfg
}

type ConfigImpl struct {
	Database DatabaseConfig
}

// Bootstrap the application Config struct with the default config
func NewConfig() ConfigImpl {
	return ConfigImpl{
		Database: DatabaseConfig{
			Driver: "mysql",
			Host:   "127.0.0.1",
			User:   "root",
			Name:   "prototype",
			Port:   3306,
		},
	}
}

type DatabaseConfig struct {
	Driver   string `env:"DB_DRIVER"`
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Port     int    `env:"DB_PORT"`
}

// get a database connection string based on the database config
// if using sqlite, DB_NAME is the path to the database file
func (db *DatabaseConfig) GetHost() (string, error) {
	conn := ""
	switch db.Driver {
	case "mysql":
		userPart := db.User
		if db.Password != "" {
			userPart += ":" + db.Password
		}
		conn = fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userPart, db.Host, db.Port, db.Name)
	case "sqlite3":
		conn = db.Name
	default:
		return conn, fmt.Errorf("unknown database driver %s", db.Driver)
	}
	return conn, nil
}
