import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { IconTrashX } from '@tabler/icons-react'
import { Button } from '@/components/ui/button'
import { Label } from '@/components/ui/label'

export function EnvironmentsTab() {
  return (
    <div className='w-full'>
      <Card>
        <CardHeader>
          <CardTitle>Environments</CardTitle>
          <CardDescription>
            In order to provide your Deployment with Environment Variables at
            Build and Runtime, you may enter them right here, for the
            Environment of your choice.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className=''>
            <div className='grid items-end w-full grid-cols-11 gap-2 pr-1'>
              <div className='col-span-5'>
                <div className='flex flex-col gap-2'>
                  <Label>Name</Label>
                  <Input placeholder='NEXT_PUBLIC_API_URL' />
                </div>
              </div>
              <div className='col-span-5'>
                <div className='flex flex-col gap-2'>
                  <Label>Value</Label>
                  <Input placeholder='https://api.example.com' />
                </div>
              </div>
              <div className='col-span-1'>
                <Button variant='outline'>Add</Button>
              </div>
            </div>

            <div className='mt-2'>
              <p className='text-sm text-muted-foreground'>
                <span className='font-medium'>TIP:</span> Paste a .env above to
                populate the form
              </p>
            </div>

            <div className='grid grid-cols-11 gap-2 p-2 mt-3 text-sm border rounded-md text-muted-foreground'>
              <div className='col-span-5'>Key</div>
              <div className='col-span-5'>Value</div>
            </div>

            <div className='grid grid-cols-11 gap-2 px-1 mt-2'>
              <div className='col-span-5'>
                <Input placeholder='NEXT_PUBLIC_API_URL' />
              </div>
              <div className='col-span-5'>
                <Input placeholder='https://api.example.com' />
              </div>
              <div className='col-span-1'>
                <Button variant='outline' className='w-full'>
                  <IconTrashX className='w-4 h-4' />
                </Button>
              </div>
            </div>
          </div>
        </CardContent>

        <CardFooter className='flex justify-end p-3 pr-6 border-t rounded-b-xl bg-muted dark:bg-muted/40'>
          <Button>Save</Button>
        </CardFooter>
      </Card>
    </div>
  )
}
