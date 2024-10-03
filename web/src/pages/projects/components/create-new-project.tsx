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

interface CreateNewProjectProps {
  setStep: (step: '1' | '2') => void
}

export function CreateNewProject({ setStep }: CreateNewProjectProps) {
  const [open, setOpen] = useState(false)
  const [gitAccount, setGitAccount] = useState<GitAccount | null>(null)
  const [gitAccounts, setGitAccounts] = useState<GitAccount[]>([])
  const [projects, setProjects] = useState<Project[]>([])

  const fetchGitAccounts = () => {
    axios.get<BaseResponse<GitAccount[]>>('/api/v1/git/list').then((res) => {
      setGitAccounts(res.data.data || [])
    })
  }

  const fetchProjects = () => {
    axios
      .get<
        BaseResponse<Project[]>
      >(`/api/v1/git-internal/list/${gitAccount?.id}`)
      .then((res) => {
        setProjects(res.data.data || [])
      })
  }

  useEffect(() => {
    fetchGitAccounts()
  }, [])

  useEffect(() => {
    if (gitAccount) {
      fetchProjects()
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [gitAccount])

  return (
    <div className='grid w-full grid-cols-12 gap-4 mt-4'>
      <Card className='col-span-8'>
        <CardHeader>
          <CardTitle>Import Git Repository</CardTitle>
        </CardHeader>
        <CardContent>
          <div>
            <div className='flex items-center w-full gap-3'>
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
                            <IconBrandGithub className='w-4 h-4 mr-2' />
                          )}
                          {gitAccount.type == 'gitlab' && (
                            <IconBrandGitlab className='w-4 h-4 mr-2' />
                          )}
                          {gitAccount.type == 'bitbucket' && (
                            <IconBrandBitbucket className='w-4 h-4 mr-2' />
                          )}
                          {gitAccount.name}
                          <IconChevronDown className='absolute w-4 h-4 opacity-50 right-2 shrink-0' />
                        </div>
                      )}
                      {!gitAccount && (
                        <>
                          {'Select Git Account...'}
                          <IconChevronDown className='absolute w-4 h-4 opacity-50 right-2 shrink-0' />
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
                                <IconBrandGithub className='w-4 h-4 mr-2' />
                              )}
                              {account.type == 'gitlab' && (
                                <IconBrandGitlab className='w-4 h-4 mr-2' />
                              )}
                              {account.type == 'bitbucket' && (
                                <IconBrandBitbucket className='w-4 h-4 mr-2' />
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
                              className='flex items-center w-full'
                            >
                              <IconPlus className='w-4 h-4 mr-2' />
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
                <IconSearch className='absolute w-4 h-4 -translate-y-1/2 left-2 top-1/2 text-muted-foreground' />
                <Input placeholder='Repository URL' className='pl-8' />
              </div>
            </div>
            <Card className='mt-3 !p-0'>
              {projects.map((item, index) => (
                <ProjectCardItem key={index} item={item} setStep={setStep} />
              ))}
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
