'use client'

import { useState, useEffect } from 'react'
import axios from 'axios'
import { Project, ApiResponse, ProjectCategory } from '@/types'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

export const useProjects = (category: ProjectCategory = 'all') => {
  const [projects, setProjects] = useState<Project[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        setLoading(true)
        setError(null)
        
        const url = category === 'all' 
          ? `${API_URL}/api/v1/projects`
          : `${API_URL}/api/v1/projects/category/${category}`
        
        const response = await axios.get<ApiResponse<Project[]>>(url)
        
        if (response.data.success) {
          setProjects(response.data.data)
        } else {
          setError('Erreur lors du chargement des projets')
        }
      } catch (err) {
        console.error('Erreur fetch projects:', err)
        setError('Impossible de charger les projets')
        // Données de test en cas d'erreur API
        setProjects([
          {
            id: '1',
            title: 'DeFi Trading Bot',
            description: 'Bot de trading automatisé pour protocoles DeFi avec stratégies d\'arbitrage',
            technologies: ['Solidity', 'Web3.js', 'Node.js', 'TypeScript'],
            category: 'blockchain',
            github_url: 'https://github.com',
            featured: true,
            created_at: '2024-01-01',
            updated_at: '2024-01-01'
          },
          {
            id: '2',
            title: 'Rust Web Server',
            description: 'Serveur web haute performance en Rust avec gestion async',
            technologies: ['Rust', 'Tokio', 'Actix-web', 'PostgreSQL'],
            category: 'rust',
            github_url: 'https://github.com',
            featured: true,
            created_at: '2024-01-01',
            updated_at: '2024-01-01'
          }
        ])
      } finally {
        setLoading(false)
      }
    }

    fetchProjects()
  }, [category])

  return { projects, loading, error }
}
