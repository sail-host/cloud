import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Link } from 'react-router-dom'
import { IconBrandGithub, IconGitBranch } from '@tabler/icons-react'

export function StepTwoRightCard() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Project Preview</CardTitle>
      </CardHeader>
      <CardContent>
        <div className='flex flex-col gap-2'>
          <div>
            <p className='text-sm text-muted-foreground'>Project Name</p>
            <p className='text-lg font-medium'>my-project</p>
          </div>
          <div>
            <p className='text-sm text-muted-foreground'>Framework</p>
            <p className='text-lg font-medium'>Next.js</p>
          </div>
          <div>
            <p className='text-sm text-muted-foreground'>Git Repository</p>
            <Link
              to='https://github.com/user/my-project.git'
              target='_blank'
              className='flex items-center gap-1 hover:underline'
            >
              <IconBrandGithub className='w-5 h-5' />
              user/my-project
            </Link>
          </div>
          <div>
            <p className='text-sm text-muted-foreground'>Branch</p>
            <p className='flex items-center gap-1 text-lg font-medium'>
              <IconGitBranch className='w-5 h-5' />
              master
            </p>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
