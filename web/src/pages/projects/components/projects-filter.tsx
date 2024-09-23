import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Button } from '@/components/ui/button'
import { IconLayoutGrid, IconList, IconPlus } from '@tabler/icons-react'
import { Link } from 'react-router-dom'

export function ProjectsFilter() {
  return (
    <div className='-mx-4 flex items-center gap-2.5 px-4 py-1'>
      <div className='w-full'>
        <Input placeholder='Search projects' />
      </div>
      <div className='w-[150px]'>
        <Select>
          <SelectTrigger className='w-full'>
            <SelectValue placeholder='Sort by' />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectItem value='name'>Sort by Name</SelectItem>
              <SelectItem value='domain'>Sort by Domain</SelectItem>
              <SelectItem value='last_commit'>Sort by Last Commit</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>
      <div className=''>
        <div className='inline-flex items-center gap-x-1 rounded-lg border p-0.5'>
          <Button
            variant='secondary'
            className='h-auto w-auto px-2'
            type='button'
          >
            <IconLayoutGrid className='h-4 w-4' />
          </Button>
          <Button variant='ghost' className='h-auto w-auto px-2' type='button'>
            <IconList className='h-4 w-4' />
          </Button>
        </div>
      </div>
      <div className=''>
        <Button asChild variant='default'>
          <Link to='/projects/create'>
            <IconPlus className='mr-2' />
            Add Project
          </Link>
        </Button>
      </div>
    </div>
  )
}
