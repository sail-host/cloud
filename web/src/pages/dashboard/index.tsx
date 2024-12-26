import { Layout } from '@/components/custom/layout'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Search } from '@/components/search'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { RecentSales } from './components/recent-sales'
import { Overview } from './components/overview'
import {
  IconCpu,
  IconDatabase,
  IconGitBranch,
  IconSitemap,
} from '@tabler/icons-react'
import { Analytics } from './components/analytics'

export default function Dashboard() {
  return (
    <Layout>
      {/* ===== Top Heading ===== */}
      <Layout.Header>
        <div className='flex items-center ml-auto space-x-4'>
          <Search />
          <ThemeSwitch />
          <UserNav />
        </div>
      </Layout.Header>

      {/* ===== Main ===== */}
      <Layout.Body>
        <div className='flex items-center justify-between mb-2 space-y-2'>
          <h1 className='text-2xl font-bold tracking-tight'>Dashboard</h1>
        </div>
        <Tabs
          orientation='vertical'
          defaultValue='overview'
          className='space-y-4'
        >
          <div className='w-full pb-2 overflow-x-auto'>
            <TabsList>
              <TabsTrigger value='overview'>Overview</TabsTrigger>
              <TabsTrigger value='analytics'>Analytics</TabsTrigger>
              <TabsTrigger value='reports'>Reports</TabsTrigger>
              <TabsTrigger value='notifications'>Notifications</TabsTrigger>
            </TabsList>
          </div>
          <TabsContent value='overview' className='space-y-4'>
            <div className='grid gap-4 sm:grid-cols-2 lg:grid-cols-4'>
              <Card>
                <CardHeader className='flex flex-row items-center justify-between pb-2 space-y-0'>
                  <CardTitle className='text-sm font-medium'>
                    Projects
                  </CardTitle>
                  <IconSitemap className='w-4 h-4 text-muted-foreground' />
                </CardHeader>
                <CardContent>
                  <div className='text-2xl font-bold'>23/23</div>
                  <p className='text-xs text-muted-foreground'>
                    +2 projects from last month
                  </p>
                </CardContent>
              </Card>
              <Card>
                <CardHeader className='flex flex-row items-center justify-between pb-2 space-y-0'>
                  <CardTitle className='text-sm font-medium'>
                    Git Accounts
                  </CardTitle>
                  <IconGitBranch className='w-4 h-4 text-muted-foreground' />
                </CardHeader>
                <CardContent>
                  <div className='text-2xl font-bold'>2/2</div>
                  <p className='text-xs text-muted-foreground'>
                    +2 accounts from last month
                  </p>
                </CardContent>
              </Card>
              <Card>
                <CardHeader className='flex flex-row items-center justify-between pb-2 space-y-0'>
                  <CardTitle className='text-sm font-medium'>
                    Disk Usage
                  </CardTitle>
                  <IconDatabase className='w-4 h-4 text-muted-foreground' />
                </CardHeader>
                <CardContent>
                  <div className='text-2xl font-bold'>63/100GB</div>
                  <p className='text-xs text-muted-foreground'>
                    +19% from last month
                  </p>
                </CardContent>
              </Card>
              <Card>
                <CardHeader className='flex flex-row items-center justify-between pb-2 space-y-0'>
                  <CardTitle className='text-sm font-medium'>
                    Ram Usage
                  </CardTitle>
                  <IconCpu className='w-4 h-4 text-muted-foreground' />
                </CardHeader>
                <CardContent>
                  <div className='text-2xl font-bold'>3.5/8GB</div>
                  <p className='text-xs text-muted-foreground'>
                    +2% since last hour
                  </p>
                </CardContent>
              </Card>
            </div>
            <div className='grid grid-cols-1 gap-4 lg:grid-cols-7'>
              <Card className='col-span-1 lg:col-span-4'>
                <CardHeader>
                  <CardTitle>Disk Usage</CardTitle>
                </CardHeader>
                <CardContent className='pl-2'>
                  <Overview />
                </CardContent>
              </Card>
              <Card className='col-span-1 lg:col-span-3'>
                <CardHeader>
                  <CardTitle>Last 5 Deployments</CardTitle>
                  <CardDescription>
                    You made 265 deployments this month.
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <RecentSales />
                </CardContent>
              </Card>
            </div>
          </TabsContent>
          <TabsContent value='analytics'>
            <Analytics />
          </TabsContent>
        </Tabs>
      </Layout.Body>
    </Layout>
  )
}
