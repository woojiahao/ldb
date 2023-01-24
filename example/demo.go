package demo

import (
	"fmt"
	"github.com/woojiahao/ldb"
)

func Demo() {
	config := ldb.Configuration{
		Driver:   "postgres",
		Username: "postgres",
		Password: "root",
		Host:     "localhost",
		Name:     "ldb",
		Port:     5432,
		SSLMode:  ldb.DISABLE,
	}
	fmt.Println(config)
}
