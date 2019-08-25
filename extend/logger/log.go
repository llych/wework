package logger

import (
	"os"
	"strings"
	"time"
	"wework/extend/conf"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Setup 日志初始化设置
func Setup() {
	switch strings.ToLower(conf.LoggerConf.Level) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    !conf.LoggerConf.Color,
		TimeFormat: time.RFC3339,
	})
	//.With().Caller().Logger()

}
