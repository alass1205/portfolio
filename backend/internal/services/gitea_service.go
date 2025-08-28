package services

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

type GiteaRepo struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    FullName    string `json:"full_name"`
    Description string `json:"description"`
    Language    string `json:"language"`
    UpdatedAt   string `json:"updated_at"`
    HTMLURL     string `json:"html_url"`
    Private     bool   `json:"private"`
}

type GiteaService struct {
    baseURL string
    client  *http.Client
}

func NewGiteaService() *GiteaService {
    return &GiteaService{
        baseURL: "https://learn.zone01dakar.sn/api/v1",
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (g *GiteaService) GetUserRepos(username string) ([]GiteaRepo, error) {
    url := fmt.Sprintf("%s/users/%s/repos", g.baseURL, username)
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := g.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var repos []GiteaRepo
    if err := json.Unmarshal(body, &repos); err != nil {
        return nil, err
    }

    return repos, nil
}

func (g *GiteaService) MapToCategory(language string) string {
    language = strings.ToLower(language)
    
    switch {
    case strings.Contains(language, "solidity") || strings.Contains(language, "javascript") && strings.Contains(language, "web3"):
        return "blockchain"
    case language == "rust":
        return "rust"
    case language == "javascript" || language == "typescript" || language == "html" || language == "css":
        return "javascript"
    case language == "go":
        return "go"
    case language == "java":
        return "java"
    default:
        return "other"
    }
}
