import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogTitle,
  DialogHeader,
  DialogContent,
  DialogFooter,
} from '@/components/ui/dialog'
import {
  Form,
  FormDescription,
  FormField,
  FormItem,
  FormControl,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { useProjectStore } from '@/store/project-store'
import { Domain } from '@/types/model'
import { zodResolver } from '@hookform/resolvers/zod'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

const FormSchema = z.object({
  domain: z.string({
    required_error: 'Please select a domain.',
  }),
  name: z
    .string({
      required_error: 'Please enter a name.',
    })
    .regex(
      /^[@\-_a-z0-9]+$/,
      'Only @, -, _,0-9 and lowercase letters are allowed.'
    ),
})

interface AddNewDomainResponse {
  ip: string
  domain: string
  type: string
  full_domain: string
}

export function NewDomainModal({
  fetchDomains: update,
}: {
  fetchDomains: () => void
}) {
  const [isOpen, setIsOpen] = useState(false)
  const [domains, setDomains] = useState<Domain[]>([])
  const [loading, setLoading] = useState(false)
  const { project } = useProjectStore()
  const [isAdding, setIsAdding] = useState(false)
  const [data, setData] = useState<AddNewDomainResponse | null>(null)

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  })

  const fetchDomains = () => {
    axios.get('/api/v1/domain/list').then((res) => {
      setDomains(res.data.data)
    })
  }

  function onSubmit(data: z.infer<typeof FormSchema>) {
    setLoading(true)
    axios
      .post(`/api/v1/project-setting/add-domain/${project?.name}`, {
        domain: data.name,
        domain_id: parseInt(data.domain),
      })
      .then((res) => {
        if (res.data.status === 'success') {
          setIsOpen(false)
          toast.success(res.data.message)
          if (res.data.data) {
            setData(res.data.data)
            setIsAdding(true)
          }
          update()
        } else {
          toast.error(res.data?.message || 'Something went wrong')
        }
      })
      .catch((err) => {
        toast.error(err.response?.data?.message || 'Something went wrong')
      })
      .finally(() => {
        setLoading(false)
      })
  }

  useEffect(() => {
    fetchDomains()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return (
    <>
      <Button onClick={() => setIsOpen(true)} className='absolute right-5'>
        Add Domain
      </Button>
      <Dialog open={isOpen} onOpenChange={setIsOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Add Domain</DialogTitle>
          </DialogHeader>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-4'>
              <FormField
                control={form.control}
                name='domain'
                render={({ field }) => (
                  <FormItem>
                    <Label>Domain</Label>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue placeholder='Select a domain' />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        {domains.map((domain) => (
                          <SelectItem
                            key={domain.id}
                            value={domain.id.toString()}
                          >
                            {domain.domain}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name='name'
                render={({ field }) => (
                  <FormItem>
                    <Label>Name</Label>
                    <Input {...field} placeholder='Subdomain name' />
                    <FormDescription>
                      For main domain, use <span className='font-bold'>@</span>
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <DialogFooter>
                <Button type='submit' loading={loading}>
                  Add
                </Button>
                <Button
                  variant='outline'
                  onClick={() => setIsOpen(false)}
                  disabled={loading}
                  type='button'
                >
                  Cancel
                </Button>
              </DialogFooter>
            </form>
          </Form>
        </DialogContent>
      </Dialog>

      <Dialog open={isAdding} onOpenChange={setIsAdding}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Domain Added</DialogTitle>
          </DialogHeader>
          <div className='flex flex-col items-center justify-center'>
            <div className=''>
              <p className='text-sm font-semibold'>Manual DNS Configuration:</p>
              <p className='text-sm'>1. Add an A record for {data?.domain}</p>
              <p className='text-sm'>2. Point it to {data?.ip}</p>
            </div>
            <p className='mt-3 text-lg font-bold'>{data?.full_domain}</p>
            <p className='text-sm text-gray-500'>IP: {data?.ip}</p>
            <p className='text-sm text-gray-500'>Type: {data?.type}</p>
          </div>
        </DialogContent>
      </Dialog>
    </>
  )
}
