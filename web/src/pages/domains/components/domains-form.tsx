import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'

import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

export const formSchema = z
  .object({
    name: z.string().min(2, {
      message: 'Name must be at least 2 characters.',
    }),
    dnsProvider: z.enum(['cloudflare', 'manual']),
    cloudflareAccountId: z.string().optional(),
    cloudflareApiToken: z.string().optional(),
  })
  .refine(
    (data) => {
      if (data.dnsProvider === 'cloudflare') {
        return !!data.cloudflareAccountId && !!data.cloudflareApiToken
      }
      return true
    },
    {
      message: 'Cloudflare Account ID and API Token are required.',
      path: ['cloudflareAccountId'],
    }
  )
  .refine(
    (data) => {
      if (data.dnsProvider === 'cloudflare') {
        return !!data.cloudflareAccountId && !!data.cloudflareApiToken
      }
      return true
    },
    {
      message: 'Cloudflare Account ID and API Token are required.',
      path: ['cloudflareApiToken'],
    }
  )

interface DomainsFormProps {
  defaultValues?: z.infer<typeof formSchema>
  onSubmit: (values: z.infer<typeof formSchema>) => void
}

export function DomainsForm({ defaultValues, onSubmit }: DomainsFormProps) {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: defaultValues ?? {
      name: '',
      dnsProvider: 'cloudflare',
      cloudflareAccountId: '',
      cloudflareApiToken: '',
    },
  })

  return (
    <Card>
      <CardHeader>
        <CardTitle>Domain</CardTitle>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-4'>
            <FormField
              control={form.control}
              name='name'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Domain Name</FormLabel>
                  <FormControl>
                    <Input placeholder='example.com' {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='dnsProvider'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>DNS Provider</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder='Select a DNS provider' />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value='cloudflare'>Cloudflare</SelectItem>
                      <SelectItem value='manual'>Manual</SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />

            {form.watch('dnsProvider') === 'cloudflare' && (
              <>
                <FormField
                  control={form.control}
                  name='cloudflareAccountId'
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Cloudflare Account ID</FormLabel>
                      <FormControl>
                        <Input
                          placeholder='Your Cloudflare Account ID'
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
                <FormField
                  control={form.control}
                  name='cloudflareApiToken'
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Cloudflare API Token</FormLabel>
                      <FormControl>
                        <Input
                          placeholder='Your Cloudflare API Token'
                          {...field}
                        />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </>
            )}
            <div className='space-x-3'>
              <Button type='submit'>Submit</Button>
              <Button type='button' variant='outline'>
                Test Connection
              </Button>
            </div>
          </form>
        </Form>
      </CardContent>
    </Card>
  )
}
