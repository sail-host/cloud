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
import { useState } from 'react'
import axios from 'axios'
import { BaseResponse } from '@/types/base'
import { toast } from 'sonner'

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
  const [isCheckLoading, setIsCheckLoading] = useState(false)
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: defaultValues ?? {
      name: '',
      gitUrl: '',
      type: 'github',
      token: '',
    },
  })

  const handleCheck = () => {
    form.clearErrors()
    if (!form.getValues('gitUrl')) {
      form.setError('gitUrl', {
        message: 'Git URL is required.',
      })
      return
    }
    if (!form.getValues('token')) {
      form.setError('token', {
        message: 'Git Token is required.',
      })
      return
    }
    if (!form.getValues('type')) {
      form.setError('type', {
        message: 'Git Type is required.',
      })
      return
    }
    if (!form.getValues('name')) {
      form.setError('name', {
        message: 'Git Name is required.',
      })
      return
    }

    setIsCheckLoading(true)
    axios
      .post<BaseResponse>('/api/v1/git/check-account', {
        url: form.getValues('gitUrl'),
        token: form.getValues('token'),
        type: form.getValues('type'),
        name: form.getValues('name'),
      })
      .then((res) => {
        if (res.data.status === 'success') {
          toast.success(res.data.message)
        } else {
          toast.error(res.data.message)
        }
      })
      .catch((err) => {
        toast.error(err.response?.data?.message || 'Failed to check account')
      })
      .finally(() => setIsCheckLoading(false))
  }

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
              <Button
                type='button'
                variant='outline'
                onClick={handleCheck}
                loading={isCheckLoading}
              >
                Test Account
              </Button>
            </div>
          </form>
        </Form>
      </CardContent>
    </Card>
  )
}
