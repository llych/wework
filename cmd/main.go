package main

import (
	"fmt"
	"time"
	"wework/extend/cache"
	"wework/extend/conf"
	"wework/extend/logger"
	"wework/router"

	"github.com/rs/zerolog/log"
)

func main() {
	// 初始化日志
	logger.Setup()
	log.Info().Msgf("init config")
	conf.Setup()
	log.Info().Msgf("app start :%d", conf.HttpConf.Port)
	cache.InitCache()
	//wework.TextMsg("hello", "", "1")
	router := router.InitRouter()
	for {
		if err := router.Run(fmt.Sprintf(":%d", conf.HttpConf.Port)); err != nil {
			log.Error().Msg(err.Error())
		}
		time.Sleep(20 * time.Second)
	}

}
