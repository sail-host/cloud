import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { SelectFramework } from './select-framework'
import { Button } from '@/components/ui/button'
import { useEffect, useState } from 'react'
import { useProjectStore } from '@/store/project-create-store'
import axios from 'axios'
import { debounce } from 'lodash'

export function ProjectInformation() {
  const {
    projectName,
    projectFramework,
    rootDir,
    setProjectFramework,
    setRootDir,
    setProjectName,
  } = useProjectStore()
  const [editRootDir, setEditRootDir] = useState(false)
  const [projectNameAvailable, setProjectNameAvailable] = useState(true)

  useEffect(() => {
    const checkProjectName = debounce(() => {
      axios.get(`/api/v1/project/check?name=${projectName}`).then((res) => {
        setProjectNameAvailable(res.data.data)
      })
    }, 300)

    checkProjectName()

    return () => {
      checkProjectName.cancel()
    }
  }, [projectName])

  return (
    <>
      <div className='flex flex-col gap-2'>
        <Label>Project Name</Label>
        <Input
          placeholder='my-project'
          value={projectName}
          onChange={(e) => setProjectName(e.target.value)}
        />
        {!projectNameAvailable && (
          <p className='text-sm text-red-500'>Project name is already used</p>
        )}
      </div>

      <div className='flex flex-col gap-2'>
        <Label>Project Framework</Label>
        <SelectFramework
          selectedFramework={projectFramework}
          setSelectedFramework={setProjectFramework}
        />
      </div>

      <div className='flex flex-col gap-2'>
        <Label>Root Directory</Label>
        <div className='grid grid-cols-12 gap-2'>
          <Input
            placeholder='src'
            className='col-span-11'
            disabled={!editRootDir}
            value={rootDir}
            onChange={(e) => setRootDir(e.target.value)}
          />
          <Button
            variant='outline'
            className='col-span-1'
            onClick={() => {
              if (editRootDir) {
                setRootDir(rootDir)
              }
              setEditRootDir(!editRootDir)
            }}
            type='button'
          >
            {editRootDir ? 'Save' : 'Edit'}
          </Button>
        </div>
      </div>
    </>
  )
}
