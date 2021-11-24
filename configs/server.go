package configs

import (
	"os"
	"strconv"
)

type server struct {
	Port int
}

func setupServer() *server {
	var port int
	portEnv := os.Getenv("PORT")
	portInt, err := strconv.Atoi(portEnv)
	if err == nil {
		port = portInt
	}

	v := server{
		Port: port,
	}

	if v.Port == 0 {
		v.Port = 8080
	}

	return &v
}
