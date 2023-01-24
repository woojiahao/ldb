package ldb

import (
	"fmt"
	"log"
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

const (
	DISABLE     SSLMode = "disable"
	ALLOW       SSLMode = "allow"
	PREFER      SSLMode = "prefer"
	REQUIRE     SSLMode = "require"
	VERIFY_CA   SSLMode = "verify-ca"
	VERIFY_FULL SSLMode = "verify-full"
)

const (
	POSTGRESQL Type = "postgresql"
	MYSQL      Type = "mysql"
)

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
