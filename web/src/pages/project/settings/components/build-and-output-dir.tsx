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

export function BuildAndOutputDir() {
  const { project } = useProjectStore()
  const [open, setOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const [selectedFramework, setSelectedFramework] = useState(
    project?.project_framework || ''
  )
  const [buildCommand, setBuildCommand] = useState(project?.build_command || '')
  const [overrideBuildCommand, setOverrideBuildCommand] = useState(
    project?.build_command ? true : false
  )
  const [outputDir, setOutputDir] = useState(project?.output_dir || '')
  const [overrideOutputDir, setOverrideOutputDir] = useState(
    project?.output_dir ? true : false
  )
  const [installCommand, setInstallCommand] = useState(
    project?.install_command || ''
  )
  const [overrideInstallCommand, setOverrideInstallCommand] = useState(
    project?.install_command ? true : false
  )

  const handleSave = () => {
    setLoading(true)

    axios
      .put(
        `/api/v1/project-setting/update-build-and-output-dir/${project?.name}`,
        {
          framework: selectedFramework,
          build_command: overrideBuildCommand ? buildCommand : '',
          output_dir: overrideOutputDir ? outputDir : '',
          install_command: overrideInstallCommand ? installCommand : '',
        }
      )
      .then((res) => {
        if (res.status === 200) {
          toast.success('Build and output settings updated')
          toast.info('Please redeploy your project to apply the changes.')
        } else {
          toast.error(res.data?.message)
        }
      })
      .catch((error) => {
        toast.error(error.response?.data?.message || 'Something went wrong')
      })
      .finally(() => {
        setLoading(false)
      })
  }

  return (
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
                            (framework) => framework.value === selectedFramework
                          )?.image
                        }
                        alt={
                          frameworks.find(
                            (framework) => framework.value === selectedFramework
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
                      The command your frontend framework provides for compiling
                      your code.
                    </p>
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>
            </Label>
            <div className='grid grid-cols-7 gap-2'>
              <Input
                placeholder='npm run build or yarn build'
                className='col-span-6'
                disabled={!overrideBuildCommand}
                value={buildCommand}
                onChange={(e) => setBuildCommand(e.target.value)}
              />
              <div className='flex items-center justify-end col-span-1 space-x-2'>
                <Label htmlFor='override-build-command'>Override</Label>{' '}
                <Switch
                  id='override-build-command'
                  checked={overrideBuildCommand}
                  onCheckedChange={(checked) =>
                    setOverrideBuildCommand(checked)
                  }
                />
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
                disabled={!overrideOutputDir}
                value={outputDir}
                onChange={(e) => setOutputDir(e.target.value)}
              />
              <div className='flex items-center justify-end col-span-1 space-x-2'>
                <Label htmlFor='override-output-directory'>Override</Label>{' '}
                <Switch
                  id='override-output-directory'
                  checked={overrideOutputDir}
                  onCheckedChange={(checked) => setOverrideOutputDir(checked)}
                />
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
                disabled={!overrideInstallCommand}
                value={installCommand}
                onChange={(e) => setInstallCommand(e.target.value)}
              />
              <div className='flex items-center justify-end col-span-1 space-x-2'>
                <Label htmlFor='override-install-command'>Override</Label>{' '}
                <Switch
                  id='override-install-command'
                  checked={overrideInstallCommand}
                  onCheckedChange={(checked) =>
                    setOverrideInstallCommand(checked)
                  }
                />
              </div>
            </div>
          </div>
        </div>
      </CardContent>
      <CardFooter className='flex justify-end p-3 pr-6 border-t rounded-b-xl bg-muted dark:bg-muted/40'>
        <Button loading={loading} onClick={handleSave} type='button'>
          Save
        </Button>
      </CardFooter>
    </Card>
  )
}
