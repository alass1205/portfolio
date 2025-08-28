package models

import (
    "time"
    "github.com/google/uuid"
)

type Project struct {
    ID          uuid.UUID `json:"id" db:"id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    Technologies []string `json:"technologies" db:"technologies"`
    Category    string    `json:"category" db:"category"`
    GitHubURL   *string   `json:"github_url" db:"github_url"`      // Pointeur pour NULL
    LiveURL     *string   `json:"live_url" db:"live_url"`          // Pointeur pour NULL
    ImageURL    *string   `json:"image_url" db:"image_url"`        // Pointeur pour NULL
    Featured    bool      `json:"featured" db:"featured"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateProjectRequest struct {
    Title        string   `json:"title"`
    Description  string   `json:"description"`
    Technologies []string `json:"technologies"`
    Category     string   `json:"category"`
    GitHubURL    *string  `json:"github_url"`    // Pointeur pour optionnel
    LiveURL      *string  `json:"live_url"`      // Pointeur pour optionnel
    ImageURL     *string  `json:"image_url"`     // Pointeur pour optionnel
    Featured     bool     `json:"featured"`
}
