package main

import (
	"log"
	"os"
)

func main() {
	getEnv("PG_PASSWORD") // test for a password being set
	psql := PostgreSQL{
		Host:     getEnv("PG_HOST"),
		Port:     getEnv("PG_PORT"),
		DB:       getEnv("PG_DB"),
		Username: getEnv("PG_USER"),
		Options:  nil,
	}

	dumpFilePath, err := psql.Dump()
	if err != nil {
		log.Fatalln("error on db dump", err)
	}
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal("Environment variable", key, "must be specified!")
	}
	return val
}
