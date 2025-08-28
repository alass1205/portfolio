package projects

import (
    "database/sql"
    "encoding/json"
    "log"
    "portfolio-backend/internal/models"
    "time"

    "github.com/go-redis/redis/v8"
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
)

type Handler struct {
    db  *sql.DB
    rdb *redis.Client
}

// GetAllProjects récupère tous les projets
func (h *Handler) GetAllProjects(c *fiber.Ctx) error {
    query := `
        SELECT id, title, description, technologies, category, 
               github_url, live_url, image_url, featured, created_at, updated_at
        FROM projects 
        ORDER BY created_at DESC`

    rows, err := h.db.Query(query)
    if err != nil {
        log.Printf("Erreur GetAllProjects: %v", err)
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "Erreur lors de la récupération des projets",
        })
    }
    defer rows.Close()

    var projects []models.Project

    for rows.Next() {
        var project models.Project
        var techsJSON []byte

        err := rows.Scan(
            &project.ID, &project.Title, &project.Description,
            &techsJSON, &project.Category, &project.GitHubURL,
            &project.LiveURL, &project.ImageURL, &project.Featured,
            &project.CreatedAt, &project.UpdatedAt,
        )
        if err != nil {
            log.Printf("Erreur scan: %v", err)
            continue
        }

        // Désérialiser les technologies JSON
        if err := json.Unmarshal(techsJSON, &project.Technologies); err != nil {
            log.Printf("Erreur unmarshal technologies: %v", err)
            project.Technologies = []string{}
        }

        projects = append(projects, project)
    }

    return c.JSON(fiber.Map{
        "success": true,
        "data":    projects,
        "count":   len(projects),
    })
}

// GetProjectByID récupère un projet par ID
func (h *Handler) GetProjectByID(c *fiber.Ctx) error {
    idParam := c.Params("id")
    projectID, err := uuid.Parse(idParam)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "success": false,
            "message": "ID de projet invalide",
        })
    }

    query := `
        SELECT id, title, description, technologies, category,
               github_url, live_url, image_url, featured, created_at, updated_at
        FROM projects WHERE id = $1`

    var project models.Project
    var techsJSON []byte

    err = h.db.QueryRow(query, projectID).Scan(
        &project.ID, &project.Title, &project.Description,
        &techsJSON, &project.Category, &project.GitHubURL,
        &project.LiveURL, &project.ImageURL, &project.Featured,
        &project.CreatedAt, &project.UpdatedAt,
    )

    if err == sql.ErrNoRows {
        return c.Status(404).JSON(fiber.Map{
            "success": false,
            "message": "Projet non trouvé",
        })
    }
    if err != nil {
        log.Printf("Erreur GetProjectByID: %v", err)
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "Erreur serveur",
        })
    }

    // Désérialiser les technologies
    if err := json.Unmarshal(techsJSON, &project.Technologies); err != nil {
        project.Technologies = []string{}
    }

    return c.JSON(fiber.Map{
        "success": true,
        "data":    project,
    })
}

// CreateProject crée un nouveau projet
func (h *Handler) CreateProject(c *fiber.Ctx) error {
    var req models.CreateProjectRequest

    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "success": false,
            "message": "Données invalides",
        })
    }

    // Valider les données
    if req.Title == "" || req.Description == "" || req.Category == "" {
        return c.Status(400).JSON(fiber.Map{
            "success": false,
            "message": "Title, description et category sont requis",
        })
    }

    // Sérialiser les technologies en JSON
    techsJSON, err := json.Marshal(req.Technologies)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "Erreur lors du traitement des technologies",
        })
    }

    // Insérer en base
    projectID := uuid.New()
    now := time.Now()

    query := `
        INSERT INTO projects (id, title, description, technologies, category,
                            github_url, live_url, image_url, featured, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id`

    err = h.db.QueryRow(query,
        projectID, req.Title, req.Description, techsJSON, req.Category,
        req.GitHubURL, req.LiveURL, req.ImageURL, req.Featured, now, now,
    ).Scan(&projectID)

    if err != nil {
        log.Printf("Erreur CreateProject: %v", err)
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "Erreur lors de la création du projet",
        })
    }

    return c.Status(201).JSON(fiber.Map{
        "success": true,
        "message": "Projet créé avec succès",
        "data": fiber.Map{
            "id": projectID,
        },
    })
}

// GetProjectsByCategory récupère les projets par catégorie
func (h *Handler) GetProjectsByCategory(c *fiber.Ctx) error {
    category := c.Params("category")

    query := `
        SELECT id, title, description, technologies, category,
               github_url, live_url, image_url, featured, created_at, updated_at
        FROM projects 
        WHERE category = $1 
        ORDER BY created_at DESC`

    rows, err := h.db.Query(query, category)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "success": false,
            "message": "Erreur serveur",
        })
    }
    defer rows.Close()

    var projects []models.Project

    for rows.Next() {
        var project models.Project
        var techsJSON []byte

        err := rows.Scan(
            &project.ID, &project.Title, &project.Description,
            &techsJSON, &project.Category, &project.GitHubURL,
            &project.LiveURL, &project.ImageURL, &project.Featured,
            &project.CreatedAt, &project.UpdatedAt,
        )
        if err != nil {
            log.Printf("Erreur scan: %v", err)
            continue
        }

        json.Unmarshal(techsJSON, &project.Technologies)
        projects = append(projects, project)
    }

    return c.JSON(fiber.Map{
        "success": true,
        "data":    projects,
        "count":   len(projects),
        "category": category,
    })
}

// UpdateProject et DeleteProject (simplifié pour l'instant)
func (h *Handler) UpdateProject(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "Update à implémenter"})
}

func (h *Handler) DeleteProject(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "Delete à implémenter"})
}
