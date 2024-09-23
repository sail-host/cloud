import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
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
import { useState } from 'react'
import {
  IconBrandGithub,
  IconCheck,
  IconChevronDown,
  IconGitBranch,
  IconHelpCircle,
  IconTrashX,
} from '@tabler/icons-react'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import { Switch } from '@/components/ui/switch'
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip'
import { Link } from 'react-router-dom'

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

export function ConfigureNewProject() {
  const [open, setOpen] = useState(false)
  const [selectedFramework, setSelectedFramework] = useState<
    Framework['value'] | null
  >(null)

  return (
    <div className='grid w-full grid-cols-12 gap-4 mt-4'>
      <Card className='col-span-8'>
        <CardHeader>
          <CardTitle>Configure Project</CardTitle>
        </CardHeader>
        <CardContent>
          <div className='w-full pt-6 space-y-4 border-t'>
            <div className='flex flex-col gap-2'>
              <Label>Project Name</Label>
              <Input placeholder='my-project' />
            </div>

            <div className='flex flex-col gap-2'>
              <Label>Project Name</Label>
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
              <Label>Root Directory</Label>
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
            </div>

            <Accordion
              type='single'
              collapsible
              className='px-3 border rounded-md'
            >
              <AccordionItem value='item-1'>
                <AccordionTrigger>Build and Output Settings</AccordionTrigger>
                <AccordionContent>
                  <div className='mt-3 space-y-4'>
                    <div className='flex flex-col gap-2'>
                      <Label className='flex items-center gap-1'>
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
                      <div className='grid grid-cols-12 gap-2'>
                        <Input
                          placeholder='npm run build or yarn build'
                          className='col-span-10'
                          disabled
                        />
                        <div className='flex items-center space-x-2'>
                          <Label htmlFor='override-build-command'>
                            Override
                          </Label>{' '}
                          <Switch id='override-build-command' />
                        </div>
                      </div>
                    </div>

                    <div className='flex flex-col gap-2'>
                      <Label className='flex items-center gap-1'>
                        Output Directory
                        <TooltipProvider>
                          <Tooltip>
                            <TooltipTrigger>
                              <IconHelpCircle className='w-5 h-5' />
                            </TooltipTrigger>
                            <TooltipContent className='border bg-background'>
                              <p className='max-w-[300px] text-center'>
                                The directory in which your compiled frontend
                                will be located.
                              </p>
                            </TooltipContent>
                          </Tooltip>
                        </TooltipProvider>
                      </Label>
                      <div className='grid grid-cols-12 gap-2'>
                        <Input
                          placeholder='public or dist'
                          className='col-span-10'
                          disabled
                        />
                        <div className='flex items-center space-x-2'>
                          <Label htmlFor='override-output-directory'>
                            Override
                          </Label>{' '}
                          <Switch id='override-output-directory' />
                        </div>
                      </div>
                    </div>

                    <div className='flex flex-col gap-2'>
                      <Label className='flex items-center gap-1'>
                        Install Command
                        <TooltipProvider>
                          <Tooltip>
                            <TooltipTrigger>
                              <IconHelpCircle className='w-5 h-5' />
                            </TooltipTrigger>
                            <TooltipContent className='border bg-background'>
                              <p className='max-w-[300px] text-center'>
                                The command that is used to install your
                                Project's software dependencies. If you don't
                                need to install dependencies, override this
                                field and leave it empty.
                              </p>
                            </TooltipContent>
                          </Tooltip>
                        </TooltipProvider>
                      </Label>
                      <div className='grid grid-cols-12 gap-2'>
                        <Input
                          placeholder='npm install, yarn install, pnpm install, bun install, etc.'
                          className='col-span-10'
                          disabled
                        />
                        <div className='flex items-center space-x-2'>
                          <Label htmlFor='override-install-command'>
                            Override
                          </Label>{' '}
                          <Switch id='override-install-command' />
                        </div>
                      </div>
                    </div>
                  </div>
                </AccordionContent>
              </AccordionItem>
            </Accordion>

            <Accordion
              type='single'
              collapsible
              className='px-3 border rounded-md'
            >
              <AccordionItem value='item-1'>
                <AccordionTrigger>Environment Variables</AccordionTrigger>
                <AccordionContent>
                  <div className='mt-3'>
                    <div className='grid items-end w-full grid-cols-11 gap-2 pr-1'>
                      <div className='col-span-5'>
                        <div className='flex flex-col gap-2'>
                          <Label>Name</Label>
                          <Input placeholder='NEXT_PUBLIC_API_URL' />
                        </div>
                      </div>
                      <div className='col-span-5'>
                        <div className='flex flex-col gap-2'>
                          <Label>Value</Label>
                          <Input placeholder='https://api.example.com' />
                        </div>
                      </div>
                      <div className='col-span-1'>
                        <Button variant='outline'>Add</Button>
                      </div>
                    </div>

                    <div className='mt-2'>
                      <p className='text-sm text-muted-foreground'>
                        <span className='font-medium'>TIP:</span> Paste a .env
                        above to populate the form
                      </p>
                    </div>

                    <div className='grid grid-cols-11 gap-2 p-2 mt-3 text-sm border rounded-md text-muted-foreground'>
                      <div className='col-span-5'>Key</div>
                      <div className='col-span-5'>Value</div>
                    </div>

                    <div className='grid grid-cols-11 gap-2 px-1 mt-2'>
                      <div className='col-span-5'>
                        <Input placeholder='NEXT_PUBLIC_API_URL' />
                      </div>
                      <div className='col-span-5'>
                        <Input placeholder='https://api.example.com' />
                      </div>
                      <div className='col-span-1'>
                        <Button variant='outline' className='w-full'>
                          <IconTrashX className='w-4 h-4' />
                        </Button>
                      </div>
                    </div>
                  </div>
                </AccordionContent>
              </AccordionItem>
            </Accordion>

            <div className='flex justify-end'>
              <Button>Deploy Project</Button>
            </div>
          </div>
        </CardContent>
      </Card>
      <div className='col-span-4'>
        <Card>
          <CardHeader>
            <CardTitle>Project Preview</CardTitle>
          </CardHeader>
          <CardContent>
            <div className='flex flex-col gap-2'>
              <div>
                <p className='text-sm text-muted-foreground'>Project Name</p>
                <p className='text-lg font-medium'>my-project</p>
              </div>
              <div>
                <p className='text-sm text-muted-foreground'>Framework</p>
                <p className='text-lg font-medium'>Next.js</p>
              </div>
              <div>
                <p className='text-sm text-muted-foreground'>Git Repository</p>
                <Link
                  to='https://github.com/user/my-project.git'
                  target='_blank'
                  className='flex items-center gap-1 hover:underline'
                >
                  <IconBrandGithub className='w-5 h-5' />
                  user/my-project
                </Link>
              </div>
              <div>
                <p className='text-sm text-muted-foreground'>Branch</p>
                <p className='flex items-center gap-1 text-lg font-medium'>
                  <IconGitBranch className='w-5 h-5' />
                  master
                </p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
