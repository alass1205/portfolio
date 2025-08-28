'use client'

import { motion } from 'framer-motion'
import { Code, Database, Globe, Zap } from 'lucide-react'

const SkillsSection = () => {
  const skillCategories = [
    {
      icon: Zap,
      title: 'Blockchain',
      color: 'from-orange-500 to-red-600',
      skills: [
        { name: 'Solidity', level: 90 },
        { name: 'Web3.js', level: 85 },
        { name: 'Ethereum', level: 88 },
        { name: 'DeFi Protocols', level: 80 }
      ]
    },
    {
      icon: Code,
      title: 'Rust',
      color: 'from-orange-600 to-red-700',
      skills: [
        { name: 'Systems Programming', level: 85 },
        { name: 'Tokio/Async', level: 80 },
        { name: 'Actix-web', level: 82 },
        { name: 'WebAssembly', level: 75 }
      ]
    },
    {
      icon: Globe,
      title: 'JavaScript',
      color: 'from-yellow-500 to-orange-600',
      skills: [
        { name: 'React/Next.js', level: 92 },
        { name: 'TypeScript', level: 88 },
        { name: 'Node.js', level: 85 },
        { name: 'Three.js', level: 70 }
      ]
    },
    {
      icon: Database,
      title: 'Go',
      color: 'from-blue-500 to-cyan-600',
      skills: [
        { name: 'Fiber/Gin', level: 85 },
        { name: 'Microservices', level: 80 },
        { name: 'PostgreSQL', level: 88 },
        { name: 'Docker', level: 85 }
      ]
    }
  ]

  return (
    <section id="skills" className="py-20 px-4">
      <div className="max-w-7xl mx-auto">
        <motion.div
          initial={{ opacity: 0, y: 50 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          viewport={{ once: true }}
          className="text-center mb-16"
        >
          <h2 className="text-5xl font-bold mb-6">
            <span className="gradient-text">Compétences</span>
          </h2>
          <p className="text-xl text-gray-300 max-w-3xl mx-auto">
            Technologies maîtrisées à travers mes projets et expériences professionnelles
          </p>
        </motion.div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {skillCategories.map((category, categoryIndex) => (
            <motion.div
              key={category.title}
              initial={{ opacity: 0, y: 50 }}
              whileInView={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.6, delay: categoryIndex * 0.1 }}
              viewport={{ once: true }}
              whileHover={{ y: -10, scale: 1.02 }}
              className="glass rounded-2xl p-6 text-center group"
            >
              <div className={`w-16 h-16 mx-auto mb-4 rounded-full bg-gradient-to-r ${category.color} flex items-center justify-center group-hover:scale-110 transition-transform duration-300`}>
                <category.icon className="w-8 h-8 text-white" />
              </div>
              
              <h3 className="text-2xl font-bold text-white mb-6">{category.title}</h3>
              
              <div className="space-y-4">
                {category.skills.map((skill, skillIndex) => (
                  <div key={skill.name} className="text-left">
                    <div className="flex justify-between items-center mb-2">
                      <span className="text-gray-300 text-sm font-medium">{skill.name}</span>
                      <span className="text-gray-400 text-xs">{skill.level}%</span>
                    </div>
                    <div className="w-full bg-gray-700 rounded-full h-2 overflow-hidden">
                      <motion.div
                        initial={{ width: 0 }}
                        whileInView={{ width: `${skill.level}%` }}
                        transition={{ duration: 1, delay: categoryIndex * 0.1 + skillIndex * 0.1 }}
                        viewport={{ once: true }}
                        className={`h-full bg-gradient-to-r ${category.color} rounded-full`}
                      />
                    </div>
                  </div>
                ))}
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  )
}

export default SkillsSection
