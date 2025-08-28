package services

import (
    "database/sql"
    "encoding/json"
    "log"
    "strings"
    "time"

    "github.com/google/uuid"
)

type SchedulerService struct {
    db            *sql.DB
    githubService *GitHubService
    ticker        *time.Ticker
    done          chan bool
}

func NewSchedulerService(db *sql.DB) *SchedulerService {
    return &SchedulerService{
        db:            db,
        githubService: NewGitHubService(),
        done:          make(chan bool),
    }
}

func (s *SchedulerService) StartAutoSync() {
    log.Println("Démarrage de la synchronisation automatique GitHub (toutes les heures)")
    
    // Synchronisation immédiate au démarrage
    go s.syncGitHubRepos()
    
    // Puis toutes les heures
    s.ticker = time.NewTicker(1 * time.Hour)
    
    go func() {
        for {
            select {
            case <-s.ticker.C:
                log.Println("Synchronisation automatique programmée...")
                s.syncGitHubRepos()
            case <-s.done:
                log.Println("Arrêt de la synchronisation automatique")
                return
            }
        }
    }()
}

func (s *SchedulerService) StopAutoSync() {
    if s.ticker != nil {
        s.ticker.Stop()
    }
    s.done <- true
}

func (s *SchedulerService) syncGitHubRepos() {
    username := "alass1205"
    
    repos, err := s.githubService.GetUserRepos(username)
    if err != nil {
        log.Printf("Erreur sync auto GitHub: %v", err)
        return
    }
    
    filteredRepos := s.githubService.FilterInterestingRepos(repos)
    
    // Vérifier s'il y a de nouveaux repos
    var existingCount int
    err = s.db.QueryRow("SELECT COUNT(*) FROM projects WHERE github_url LIKE 'https://github.com/alass1205/%'").Scan(&existingCount)
    if err != nil {
        log.Printf("Erreur vérification projets existants: %v", err)
        return
    }
    
    if len(filteredRepos) > existingCount {
        log.Printf("Nouveaux repos détectés: %d vs %d existants", len(filteredRepos), existingCount)
        
        // Nettoyer et resync tous les projets GitHub
        _, err = s.db.Exec("DELETE FROM projects WHERE github_url LIKE 'https://github.com/alass1205/%'")
        if err != nil {
            log.Printf("Erreur nettoyage: %v", err)
            return
        }
        
        var syncCount int
        for _, repo := range filteredRepos {
            if s.insertProject(repo) {
                syncCount++
            }
        }
        
        log.Printf("Sync auto réussie: %d projets mis à jour", syncCount)
    } else {
        log.Printf("Aucun nouveau repo (actuel: %d)", existingCount)
    }
}

func (s *SchedulerService) insertProject(repo GitHubRepo) bool {
    category := s.githubService.MapToCategory(repo.Language, repo.Topics, repo.Name, repo.Description)
    description := s.githubService.GenerateSmartDescription(repo, category)
    technologies := s.githubService.GenerateSmartTechnologies(repo, category)
    featured := s.githubService.ShouldBeFeatured(repo, category)
    
    // Créer titre propre
    title := strings.ReplaceAll(repo.Name, "-", " ")
    title = strings.Title(title)
    
    // Sérialiser technologies
    techsJSON, err := json.Marshal(technologies)
    if err != nil {
        log.Printf("Erreur sérialisation technologies pour %s: %v", repo.Name, err)
        return false
    }
    
    query := `
        INSERT INTO projects (id, title, description, technologies, category, 
                            github_url, featured, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
    
    _, err = s.db.Exec(query,
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
        return false
    }
    
    return true
}
