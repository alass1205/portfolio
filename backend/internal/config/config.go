package config

import "os"

type Config struct {
    DBHost        string
    DBPort        string
    DBName        string
    DBUser        string
    DBPassword    string
    RedisURL      string
    JWTSecret     string
    Port          string
    Env           string
    WebhookSecret string
    AutoSyncHours int
}

func Load() *Config {
    return &Config{
        DBHost:        getEnv("DB_HOST", "localhost"),
        DBPort:        getEnv("DB_PORT", "5432"),
        DBName:        getEnv("DB_NAME", "portfolio"),
        DBUser:        getEnv("DB_USER", "portfolio_user"),
        DBPassword:    getEnv("DB_PASSWORD", "portfolio_pass"),
        RedisURL:      getEnv("REDIS_URL", "localhost:6379"),
        JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
        Port:          getEnv("PORT", "8080"),
        Env:           getEnv("ENV", "development"),
        WebhookSecret: getEnv("GITHUB_WEBHOOK_SECRET", ""),
        AutoSyncHours: getEnvInt("AUTO_SYNC_HOURS", 1),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
    // Implementation pour convertir en int
    return defaultValue
}
