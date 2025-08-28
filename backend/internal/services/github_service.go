package services

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

type GitHubRepo struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    FullName    string   `json:"full_name"`
    Description string   `json:"description"`
    Language    string   `json:"language"`
    UpdatedAt   string   `json:"updated_at"`
    HTMLURL     string   `json:"html_url"`
    Private     bool     `json:"private"`
    Fork        bool     `json:"fork"`
    Stargazers  int      `json:"stargazers_count"`
    Topics      []string `json:"topics"`
    Size        int      `json:"size"`
}

type GitHubService struct {
    client *http.Client
}

func NewGitHubService() *GitHubService {
    return &GitHubService{
        client: &http.Client{
            Timeout: 15 * time.Second,
        },
    }
}

func (g *GitHubService) GetUserRepos(username string) ([]GitHubRepo, error) {
    url := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=updated&per_page=100", username)
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("User-Agent", "Portfolio-App/1.0")
    req.Header.Set("Accept", "application/vnd.github.v3+json")

    resp, err := g.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("GitHub API error: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var repos []GitHubRepo
    if err := json.Unmarshal(body, &repos); err != nil {
        return nil, err
    }

    return repos, nil
}

func (g *GitHubService) MapToCategory(language string, topics []string, name string, description string) string {
    language = strings.ToLower(language)
    nameAndDesc := strings.ToLower(name + " " + description)
    
    // Blockchain keywords
    blockchainKeywords := []string{
        "blockchain", "defi", "nft", "smart", "contract", "ethereum", 
        "solana", "web3", "dapp", "crypto", "trading", "payment", 
        "channel", "token", "wallet", "marketplace", "finance",
    }
    
    for _, keyword := range blockchainKeywords {
        if strings.Contains(nameAndDesc, keyword) {
            return "blockchain"
        }
    }
    
    // Check topics
    for _, topic := range topics {
        topic = strings.ToLower(topic)
        if topic == "blockchain" || topic == "ethereum" || topic == "solidity" || 
           topic == "defi" || topic == "nft" || topic == "web3" || topic == "solana" {
            return "blockchain"
        }
    }
    
    // Map by language
    switch language {
    case "solidity":
        return "blockchain"
    case "rust":
        return "rust"
    case "go":
        return "go"
    case "javascript", "typescript":
        return "javascript"
    case "python":
        return "python"
    case "java":
        return "java"
    case "c++", "cpp":
        return "cpp"
    case "c":
        return "c"
    default:
        return "other"
    }
}

func (g *GitHubService) GenerateSmartDescription(repo GitHubRepo, category string) string {
    if repo.Description != "" && len(repo.Description) > 10 {
        return repo.Description
    }
    
    name := strings.ToLower(repo.Name)
    
    // Blockchain descriptions
    if category == "blockchain" {
        if strings.Contains(name, "nft") {
            return "Plateforme NFT avec smart contracts pour création, trading et gestion de tokens non-fongibles"
        }
        if strings.Contains(name, "trading") {
            return "Plateforme de trading décentralisée avec orderbook et fonctionnalités DeFi avancées"
        }
        if strings.Contains(name, "payment") || strings.Contains(name, "channel") {
            return "Système de payment channels pour transactions off-chain rapides et économiques"
        }
        if strings.Contains(name, "marketplace") {
            return "Marketplace décentralisée pour échange d'actifs numériques avec smart contracts"
        }
        if strings.Contains(name, "financial") || strings.Contains(name, "defi") {
            return "Protocole DeFi avec instruments financiers décentralisés et yield farming"
        }
        if strings.Contains(name, "solana") {
            return "Application décentralisée développée sur la blockchain Solana"
        }
        return "Application blockchain décentralisée avec smart contracts et intégration Web3"
    }
    
    // Other categories
    switch category {
    case "rust":
        if strings.Contains(name, "server") || strings.Contains(name, "http") {
            return "Serveur HTTP haute performance développé en Rust avec gestion asynchrone"
        }
        if strings.Contains(name, "filler") {
            return "Algorithme de remplissage optimisé avec IA et focus sur les performances"
        }
        if strings.Contains(name, "smart") {
            return "Système intelligent développé en Rust pour applications IoT et embarquées"
        }
        return "Application système haute performance développée en Rust avec gestion mémoire optimisée"
        
    case "go":
        if strings.Contains(name, "server") {
            return "Serveur web en Go avec architecture microservices et API REST"
        }
        if strings.Contains(name, "dashboard") {
            return "Dashboard de monitoring temps réel avec WebSockets et métriques système"
        }
        if strings.Contains(name, "api") {
            return "API REST robuste développée en Go avec authentification et base de données"
        }
        return "Application backend Go avec architecture scalable et performance optimisée"
        
    case "javascript":
        if strings.Contains(name, "graphql") {
            return "API GraphQL moderne avec résolveurs avancés, cache intelligent et subscriptions"
        }
        if strings.Contains(name, "react") {
            return "Application React avec interface utilisateur moderne et state management"
        }
        if strings.Contains(name, "node") {
            return "Application Node.js backend avec Express et intégration base de données"
        }
        if strings.Contains(name, "tracker") {
            return "Application de géolocalisation avec cartes interactives et tracking temps réel"
        }
        return "Application web JavaScript moderne avec interface utilisateur interactive"
        
    default:
        return fmt.Sprintf("Projet %s développé avec %s", category, strings.Title(repo.Language))
    }
}

func (g *GitHubService) GenerateSmartTechnologies(repo GitHubRepo, category string) []string {
    var technologies []string
    name := strings.ToLower(repo.Name)
    
    // Add main language
    if repo.Language != "" {
        technologies = append(technologies, strings.Title(repo.Language))
    }
    
    // Category-specific technologies
    switch category {
    case "blockchain":
        if strings.Contains(name, "solana") {
            technologies = append(technologies, "Solana", "Anchor", "Rust")
        } else {
            technologies = append(technologies, "Smart Contracts", "Web3.js")
            if repo.Language == "JavaScript" || repo.Language == "TypeScript" {
                technologies = append(technologies, "Ethereum", "DeFi")
            }
        }
        if strings.Contains(name, "nft") {
            technologies = append(technologies, "NFT", "IPFS")
        }
        if strings.Contains(name, "trading") {
            technologies = append(technologies, "DeFi", "AMM")
        }
        
    case "rust":
        technologies = append(technologies, "Tokio", "Async")
        if strings.Contains(name, "server") || strings.Contains(name, "http") {
            technologies = append(technologies, "HTTP", "Actix-web")
        }
        if strings.Contains(name, "smart") {
            technologies = append(technologies, "IoT", "Embedded")
        }
        
    case "go":
        technologies = append(technologies, "HTTP", "API")
        if strings.Contains(name, "dashboard") {
            technologies = append(technologies, "WebSocket", "Real-time")
        }
        technologies = append(technologies, "PostgreSQL", "Docker")
        
    case "javascript":
        if strings.Contains(name, "graphql") {
            technologies = append(technologies, "GraphQL", "Apollo", "Node.js")
        }
        if strings.Contains(name, "react") {
            technologies = append(technologies, "React", "Frontend")
        }
        technologies = append(technologies, "Node.js", "Express")
    }
    
    // Ensure at least 2 technologies
    if len(technologies) < 2 && repo.Language != "" {
        switch strings.ToLower(repo.Language) {
        case "javascript":
            technologies = append(technologies, "Node.js", "Frontend")
        case "go":
            technologies = append(technologies, "Backend", "API")
        case "rust":
            technologies = append(technologies, "Systems", "Performance")
        case "python":
            technologies = append(technologies, "Backend", "Data")
        }
    }
    
    return technologies
}

func (g *GitHubService) ShouldBeFeatured(repo GitHubRepo, category string) bool {
    // Featured if has stars
    if repo.Stargazers > 0 {
        return true
    }
    
    // Featured if blockchain or rust
    if category == "blockchain" || category == "rust" {
        return true
    }
    
    // Featured if long description
    if len(repo.Description) > 60 {
        return true
    }
    
    // Featured if large repo
    if repo.Size > 1000 {
        return true
    }
    
    // Featured if important keywords
    nameAndDesc := strings.ToLower(repo.Name + " " + repo.Description)
    importantKeywords := []string{
        "trading", "marketplace", "dashboard", "api", "smart", "defi", 
        "nft", "graphql", "server", "platform",
    }
    
    for _, keyword := range importantKeywords {
        if strings.Contains(nameAndDesc, keyword) {
            return true
        }
    }
    
    return false
}

func (g *GitHubService) FilterInterestingRepos(repos []GitHubRepo) []GitHubRepo {
    var filtered []GitHubRepo
    
    for _, repo := range repos {
        if repo.Fork || repo.Private {
            continue
        }
        
        if repo.Language == "" && repo.Description == "" {
            continue
        }
        
        name := strings.ToLower(repo.Name)
        skipKeywords := []string{"template", "test", "example", "hello", "world"}
        shouldSkip := false
        
        for _, keyword := range skipKeywords {
            if strings.Contains(name, keyword) {
                shouldSkip = true
                break
            }
        }
        
        if shouldSkip {
            continue
        }
        
        filtered = append(filtered, repo)
    }
    
    return filtered
}
