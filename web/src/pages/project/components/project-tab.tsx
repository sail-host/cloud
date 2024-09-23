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

export function ProjectTab() {
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
            <Button variant='outline'>
              <IconGitBranch className='mr-2 h-4 w-4' />
              Repository
            </Button>
            <Button variant='outline'>
              <IconReload className='mr-2 h-4 w-4' />
              Rebuild
            </Button>
            <Button>Visit</Button>
          </div>
        </CardHeader>
        <CardContent className='grid w-full grid-cols-11 gap-4 text-sm'>
          <div className='col-span-4'>
            <Skeleton className='h-full w-full' />
          </div>
          <div className='col-span-7 flex flex-col gap-4'>
            <div>
              <p className='mb-1 font-light text-muted-foreground'>
                Deployment URL
              </p>
              <a href='#' className='hover:underline'>
                https://solvie-dashboard.vercel.app
              </a>
            </div>

            <div>
              <p className='mb-1 font-light text-muted-foreground'>Domains</p>
              <div className='flex flex-wrap gap-2'>
                <a
                  href='https://solvie-dashboard.vercel.app'
                  target='_blank'
                  className='flex items-center hover:underline'
                >
                  solvie-dashboard.vercel.app
                  <IconExternalLink className='ml-1 h-4 w-4' />
                </a>
                <a
                  href='https://solvie-dashboard.vercel.app'
                  target='_blank'
                  className='flex items-center hover:underline'
                >
                  solvie-dashboard.vercel.app
                  <IconExternalLink className='ml-1 h-4 w-4' />
                </a>
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
                  <div className='h-2 w-2 rounded-full bg-green-500' />
                  <span className='ml-1'>Active</span>
                </div>
              </div>
              <div className='col-span-7'>April 9, 2024 at 11:34 AM</div>
            </div>

            <div>
              <p className='mb-1 font-light text-muted-foreground'>Source</p>
              <div className='flex flex-col gap-0.5'>
                <div className='flex items-center gap-x-2'>
                  <IconGitBranch className='h-4 w-4' />
                  <a href='#' className='hover:underline'>
                    <span>master</span>
                  </a>
                </div>
                <div className='flex items-center gap-x-2'>
                  <IconGitCommit className='h-4 w-4 rotate-90' />
                  <a href='#' className='space-x-3 hover:underline'>
                    <span>d093751</span>
                    <span>Update npmrc and package.json scripts</span>
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
