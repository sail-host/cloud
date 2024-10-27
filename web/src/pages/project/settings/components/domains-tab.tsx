import Loader from '@/components/loader'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Table, TableBody, TableCell, TableRow } from '@/components/ui/table'
import { useProjectStore } from '@/store/project-store'
import { BaseResponse } from '@/types/base'
import { IconCheck, IconGitBranch } from '@tabler/icons-react'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { NewDomainModal } from './new-domain-modal'

interface Domain {
  id: number
  createdAt: string
  updatedAt: string
  project_id: number
  domain_id: number
  domain: string
  valid: boolean
  configured: boolean
  // env: string
  // branch: string
}

export function DomainsTab() {
  const { project } = useProjectStore()
  const [domains, setDomains] = useState<Domain[]>([])
  const [loading, setLoading] = useState(true)

  const fetchDomains = () => {
    axios
      .get<BaseResponse<Domain[]>>(
        `/api/v1/project-setting/domains/${project?.name}`
      )
      .then((res) => {
        if (res.data.status === 'success') {
          setDomains(res.data.data)
        }
      })
      .catch((err) => {
        console.log(err)
      })
      .finally(() => {
        setLoading(false)
      })
  }

  useEffect(() => {
    fetchDomains()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return (
    <div className='w-full'>
      <Card>
        <CardHeader className='relative'>
          <CardTitle>Domains</CardTitle>
          <CardDescription>
            These domains are assigned to your Production Deployments.
          </CardDescription>
          <NewDomainModal fetchDomains={fetchDomains} />
        </CardHeader>
        <CardContent>
          <div>
            <Table>
              <TableBody>
                {domains.map((domain) => (
                  <TableRow key={domain.id}>
                    <TableCell className=''>
                      <div className='flex items-center gap-2'>
                        <p>{domain.domain}</p>
                        <Badge className='px-2 py-0 font-light capitalize bg-blue-500 rounded-full hover:bg-blue-500'>
                          Production
                        </Badge>
                      </div>

                      <div className='flex items-center mt-3 gap-x-4'>
                        <div className='flex items-center gap-x-2'>
                          <div className='flex items-center justify-center w-5 h-5 bg-green-500 rounded-full'>
                            <IconCheck size={16} />
                          </div>
                          <p className='text-sm font-light'>
                            Configuration is valid
                          </p>
                        </div>
                        <div className='flex items-center gap-x-2'>
                          <div className='flex items-center justify-center w-5 h-5 bg-green-500 rounded-full'>
                            <IconCheck size={16} />
                          </div>
                          <p className='text-sm font-light'>
                            DNS is configured
                          </p>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell className=''>
                      <div className='flex items-center text-sm font-light gap-x-2 text-muted-foreground'>
                        <IconGitBranch size={16} />
                        <p>
                          Active branch:{' '}
                          <span className='font-medium text-foreground'>
                            master
                          </span>
                        </p>
                      </div>
                    </TableCell>
                    <TableCell className='space-x-2 text-right'>
                      <Button variant='outline'>Edit</Button>
                      <Button variant='destructive'>Delete</Button>
                    </TableCell>
                  </TableRow>
                ))}

                {loading && (
                  <TableRow>
                    <TableCell colSpan={3} className='h-24 text-center'>
                      <Loader />
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
