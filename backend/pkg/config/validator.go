package config

import (
	"fmt"
	"strings"
)

// ValidateConfig 验证配置是否有效
func ValidateConfig(config *Config) error {
	// 验证服务器配置
	if config.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	
	// 验证数据库配置
	if config.Database.Driver == "" {
		return fmt.Errorf("database driver is required")
	}
	
	if config.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}
	
	if config.Database.Port == "" {
		return fmt.Errorf("database port is required")
	}
	
	if config.Database.Username == "" {
		return fmt.Errorf("database username is required")
	}
	
	if config.Database.DBName == "" {
		return fmt.Errorf("database name is required")
	}
	
	// 验证数据库驱动类型
	if !strings.EqualFold(config.Database.Driver, "mysql") {
		return fmt.Errorf("unsupported database driver: %s", config.Database.Driver)
	}
	
	// 验证认证配置
	if config.Auth.JWTSecret == "" {
		return fmt.Errorf("JWT secret is required")
	}
	
	if config.Auth.JWTDuration <= 0 {
		return fmt.Errorf("JWT duration must be positive")
	}
	
	return nil
}