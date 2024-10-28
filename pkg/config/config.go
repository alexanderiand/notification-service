package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-yaml/yaml"
)

// environment variables
const (
	cfgFilePath = "CONFIG_FILE_PATH"
	httpSrvAddr = "HTTP_SERVER_ADDR"
	dbFilePath  = "DATABASE_FILE_PATH"
)

// errors
var (
	ErrConfigFileNotExists     = errors.New("error, the config file does not exists")
	ErrInvalidConfigFilePath   = errors.New("error, invalid config file path")
	ErrInvalidConfigParams     = errors.New("error, invalid config params")
	ErrHTTPServerAddrNotExists = errors.New("error, http server address does not exists")
	ErrDBFilePathNotExists     = errors.New("error, the database file path (for sqlite3) does not exists")
)

// Configuration structures
// Config
type (
	Config struct {
		Env              string     `yaml:"env"`
		Service          Service    `yaml:"service"`
		HTTPServer       HTTPserver `yaml:"http_server"`
		DatabaseFilePath string     `env:"DATABASE_FILE_PATH"`
	}

	Service struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTPserver struct {
		Addr          string        `env:"HTTP_SERVER_ADDR"`
		RWTimeout     time.Duration `yaml:"rw_timeouts"`
		IdleTimeout   time.Duration `yaml:"idle_timeout"`
		MaxHeaderSize int           `yaml:"max_header_size"`
	}
)

// InitConfig load config files, parse, validate, and set config params into *Config, return *Config, error
// If config file path is empty or not set in .env, return ErrInvalidConfigFilePath
// If config files does not exists, return ErrConfigFileNotExists
// If config params invalid, return ErrInvalidConfigParams
func InitConfig() (*Config, error) {
	cfg := &Config{}

	// load config file
	cfgFile, err := loadConfigFile()
	if err != nil {
		return nil, err
	}

	// parse this file
	if cfg, err = parseConfigFile(cfgFile, cfg); err != nil {
		return nil, err
	}

	// load environment variables, and set cfg params
	if err := loadAndSetEnv(cfg); err != nil {
		return nil, err
	}

	// validate config params
	if err := simpleCfgParamsValidator(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// loadAndSetEnv
func loadAndSetEnv(cfg *Config) error {
	op := "pkg.config.loadAndSet"

	addr, ok := os.LookupEnv(httpSrvAddr)
	if !ok {
		return fmt.Errorf("%s %w", op, ErrHTTPServerAddrNotExists)
	}
	cfg.HTTPServer.Addr = addr

	dbPath, ok := os.LookupEnv(dbFilePath)
	if !ok {
		return fmt.Errorf("%s %w", op, ErrDBFilePathNotExists)
	}
	cfg.DatabaseFilePath = dbPath

	return nil
}

// loadConfigFile get a config file path from env, and load the configs/local.yaml,
// return *os.FIle, error
// If config file path is empty or not set in .env, return ErrConfigFilePathNotExists
// If config files does not exists, return ErrConfigFileNotExists
// If something other going wrong, return the unknown error
func loadConfigFile() (cfgFile *os.File, err error) {
	op := "pkg.config.loadConfigFile"

	cfgPath, ok := os.LookupEnv(cfgFilePath)
	if !ok {
		return nil, fmt.Errorf("%s %w", op, ErrInvalidConfigFilePath)
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s %w", op, ErrConfigFileNotExists)
	}

	cfgFile, err = os.Open(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return cfgFile, nil
}

// parseConfigFile unmarshal local.yaml file, set the configs params into *Config,
// return *Config, error
// If the config params is invalid, return ErrInvalidConfigParams
// If while parsing accrued something going wrong, return the unknown error
func parseConfigFile(cfgFile *os.File, cfg *Config) (*Config, error) {
	op := "pkg.config.parseConfigFile"

	if err := yaml.NewDecoder(cfgFile).Decode(cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return cfg, nil
}

// isValidCfgParams , simple config params validator
func simpleCfgParamsValidator(cfg *Config) error {
	op := "pkg.config.isValidCfgParams"

	if cfg == nil {
		return fmt.Errorf("%s %w", op, ErrConfigFileNotExists)
	}

	// env
	if err := detectEmptyParam(cfg.Env, "env"); err != nil {
		return err
	}

	cfgVars := map[string]bool{"local": true, "dev": true, "prod": true}
	if cfg.Env != "" {
		if _, ext := cfgVars[cfg.Env]; !ext {
			return fmt.Errorf("%s %w", op, errors.New("error, the env config param, does not exists"))
		}
	}

	// service
	if cfg.Service.Name == "" || cfg.Service.Version == "" {
		return fmt.Errorf("%s %w", op, errors.New("error, service name or version invalid"))
	}

	// httpServer
	// addr
	if err := detectEmptyParam(cfg.HTTPServer.Addr, "http_server_addr"); err != nil {
		return err
	}

	if cfg.HTTPServer.IdleTimeout <= 0 || cfg.HTTPServer.RWTimeout <= 0 {
		return fmt.Errorf("%s %w", op, errors.New("error, invalid http server timeouts"))
	}

	// dbFilePath
	if err := detectEmptyParam(cfg.DatabaseFilePath, "the database file path"); err != nil {
		return err
	}

	return nil
}

func detectEmptyParam(param, paramName string) error {
	op := "pkg.config.detectEmptyParam"

	errMsg := fmt.Sprintf("error, the %s param is empty", paramName)
	if param == "" {
		return fmt.Errorf("%s %w", op, errors.New(errMsg))
	}

	return nil
}
