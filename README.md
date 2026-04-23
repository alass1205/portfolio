# 🚀 Portfolio — Alassane Sall

Portfolio fullstack avec synchronisation automatique GitHub.

**Stack :** Next.js 14 · Go (Fiber) · PostgreSQL · Redis

---

## 📁 Structure

```
portfolio/
├── backend/      # API Go/Fiber
└── frontend/     # Next.js 14
```

---

## ⚙️ Prérequis

- Go 1.21+
- Node.js 18+
- PostgreSQL 16
- Redis

---

## 🏃 Lancer en local

### 1. Cloner le repo

```bash
git clone https://github.com/alass1205/portfolio.git
cd portfolio
```

---

### 2. Backend

```bash
cd backend
```

Créer le fichier `.env` :

```env
DB_HOST=localhost
DB_PORT=5433
DB_NAME=portfolio
DB_USER=portfolio_user
DB_PASSWORD=portfolio_pass
REDIS_URL=localhost:6379
PORT=8080
ENV=development
```

Créer la table PostgreSQL :

```bash
psql -h localhost -p 5433 -U portfolio_user -d portfolio -c "
CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";
CREATE TABLE IF NOT EXISTS projects (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    technologies TEXT[] DEFAULT '{}',
    category    VARCHAR(100),
    github_url  VARCHAR(500),
    live_url    VARCHAR(500),
    image_url   VARCHAR(500),
    featured    BOOLEAN DEFAULT false,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);"
```

Lancer le backend :

```bash
go mod tidy
go run main.go
```

→ API disponible sur `http://localhost:8080`

---

### 3. Frontend

```bash
cd frontend
npm install
npm run dev
```

→ App disponible sur `http://localhost:3000`

---

## 🌐 Endpoints API

| Méthode | Endpoint | Description |
|---------|----------|-------------|
| GET | `/health` | Statut de l'API |
| GET | `/api/v1/projects` | Liste tous les projets |
| GET | `/api/v1/projects/category/:cat` | Projets par catégorie |
| POST | `/api/v1/sync/sync-github` | Sync manuelle GitHub |

**Catégories disponibles :** `blockchain` · `go` · `rust` · `javascript`

### Forcer une synchronisation GitHub

```bash
curl -X POST http://localhost:8080/api/v1/sync/sync-github
```

---

## 🔄 Synchronisation GitHub

Le backend sync automatiquement les repos publics de [alass1205](https://github.com/alass1205) :

- **Au démarrage** → sync immédiate
- **Toutes les heures** → resync si nouveaux repos détectés
- Les projets sont catégorisés et décrits automatiquement

---

## 🚀 Déploiement (gratuit)

| Composant | Plateforme | URL |
|-----------|-----------|-----|
| Frontend | Vercel | `alassane-sall.vercel.app` |
| Backend | Render | `portfolio-backend-xxxx.onrender.com` |
| PostgreSQL | Neon.tech | — |
| Redis | Upstash | — |

### Variables d'environnement en production (Render)

```env
DB_HOST=<neon-host>
DB_PORT=5432
DB_NAME=portfolio
DB_USER=<neon-user>
DB_PASSWORD=<neon-password>
REDIS_URL=<upstash-url>
PORT=8080
ENV=production
```

### Variables d'environnement en production (Vercel)

```env
NEXT_PUBLIC_API_URL=https://portfolio-backend-xxxx.onrender.com
```

---

## 👤 Auteur

**Alassane Sall** — Backend & Blockchain Developer  
[github.com/alass1205](https://github.com/alass1205)