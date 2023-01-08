package basic

import (
	"trading-bot/internal/logger"
	"trading-bot/internal/platform"
)

type Strategy struct {
	platform platform.Platform
	logger   logger.Logger
}

// TODO
func New(platform platform.Platform, logger logger.Logger) *Strategy {
	return &Strategy{
		platform: platform,
		logger:   logger,
	}
}

func (b *Strategy) Run() <-chan error {
	ch := make(chan error)
	return ch
}
