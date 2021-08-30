package external_server_ticker

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type ServerTickerConfig struct {
	Url         string `yaml:"url" env:"SERVER_TICKER_URL"`
	IntervalSec int    `yaml:"interval_sec" env:"SERVER_TICKER_INTERVAL"`
}

type ServerTicker struct {
	config *ServerTickerConfig
	logger *zap.Logger
}

func NewServerTicker(config *ServerTickerConfig, logger *zap.Logger) *ServerTicker {
	return &ServerTicker{config: config, logger: logger}
}

func (ticker *ServerTicker) ping() (string, error) {
	res, err := http.Get(ticker.config.Url)
	if err != nil {
		return "", err
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("HTTP status code error: %d %s", res.StatusCode, res.Status))
	}

	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(all), nil
}

func (ticker *ServerTicker) tick() {
	serverAnswer, err := ticker.ping()
	if err != nil {
		ticker.logger.Info("Can't ping external server", zap.Error(err))
	} else if serverAnswer != "pong" {
		ticker.logger.Info("External server's ping return bad answer", zap.String("server_answer", serverAnswer))
	} else {
		ticker.logger.Info("External server's ping ok")
	}
}

func (ticker *ServerTicker) Start() {
	if ticker.config.IntervalSec > 0 && ticker.config.Url != "" {
		go func() {
			t := time.NewTicker(time.Second * time.Duration(ticker.config.IntervalSec))
			ticker.logger.Info("Start external server ticker ...")
			for range t.C {
				ticker.tick()
			}
		}()
	} else {
		ticker.logger.Info("External server ticker disable")
	}
}
