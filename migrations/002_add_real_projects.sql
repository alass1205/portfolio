-- Supprimer les anciens projets de test
DELETE FROM projects;

-- Insérer tes vrais projets
INSERT INTO projects (title, description, technologies, category, github_url, featured) VALUES
-- Projets Blockchain
('Financial Instruments', 'Système d''instruments financiers décentralisés avec smart contracts avancés', '["JavaScript", "Solidity", "Web3.js", "Node.js"]', 'blockchain', 'https://learn.zone01dakar.sn/git/alassall/financial-instruments', true),

('Payment Channel', 'Implémentation de channels de paiement pour transactions hors-chaîne avec signatures cryptographiques', '["JavaScript", "Solidity", "Ethereum", "Web3"]', 'blockchain', 'https://learn.zone01dakar.sn/git/alassall/payment-channel', true),

('Mini Payment Channel', 'Version simplifiée des payment channels pour transactions rapides et économiques', '["Solidity", "Smart Contracts", "Ethereum"]', 'blockchain', 'https://learn.zone01dakar.sn/git/alassall/mini-payment-channel', false),

('NFT Marketplace', 'Marketplace décentralisée pour l''échange de NFTs avec interface moderne', '["JavaScript", "Solidity", "React", "Web3.js"]', 'blockchain', 'https://learn.zone01dakar.sn/git/alassall/nft-marketplace', true),

('Non-Fungible Cats', 'Collection de NFTs de chats avec génération procédurale et rareté', '["HTML", "JavaScript", "Solidity", "NFT"]', 'blockchain', 'https://learn.zone01dakar.sn/git/alassall/non-fungible-cats', false),

('Solana Counter Example', 'Premier projet Solana avec programme de compteur décentralisé', '["JavaScript", "Rust", "Solana", "Anchor"]', 'blockchain', 'https://learn.zone01dakar.sn/git/alassall/solana-counter-example', false),

-- Projets Rust
('Smart Road', 'Premier projet Rust - Système de gestion intelligente de routes avec IoT', '["Rust", "IoT", "Systems Programming"]', 'rust', 'https://learn.zone01dakar.sn/git/alassall/smart-road', true),

('Filler', 'Algorithme de remplissage optimisé en Rust avec performance élevée', '["Rust", "Algorithms", "Performance"]', 'rust', 'https://learn.zone01dakar.sn/git/alassall/filler', false),

('Localhost', 'Serveur HTTP local développé en Rust avec gestion asynchrone', '["Rust", "HTTP", "Async", "Networking"]', 'rust', 'https://learn.zone01dakar.sn/git/alassall/localhost', false),

-- Projets JavaScript/Web
('GraphQL API', 'API GraphQL moderne avec résolveurs avancés et cache intelligent', '["JavaScript", "GraphQL", "Node.js", "Apollo"]', 'javascript', 'https://learn.zone01dakar.sn/git/alassall/graphql', true),

('Groupie Tracker Geolocation', 'Application de tracking de groupes musicaux avec géolocalisation avancée', '["CSS", "JavaScript", "Maps API", "Frontend"]', 'javascript', 'https://learn.zone01dakar.sn/git/alassall/groupie-tracker-geolocalization', false),

-- Projets Go
('Node Dashboard', 'Dashboard de monitoring pour réseaux privés avec métriques temps réel', '["Go", "Dashboard", "Monitoring", "WebSocket"]', 'go', 'https://learn.zone01dakar.sn/git/alassall/node-dashboard', true),

('ASCII Art Web Stylize', 'Générateur d''art ASCII web avec styles personnalisables', '["Go", "Web", "ASCII", "HTTP"]', 'go', 'https://learn.zone01dakar.sn/git/alassall/ascii-art-web-stylize', false),

-- Projets Java
('Piscine Java', 'Collection de projets Java avancés développés pendant la piscine', '["Java", "OOP", "Spring", "Maven"]', 'java', 'https://learn.zone01dakar.sn/git/alassall/piscine-java', false);
