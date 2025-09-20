package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Services     map[string]string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	ProxyTimeout time.Duration
}

func parseServices(raw string) map[string]string {
	out := make(map[string]string)

	if raw == "" {
		log.Printf("[ERROR]: raw none")
	}

	parts := strings.Split(raw, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		kv := strings.SplitN(p, "=", 2)
		if len(kv) != 2 {
			continue
		}

		out[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}

	return out
}

func Load() *Config {

	_ = godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	srv := parseServices(os.Getenv("SERVICES"))

	readSec := getIntEnv("READ_TIMEOUT_SEC", 15)
	writeSec := getIntEnv("WRITE_TIMEOUT_SEC", 15)
	proxySec := getIntEnv("PROXY_TIMEOUT_SEC", 10)

	return &Config{
		Port:         port,
		Services:     srv,
		ReadTimeout:  time.Duration(readSec) * time.Second,
		WriteTimeout: time.Duration(writeSec) * time.Second,
		ProxyTimeout: time.Duration(proxySec) * time.Second,
	}
}

func getIntEnv(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}
