import { Project } from '@/pages/projects/components/project-card-item'
import { GitAccount } from '@/types/model'
import { create } from 'zustand'

interface Environment {
  id: string
  name: string
  value: string
}

interface ProjectStore {
  setProject: (project: Project | null) => void
  project: Project | null
  projectName: string
  projectFramework: string
  rootDir: string
}

interface ProjectCreateStore {
  gitAccount: GitAccount | null
  setGitAccount: (gitAccount: GitAccount | null) => void
  step: '1' | '2'
  setStep: (step: '1' | '2') => void
}

interface ProjectSettingStore {
  projectName: string
  projectFramework: string
  rootDir: string
  buildCommand: string
  outputDir: string
  installCommand: string
  environments: Environment[]
  setProjectName: (projectName: string) => void
  setProjectFramework: (projectFramework: string) => void
  setRootDir: (rootDir: string) => void
  setBuildCommand: (buildCommand: string) => void
  setOutputDir: (outputDir: string) => void
  setInstallCommand: (installCommand: string) => void
  setEnvironments: (environments: Environment[]) => void
}

export const useProjectStore = create<ProjectStore>((set) => ({
  project: null,
  setProject: (project) => set({ project }),
  projectName: '',
  projectFramework: '',
  rootDir: '',
}))

export const useProjectCreateStore = create<ProjectCreateStore>((set) => ({
  gitAccount: null,
  setGitAccount: (gitAccount) => set({ gitAccount }),
  step: '1',
  setStep: (step) => set({ step }),
}))

export const useProjectSettingStore = create<ProjectSettingStore>((set) => ({
  installCommand: '',
  environments: [],
  setProjectName: (projectName) => set({ projectName }),
  setProjectFramework: (projectFramework) => set({ projectFramework }),
  setRootDir: (rootDir) => set({ rootDir }),
  setBuildCommand: (buildCommand) => set({ buildCommand }),
  setOutputDir: (outputDir) => set({ outputDir }),
  setInstallCommand: (installCommand) => set({ installCommand }),
  setEnvironments: (environments) => set({ environments }),
  projectName: '',
  projectFramework: '',
  rootDir: '',
  buildCommand: '',
  outputDir: '',
}))
