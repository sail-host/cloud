import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import {
  IconCircleArrowUp,
  IconDotsVertical,
  IconGitCommit,
  IconGitPullRequest,
  IconReload,
} from '@tabler/icons-react'
import { Table, TableBody, TableCell, TableRow } from '@/components/ui/table'
import { Badge } from '@/components/ui/badge'
import { formatBytes, formatDistanceToNow, formatTime } from '@/lib/utils'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { BaseResponse } from '@/types/base'
import { Loading } from '@/components/custom/loading'
import { toast } from 'sonner'

interface Deployment {
  id: number
  status: 'pending' | 'building' | 'deploying' | 'running' | 'error' | 'success'
  created_at: string
  git_hash: string
  git_commit: string
  git_branch: string
  git_date: string
  git_url: string
  is_current: boolean
  size: number
  user: string
  deployment_time: number
}

export function DeploymentsTab({ uuid }: { uuid?: string }) {
  const [loading, setLoading] = useState(true)
  const [deployments, setDeployments] = useState<Deployment[]>([])

  const fetchDeployments = () => {
    axios
      .get<BaseResponse<Deployment[]>>(`/api/v1/project/deployments/${uuid}`)
      .then((res) => {
        if (res.data.status === 'success') {
          setDeployments(res.data.data || [])
        } else {
          toast.error(res.data.message)
        }
      })
      .catch((err) => {
        console.error(err)
      })
      .finally(() => {
        setLoading(false)
      })
  }

  useEffect(() => {
    fetchDeployments()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return (
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
            <IconReload className='w-4 h-4 mr-2' />
            Redeploy
          </Button>
        </div>
      </CardHeader>
      <CardContent className='text-sm'>
        <div>
          <Table className='w-full'>
            <TableBody>
              {loading ? (
                <TableRow>
                  <TableCell colSpan={5} className='text-center'>
                    <Loading loading={loading} />
                  </TableCell>
                </TableRow>
              ) : (
                deployments?.map((deployment) => (
                  <TableRow key={deployment.id}>
                    <TableCell className=''>
                      <div className='font-medium'>
                        {deployment.git_hash.slice(0, 10)}
                      </div>
                      <div className='flex items-center capitalize gap-x-2 text-muted-foreground'>
                        {deployment.git_branch}

                        {deployment.is_current && (
                          <Badge className='px-1 py-0 text-xs rounded-full bg-blue-500/40 hover:bg-blue-500/40'>
                            <IconCircleArrowUp className='mr-0.5 inline-block h-3 w-3' />
                            Current
                          </Badge>
                        )}
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className='flex flex-col'>
                        <div className='flex items-center'>
                          {deployment.status === 'success' && (
                            <>
                              <div className='w-2 h-2 bg-green-500 rounded-full' />
                              <span className='ml-1'>Active</span>
                            </>
                          )}
                          {deployment.status === 'error' && (
                            <>
                              <div className='w-2 h-2 bg-red-500 rounded-full' />
                              <span className='ml-1'>Error</span>
                            </>
                          )}
                          {deployment.status === 'pending' && (
                            <>
                              <div className='w-2 h-2 bg-yellow-500 rounded-full' />
                              <span className='ml-1'>Pending</span>
                            </>
                          )}
                          {deployment.status === 'building' && (
                            <>
                              <div className='w-2 h-2 bg-blue-500 rounded-full' />
                              <span className='ml-1'>Building</span>
                            </>
                          )}
                          {deployment.status === 'deploying' && (
                            <>
                              <div className='w-2 h-2 bg-purple-500 rounded-full' />
                              <span className='ml-1'>Deploying</span>
                            </>
                          )}
                        </div>
                        <div className='text-muted-foreground'>
                          {deployment.status === 'success' ? (
                            <p>
                              {formatTime(deployment.deployment_time)} (
                              {formatBytes(deployment.size).toString()})
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
                          <IconGitPullRequest className='w-4 h-4' />
                          <a
                            href={`${deployment.git_url}/tree/${deployment.git_branch}`}
                            className='font-light hover:underline'
                            target='_blank'
                            rel='noreferrer'
                          >
                            {deployment.git_branch}
                          </a>
                        </div>
                        <div className='flex items-center gap-x-2'>
                          <IconGitCommit className='w-4 h-4' />
                          <a
                            href={`${deployment.git_url}/commit/${deployment.git_hash}`}
                            className='flex items-center font-light gap-x-2 hover:underline'
                            target='_blank'
                            rel='noreferrer'
                          >
                            <span>{deployment.git_hash.slice(0, 7)}</span>
                            <span className='text-muted-foreground'>
                              {deployment.git_commit.length > 30
                                ? `${deployment.git_commit.slice(0, 30)}...`
                                : deployment.git_commit}
                            </span>
                          </a>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell className='text-right'>
                      <p className='font-light text-muted-foreground'>
                        {formatDistanceToNow(new Date(deployment.created_at))}{' '}
                        by{' '}
                        <span className='font-medium text-gray-800 hover:underline dark:text-white/90'>
                          {deployment.user}
                        </span>
                      </p>
                    </TableCell>
                    <TableCell className='text-right'>
                      <Button variant='ghost' size='icon'>
                        <IconDotsVertical className='w-4 h-4' />
                      </Button>
                    </TableCell>
                  </TableRow>
                ))
              )}
            </TableBody>
          </Table>
        </div>
      </CardContent>
    </Card>
  )
}
