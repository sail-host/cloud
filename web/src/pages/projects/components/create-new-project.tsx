import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useState } from 'react'
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
  IconLock,
  IconPlus,
  IconPointFilled,
  IconSearch,
} from '@tabler/icons-react'
import { Link } from 'react-router-dom'
import { Input } from '@/components/ui/input'
import { projects } from '../data/projects'

interface Account {
  type: 'github' | 'gitlab' | 'bitbucket'
  url: string
  username: string
}

const accounts: Account[] = [
  {
    type: 'github',
    url: 'https://github.com/johndoe',
    username: 'johndoe',
  },
  {
    type: 'gitlab',
    url: 'https://gitlab.com/janedoe',
    username: 'janedoe',
  },
  {
    type: 'bitbucket',
    url: 'https://bitbucket.org/alexsmith',
    username: 'alexsmith',
  },
  {
    type: 'github',
    url: 'https://github.com/sarahbrown',
    username: 'sarahbrown',
  },
]

interface CreateNewProjectProps {
  setStep: (step: '1' | '2') => void
}

export function CreateNewProject({ setStep }: CreateNewProjectProps) {
  const [open, setOpen] = useState(false)
  const [value, setValue] = useState('')

  const handleImport = () => {
    setStep('2')
  }

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
                      {value && (
                        <div className='flex items-center'>
                          {accounts.find(
                            (account) => account.username === value
                          )?.type == 'github' && (
                            <IconBrandGithub className='mr-2 h-4 w-4' />
                          )}
                          {accounts.find(
                            (account) => account.username === value
                          )?.type == 'gitlab' && (
                            <IconBrandGitlab className='mr-2 h-4 w-4' />
                          )}
                          {accounts.find(
                            (account) => account.username === value
                          )?.type == 'bitbucket' && (
                            <IconBrandBitbucket className='mr-2 h-4 w-4' />
                          )}
                          {value}
                          <IconChevronDown className='absolute right-2 h-4 w-4 shrink-0 opacity-50' />
                        </div>
                      )}
                      {!value && (
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
                          {accounts.map((account) => (
                            <CommandItem
                              key={account.username}
                              value={account.username}
                              onSelect={(currentValue) => {
                                setValue(
                                  currentValue === value ? '' : currentValue
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
                              {account.username}
                              <IconCheck
                                className={cn(
                                  'absolute right-2 h-4 w-4',
                                  value === account.username
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
                <div
                  key={index}
                  className='flex items-center justify-between border-b last:border-none'
                >
                  <div className='flex items-center gap-4 p-4'>
                    <img
                      src='https://api-frameworks.vercel.sh/framework-logos/vite.svg'
                      alt={item.name}
                      className='h-6 w-6 rounded-full'
                    />
                    <div className='flex items-center gap-1'>
                      <h3 className='text-sm font-medium'>{item.name}</h3>
                      <IconLock className='h-4 w-4 text-muted-foreground' />
                      <IconPointFilled className='h-2 w-2 text-muted-foreground/60' />
                      <span className='text-sm text-muted-foreground'>
                        {item.git.last_commit.date}
                      </span>
                    </div>
                  </div>
                  <div className='p-4'>
                    <Button
                      className='h-auto px-3 py-1.5 text-sm'
                      onClick={handleImport}
                      type='button'
                    >
                      Import
                    </Button>
                  </div>
                </div>
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
