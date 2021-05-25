package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Configuration struct {
	APIPort  int
	AuthSeed string
	AuthKey  string
	ShowAuth bool

	Author     string
	Version    string
	SnakeColor string
	SnakeHead  string
	SnakeTail  string
	Shout      string
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

	// auth
	Config.AuthSeed = envHelper("BS_AUTH_SEED", generateRandomToken(""))
	key := envHelper("BS_AUTH_KEY", "")
	if key == "" {
		key = generateRandomToken(Config.AuthSeed)
		Config.ShowAuth = true
	}
	Config.AuthKey = key

	Config.Author = envHelper("BS_AUTHOR", "Someone Online")
	Config.Version = envHelper("BS_VERSION", "v0.0.1")

	// snake options; if not passed in, we make it random
	Config.SnakeColor = envHelper("BS_SNAKE_COLOR", getRandomColorHex())
	Config.SnakeHead = envHelper("BS_SNAKE_HEAD", getRandomHead())
	Config.SnakeTail = envHelper("BS_SNAKE_TAIL", getRandomTail())

	Config.Shout = envHelper("BS_SHOUT", getRandomShout())
}

func envHelper(variable, defaultValue string) string {
	found := os.Getenv(variable)
	if found == "" {
		found = defaultValue
	}
	return found
}

// generateRandomToken generates a new token for various uses
func generateRandomToken(seed string) (token string) {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Int63n(999999999999)
	r2 := rand.Int63n(999999999999)
	r3 := getRandomColorHex()
	r4 := getRandomHead()
	r5 := getRandomTail()
	str := fmt.Sprintf("%d-%s-%d-%s-%s-%s", r1, seed, r2, r3, r4, r5)
	hasher := md5.New()
	hasher.Write([]byte(str))
	token = hex.EncodeToString(hasher.Sum(nil))
	return token
}
