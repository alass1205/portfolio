export interface Project {
  id: string
  title: string
  description: string
  technologies: string[]
  category: string
  github_url?: string
  live_url?: string
  image_url?: string
  featured: boolean
  created_at: string
  updated_at: string
}

export interface ApiResponse<T> {
  success: boolean
  data: T
  count?: number
  message?: string
}

export type ProjectCategory = 'blockchain' | 'rust' | 'javascript' | 'go' | 'all'
