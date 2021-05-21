package api

import (
	"os"
	"strconv"
	"strings"
)

type Configuration struct {
	BS_API_PORT int
}

var Config *Configuration

func ConfigSetup() {
	if Config != nil {
		return
	}
	Config = &Configuration{}
	port := strings.Trim(envHelper("BS_API_PORT", "7000"), ":")
	parsedPort, err := strconv.ParseInt(port, 10, 64)
	if err == nil {
		// they passed a valid port
		Config.BS_API_PORT = int(parsedPort)
	}

}

func envHelper(variable, defaultValue string) string {
	found := os.Getenv(variable)
	if found == "" {
		found = defaultValue
	}
	return found
}
