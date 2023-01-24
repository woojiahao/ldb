package ldb

import (
	"reflect"
	"testing"
)

// TODO: Expand test cases for MySQL
func TestConfiguration_ConnectionString(t *testing.T) {
	config := LoadFromEnv("postgres", POSTGRESQL, ".env.example")
	expected := "postgres://postgres:root@localhost:5432/ldb?sslmode=disable"
	if config.ConnectionString() != expected {
		t.Errorf("Invalid connection string generated")
		t.Errorf("Expected:\n%s", expected)
		t.Errorf("Got:\n%s", config.ConnectionString())
	}
}

func TestLoadFromEnv(t *testing.T) {
	config := LoadFromEnv("postgres", POSTGRESQL, ".env.example")
	expected := &Configuration{
		Driver:   "postgres",
		Type:     POSTGRESQL,
		Username: "postgres",
		Password: "root",
		Host:     "localhost",
		Database: "ldb",
		Port:     5432,
		SSLMode:  DISABLE,
	}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Invalid configuration given")
		t.Errorf("Expected:\n%v", expected)
		t.Errorf("Got:\n%v", config)
	}
}
