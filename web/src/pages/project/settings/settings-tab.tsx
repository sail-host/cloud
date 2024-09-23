import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from '@/components/ui/card'
import { TabsContent } from '@/components/ui/tabs'
import { useState } from 'react'
import { SettingsTabSidebar } from './components/sidebar'
import { SettingsTab as SettingsTabComponent } from './components/settings-tab'

export function SettingsTab() {
  const [activeTab, setActiveTab] = useState('general')

  return (
    <TabsContent value='settings' className='space-y-4'>
      <Card>
        <CardHeader className='flex flex-row justify-between space-y-0'>
          <div>
            <CardTitle className='mb-1 text-xl'>Settings</CardTitle>
            <CardDescription>Configure your project settings.</CardDescription>
          </div>
        </CardHeader>
        <CardContent className='pt-4 border-t'>
          <div className='flex flex-col flex-1 space-y-4 md:space-y-2 md:overflow-hidden lg:flex-row lg:space-x-8 lg:space-y-0'>
            <aside className='top-0 lg:sticky lg:w-1/5'>
              <SettingsTabSidebar
                activeTab={activeTab}
                setActiveTab={setActiveTab}
              />
            </aside>
            <div className='flex w-full p-1 md:overflow-y-hidden'>
              <SettingsTabComponent activeTab={activeTab} />
            </div>
          </div>
        </CardContent>
      </Card>
    </TabsContent>
  )
}
