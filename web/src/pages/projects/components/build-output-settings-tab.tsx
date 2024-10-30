import { IconHelpCircle } from '@tabler/icons-react'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import { Switch } from '@/components/ui/switch'
import { Label } from '@/components/ui/label'
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip'
import { Input } from '@/components/ui/input'
import { useProjectSettingStore } from '@/store/project-create-store'
import { useState } from 'react'

export function BuildOutputSettingsTab() {
  const {
    buildCommand,
    setBuildCommand,
    outputDir,
    setOutputDir,
    installCommand,
    setInstallCommand,
  } = useProjectSettingStore()
  const [buildEnabled, setBuildEnabled] = useState(false)
  const [outputDirEnabled, setOutputDirEnabled] = useState(false)
  const [installCommandEnabled, setInstallCommandEnabled] = useState(false)

  return (
    <Accordion type='single' collapsible className='rounded-md border px-3'>
      <AccordionItem value='item-1'>
        <AccordionTrigger>Build and Output Settings</AccordionTrigger>
        <AccordionContent>
          <div className='mt-3 space-y-4 pl-1'>
            <div className='flex flex-col gap-2'>
              <Label className='flex items-center gap-1'>
                Build Command
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger>
                      <IconHelpCircle className='h-5 w-5' />
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
                  disabled={!buildEnabled}
                  value={buildCommand}
                  onChange={(e) => setBuildCommand(e.target.value)}
                />
                <div className='flex items-center space-x-2'>
                  <Label htmlFor='override-build-command'>Override</Label>
                  <Switch
                    id='override-build-command'
                    checked={buildEnabled}
                    onClick={() => {
                      if (buildEnabled) {
                        setBuildCommand('')
                      }
                      setBuildEnabled(!buildEnabled)
                    }}
                  />
                </div>
              </div>
            </div>

            <div className='flex flex-col gap-2'>
              <Label className='flex items-center gap-1'>
                Output Directory
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger>
                      <IconHelpCircle className='h-5 w-5' />
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
              <div className='grid grid-cols-12 gap-2'>
                <Input
                  placeholder='public or dist'
                  className='col-span-10'
                  disabled={!outputDirEnabled}
                  value={outputDir}
                  onChange={(e) => setOutputDir(e.target.value)}
                />
                <div className='flex items-center space-x-2'>
                  <Label htmlFor='override-output-directory'>Override</Label>
                  <Switch
                    id='override-output-directory'
                    checked={outputDirEnabled}
                    onClick={() => {
                      if (outputDirEnabled) {
                        setOutputDir('')
                      }
                      setOutputDirEnabled(!outputDirEnabled)
                    }}
                  />
                </div>
              </div>
            </div>

            <div className='flex flex-col gap-2'>
              <Label className='flex items-center gap-1'>
                Install Command
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger>
                      <IconHelpCircle className='h-5 w-5' />
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
              <div className='grid grid-cols-12 gap-2'>
                <Input
                  placeholder='npm install, yarn install, pnpm install, bun install, etc.'
                  className='col-span-10'
                  disabled={!installCommandEnabled}
                  value={installCommand}
                  onChange={(e) => setInstallCommand(e.target.value)}
                />
                <div className='flex items-center space-x-2'>
                  <Label htmlFor='override-install-command'>Override</Label>
                  <Switch
                    id='override-install-command'
                    checked={installCommandEnabled}
                    onClick={() => {
                      if (installCommandEnabled) {
                        setInstallCommand('')
                      }
                      setInstallCommandEnabled(!installCommandEnabled)
                    }}
                  />
                </div>
              </div>
            </div>
          </div>
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  )
}
