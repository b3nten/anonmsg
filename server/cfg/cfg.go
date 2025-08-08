package cfg

import (
	"log"

	"github.com/go-viper/mapstructure/v2"
	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Port                 string   `koanf:"port"`
	DatabaseURL          string   `koanf:"database_url"`
	LogLevel             string   `koanf:"log_level"`
	LogFormat            string   `koanf:"log_format"`
	EnableDebugEndpoints bool     `koanf:"enable_debug_endpoints"`
	EnableMetrics        bool     `koanf:"enable_metrics"`
	EnableDocs           bool     `koanf:"enable_docs"`
	AllowedOrigins       []string `koanf:"allowed_origins"`
	AllowedMethods       []string `koanf:"allowed_methods"`
}

func Init(configFile string) (Config, error) {
	godotenv.Load()
	k := koanf.New(".")
	if err := k.Load(file.Provider(configFile), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	k.Load(env.Provider(".", env.Opt{}), nil)
	var config Config
	return config, k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{
		DecoderConfig: &mapstructure.DecoderConfig{
			ErrorUnset: true,
		},
	})
}
