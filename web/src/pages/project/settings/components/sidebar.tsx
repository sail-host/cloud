import { buttonVariants } from '@/components/custom/button'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { cn } from '@/lib/utils'
import {
  IconClock,
  IconFileCode,
  IconFileZip,
  IconGitPullRequest,
  IconLock,
  IconSettings,
  IconWorldCog,
} from '@tabler/icons-react'

interface SidebarItem {
  id: string
  title: string
  icon: React.ReactNode
}

const items: SidebarItem[] = [
  {
    id: 'general',
    title: 'General',
    icon: <IconSettings />,
  },
  {
    id: 'domains',
    title: 'Domains',
    icon: <IconWorldCog />,
  },
  {
    id: 'environments',
    title: 'Environments',
    icon: <IconFileZip />,
  },
  {
    id: 'nginx-config',
    title: 'Nginx Config',
    icon: <IconFileCode />,
  },
  {
    id: 'crontab',
    title: 'Crontab',
    icon: <IconClock />,
  },
  {
    id: 'git',
    title: 'Git',
    icon: <IconGitPullRequest />,
  },
  {
    id: 'ssl',
    title: 'SSL',
    icon: <IconLock />,
  },
]

interface SettingsTabSidebarProps {
  activeTab: string
  setActiveTab: (tab: string) => void
}

export function SettingsTabSidebar({
  activeTab,
  setActiveTab,
}: SettingsTabSidebarProps) {
  return (
    <>
      <div className='p-1 md:hidden'>
        <Select value={activeTab} onValueChange={setActiveTab}>
          <SelectTrigger className='h-12 sm:w-48'>
            <SelectValue placeholder='Settings' />
          </SelectTrigger>
          <SelectContent>
            {items.map((item) => (
              <SelectItem key={item.id} value={item.id}>
                <div className='flex gap-x-4 px-2 py-1'>
                  <span className='scale-125'>{item.icon}</span>
                  <span className='text-md'>{item.title}</span>
                </div>
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <div className='hidden w-full overflow-x-auto bg-background px-1 py-2 md:block'>
        <nav
          className={cn('flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1')}
        >
          {items.map((item) => (
            <div
              key={item.id}
              className={cn(
                buttonVariants({ variant: 'ghost' }),
                activeTab === item.id
                  ? 'bg-muted hover:bg-muted'
                  : 'hover:bg-transparent hover:underline',
                'cursor-pointer justify-start font-light'
              )}
              onClick={() => setActiveTab(item.id)}
            >
              <span className='mr-2'>{item.icon}</span>
              {item.title}
            </div>
          ))}
        </nav>
      </div>
    </>
  )
}
