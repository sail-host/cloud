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

// eslint-disable-next-line react-refresh/only-export-components
export const formSchema = z.object({
  name: z.string().min(2, {
    message: 'Name must be at least 2 characters.',
  }),
  gitUrl: z.string().url({
    message: 'Git URL must be a valid URL.',
  }),
  type: z.enum(['github', 'gitlab', 'bitbucket', 'gitea']),
  token: z.string().min(1, {
    message: 'Token must be at least 1 character.',
  }),
})

interface GitAccountFormProps {
  defaultValues?: z.infer<typeof formSchema>
  onSubmit: (values: z.infer<typeof formSchema>) => void
  isLoading?: boolean
}

export function GitAccountForm({
  defaultValues,
  onSubmit,
  isLoading,
}: GitAccountFormProps) {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: defaultValues ?? {
      name: '',
      gitUrl: '',
      type: 'github',
      token: '',
    },
  })

  return (
    <Card>
      <CardHeader>
        <CardTitle>Git Account</CardTitle>
      </CardHeader>
      <CardContent>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-4'>
            <FormField
              control={form.control}
              name='name'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input placeholder='John Doe' {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='gitUrl'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Git URL</FormLabel>
                  <FormControl>
                    <Input
                      placeholder='https://github.com/JohnDoe123'
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='type'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Git Type</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder='Select a git type' />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value='github'>GitHub</SelectItem>
                      <SelectItem value='gitlab'>GitLab</SelectItem>
                      <SelectItem value='bitbucket'>Bitbucket</SelectItem>
                    </SelectContent>
                  </Select>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name='token'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Git Token</FormLabel>
                  <FormControl>
                    <Input placeholder='1234567890' {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div className='space-x-3'>
              <Button type='submit' loading={isLoading}>
                Submit
              </Button>
              <Button type='button' variant='outline'>
                Test Account
              </Button>
            </div>
          </form>
        </Form>
      </CardContent>
    </Card>
  )
}
