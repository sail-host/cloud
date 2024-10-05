import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { SelectFramework } from './select-framework'
import { Button } from '@/components/ui/button'
import { useState } from 'react'

export function ProjectInformation() {
  const [selectedFramework, setSelectedFramework] = useState<string | null>('')

  return (
    <>
      <div className='flex flex-col gap-2'>
        <Label>Project Name</Label>
        <Input placeholder='my-project' />
      </div>

      <div className='flex flex-col gap-2'>
        <Label>Project Framework</Label>
        <SelectFramework
          selectedFramework={selectedFramework}
          setSelectedFramework={setSelectedFramework}
        />
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
    </>
  )
}
