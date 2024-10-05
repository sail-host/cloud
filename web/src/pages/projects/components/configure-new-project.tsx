import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { BuildOutputSettingsTab } from './build-output-settings-tab'
import { EnvVariableTab } from './env-variable-tab'
import { ProjectInformation } from './project-information'
import { StepTwoRightCard } from './step-two-right-card'

export function ConfigureNewProject() {
  return (
    <div className='grid w-full grid-cols-12 gap-4 mt-4'>
      <Card className='col-span-8'>
        <CardHeader>
          <CardTitle>Configure Project</CardTitle>
        </CardHeader>
        <CardContent>
          <div className='w-full pt-6 space-y-4 border-t'>
            <ProjectInformation />

            <BuildOutputSettingsTab />

            <EnvVariableTab />

            <div className='flex justify-end'>
              <Button>Deploy Project</Button>
            </div>
          </div>
        </CardContent>
      </Card>
      <div className='col-span-4'>
        <StepTwoRightCard />
      </div>
    </div>
  )
}
