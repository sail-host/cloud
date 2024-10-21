import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import { Skeleton } from '@/components/ui/skeleton'
import { Switch } from '@/components/ui/switch'
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip'
import { cn } from '@/lib/utils'
import { useProjectStore } from '@/store/project-store'
import { IconChevronDown, IconCheck, IconHelpCircle } from '@tabler/icons-react'
import axios from 'axios'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { toast } from 'sonner'

interface Framework {
  name: string
  value: string
  image: string
}

const frameworks: Framework[] = [
  {
    name: 'Next.js',
    value: 'nextjs',
    image: '/images/frameworks/next.png',
  },
  {
    name: 'React',
    value: 'react',
    image: '/images/frameworks/react.png',
  },
  {
    name: 'Vue',
    value: 'vue',
    image: '/images/frameworks/vue.png',
  },
  {
    name: 'Vite',
    value: 'vite',
    image: '/images/frameworks/vite.png',
  },
  {
    name: 'Svelte',
    value: 'svelte',
    image: '/images/frameworks/svelte.png',
  },
  {
    name: 'Astro',
    value: 'astro',
    image: '/images/frameworks/astro.png',
  },
  {
    name: 'Nuxt',
    value: 'nuxt',
    image: '/images/frameworks/nuxt.png',
  },
  {
    name: 'Remix',
    value: 'remix',
    image: '/images/frameworks/remix.png',
  },
]

export function GeneralTab() {
  const [open, setOpen] = useState(false)
  const navigate = useNavigate()
  const [selectedFramework, setSelectedFramework] = useState<
    Framework['value'] | null
  >(null)

  const { project, setProject } = useProjectStore()

  const [projectName, setProjectName] = useState(project?.name || '')
  const [projectNameLoading, setProjectNameLoading] = useState(false)

  const handleSaveProjectName = () => {
    if (project && project.id) {
      setProjectNameLoading(true)

      axios
        .put(`/api/v1/project-setting/update-name/${project.name}`, {
          name: projectName,
        })
        .then((res) => {
          if (res.status === 200) {
            setProject({ ...project, name: projectName })
            toast.success('Project name updated')

            // Update the project name in the URL
            navigate(`/projects/${projectName}`)
          }
        })
        .catch((err) => {
          toast.error(err.response.data.message)
        })
        .finally(() => {
          setProjectNameLoading(false)
        })
    }
  }

  return (
    <div className='w-full space-y-6'>
      <Card>
        <CardHeader>
          <CardTitle className='text-xl'>Project Name</CardTitle>
          <CardDescription>Change the name of your project.</CardDescription>
        </CardHeader>
        <CardContent>
          <Input
            placeholder='Project Name'
            value={projectName}
            onChange={(e) => setProjectName(e.target.value)}
          />
        </CardContent>
        <CardFooter className='flex justify-end p-3 pr-6 border-t rounded-b-xl bg-muted dark:bg-muted/40'>
          <Button
            onClick={handleSaveProjectName}
            type='button'
            loading={projectNameLoading}
          >
            Save
          </Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className='text-xl'>Build and Output Settings</CardTitle>
          <CardDescription>
            Change the build and output settings of your project.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className='w-full space-y-4'>
            <div className='flex flex-col gap-2'>
              <Label className='font-light text-muted-foreground'>
                Project Name
              </Label>
              <Popover open={open} onOpenChange={setOpen}>
                <PopoverTrigger asChild>
                  <Button
                    variant='outline'
                    role='combobox'
                    aria-expanded={open}
                    className='justify-between w-full'
                  >
                    {selectedFramework &&
                    frameworks.find(
                      (framework) => framework.value === selectedFramework
                    ) ? (
                      <div className='flex items-center'>
                        <img
                          src={
                            frameworks.find(
                              (framework) =>
                                framework.value === selectedFramework
                            )?.image
                          }
                          alt={
                            frameworks.find(
                              (framework) =>
                                framework.value === selectedFramework
                            )?.name
                          }
                          width={20}
                          height={20}
                          className='mr-2'
                        />
                        {
                          frameworks.find(
                            (framework) => framework.value === selectedFramework
                          )?.name
                        }
                      </div>
                    ) : (
                      'Select framework...'
                    )}
                    <IconChevronDown className='w-4 h-4 ml-2 opacity-50 shrink-0' />
                  </Button>
                </PopoverTrigger>
                <PopoverContent className='w-[740px] p-0'>
                  <Command>
                    <CommandInput placeholder='Search framework...' />
                    <CommandList>
                      <CommandEmpty>No framework found.</CommandEmpty>
                      <CommandGroup>
                        {frameworks.map((framework) => (
                          <CommandItem
                            key={framework.value}
                            value={framework.value}
                            onSelect={(currentValue) => {
                              setSelectedFramework(
                                currentValue === selectedFramework
                                  ? ''
                                  : currentValue
                              )
                              setOpen(false)
                            }}
                            className='relative'
                          >
                            <img
                              src={framework.image}
                              alt={framework.name}
                              width={20}
                              height={20}
                              className='mr-2'
                            />
                            {framework.name}
                            <IconCheck
                              className={cn(
                                'absolute right-2 h-4 w-4',
                                selectedFramework === framework.value
                                  ? 'opacity-100'
                                  : 'opacity-0'
                              )}
                            />
                          </CommandItem>
                        ))}
                      </CommandGroup>
                    </CommandList>
                  </Command>
                </PopoverContent>
              </Popover>
            </div>

            <div className='flex flex-col gap-2'>
              <Label className='flex items-center gap-1 font-light text-muted-foreground'>
                Build Command
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger>
                      <IconHelpCircle className='w-5 h-5' />
                    </TooltipTrigger>
                    <TooltipContent className='border bg-background'>
                      <p className='max-w-[300px] text-center'>
                        The command your frontend framework provides for
                        compiling your code.
                      </p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
              </Label>
              <div className='grid grid-cols-7 gap-2'>
                <Input
                  placeholder='npm run build or yarn build'
                  className='col-span-6'
                  disabled
                />
                <div className='flex items-center justify-end col-span-1 space-x-2'>
                  <Label htmlFor='override-build-command'>Override</Label>{' '}
                  <Switch id='override-build-command' />
                </div>
              </div>
            </div>

            <div className='flex flex-col gap-2'>
              <Label className='flex items-center gap-1 font-light text-muted-foreground'>
                Output Directory
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger>
                      <IconHelpCircle className='w-5 h-5' />
                    </TooltipTrigger>
                    <TooltipContent className='border bg-background'>
                      <p className='max-w-[300px] text-center'>
                        The directory in which your compiled frontend will be
                        located.
                      </p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
              </Label>
              <div className='grid grid-cols-7 gap-2'>
                <Input
                  placeholder='public or dist'
                  className='col-span-6'
                  disabled
                />
                <div className='flex items-center justify-end col-span-1 space-x-2'>
                  <Label htmlFor='override-output-directory'>Override</Label>{' '}
                  <Switch id='override-output-directory' />
                </div>
              </div>
            </div>

            <div className='flex flex-col gap-2'>
              <Label className='flex items-center gap-1 font-light text-muted-foreground'>
                Install Command
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger>
                      <IconHelpCircle className='w-5 h-5' />
                    </TooltipTrigger>
                    <TooltipContent className='border bg-background'>
                      <p className='max-w-[300px] text-center'>
                        The command that is used to install your Project's
                        software dependencies. If you don't need to install
                        dependencies, override this field and leave it empty.
                      </p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
              </Label>
              <div className='grid grid-cols-7 gap-2'>
                <Input
                  placeholder='npm install, yarn install, pnpm install, bun install, etc.'
                  className='col-span-6'
                  disabled
                />
                <div className='flex items-center justify-end col-span-1 space-x-2'>
                  <Label htmlFor='override-install-command'>Override</Label>{' '}
                  <Switch id='override-install-command' />
                </div>
              </div>
            </div>
          </div>
        </CardContent>
        <CardFooter className='flex justify-end p-3 pr-6 border-t rounded-b-xl bg-muted dark:bg-muted/40'>
          <Button>Save</Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle className='text-xl'>Root Directory</CardTitle>
          <CardDescription>
            Change the root directory of your project.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className='grid grid-cols-12 gap-2'>
            <Input
              placeholder='src'
              className='col-span-11'
              disabled
              value='./'
            />
            <Button variant='outline' className='col-span-1'>
              Edit
            </Button>
          </div>
        </CardContent>
        <CardFooter className='flex justify-end p-3 pr-6 border-t rounded-b-xl bg-muted dark:bg-muted/40'>
          <Button>Save</Button>
        </CardFooter>
      </Card>

      <Card className='border-red-500 dark:border-red-500/30'>
        <CardHeader>
          <CardTitle className='text-xl'>Delete Project</CardTitle>
          <CardDescription>
            The project will be permanently deleted, including its deployments
            and domains. This action is irreversible and can not be undone.
          </CardDescription>
        </CardHeader>
        <CardContent className='border-t'>
          <div className='flex flex-row gap-4 mt-6'>
            <Skeleton className='h-20 w-36' />
            <div className='flex flex-col justify-center'>
              <h5 className='text-base font-bold'>solvie-dashboard</h5>
              <p className='text-sm text-muted-foreground'>
                Last Update: 2 days ago
              </p>
            </div>
          </div>
        </CardContent>
        <CardFooter className='flex justify-end p-3 pr-6 border-t border-red-500 rounded-b-xl bg-red-600/10 dark:border-red-500/30 dark:bg-red-500/10'>
          <Button variant='destructive'>Delete Project</Button>
        </CardFooter>
      </Card>
    </div>
  )
}
