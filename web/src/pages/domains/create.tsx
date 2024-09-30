import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { z } from 'zod'
import { DomainsForm, formSchema } from './components/domains-form'
import { useState } from 'react'
import axios from 'axios'
import { toast } from 'sonner'
import { BaseResponse } from '@/types/base'
import { useNavigate } from 'react-router-dom'

export default function CreateDomain() {
  const [isLoading, setIsLoading] = useState(false)
  const navigate = useNavigate()

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    setIsLoading(true)
    axios
      .post<BaseResponse>('/api/v1/domain/create', {
        domain: values.name,
        dns_provider: values.dnsProvider,
        cloudflare_zone_id: values.cloudflareZoneId,
        cloudflare_api_key: values.cloudflareApiKey,
      })
      .then((res) => {
        if (res.data.status === 'success') {
          toast.success('Domain created successfully')
        } else {
          toast.error(res.data.message)
        }
        navigate('/domains')
      })
      .catch((err) => {
        toast.error(err.response?.data?.message || 'Something went wrong')
      })
      .finally(() => setIsLoading(false))
  }

  return (
    <Layout>
      {/* ===== Top Heading ===== */}
      <Layout.Header sticky>
        <Search />
        <div className='flex items-center ml-auto space-x-4'>
          <ThemeSwitch />
          <UserNav />
        </div>
      </Layout.Header>

      <Layout.Body>
        <div className='flex items-center justify-between mb-2 space-y-2'>
          <div>
            <h2 className='text-2xl font-bold tracking-tight'>Create Domain</h2>
            <p className='text-muted-foreground'>Create a new domain!</p>
          </div>
        </div>
        <div className='flex-1 px-4 py-1 -mx-4 overflow-auto lg:flex-row lg:space-x-12 lg:space-y-0'>
          <DomainsForm onSubmit={onSubmit} isLoading={isLoading} />
        </div>
      </Layout.Body>
    </Layout>
  )
}
