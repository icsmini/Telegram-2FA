package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/icsmini/Telegram-2FA/internal/app/apiserver"
	"log"
	"net/http"
	"time"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")

}

func main() {
	flag.Parse()

	fmt.Println(configPath)
	// Создание конфига
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	var srv = &http.Server{
		Addr:           config.BindAddr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Создание API сервера
	s := apiserver.New(config, srv)

	// Запуск сервера
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	//

}
