# Guide de déploiement

## Railway (Recommandé - Gratuit)

### Étapes :
1. Push ton code sur GitHub
2. Va sur railway.app et connecte ton repo GitHub
3. Configure les variables d'environnement
4. Deploy automatique à chaque commit

### Variables d'environnement backend :
- `DB_HOST` - Fourni par Railway Postgres
- `DB_PORT` - Fourni par Railway Postgres  
- `DB_NAME` - Fourni par Railway Postgres
- `DB_USER` - Fourni par Railway Postgres
- `DB_PASSWORD` - Fourni par Railway Postgres

### Variables d'environnement frontend :
- `NEXT_PUBLIC_API_URL` - URL de ton backend Railway

## Render (Alternative)

Similaire à Railway, gratuit jusqu'à certaines limites.

## Vercel + PlanetScale

- Frontend sur Vercel (gratuit)
- Base de données PlanetScale (gratuit jusqu'à 1GB)
- Backend sur Railway ou Render

## Local avec Docker

```bash
docker-compose -f docker-compose.dev.yml up
```

## Avantages du déploiement automatique :

1. **Zero-downtime** - Pas d'interruption
2. **Rollback facile** - Retour arrière en un clic
3. **Monitoring intégré** - Logs et métriques
4. **HTTPS automatique** - Certificats SSL
5. **Auto-scaling** - S'adapte à la charge
