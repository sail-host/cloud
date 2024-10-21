import { create } from 'zustand'

interface ProjectStore {
  project: Project | null
  setProject: (project: Project | null) => void
}

export interface Project {
  id: number
  name: string
  status: 'pending' | 'building' | 'deploying' | 'running' | 'error' | 'success'
  created_at: string
  git_branch: string
  git_commit: string
  git_url: string
  git_hash: string
  domains: {
    id: number
    domain: string
    is_deployment: boolean
    created_at: string
  }[]
  project_framework: string
  build_command: string
  output_dir: string
  install_command: string
}

export const useProjectStore = create<ProjectStore>((set) => ({
  project: null,
  setProject: (project) => set({ project }),
}))
