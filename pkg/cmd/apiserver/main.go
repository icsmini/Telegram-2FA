package main

import (
	"flag"
	"github.com/icsmini/Telegram-2FA/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")

}

func main() {
	// Создание конфига
	config := apiserver.NewConfig()

	// Создание API сервера
	s := apiserver.New(config)

	// Запуск сервера
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	//

}
