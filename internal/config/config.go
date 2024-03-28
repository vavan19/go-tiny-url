package config

import (
	"errors"
	"flag"
)

const (
	maxRangeValue = 3521614606208
)

// Config holds the configuration for the application instance.
type Config struct {
	StartRange int // Starting range for this instance
	EndRange   int // Ending range for this instance
	MongoDBURL string
}

func Load() Config {
	var cfg Config

	flag.IntVar(&cfg.StartRange, "startRange", 0, "Starting range for this instance")
	flag.IntVar(&cfg.EndRange, "endRange", 1000, "Ending range for this instance")
	flag.StringVar(&cfg.MongoDBURL, "mongoDBURL", "mongodb://mongo:27017", "Ending mongoDBURL for this instance")
	flag.Parse()

	return cfg
}

func (c *Config) Validate() error {
	if c.StartRange < 0 || c.StartRange > maxRangeValue {
		return errors.New("invalid start range")
	}
	if c.EndRange < 0 || c.EndRange > maxRangeValue {
		return errors.New("invalid end range")
	}
	if c.StartRange > c.EndRange {
		return errors.New("start range cannot be greater than end range")
	}
	return nil
}
