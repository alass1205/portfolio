package database

import (
    "database/sql"
    "fmt"
    "portfolio-backend/internal/config"

    "github.com/go-redis/redis/v8"
    _ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}

func ConnectRedis(cfg *config.Config) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr: cfg.RedisURL,
    })

    return rdb
}
