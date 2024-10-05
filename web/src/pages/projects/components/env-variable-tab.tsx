import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { randomString } from '@/lib/utils'
import { useProjectSettingStore } from '@/store/project-create-store'
import { IconTrashX } from '@tabler/icons-react'
import { useState } from 'react'

export function EnvVariableTab() {
  const { environments, setEnvironments } = useProjectSettingStore()
  const [name, setName] = useState('')
  const [value, setValue] = useState('')

  const handleAdd = () => {
    if (name && value) {
      setEnvironments([...environments, { id: randomString(10), name, value }])
      setName('')
      setValue('')
    }
  }

  const handleDelete = (id: string) => {
    setEnvironments(environments.filter((env) => env.id !== id))
  }

  const handleChange = (id: string, name: string, value: string) => {
    setEnvironments(
      environments.map((env) =>
        env.id === id ? { ...env, [name]: value } : env
      )
    )
  }

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
                  <Input
                    placeholder='NEXT_PUBLIC_API_URL'
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                  />
                </div>
              </div>
              <div className='col-span-5'>
                <div className='flex flex-col gap-2'>
                  <Label>Value</Label>
                  <Input
                    placeholder='https://api.example.com'
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                  />
                </div>
              </div>
              <div className='col-span-1'>
                <Button variant='outline' type='button' onClick={handleAdd}>
                  Add
                </Button>
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

            {environments.length > 0 &&
              environments.map((env) => (
                <div className='grid grid-cols-11 gap-2 px-1 mt-2'>
                  <div className='col-span-5'>
                    <Input
                      placeholder='NEXT_PUBLIC_API_URL'
                      value={env.name}
                      onChange={(e) =>
                        handleChange(env.id, 'name', e.target.value)
                      }
                    />
                  </div>
                  <div className='col-span-5'>
                    <Input
                      placeholder='https://api.example.com'
                      value={env.value}
                      onChange={(e) =>
                        handleChange(env.id, 'value', e.target.value)
                      }
                    />
                  </div>
                  <div className='col-span-1'>
                    <Button
                      variant='outline'
                      className='w-full'
                      onClick={() => handleDelete(env.id)}
                    >
                      <IconTrashX className='w-4 h-4' />
                    </Button>
                  </div>
                </div>
              ))}

            {environments.length === 0 && (
              <div className='flex items-center justify-center p-2'>
                <p className='text-sm text-muted-foreground'>
                  No environment variables found
                </p>
              </div>
            )}
          </div>
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  )
}
