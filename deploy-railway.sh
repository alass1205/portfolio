#!/bin/bash

echo "Déploiement sur Railway..."

# Backend
echo "Déploying backend..."
cd backend
railway login
railway link
railway up --detach

# Variables d'environnement backend
railway variables set DB_HOST=$RAILWAY_POSTGRES_HOST
railway variables set DB_PORT=$RAILWAY_POSTGRES_PORT
railway variables set DB_NAME=$RAILWAY_POSTGRES_DATABASE
railway variables set DB_USER=$RAILWAY_POSTGRES_USER
railway variables set DB_PASSWORD=$RAILWAY_POSTGRES_PASSWORD
railway variables set GITHUB_WEBHOOK_SECRET=your-webhook-secret

# Frontend  
echo "Deploying frontend..."
cd ../frontend
railway login
railway link
railway variables set NEXT_PUBLIC_API_URL=https://your-backend-url.railway.app
railway up --detach

echo "Déploiement terminé!"
