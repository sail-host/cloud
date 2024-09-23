import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { TabsContent } from '@/components/ui/tabs'

export function LogsTab() {
  return (
    <TabsContent value='logs' className='space-y-4'>
      <Card>
        <CardHeader className='flex flex-row justify-between space-y-0'>
          <div>
            <CardTitle className='mb-1 text-xl'>Logs</CardTitle>
            <CardDescription>Build and runtime logs.</CardDescription>
          </div>
        </CardHeader>
        <CardContent className=''>logs</CardContent>
      </Card>
    </TabsContent>
  )
}
