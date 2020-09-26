package config

import (
	"github.com/maxidelgado/toolkit-go/pkg/logger"
	"github.com/maxidelgado/toolkit-go/pkg/router"
	"sync"
)

var cfg config
var once sync.Once

type config struct {
	Router router.Config
}

func Get() config {
	once.Do(func() {
		cfg = config{
			Router: router.Config{
				Logging: logger.Config{Level: "info"},
			},
		}
	})
	return cfg
}
