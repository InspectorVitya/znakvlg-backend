package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Config struct {
	JwtSecret string `envconfig:"JWT_SECRET" required:"true"`
	DataBase
	HTTP
}
type DataBase struct {
	DBURL string `envconfig:"DB_URL" required:"true"`
}

type HTTP struct {
	Port               string        `envconfig:"HTTP_PORT" required:"true"`
	ReadTimeout        time.Duration `envconfig:"HTTP_READ_TIMEOUT" required:"true"`
	WriteTimeout       time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" required:"true"`
	MaxHeaderMegabytes int           `envconfig:"HTTP_MAX_HEADER_BYTES" required:"true"`
}

func LoadConfig() *Config {

	out := zerolog.NewConsoleWriter()
	out.Out = os.Stderr
	logger := zerolog.New(out)

	for _, fileName := range []string{".env"} { //.env.local for local secrets (higher priority than .env)
		err := godotenv.Load(fileName) //in cycle cause first error in varargs prevents loading next files
		if err != nil {
			logger.Error().Err(fmt.Errorf("error loading %s fileName : %w", fileName, err)).Send()
		}
	}

	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		logger.Fatal().Err(err).Send()
	} else {
		logger.Info().Msg("Config initialized")
	}
	return &cfg
}
