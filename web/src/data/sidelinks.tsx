import {
  IconChecklist,
  IconLayoutDashboard,
  IconSettings,
  IconWorld,
  IconGitPullRequest,
} from '@tabler/icons-react'

export interface NavLink {
  title: string
  label?: string
  href: string
  icon: JSX.Element
}

export interface SideLink extends NavLink {
  sub?: NavLink[]
}

export const sidelinks: SideLink[] = [
  {
    title: 'Dashboard',
    label: '',
    href: '/',
    icon: <IconLayoutDashboard size={18} />,
  },
  {
    title: 'Projects',
    label: '',
    href: '/projects',
    icon: <IconChecklist size={18} />,
  },
  {
    title: 'Domains',
    label: '',
    href: '/domains',
    icon: <IconWorld size={18} />,
  },
  {
    title: 'Git Accounts',
    label: '',
    href: '/git-accounts',
    icon: <IconGitPullRequest size={18} />,
  },
  {
    title: 'Settings',
    label: '',
    href: '/settings',
    icon: <IconSettings size={18} />,
  },
]
