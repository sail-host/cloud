import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { IconTrashX } from '@tabler/icons-react'

export function EnvVariableTab() {
  return (
    <Accordion type='single' collapsible className='px-3 border rounded-md'>
      <AccordionItem value='item-1'>
        <AccordionTrigger>Environment Variables</AccordionTrigger>
        <AccordionContent>
          <div className='mt-3'>
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
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  )
}
