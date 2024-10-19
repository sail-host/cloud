import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { BuildOutputSettingsTab } from './build-output-settings-tab'
import { EnvVariableTab } from './env-variable-tab'
import { ProjectInformation } from './project-information'
import { StepTwoRightCard } from './step-two-right-card'
import { useState } from 'react'
import axios from 'axios'
import {
  useProjectCreateStore,
  useProjectSettingStore,
  useProjectStore,
} from '@/store/project-create-store'
import { toast } from 'sonner'

export function ConfigureNewProject() {
  const [loading, setLoading] = useState(false)
  const { gitAccount } = useProjectCreateStore()
  const { buildCommand, environments, installCommand, outputDir } =
    useProjectSettingStore()
  const { project, projectFramework, projectName, rootDir } = useProjectStore()

  const handleDeploy = () => {
    setLoading(true)
    axios
      .post('/api/v1/project/create', {
        name: projectName,
        framework: projectFramework,
        rootDir: rootDir,
        git_url: gitAccount?.url,
        git_repo: project?.name,
        git_id: gitAccount?.id,
        production_branch: project?.default_branch,
        build_command: buildCommand,
        install_command: installCommand,
        output_dir: outputDir,
        environments: environments,
      })
      .then((res) => {
        console.log(res)
        toast.success('Project created successfully')
      })
      .catch((err) => {
        toast.error(err.response?.data?.message || 'Failed to create project')
      })
      .finally(() => {
        setLoading(false)
      })
  }

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
              <Button type='button' onClick={handleDeploy} loading={loading}>
                Deploy Project
              </Button>
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
