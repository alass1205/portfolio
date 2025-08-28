package sync

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "portfolio-backend/internal/services"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type GitHubHandler struct {
    db            *sql.DB
    githubService *services.GitHubService
}

func NewGitHubHandler(db *sql.DB) *GitHubHandler {
    return &GitHubHandler{
        db:            db,
        githubService: services.NewGitHubService(),
    }
}

func (h *GitHubHandler) SyncFromGitHub(c *fiber.Ctx) error {
    username := "alass1205"
    
    log.Printf("🚀 Synchronisation GitHub pour %s...", username)
    
    repos, err := h.githubService.GetUserRepos(username)
    if err != nil {
        log.Printf("Erreur récupération repos GitHub: %v", err)
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "Erreur lors de la récupération des repos GitHub",
            "error":   err.Error(),
        })
    }
    
    log.Printf("📦 %d repos trouvés", len(repos))
    
    filteredRepos := h.githubService.FilterInterestingRepos(repos)
    log.Printf("✅ %d repos filtrés", len(filteredRepos))
    
    // Nettoyer les anciens projets GitHub
    _, err = h.db.Exec("DELETE FROM projects WHERE github_url LIKE 'https://github.com/alass1205/%'")
    if err != nil {
        log.Printf("Erreur nettoyage: %v", err)
    }
    
    var syncCount int
    var featuredCount int
    
    for _, repo := range filteredRepos {
        category := h.githubService.MapToCategory(repo.Language, repo.Topics, repo.Name, repo.Description)
        
        // Générer description intelligente
        description := h.githubService.GenerateSmartDescription(repo, category)
        
        // Générer technologies intelligentes
        technologies := h.githubService.GenerateSmartTechnologies(repo, category)
        
        // Déterminer si featured (avec paramètre category)
        featured := h.githubService.ShouldBeFeatured(repo, category)
        
        if featured {
            featuredCount++
        }
        
        // Créer titre propre
        title := strings.ReplaceAll(repo.Name, "-", " ")
        title = strings.Title(title)
        
        // Sérialiser technologies
        techsJSON, _ := json.Marshal(technologies)
        
        query := `
            INSERT INTO projects (id, title, description, technologies, category, 
                                github_url, featured, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
        
        _, err := h.db.Exec(query,
            uuid.New(),
            title,
            description,
            techsJSON,
            category,
            repo.HTMLURL,
            featured,
            time.Now(),
            time.Now(),
        )
        
        if err != nil {
            log.Printf("Erreur insertion repo %s: %v", repo.Name, err)
            continue
        }
        
        featuredIcon := ""
        if featured {
            featuredIcon = "⭐"
        }
        
        log.Printf("✅ Ajouté: %s (%s) %s", title, category, featuredIcon)
        syncCount++
    }
    
    return c.JSON(fiber.Map{
        "success": true,
        "message": fmt.Sprintf("Synchronisation réussie: %d projets ajoutés (%d featured)", 
                              syncCount, featuredCount),
        "synced_count":   syncCount,
        "featured_count": featuredCount,
        "total_found":    len(repos),
        "filtered":       len(filteredRepos),
    })
}

func SetupGitHubRoutes(app fiber.Router, db *sql.DB) {
    handler := NewGitHubHandler(db)
    app.Post("/sync-github", handler.SyncFromGitHub)
}
