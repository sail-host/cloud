import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useEffect, useState } from 'react'
import { cn } from '@/lib/utils'
import { Button } from '@/components/ui/button'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import {
  IconBrandBitbucket,
  IconBrandGithub,
  IconBrandGitlab,
  IconCheck,
  IconChevronDown,
  IconPlus,
  IconSearch,
} from '@tabler/icons-react'
import { Link } from 'react-router-dom'
import { Input } from '@/components/ui/input'
import { GitAccount } from '@/types/model'
import axios from 'axios'
import { BaseResponse } from '@/types/base'
import { Project, ProjectCardItem } from './project-card-item'
import { Loading } from '@/components/custom/loading'
import { useProjectCreateStore } from '@/store/project-create-store'

interface FetchProjectsResponse {
  status: 'success' | 'error'
  data: Project[]
  last_page: number
  next_page: number
}

export function CreateNewProject() {
  const { gitAccount, setGitAccount } = useProjectCreateStore()
  const [open, setOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const [gitAccounts, setGitAccounts] = useState<GitAccount[]>([])
  const [projects, setProjects] = useState<Project[]>([])
  const [lastPage, setLastPage] = useState<number>(0)
  const [page, setPage] = useState<number>(1)

  const fetchGitAccounts = () => {
    axios.get<BaseResponse<GitAccount[]>>('/api/v1/git/list').then((res) => {
      setGitAccounts(res.data.data || [])
    })
  }

  const fetchProjects = () => {
    setLoading(true)
    axios
      .get<FetchProjectsResponse>(
        `/api/v1/git-internal/list/${gitAccount?.id}?page=${page}`
      )
      .then((res) => {
        setProjects((i) => [...i, ...(res.data.data || [])])
        setLastPage(res.data.last_page)
      })
      .finally(() => setLoading(false))
  }

  useEffect(() => {
    fetchGitAccounts()
  }, [])

  useEffect(() => {
    if (gitAccount) {
      setPage(1)
      setProjects([])
      fetchProjects()
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [gitAccount])

  return (
    <div className='mt-4 grid w-full grid-cols-12 gap-4'>
      <Card className='col-span-8'>
        <CardHeader>
          <CardTitle>Import Git Repository</CardTitle>
        </CardHeader>
        <CardContent>
          <div>
            <div className='flex w-full items-center gap-3'>
              <div className=''>
                <Popover open={open} onOpenChange={setOpen}>
                  <PopoverTrigger asChild>
                    <Button
                      variant='outline'
                      role='combobox'
                      aria-expanded={open}
                      className='relative w-[300px] justify-between'
                    >
                      {gitAccount && (
                        <div className='flex items-center'>
                          {gitAccount.type == 'github' && (
                            <IconBrandGithub className='mr-2 h-4 w-4' />
                          )}
                          {gitAccount.type == 'gitlab' && (
                            <IconBrandGitlab className='mr-2 h-4 w-4' />
                          )}
                          {gitAccount.type == 'bitbucket' && (
                            <IconBrandBitbucket className='mr-2 h-4 w-4' />
                          )}
                          {gitAccount.name}
                          <IconChevronDown className='absolute right-2 h-4 w-4 shrink-0 opacity-50' />
                        </div>
                      )}
                      {!gitAccount && (
                        <>
                          {'Select Git Account...'}
                          <IconChevronDown className='absolute right-2 h-4 w-4 shrink-0 opacity-50' />
                        </>
                      )}
                    </Button>
                  </PopoverTrigger>
                  <PopoverContent className='w-[300px] p-0'>
                    <Command>
                      <CommandInput placeholder='Search Git Account...' />
                      <CommandList>
                        <CommandEmpty>No Git Account found.</CommandEmpty>
                        <CommandGroup>
                          {gitAccounts.map((account) => (
                            <CommandItem
                              key={account.id}
                              value={account.id}
                              onSelect={(currentValue) => {
                                setGitAccount(
                                  currentValue === gitAccount?.id
                                    ? null
                                    : account
                                )
                                setOpen(false)
                              }}
                              className='relative'
                            >
                              {account.type == 'github' && (
                                <IconBrandGithub className='mr-2 h-4 w-4' />
                              )}
                              {account.type == 'gitlab' && (
                                <IconBrandGitlab className='mr-2 h-4 w-4' />
                              )}
                              {account.type == 'bitbucket' && (
                                <IconBrandBitbucket className='mr-2 h-4 w-4' />
                              )}
                              {account.name}
                              <IconCheck
                                className={cn(
                                  'absolute right-2 h-4 w-4',
                                  gitAccount?.id === account.id
                                    ? 'opacity-100'
                                    : 'opacity-0'
                                )}
                              />
                            </CommandItem>
                          ))}
                          <CommandItem>
                            <Link
                              to='/git-accounts/create'
                              className='flex w-full items-center'
                            >
                              <IconPlus className='mr-2 h-4 w-4' />
                              Add Git Account
                            </Link>
                          </CommandItem>
                        </CommandGroup>
                      </CommandList>
                    </Command>
                  </PopoverContent>
                </Popover>
              </div>
              <div className='relative w-full'>
                <IconSearch className='absolute left-2 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground' />
                <Input placeholder='Repository URL' className='pl-8' />
              </div>
            </div>
            <Card className='mt-3 !p-0'>
              {projects.map((item, index) => (
                <ProjectCardItem key={index} item={item} />
              ))}

              {!gitAccount && (
                <div className='flex items-center justify-center p-4'>
                  <p className='text-muted-foreground'>
                    Please select a Git account to view repositories.
                  </p>
                </div>
              )}

              {gitAccount && projects.length === 0 && !loading && (
                <div className='flex items-center justify-center p-4'>
                  <p className='text-muted-foreground'>
                    No repositories found.
                  </p>
                </div>
              )}

              {loading && projects.length === 0 && (
                <div className='flex items-center justify-center p-4'>
                  <Loading loading />
                </div>
              )}

              {page < lastPage && (
                <div className='my-4 flex items-center justify-center'>
                  <Button
                    type='button'
                    onClick={() => {
                      setPage((i) => i + 1)
                      fetchProjects()
                    }}
                    loading={loading}
                  >
                    Load More
                  </Button>
                </div>
              )}
            </Card>
          </div>
        </CardContent>
      </Card>
      <div className='col-span-4'>
        <h1>Create New Project</h1>
      </div>
    </div>
  )
}
