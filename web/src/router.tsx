import { createBrowserRouter } from 'react-router-dom'
import GeneralError from './pages/errors/general-error'
import NotFoundError from './pages/errors/not-found-error'
import MaintenanceError from './pages/errors/maintenance-error'
import UnauthorisedError from './pages/errors/unauthorised-error.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    lazy: async () => {
      const BaseLayout = await import('./components/base-layout')
      return { Component: BaseLayout.default }
    },
    errorElement: <GeneralError />,
    children: [
      // Auth routes
      {
        path: '/login',
        lazy: async () => ({
          Component: (await import('./pages/auth/login.tsx')).default,
        }),
      },
      {
        path: '/register',
        lazy: async () => ({
          Component: (await import('./pages/auth/register.tsx')).default,
        }),
      },
      {
        path: '/forgot-password',
        lazy: async () => ({
          Component: (await import('./pages/auth/forgot-password')).default,
        }),
      },
      {
        path: '/otp',
        lazy: async () => ({
          Component: (await import('./pages/auth/otp')).default,
        }),
      },
    ],
  },

  // Main routes
  {
    path: '/',
    lazy: async () => {
      const AppShell = await import('./components/app-shell')
      return { Component: AppShell.default }
    },
    errorElement: <GeneralError />,
    children: [
      {
        index: true,
        lazy: async () => ({
          Component: (await import('@/pages/dashboard/index.tsx')).default,
        }),
      },
      {
        path: 'projects',
        lazy: async () => ({
          Component: (await import('@/pages/projects/list.tsx')).default,
        }),
      },
      {
        path: 'projects/create',
        lazy: async () => ({
          Component: (await import('@/pages/projects/create.tsx')).default,
        }),
      },
      {
        path: 'projects/:uuid',
        lazy: async () => ({
          Component: (await import('@/pages/project/show.tsx')).default,
        }),
      },
      {
        path: 'domains',
        lazy: async () => ({
          Component: (await import('@/pages/domains/list.tsx')).default,
        }),
      },
      {
        path: 'domains/create',
        lazy: async () => ({
          Component: (await import('@/pages/domains/create.tsx')).default,
        }),
      },
      {
        path: 'domains/edit/:id',
        lazy: async () => ({
          Component: (await import('@/pages/domains/edit.tsx')).default,
        }),
      },
      {
        path: 'git-accounts',
        lazy: async () => ({
          Component: (await import('@/pages/git-accounts/list.tsx')).default,
        }),
      },
      {
        path: 'git-accounts/create',
        lazy: async () => ({
          Component: (await import('@/pages/git-accounts/create.tsx')).default,
        }),
      },
      {
        path: 'git-accounts/edit/:id',
        lazy: async () => ({
          Component: (await import('@/pages/git-accounts/edit.tsx')).default,
        }),
      },
      // Old routes
      {
        path: 'settings',
        lazy: async () => ({
          Component: (await import('./pages/settings')).default,
        }),
        errorElement: <GeneralError />,
        children: [
          {
            index: true,
            lazy: async () => ({
              Component: (await import('./pages/settings/profile')).default,
            }),
          },
          {
            path: 'account',
            lazy: async () => ({
              Component: (await import('./pages/settings/account')).default,
            }),
          },
          {
            path: 'appearance',
            lazy: async () => ({
              Component: (await import('./pages/settings/appearance')).default,
            }),
          },
          {
            path: 'notifications',
            lazy: async () => ({
              Component: (await import('./pages/settings/notifications'))
                .default,
            }),
          },
          {
            path: 'display',
            lazy: async () => ({
              Component: (await import('./pages/settings/display')).default,
            }),
          },
          {
            path: 'error-example',
            lazy: async () => ({
              Component: (await import('./pages/settings/error-example'))
                .default,
            }),
            errorElement: <GeneralError className='h-[50svh]' minimal />,
          },
        ],
      },
    ],
  },

  // Error routes
  { path: '/500', Component: GeneralError },
  { path: '/404', Component: NotFoundError },
  { path: '/503', Component: MaintenanceError },
  { path: '/401', Component: UnauthorisedError },

  // Fallback 404 route
  { path: '*', Component: NotFoundError },
])

export default router
