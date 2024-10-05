import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Link } from 'react-router-dom'
import { IconBrandGithub, IconGitBranch } from '@tabler/icons-react'
import { useProjectStore } from '@/store/project-create-store'

export function StepTwoRightCard() {
  const { projectName, projectFramework, project } = useProjectStore()

  return (
    <Card>
      <CardHeader>
        <CardTitle>Project Preview</CardTitle>
      </CardHeader>
      <CardContent>
        <div className='flex flex-col gap-2'>
          <div>
            <p className='text-sm text-muted-foreground'>Project Name</p>
            <p className='text-lg font-medium'>{projectName}</p>
          </div>
          <div>
            <p className='text-sm text-muted-foreground'>Framework</p>
            <p className='text-lg font-medium capitalize'>{projectFramework}</p>
          </div>
          <div>
            <p className='text-sm text-muted-foreground'>Git Repository</p>
            <Link
              to={project?.url ?? ''}
              target='_blank'
              className='flex items-center gap-1 hover:underline'
            >
              <IconBrandGithub className='w-5 h-5' />
              {project?.owner}/{project?.name}
            </Link>
          </div>
          <div>
            <p className='text-sm text-muted-foreground'>Branch</p>
            <p className='flex items-center gap-1 text-lg font-medium'>
              <IconGitBranch className='w-5 h-5' />
              {project?.default_branch}
            </p>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
