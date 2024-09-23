export interface Project {
  id: number
  name: string
  domain: string
  favicon: string
  git: {
    url: string
    branch: string
    last_commit: {
      message: string
      date: string
    }
  }
}

export const projects: Project[] = [
  {
    id: 1,
    name: 'TechSphere',
    domain: 'techsphere.io',
    favicon: 'https://techsphere.io/favicon.ico',
    git: {
      url: 'https://github.com/techsphere/main-app.git',
      branch: 'develop',
      last_commit: {
        message: 'Implement user authentication',
        date: '1 day',
      },
    },
  },
  {
    id: 2,
    name: 'QuantumLeap',
    domain: 'quantumleap.tech',
    favicon: 'https://quantumleap.tech/favicon.ico',
    git: {
      url: 'https://gitlab.com/quantum-leap/core.git',
      branch: 'master',
      last_commit: {
        message: 'Optimize quantum algorithm',
        date: '1 hour',
      },
    },
  },
  {
    id: 3,
    name: 'CyberBridge',
    domain: 'cyberbridge.org',
    favicon: 'https://cyberbridge.org/favicon.ico',
    git: {
      url: 'https://github.com/cyber-bridge/platform.git',
      branch: 'main',
      last_commit: {
        message: 'Enhance security protocols',
        date: '1 week',
      },
    },
  },
  {
    id: 4,
    name: 'EcoSmart',
    domain: 'ecosmart.co',
    favicon: 'https://ecosmart.co/favicon.ico',
    git: {
      url: 'https://bitbucket.org/ecosmart/green-app.git',
      branch: 'production',
      last_commit: {
        message: 'Add carbon footprint calculator',
        date: '1 month',
      },
    },
  },
  {
    id: 5,
    name: 'ArtificialMind',
    domain: 'artificialmind.ai',
    favicon: 'https://artificialmind.ai/favicon.ico',
    git: {
      url: 'https://github.com/artificial-mind/neural-net.git',
      branch: 'experimental',
      last_commit: {
        message: 'Implement deep learning model',
        date: '1 day',
      },
    },
  },
  {
    id: 6,
    name: 'SpaceVoyager',
    domain: 'spacevoyager.com',
    favicon: 'https://spacevoyager.com/favicon.ico',
    git: {
      url: 'https://gitlab.com/space-voyager/mission-control.git',
      branch: 'stable',
      last_commit: {
        message: 'Update orbital calculations',
        date: '1 week',
      },
    },
  },
  {
    id: 7,
    name: 'HealthPulse',
    domain: 'healthpulse.med',
    favicon: 'https://healthpulse.med/favicon.ico',
    git: {
      url: 'https://github.com/health-pulse/patient-portal.git',
      branch: 'release',
      last_commit: {
        message: 'Implement telemedicine features',
        date: '1 month',
      },
    },
  },
  {
    id: 8,
    name: 'FinTechFlow',
    domain: 'fintechflow.io',
    favicon: 'https://fintechflow.io/favicon.ico',
    git: {
      url: 'https://bitbucket.org/fintech-flow/transactions.git',
      branch: 'main',
      last_commit: {
        message: 'Implement blockchain integration',
        date: '1 month',
      },
    },
  },
  {
    id: 9,
    name: 'EduQuest',
    domain: 'eduquest.edu',
    favicon: 'https://eduquest.edu/favicon.ico',
    git: {
      url: 'https://github.com/eduquest/learning-platform.git',
      branch: 'develop',
      last_commit: {
        message: 'Add interactive quiz module',
        date: '1 day',
      },
    },
  },
  {
    id: 10,
    name: 'GreenEnergy',
    domain: 'greenenergy.eco',
    favicon: 'https://greenenergy.eco/favicon.ico',
    git: {
      url: 'https://gitlab.com/green-energy/solar-tracker.git',
      branch: 'master',
      last_commit: {
        message: 'Optimize energy distribution algorithm',
        date: '1 week',
      },
    },
  },
]
