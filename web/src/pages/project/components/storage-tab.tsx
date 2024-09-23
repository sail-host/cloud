import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from '@/components/ui/card'
import { TabsContent } from '@/components/ui/tabs'

export function StorageTab() {
  return (
    <TabsContent value='storage' className='space-y-4'>
      <Card>
        <CardHeader className='flex flex-row justify-between space-y-0'>
          <div>
            <CardTitle className='mb-1 text-xl'>Storage</CardTitle>
            <CardDescription>
              Create databases and stores that you can connect to your project.
            </CardDescription>
          </div>
        </CardHeader>
        <CardContent className=''>storage</CardContent>
      </Card>
    </TabsContent>
  )
}
