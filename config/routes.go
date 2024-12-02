package config

import (
	"os"
)

func Routes() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return port
}
