package basic

import (
	"trading-bot/internal/logger"
	"trading-bot/internal/platform"
)

type Strategy struct {
	platform platform.Interface
	logger   logger.Interface
}

// TODO
func New(platform platform.Interface, logger logger.Interface) *Strategy {
	return &Strategy{
		platform: platform,
		logger:   logger,
	}
}

func (b *Strategy) Run() error {
	return nil
}
