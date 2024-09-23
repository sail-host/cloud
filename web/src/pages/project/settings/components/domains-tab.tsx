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
import { IconCheck, IconGitBranch } from '@tabler/icons-react'

interface Domain {
  name: string
  url: string
  branch: string
  env: string
  isDNS: boolean
  isActive: boolean
  isValid: boolean
  createdAt: string
}

const domains: Domain[] = [
  {
    name: 'techwave.io',
    url: 'https://techwave.io',
    branch: 'main',
    env: 'production',
    isDNS: true,
    isActive: true,
    isValid: true,
    createdAt: '2024-03-01T12:30:00Z',
  },
  {
    name: 'cloudpulse.net',
    url: 'https://cloudpulse.net',
    branch: 'develop',
    env: 'staging',
    isDNS: true,
    isActive: false,
    isValid: true,
    createdAt: '2024-02-28T09:15:00Z',
  },
  {
    name: 'devforge.com',
    url: 'https://devforge.com',
    branch: 'feature/new-ui',
    env: 'development',
    isDNS: false,
    isActive: true,
    isValid: false,
    createdAt: '2024-02-25T18:45:00Z',
  },
  {
    name: 'codestream.org',
    url: 'https://codestream.org',
    branch: 'main',
    env: 'production',
    isDNS: true,
    isActive: true,
    isValid: true,
    createdAt: '2024-02-20T14:00:00Z',
  },
  {
    name: 'byteburst.app',
    url: 'https://byteburst.app',
    branch: 'release/v2.0',
    env: 'staging',
    isDNS: false,
    isActive: true,
    isValid: true,
    createdAt: '2024-03-05T10:20:00Z',
  },
]

export function DomainsTab() {
  return (
    <div className='w-full'>
      <Card>
        <CardHeader>
          <CardTitle>Domains</CardTitle>
          <CardDescription>
            These domains are assigned to your Production Deployments.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div>
            <Table>
              <TableBody>
                {domains.map((domain) => (
                  <TableRow key={domain.name}>
                    <TableCell className=''>
                      <div className='flex items-center gap-2'>
                        <p>{domain.name}</p>
                        <Badge className='px-2 py-0 font-light capitalize bg-blue-500 rounded-full hover:bg-blue-500'>
                          {domain.env}
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
                            {domain.branch}
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
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
