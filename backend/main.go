package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "portfolio-backend/internal/config"
    "portfolio-backend/internal/database"
    "portfolio-backend/internal/handlers/projects"
    "portfolio-backend/internal/handlers/sync"
    "portfolio-backend/internal/middleware"
    "portfolio-backend/internal/services"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
)

func main() {
    // Charger variables d'environnement
    if err := godotenv.Load(); err != nil {
        log.Println("Pas de fichier .env trouvé")
    }

    // Configuration
    cfg := config.Load()

    // Connexion base de données
    db, err := database.Connect(cfg)
    if err != nil {
        log.Fatal("Erreur connexion DB:", err)
    }
    defer db.Close()

    // Connexion Redis
    rdb := database.ConnectRedis(cfg)

    // Démarrer le service de synchronisation automatique
    scheduler := services.NewSchedulerService(db)
    scheduler.StartAutoSync()
    defer scheduler.StopAutoSync()

    // Créer app Fiber
    app := fiber.New(fiber.Config{
        ErrorHandler: middleware.ErrorHandler,
    })

    // Middleware globaux
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
    }))
    app.Use(logger.New())

    // Routes API
    api := app.Group("/api/v1")

    // Routes projets
    projectsGroup := api.Group("/projects")
    projects.SetupRoutes(projectsGroup, db, rdb)

    // Routes synchronisation GitHub
    syncGroup := api.Group("/sync")
    sync.SetupGitHubRoutes(syncGroup, db)

    // Route santé
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status": "ok",
            "message": "Portfolio API is running",
        })
    })

    // Gestionnaire de signaux pour arrêt propre
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    go func() {
        <-c
        log.Println("Arrêt du serveur en cours...")
        scheduler.StopAutoSync()
        app.Shutdown()
    }()

    // Démarrer serveur
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Serveur démarré sur le port %s", port)
    log.Printf("Synchronisation automatique GitHub activée")
    log.Printf("Endpoints disponibles:")
    log.Printf("  GET  /api/v1/projects")
    log.Printf("  POST /api/v1/sync/sync-github")
    log.Fatal(app.Listen(":" + port))
}
