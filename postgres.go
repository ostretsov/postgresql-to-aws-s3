package main

import (
	"fmt"
	"os/exec"
	"time"
)

// PostgreSQL is an `Exporter` interface that backs up a PostgreSQL database via the `pg_dump` command
type PostgreSQL struct {
	// DB Host (e.g. 127.0.0.1)
	Host string
	// DB Port (e.g. 5432)
	Port string
	// DB Name
	DB string
	// Connection Username
	Username string
	// Extra pg_dump options
	// e.g []string{"--inserts"}
	Options []string
}

func (x PostgreSQL) Dump() (path string, err error) {
	path = fmt.Sprintf(`bu_%v_%v.sql.tar.gz`, x.DB, time.Now().Unix())
	options := append(x.dumpOptions(), "-Fc", fmt.Sprintf(`-f%v`, path))
	out, err := exec.Command("pg_dump", options...).Output()
	if err != nil {
		return path, fmt.Errorf("error: %s; output: %s", err.Error(), string(out))
	}
	return path, nil
}

func (x PostgreSQL) dumpOptions() []string {
	options := x.Options

	if x.DB != "" {
		options = append(options, fmt.Sprintf(`-d%v`, x.DB))
	}

	if x.Host != "" {
		options = append(options, fmt.Sprintf(`-h%v`, x.Host))
	}

	if x.Port != "" {
		options = append(options, fmt.Sprintf(`-p%v`, x.Port))
	}

	if x.Username != "" {
		options = append(options, fmt.Sprintf(`-U%v`, x.Username))
	}

	return options
}
