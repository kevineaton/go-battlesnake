package api

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Configuration struct {
	APIPort int

	Author     string
	Version    string
	SnakeColor string
	SnakeHead  string
	SnakeTail  string
}

var Config *Configuration

func ConfigSetup() {
	if Config != nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	Config = &Configuration{}
	port := strings.Trim(envHelper("BS_API_PORT", "7000"), ":")
	parsedPort, err := strconv.ParseInt(port, 10, 64)
	if err == nil {
		// they passed a valid port
		Config.APIPort = int(parsedPort)
	}

	Config.Author = envHelper("BS_AUTHOR", "Someone Online")
	Config.Version = envHelper("BS_VERSION", "v0.0.1")

	// snake options; if not passed in, we make it random
	Config.SnakeColor = envHelper("BS_SNAKE_COLOR", getRandomColorHex())
	Config.SnakeHead = envHelper("BS_SNAKE_HEAD", getRandomHead())
	Config.SnakeTail = envHelper("BS_SNAKE_TAIL", getRandomTail())
}

func envHelper(variable, defaultValue string) string {
	found := os.Getenv(variable)
	if found == "" {
		found = defaultValue
	}
	return found
}
