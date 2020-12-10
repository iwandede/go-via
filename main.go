package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/iwandede/go-via/middleware"
	"github.com/iwandede/go-via/server"

	AppConfig "github.com/iwandede/go-via/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		env    = flag.String("env", "production", "environment variable")
		config = flag.String("config", "./config/production.yaml", "config file")
	)
	flag.Parse()

	file, err := os.Open(*config)
	if err != nil {
		log.Fatalf("failed to open config file %v: %s", config, err)
		os.Exit(1)
	}
	defer file.Close()

	conf, err := AppConfig.NewConfigFromYAML(file)
	if err != nil {
		log.Fatalf("Error parse file %v", err)
	}

	servers := server.NewAppHttpServer(conf)
	router := servers.InitRouter()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", servers.Server.Port),
		Handler:      middleware.WrapHandler(router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Service Started on Port %s & Environment %v", s.Addr, *env)
	log.Debug(s.ListenAndServe())
}
