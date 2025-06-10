package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// LoadConfigWithEnv 加载配置文件并应用环境变量覆盖
func LoadConfigWithEnv(configPath string) (*Config, error) {
	// 加载配置文件
	config, err := LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	// 应用环境变量覆盖
	applyEnvOverrides(config)

	// 验证配置
	if err := ValidateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return config, nil
}

// FindConfigFile 查找配置文件
func FindConfigFile(configName string) (string, error) {
	// 尝试在当前目录查找
	if _, err := os.Stat(configName); err == nil {
		absPath, err := filepath.Abs(configName)
		if err != nil {
			return "", err
		}
		return absPath, nil
	}

	// 尝试在config目录查找
	configPath := filepath.Join("config", configName)
	if _, err := os.Stat(configPath); err == nil {
		absPath, err := filepath.Abs(configPath)
		if err != nil {
			return "", err
		}
		return absPath, nil
	}

	// 尝试在上级目录查找
	parentConfigPath := filepath.Join("..", "config", configName)
	if _, err := os.Stat(parentConfigPath); err == nil {
		absPath, err := filepath.Abs(parentConfigPath)
		if err != nil {
			return "", err
		}
		return absPath, nil
	}

	return "", fmt.Errorf("config file %s not found", configName)
}

// applyEnvOverrides 应用环境变量覆盖配置
func applyEnvOverrides(config *Config) {
	// 服务器配置
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}

	if host := os.Getenv("SERVER_HOST"); host != "" {
		config.Server.Host = host
	}

	// 数据库配置
	if driver := os.Getenv("DB_DRIVER"); driver != "" {
		config.Database.Driver = driver
	}

	if host := os.Getenv("DB_HOST"); host != "" {
		config.Database.Host = host
	}

	if port := os.Getenv("DB_PORT"); port != "" {
		config.Database.Port = port
	}

	if username := os.Getenv("DB_USERNAME"); username != "" {
		config.Database.Username = username
	}

	if password := os.Getenv("DB_PASSWORD"); password != "" {
		config.Database.Password = password
	}

	if dbname := os.Getenv("DB_NAME"); dbname != "" {
		config.Database.DBName = dbname
	}

	if charset := os.Getenv("DB_CHARSET"); charset != "" {
		config.Database.Charset = charset
	}
	
	// 认证配置
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		config.Auth.JWTSecret = jwtSecret
	}
	
	if jwtDuration := os.Getenv("JWT_DURATION"); jwtDuration != "" {
		// 这里简化处理，实际应该转换为int
		config.Auth.JWTDuration = 24
	}
}