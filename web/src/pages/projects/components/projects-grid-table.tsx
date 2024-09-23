import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Project } from '../data/projects'
import {
  IconDotsVertical,
  IconFolderCode,
  IconGitBranch,
} from '@tabler/icons-react'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Button } from '@/components/ui/button'
import { Link } from 'react-router-dom'

interface Props {
  projects: Project[]
}

export function ProjectsGridTable({ projects }: Props) {
  return (
    <div className='grid grid-cols-12 gap-4 mt-4'>
      {projects.map((project) => (
        <div
          key={project.id}
          className='col-span-12 md:col-span-6 lg:col-span-4'
        >
          <ProjectCard project={project} />
        </div>
      ))}
    </div>
  )
}

function ProjectCard({ project }: { project: Project }) {
  return (
    <Card className='col-span-3 text-sm'>
      <CardHeader className='flex flex-row items-center justify-between p-3 pb-2 space-y-0'>
        <CardTitle className='inline-flex text-sm font-medium'>
          <IconFolderCode size={28} className='mt-1 mr-2' />
          <div className='flex flex-col -space-y-0.5 text-sm'>
            <Link
              to={`/projects/${project.name.toLocaleLowerCase()}`}
              className='font-medium cursor-pointer hover:underline'
            >
              {project.name}
            </Link>
            <a href='#' className='font-extralight hover:underline'>
              {project.domain}
            </a>
          </div>
        </CardTitle>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant='ghost' className='w-auto h-auto px-2'>
              <IconDotsVertical size={20} />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Profile</DropdownMenuItem>
            <DropdownMenuItem>Billing</DropdownMenuItem>
            <DropdownMenuItem>Team</DropdownMenuItem>
            <DropdownMenuItem>Subscription</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </CardHeader>
      <CardContent className='p-3 pt-0'>
        <div className='text-sm font-light'>
          {project.git.last_commit.message}
        </div>
        <p className='flex items-center font-light text-muted-foreground'>
          {project.git.last_commit.date} ago on
          <span className='ml-1 font-semibold text-foreground'>
            <IconGitBranch size={16} className='-mt-0.5 mr-1 inline-block' />
            {project.git.branch}
          </span>
        </p>
      </CardContent>
    </Card>
  )
}
