package ldb

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type (
	SSLMode string
	Type    string

	Configuration struct {
		Driver   string
		Type     string
		Username string
		Password string
		Host     string
		Database string
		Port     int
		SSLMode
	}
)

// SSL Modes available
const (
	DISABLE     SSLMode = "disable"
	ALLOW       SSLMode = "allow"
	PREFER      SSLMode = "prefer"
	REQUIRE     SSLMode = "require"
	VERIFY_CA   SSLMode = "verify-ca"
	VERIFY_FULL SSLMode = "verify-full"
)

func (s SSLMode) isValid() bool {
	switch s {
	case DISABLE:
	case ALLOW:
	case PREFER:
	case REQUIRE:
	case VERIFY_FULL:
	case VERIFY_CA:
		return true
	default:
		return false
	}
	return false
}

// Databases available
const (
	POSTGRESQL Type = "postgresql"
	MYSQL      Type = "mysql"
)

// LoadFromEnv loads a given configuration from an environment file that is read from the root of the project. Database
// driver and database type are specified as parameters.
// Expects DATABASE_{property} when reading a .env file
func LoadFromEnv(driver, databaseType, envFile string) Configuration {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	username, ok := os.LookupEnv("DATABASE_USERNAME")
	if !ok {
		log.Fatalf("Missing field 'DATABASE_USERNAME'")
	}

	password, ok := os.LookupEnv("DATABASE_PASSWORD")
	if !ok {
		log.Fatalf("Missing field 'DATABASE_PASSWORD'")
	}

	host, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		log.Fatalf("Missing field 'DATABASE_HOST'")
	}

	database, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		log.Fatalf("Missing field 'DATABASE_NAME'")
	}

	port, ok := os.LookupEnv("DATABASE_PORT")
	if !ok {
		log.Fatalf("Missing field 'DATABASE_PORT'")
	}
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Invalid 'DATABASE_PORT' provided. Must be integer value.")
	}

	sslMode, ok := os.LookupEnv("DATABASE_SSL")
	ssl := SSLMode(sslMode)
	if !ok || !ssl.isValid() {
		ssl = DISABLE
	}

	return Configuration{
		Driver:   driver,
		Type:     databaseType,
		Username: username,
		Password: password,
		Host:     host,
		Database: database,
		Port:     portNumber,
		SSLMode:  ssl,
	}
}

// ConnectionString generates the appropriate connection URI for a database configuration
func (c Configuration) ConnectionString() string {
	switch c.Type {
	case POSTGRESQL:
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.Database, c.SSLMode)
	case MYSQL:
		return fmt.Sprintf("mysql://%s:%s@%s:%d/%s?ssl-mode=%s", c.Username, c.Password, c.Host, c.Port, c.Database, c.SSLMode)
	default:
		log.Fatalf("Invalid type of database. Only PostgreSQL and MySQL are supported currently.")
		return ""
	}
}
