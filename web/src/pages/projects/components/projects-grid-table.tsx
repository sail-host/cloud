import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
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
import { Project } from '../list'
import { formatDistanceToNow } from '@/lib/utils'

interface Props {
  projects: Project[]
}

export function ProjectsGridTable({ projects }: Props) {
  return (
    <div className='mt-4 grid grid-cols-12 gap-4'>
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
      <CardHeader className='flex flex-row items-center justify-between space-y-0 p-3 pb-2'>
        <CardTitle className='inline-flex text-sm font-medium'>
          <IconFolderCode size={28} className='mr-2 mt-1' />
          <div className='flex flex-col -space-y-0.5 text-sm'>
            <Link
              to={`/projects/${project.name.toLocaleLowerCase()}`}
              className='cursor-pointer font-medium hover:underline'
            >
              {project.name}
            </Link>
            <a
              href={`https://${project.domain}`}
              target='_blank'
              rel='noopener noreferrer'
              className='font-extralight hover:underline'
            >
              {project.domain}
            </a>
          </div>
        </CardTitle>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant='ghost' className='h-auto w-auto px-2'>
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
        <div className='text-sm font-light'>{project.git_commit}</div>
        <p className='flex items-center font-light text-muted-foreground'>
          {formatDistanceToNow(new Date(project.git_date))} ago on
          <span className='ml-1 font-semibold text-foreground'>
            <IconGitBranch size={16} className='-mt-0.5 mr-1 inline-block' />
            {project.git_branch}
          </span>
        </p>
      </CardContent>
    </Card>
  )
}
