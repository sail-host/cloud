import {
  IconBarrierBlock,
  IconChecklist,
  IconError404,
  IconExclamationCircle,
  IconLayoutDashboard,
  IconServerOff,
  IconSettings,
  IconLock,
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
  // Old routes
  {
    title: 'Error Pages',
    label: '',
    href: '',
    icon: <IconExclamationCircle size={18} />,
    sub: [
      {
        title: 'Not Found',
        label: '',
        href: '/404',
        icon: <IconError404 size={18} />,
      },
      {
        title: 'Internal Server Error',
        label: '',
        href: '/500',
        icon: <IconServerOff size={18} />,
      },
      {
        title: 'Maintenance Error',
        label: '',
        href: '/503',
        icon: <IconBarrierBlock size={18} />,
      },
      {
        title: 'Unauthorised Error',
        label: '',
        href: '/401',
        icon: <IconLock size={18} />,
      },
    ],
  },
  {
    title: 'Settings',
    label: '',
    href: '/settings',
    icon: <IconSettings size={18} />,
  },
]
