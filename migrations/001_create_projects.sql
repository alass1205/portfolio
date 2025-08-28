-- Table des projets
CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    technologies JSONB NOT NULL DEFAULT '[]',
    category VARCHAR(100) NOT NULL,
    github_url VARCHAR(500),
    live_url VARCHAR(500),
    image_url VARCHAR(500),
    featured BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index pour les catégories
CREATE INDEX IF NOT EXISTS idx_projects_category ON projects(category);
CREATE INDEX IF NOT EXISTS idx_projects_featured ON projects(featured);
CREATE INDEX IF NOT EXISTS idx_projects_created_at ON projects(created_at DESC);

-- Données de test
INSERT INTO projects (title, description, technologies, category, github_url, featured) VALUES
('DeFi Trading Bot', 'Bot de trading automatisé pour protocoles DeFi avec stratégies d''arbitrage', '["Solidity", "Web3.js", "Node.js", "TypeScript"]', 'blockchain', 'https://github.com/username/defi-bot', true),
('Rust Web Server', 'Serveur web haute performance en Rust avec gestion async', '["Rust", "Tokio", "Actix-web", "PostgreSQL"]', 'rust', 'https://github.com/username/rust-server', true),
('React Dashboard', 'Dashboard d''analytics avec graphiques interactifs', '["React", "TypeScript", "D3.js", "TailwindCSS"]', 'javascript', 'https://github.com/username/react-dashboard', false),
('Go Microservice', 'API REST en Go avec architecture microservices', '["Go", "Fiber", "Docker", "PostgreSQL"]', 'go', 'https://github.com/username/go-api', false);
