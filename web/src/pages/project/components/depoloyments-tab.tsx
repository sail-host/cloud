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
  IconCircleArrowUp,
  IconDotsVertical,
  IconGitCommit,
  IconGitPullRequest,
  IconReload,
} from '@tabler/icons-react'
import { Table, TableBody, TableCell, TableRow } from '@/components/ui/table'
import { Badge } from '@/components/ui/badge'
import { cn, formatDistanceToNow } from '@/lib/utils'

interface Deployment {
  id: string
  environment: string
  commit: {
    message: string
    author: string
    sha: string
    url: string
    createdAt: Date
  }
  deploy: {
    time: number
    status: 'success' | 'error' | 'pending'
  }
  isCurrent: boolean
  isActive: boolean
  createdAt: string
}

const deployments: Deployment[] = [
  {
    id: '550e8400-e29b-41d4-a716-446655440000',
    environment: 'production',
    commit: {
      message: 'Update user authentication',
      author: 'Alice Johnson',
      sha: 'a1b2c3d4e5',
      url: 'https://github.com/project/commit/a1b2c3d4e5',
      createdAt: new Date('2023-04-15T10:30:00Z'),
    },
    deploy: {
      time: 8,
      status: 'success',
    },
    isCurrent: true,
    isActive: true,
    createdAt: '2023-04-15T10:38:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440001',
    environment: 'staging',
    commit: {
      message: 'Implement new feature',
      author: 'Bob Smith',
      sha: 'f6g7h8i9j0',
      url: 'https://github.com/project/commit/f6g7h8i9j0',
      createdAt: new Date('2023-04-14T15:45:00Z'),
    },
    deploy: {
      time: 12,
      status: 'success',
    },
    isCurrent: false,
    isActive: true,
    createdAt: '2024-04-14T15:57:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440002',
    environment: 'development',
    commit: {
      message: 'Fix bug in login process',
      author: 'Charlie Brown',
      sha: 'k1l2m3n4o5',
      url: 'https://github.com/project/commit/k1l2m3n4o5',
      createdAt: new Date('2023-04-13T09:20:00Z'),
    },
    deploy: {
      time: 6,
      status: 'success',
    },
    isCurrent: false,
    isActive: true,
    createdAt: '2024-08-13T09:26:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440003',
    environment: 'production',
    commit: {
      message: 'Optimize database queries',
      author: 'Diana Prince',
      sha: 'p6q7r8s9t0',
      url: 'https://github.com/project/commit/p6q7r8s9t0',
      createdAt: new Date('2023-04-12T14:10:00Z'),
    },
    deploy: {
      time: 15,
      status: 'error',
    },
    isCurrent: false,
    isActive: false,
    createdAt: '2023-04-12T14:25:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440004',
    environment: 'staging',
    commit: {
      message: 'Add new API endpoints',
      author: 'Ethan Hunt',
      sha: 'u1v2w3x4y5',
      url: 'https://github.com/project/commit/u1v2w3x4y5',
      createdAt: new Date('2023-04-11T11:55:00Z'),
    },
    deploy: {
      time: 10,
      status: 'success',
    },
    isCurrent: false,
    isActive: true,
    createdAt: '2023-04-11T12:05:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440005',
    environment: 'development',
    commit: {
      message: 'Refactor authentication module',
      author: 'Fiona Gallagher',
      sha: 'z1a2b3c4d5',
      url: 'https://github.com/project/commit/z1a2b3c4d5',
      createdAt: new Date('2023-04-10T16:30:00Z'),
    },
    deploy: {
      time: 7,
      status: 'success',
    },
    isCurrent: false,
    isActive: true,
    createdAt: '2023-04-10T16:37:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440006',
    environment: 'production',
    commit: {
      message: 'Update dependencies',
      author: 'George Lucas',
      sha: 'e6f7g8h9i0',
      url: 'https://github.com/project/commit/e6f7g8h9i0',
      createdAt: new Date('2023-04-09T13:15:00Z'),
    },
    deploy: {
      time: 20,
      status: 'pending',
    },
    isCurrent: false,
    isActive: false,
    createdAt: '2023-04-09T13:35:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440007',
    environment: 'staging',
    commit: {
      message: 'Implement user roles and permissions',
      author: 'Hannah Montana',
      sha: 'j1k2l3m4n5',
      url: 'https://github.com/project/commit/j1k2l3m4n5',
      createdAt: new Date('2023-04-08T10:05:00Z'),
    },
    deploy: {
      time: 14,
      status: 'success',
    },
    isCurrent: false,
    isActive: true,
    createdAt: '2023-04-08T10:19:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440008',
    environment: 'development',
    commit: {
      message: 'Add unit tests for new features',
      author: 'Ian Malcolm',
      sha: 'o6p7q8r9s0',
      url: 'https://github.com/project/commit/o6p7q8r9s0',
      createdAt: new Date('2023-04-07T17:40:00Z'),
    },
    deploy: {
      time: 9,
      status: 'success',
    },
    isCurrent: false,
    isActive: true,
    createdAt: '2023-04-07T17:49:00Z',
  },
  {
    id: '550e8400-e29b-41d4-a716-446655440009',
    environment: 'production',
    commit: {
      message: 'Implement performance optimizations',
      author: 'Julia Roberts',
      sha: 't1u2v3w4x5',
      url: 'https://github.com/project/commit/t1u2v3w4x5',
      createdAt: new Date('2023-04-06T12:25:00Z'),
    },
    deploy: {
      time: 18,
      status: 'success',
    },
    isCurrent: false,
    isActive: true,
    createdAt: '2023-04-06T12:43:00Z',
  },
]

export function DeploymentsTab() {
  return (
    <TabsContent value='deployments' className='space-y-4'>
      <Card>
        <CardHeader className='flex flex-row justify-between space-y-0'>
          <div>
            <CardTitle className='mb-1 text-xl'>Deployments</CardTitle>
            <CardDescription>
              The deployments that are available to your visitors.
            </CardDescription>
          </div>
          <div className='flex items-center gap-x-4'>
            <Button variant='outline'>
              <IconReload className='mr-2 h-4 w-4' />
              Redeploy
            </Button>
          </div>
        </CardHeader>
        <CardContent className='text-sm'>
          <div>
            <Table className='w-full'>
              {/* <TableHeader>
                <TableRow>
                  <TableHead className='w-[100px]'>Invoice</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead>Method</TableHead>
                  <TableHead className='text-right'>Amount</TableHead>
                </TableRow>
              </TableHeader> */}
              <TableBody>
                {deployments.map((deployment) => (
                  <TableRow key={deployment.id}>
                    <TableCell className=''>
                      <div className='font-medium'>
                        {deployment.commit.sha.slice(0, 10)}
                      </div>
                      <div className='flex items-center gap-x-2 capitalize text-muted-foreground'>
                        {deployment.environment}

                        {deployment.isCurrent && (
                          <Badge className='rounded-full bg-blue-500/40 px-1 py-0 text-xs hover:bg-blue-500/40'>
                            <IconCircleArrowUp className='mr-0.5 inline-block h-3 w-3' />
                            Current
                          </Badge>
                        )}
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className='flex flex-col'>
                        <div className='flex items-center gap-x-2'>
                          <div
                            className={cn(
                              'h-3 w-3 rounded-full',
                              deployment.deploy.status === 'success'
                                ? 'bg-green-500'
                                : deployment.deploy.status === 'error'
                                  ? 'bg-red-500'
                                  : 'bg-yellow-500'
                            )}
                          ></div>
                          <span className='capitalize'>
                            {deployment.deploy.status === 'success'
                              ? 'Active'
                              : deployment.deploy.status}
                          </span>
                        </div>
                        <div className='text-muted-foreground'>
                          {deployment.deploy.status === 'success' ? (
                            <p>
                              {deployment.deploy.time}s (
                              {deployment.deploy.time * 1.23} MB)
                            </p>
                          ) : (
                            <p>Failed</p>
                          )}
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className='flex flex-col'>
                        <div className='flex items-center gap-x-2'>
                          <IconGitPullRequest className='h-4 w-4' />
                          <a href='#' className='font-light hover:underline'>
                            master
                          </a>
                        </div>
                        <div className='flex items-center gap-x-2'>
                          <IconGitCommit className='h-4 w-4' />
                          <a
                            href={deployment.commit.url}
                            className='flex items-center gap-x-2 font-light hover:underline'
                          >
                            <span>{deployment.commit.sha.slice(0, 7)}</span>
                            <span className='text-muted-foreground'>
                              {deployment.commit.message.length > 30
                                ? `${deployment.commit.message.slice(0, 30)}...`
                                : deployment.commit.message}
                            </span>
                          </a>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell className='text-right'>
                      <p className='font-light text-muted-foreground'>
                        {formatDistanceToNow(deployment.commit.createdAt)} by{' '}
                        <span className='font-medium text-gray-800 hover:underline dark:text-white/90'>
                          {deployment.commit.author}
                        </span>
                      </p>
                    </TableCell>
                    <TableCell className='text-right'>
                      <Button variant='ghost' size='icon'>
                        <IconDotsVertical className='h-4 w-4' />
                      </Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>
    </TabsContent>
  )
}
