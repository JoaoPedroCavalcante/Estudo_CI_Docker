package database

import (
    "fmt"
    "log"
    "os"

    "github.com/guilhermeonrails/api-go-gin/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConectaComBancoDeDados() error {
    host := getEnv("DB_HOST", "localhost")
    user := getEnv("DB_USER", "root")
    password := getEnv("DB_PASSWORD", "root")
    dbname := getEnv("DB_NAME", "root")
    port := getEnv("DB_PORT", "5432")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbname, port)

    db, err := gorm.Open(postgres.Open(dsn))
    if err != nil {
        log.Printf("[error] failed to initialize database, got error %v", err)
        return err
    }

    DB = db
    if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
        return err
    }
    return nil
}

func getEnv(key, defaultVal string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return defaultVal
}
