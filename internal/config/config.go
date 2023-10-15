package config

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
	"time"
)

const (
	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
	defaultLimiterRPS             = 10
	defaultLimiterBurst           = 2
	defaultLimiterTTL             = 10 * time.Minute
	defaultVerificationCodeLength = 8

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		Enviromnet string
		Postgres   PostgresConfig
		HTTP       HTTPConfig
		Auth       AuthConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}

	AuthConfig struct {
		JWT                    JWTConfig
		PasswordSalt           string
		VerificationCodeLength int `mapstructure:"verificationCodeLength"`
	}

	PostgresConfig struct {
		Host     string
		User     string
		Port     string
		Password string
		Name     string
		SSLMode  string
	}
)

func Init(configsDir string) (*Config, error) {
	populateDefaults()

	if err := gotenv.OverLoad(); err != nil {
		return nil, err
	}

	if err := parseConfigFile(configsDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth.verificationCodeLength", &cfg.Auth.VerificationCodeLength); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")

	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.Name = os.Getenv("POSTGRES_Name")
	cfg.Postgres.SSLMode = os.Getenv("POSTGRES_SSLMODE")

	cfg.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	/*viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
	viper.SetDefault("auth.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("auth.refreshTokenTTL", defaultRefreshTokenTTL)
	viper.SetDefault("auth.verificationCodeLength", defaultVerificationCodeLength)*/
}
