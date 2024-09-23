import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { ProjectTab } from './components/project-tab'
import { DeploymentsTab } from './components/depoloyments-tab'
import { LogsTab } from './components/logs-tab'
import { StorageTab } from './components/storage-tab'
import { SettingsTab } from './settings/settings-tab'

export default function ProjectShow() {
  return (
    <Layout>
      {/* ===== Top Heading ===== */}
      <Layout.Header sticky>
        <Search />
        <div className='flex items-center ml-auto space-x-4'>
          <ThemeSwitch />
          <UserNav />
        </div>
      </Layout.Header>

      <Layout.Body>
        <div className='flex items-center justify-between mb-2 space-y-2'>
          <h1 className='text-2xl font-bold tracking-tight'>
            solvie-dashboard
          </h1>
        </div>
        <div className='mt-4'>
          <Tabs
            orientation='vertical'
            defaultValue='project'
            className='space-y-4'
          >
            <div className='w-full pb-2 overflow-x-auto'>
              <TabsList>
                <TabsTrigger value='project'>Project</TabsTrigger>
                <TabsTrigger value='deployments'>Deployments</TabsTrigger>
                <TabsTrigger value='logs'>Logs</TabsTrigger>
                <TabsTrigger value='storage'>Storage</TabsTrigger>
                <TabsTrigger value='settings'>Settings</TabsTrigger>
              </TabsList>
            </div>
            <ProjectTab />
            <DeploymentsTab />
            <LogsTab />
            <StorageTab />
            <SettingsTab />
          </Tabs>
        </div>
      </Layout.Body>
    </Layout>
  )
}
