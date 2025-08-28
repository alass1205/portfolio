'use client'

import { useState } from 'react'
import { motion, AnimatePresence } from 'framer-motion'
import { ExternalLink, Github, Filter } from 'lucide-react'
import { useProjects } from '@/hooks/useProjects'
import { ProjectCategory } from '@/types'

const ProjectsSection = () => {
  const [selectedCategory, setSelectedCategory] = useState<ProjectCategory>('all')
  const { projects, loading, error } = useProjects(selectedCategory)

  const categories = [
    { key: 'all' as ProjectCategory, label: 'Tous', color: 'from-gray-600 to-gray-800' },
    { key: 'blockchain' as ProjectCategory, label: 'Blockchain', color: 'from-orange-500 to-red-600' },
    { key: 'rust' as ProjectCategory, label: 'Rust', color: 'from-orange-600 to-red-700' },
    { key: 'javascript' as ProjectCategory, label: 'JavaScript', color: 'from-yellow-500 to-orange-600' },
    { key: 'go' as ProjectCategory, label: 'Go', color: 'from-blue-500 to-cyan-600' }
  ]

  const getCategoryColor = (category: string) => {
    const colors = {
      blockchain: 'bg-gradient-to-r from-orange-500 to-red-600',
      rust: 'bg-gradient-to-r from-orange-600 to-red-700',
      javascript: 'bg-gradient-to-r from-yellow-500 to-orange-600',
      go: 'bg-gradient-to-r from-blue-500 to-cyan-600'
    }
    return colors[category as keyof typeof colors] || 'bg-gradient-to-r from-gray-500 to-gray-700'
  }

  return (
    <section id="projects" className="py-20 px-4">
      <div className="max-w-7xl mx-auto">
        {/* Header */}
        <motion.div
          initial={{ opacity: 0, y: 50 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          viewport={{ once: true }}
          className="text-center mb-16"
        >
          <h2 className="text-5xl font-bold mb-6">
            <span className="gradient-text">Mes Projets</span>
          </h2>
          <p className="text-xl text-gray-300 max-w-3xl mx-auto">
            Découvrez mes réalisations en blockchain, développement système avec Rust,
            applications web JavaScript et APIs Go performantes.
          </p>
        </motion.div>

        {/* Filtres */}
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6, delay: 0.2 }}
          viewport={{ once: true }}
          className="flex flex-wrap justify-center gap-4 mb-12"
        >
          {categories.map((category) => (
            <motion.button
              key={category.key}
              onClick={() => setSelectedCategory(category.key)}
              whileHover={{ scale: 1.05, y: -2 }}
              whileTap={{ scale: 0.95 }}
              className={`px-6 py-3 rounded-full font-semibold transition-all duration-300 ${
                selectedCategory === category.key
                  ? `bg-gradient-to-r ${category.color} text-white shadow-lg`
                  : 'glass text-gray-300 hover:text-white hover:bg-white/10'
              }`}
            >
              <div className="flex items-center space-x-2">
                <Filter size={16} />
                <span>{category.label}</span>
              </div>
            </motion.button>
          ))}
        </motion.div>

        {/* Projets Grid */}
        {loading ? (
          <div className="flex justify-center items-center h-64">
            <div className="animate-spin rounded-full h-16 w-16 border-t-2 border-primary-500"></div>
          </div>
        ) : error ? (
          <div className="text-center text-red-400 py-8">
            <p>{error}</p>
          </div>
        ) : (
          <AnimatePresence mode="wait">
            <motion.div
              key={selectedCategory}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: -20 }}
              transition={{ duration: 0.5 }}
              className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"
            >
              {projects.map((project, index) => (
                <motion.div
                  key={project.id}
                  initial={{ opacity: 0, y: 50 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.5, delay: index * 0.1 }}
                  whileHover={{ y: -10, scale: 1.02 }}
                  className="glass rounded-2xl overflow-hidden group cursor-pointer"
                >
                  {/* Image placeholder */}
                  <div className={`h-48 ${getCategoryColor(project.category)} relative overflow-hidden`}>
                    <div className="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
                    <div className="absolute top-4 right-4">
                      <span className="px-3 py-1 bg-black/30 backdrop-blur-sm rounded-full text-sm font-medium text-white">
                        {project.category}
                      </span>
                    </div>
                    {project.featured && (
                      <div className="absolute top-4 left-4">
                        <span className="px-3 py-1 bg-yellow-500/20 backdrop-blur-sm rounded-full text-sm font-medium text-yellow-300">
                          ⭐ Featured
                        </span>
                      </div>
                    )}
                  </div>

                  <div className="p-6">
                    <h3 className="text-xl font-bold text-white mb-3 group-hover:text-primary-400 transition-colors duration-300">
                      {project.title}
                    </h3>
                    
                    <p className="text-gray-300 text-sm mb-4 line-clamp-3">
                      {project.description}
                    </p>

                    {/* Technologies */}
                    <div className="flex flex-wrap gap-2 mb-6">
                      {project.technologies.slice(0, 3).map((tech) => (
                        <span
                          key={tech}
                          className="px-3 py-1 bg-white/10 rounded-full text-xs font-medium text-gray-300"
                        >
                          {tech}
                        </span>
                      ))}
                      {project.technologies.length > 3 && (
                        <span className="px-3 py-1 bg-white/10 rounded-full text-xs font-medium text-gray-400">
                          +{project.technologies.length - 3}
                        </span>
                      )}
                    </div>

                    {/* Actions */}
                    <div className="flex space-x-4">
                      {project.github_url && (
                        <motion.a
                          href={project.github_url}
                          target="_blank"
                          rel="noopener noreferrer"
                          whileHover={{ scale: 1.1 }}
                          whileTap={{ scale: 0.9 }}
                          className="flex items-center space-x-2 text-gray-300 hover:text-white transition-colors duration-300"
                        >
                          <Github size={16} />
                          <span className="text-sm">Code</span>
                        </motion.a>
                      )}
                      
                      {project.live_url && (
                        <motion.a
                          href={project.live_url}
                          target="_blank"
                          rel="noopener noreferrer"
                          whileHover={{ scale: 1.1 }}
                          whileTap={{ scale: 0.9 }}
                          className="flex items-center space-x-2 text-gray-300 hover:text-primary-400 transition-colors duration-300"
                        >
                          <ExternalLink size={16} />
                          <span className="text-sm">Demo</span>
                        </motion.a>
                      )}
                    </div>
                  </div>
                </motion.div>
              ))}
            </motion.div>
          </AnimatePresence>
        )}

        {/* Empty state */}
        {!loading && !error && projects.length === 0 && (
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            className="text-center py-16"
          >
            <div className="text-6xl mb-4">🚀</div>
            <h3 className="text-2xl font-bold text-white mb-2">Aucun projet trouvé</h3>
            <p className="text-gray-400">Essayez de sélectionner une autre catégorie</p>
          </motion.div>
        )}
      </div>
    </section>
  )
}

export default ProjectsSection
