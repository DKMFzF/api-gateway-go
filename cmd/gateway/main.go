package main

import (
	"log"

	"api-gateway/internal/config"
	"api-gateway/internal/proxy"
	"api-gateway/internal/router"
)

func main() {
	cfg := config.Load()

	proxies := make(map[string]*proxy.ProxyWrap)
	for name, url := range cfg.Services {
		p, err := proxy.NewProxy(url, cfg.ProxyTimeout)
		if err != nil {
			log.Fatalf("[Error] create proxy for %s: %v", name, err)
		}

		proxies[name] = p
	}

	r := router.New(cfg, proxies)

	log.Printf("api-gateway running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
