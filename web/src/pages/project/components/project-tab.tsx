import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { TabsContent } from '@/components/ui/tabs'
import {
  IconExternalLink,
  IconGitBranch,
  IconGitCommit,
  IconReload,
} from '@tabler/icons-react'
import { Skeleton } from '@/components/ui/skeleton'
import { useProjectStore } from '@/store/project-store'

export function ProjectTab() {
  const { project } = useProjectStore()

  const deploymentUrl =
    project?.domains && project.domains.length > 0
      ? `https://${project.domains[0].domain}`
      : ''

  return (
    <TabsContent value='project' className=''>
      <Card>
        <CardHeader className='flex flex-row justify-between space-y-0'>
          <div>
            <CardTitle className='mb-1 text-xl'>
              Production Deployment
            </CardTitle>
            <CardDescription>
              The deployment that is available to your visitors.
            </CardDescription>
          </div>
          <div className='flex items-center gap-x-4'>
            <Button variant='outline' asChild>
              <a href={project?.git_url} target='_blank' rel='noreferrer'>
                <IconGitBranch className='w-4 h-4 mr-2' />
                Repository
              </a>
            </Button>
            <Button variant='outline'>
              <IconReload className='w-4 h-4 mr-2' />
              Rebuild
            </Button>
            <Button>Visit</Button>
          </div>
        </CardHeader>
        <CardContent className='grid w-full grid-cols-11 gap-4 text-sm'>
          <div className='col-span-4'>
            <Skeleton className='w-full h-full' />
          </div>
          <div className='flex flex-col col-span-7 gap-4'>
            <div>
              <p className='mb-1 font-light text-muted-foreground'>
                Deployment URL
              </p>
              <a href={deploymentUrl} className='hover:underline'>
                {deploymentUrl}
              </a>
            </div>

            <div>
              <p className='mb-1 font-light text-muted-foreground'>Domains</p>
              <div className='flex flex-wrap gap-2'>
                {project?.domains.map((domain) => (
                  <a
                    href={`https://${domain.domain}`}
                    target='_blank'
                    className='flex items-center hover:underline'
                  >
                    {domain.domain}
                    <IconExternalLink className='w-4 h-4 ml-1' />
                  </a>
                ))}
              </div>
            </div>

            <div className='grid grid-cols-8 gap-x-2'>
              <div className='col-span-1'>
                <p className='mb-1 font-light text-muted-foreground'>Status</p>
              </div>
              <div className='col-span-7'>
                <p className='mb-1 font-light text-muted-foreground'>Created</p>
              </div>

              <div className='col-span-1'>
                <div className='flex items-center'>
                  {project?.status === 'success' && (
                    <>
                      <div className='w-2 h-2 bg-green-500 rounded-full' />
                      <span className='ml-1'>Active</span>
                    </>
                  )}
                  {project?.status === 'error' && (
                    <>
                      <div className='w-2 h-2 bg-red-500 rounded-full' />
                      <span className='ml-1'>Error</span>
                    </>
                  )}
                  {project?.status === 'pending' && (
                    <>
                      <div className='w-2 h-2 bg-yellow-500 rounded-full' />
                      <span className='ml-1'>Pending</span>
                    </>
                  )}
                  {project?.status === 'building' && (
                    <>
                      <div className='w-2 h-2 bg-blue-500 rounded-full' />
                      <span className='ml-1'>Building</span>
                    </>
                  )}
                  {project?.status === 'deploying' && (
                    <>
                      <div className='w-2 h-2 bg-purple-500 rounded-full' />
                      <span className='ml-1'>Deploying</span>
                    </>
                  )}
                </div>
              </div>
              <div className='col-span-7'>
                {project?.created_at
                  ? new Date(project.created_at).toLocaleString()
                  : ''}
              </div>
            </div>

            <div>
              <p className='mb-1 font-light text-muted-foreground'>Source</p>
              <div className='flex flex-col gap-0.5'>
                <div className='flex items-center gap-x-2'>
                  <IconGitBranch className='w-4 h-4' />
                  <a href='#' className='hover:underline'>
                    {project?.git_branch}
                  </a>
                </div>
                <div className='flex items-center gap-x-2'>
                  <IconGitCommit className='w-4 h-4 rotate-90' />
                  <a href='#' className='space-x-3 hover:underline'>
                    <span>{project?.git_hash?.substring(0, 10)}</span>
                    <span>{project?.git_commit}</span>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </TabsContent>
  )
}
