import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Skeleton } from '@/components/ui/skeleton'
import { useProjectStore } from '@/store/project-store'
import axios from 'axios'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { toast } from 'sonner'
import { BuildAndOutputDir } from './build-and-output-dir'

export function GeneralTab() {
  const { project, setProject } = useProjectStore()
  const navigate = useNavigate()

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

      <BuildAndOutputDir />

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
