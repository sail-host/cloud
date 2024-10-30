import { Button } from '@/components/ui/button'
import { formatDistanceToNow, slugify } from '@/lib/utils'
import {
  useProjectCreateStore,
  useProjectStore,
} from '@/store/project-create-store'
import { IconLock, IconPointFilled } from '@tabler/icons-react'

interface ProjectCardItemProps {
  item: Project
}

export interface Project {
  id: number
  name: string
  owner: string
  full_name: string
  description: string
  url: string
  default_branch: string
  created_at: string
  updated_at: string
  clone_url: string
  private: boolean
  framework: framework
}

type framework =
  | 'react'
  | 'vue'
  | 'next'
  | 'svelte'
  | 'angular'
  | 'lit'
  | 'ember'
  | 'vanilla'
  | 'remix'
  | 'nuxt'
  | 'vite'

const iconFramework = {
  react: 'https://api-frameworks.vercel.sh/framework-logos/react.svg',
  vue: 'https://api-frameworks.vercel.sh/framework-logos/vue.svg',
  next: 'https://api-frameworks.vercel.sh/framework-logos/next.svg',
  svelte: 'https://api-frameworks.vercel.sh/framework-logos/svelte.svg',
  angular: 'https://api-frameworks.vercel.sh/framework-logos/angular.svg',
  lit: 'https://api-frameworks.vercel.sh/framework-logos/lit.svg',
  ember: 'https://api-frameworks.vercel.sh/framework-logos/ember.svg',
  vanilla: 'https://api-frameworks.vercel.sh/framework-logos/vanilla.svg',
  remix: 'https://api-frameworks.vercel.sh/framework-logos/remix.svg',
  nuxt: 'https://api-frameworks.vercel.sh/framework-logos/nuxt.svg',
  vite: 'https://api-frameworks.vercel.sh/framework-logos/vite.svg',
}

export function ProjectCardItem({ item }: ProjectCardItemProps) {
  const { setStep } = useProjectCreateStore()
  const { setProject, setProjectFramework, setProjectName } = useProjectStore()

  const handleImport = () => {
    setProject(item)
    setProjectFramework(item.framework)
    setProjectName(slugify(item.name))
    setStep('2')
  }

  return (
    <div className='flex items-center justify-between border-b last:border-none'>
      <div className='flex items-center gap-4 p-4'>
        {item.framework && iconFramework[item.framework as framework] ? (
          <img
            src={iconFramework[item.framework]}
            alt={item.name}
            className='h-6 w-6 rounded-full'
          />
        ) : (
          <div className='h-6 w-6 rounded-full bg-gray-200'></div>
        )}
        <div className='flex items-center gap-1'>
          <h3 className='text-sm font-medium'>{item.name}</h3>
          <IconLock className='h-4 w-4 text-muted-foreground' />
          <IconPointFilled className='h-2 w-2 text-muted-foreground/60' />
          <span className='text-sm text-muted-foreground'>
            {formatDistanceToNow(new Date(item.updated_at))}
          </span>
        </div>
      </div>
      <div className='p-4'>
        <Button
          className='h-auto px-3 py-1.5 text-sm'
          onClick={handleImport}
          type='button'
        >
          Import
        </Button>
      </div>
    </div>
  )
}
