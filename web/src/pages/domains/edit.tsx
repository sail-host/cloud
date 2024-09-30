import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { z } from 'zod'
import { useNavigate, useParams } from 'react-router-dom'
import { DomainsForm, formSchema } from './components/domains-form'
import { useEffect, useState } from 'react'
import { Domain } from '@/types/model'
import axios from 'axios'
import { BaseResponse } from '@/types/base'
import { Loading } from '@/components/custom/loading'
import { toast } from 'sonner'

export default function EditDomain() {
  const [isLoading, setIsLoading] = useState(false)
  const [dataLoading, setDataLoading] = useState(true)
  const [domain, setDomain] = useState<Domain | null>(null)
  const { id } = useParams()
  const navigate = useNavigate()

  const fetchData = () => {
    setDataLoading(true)
    axios
      .get<BaseResponse<Domain>>(`/api/v1/domain/show/${id}`)
      .then((res) => {
        setDomain(res.data.data)
      })
      .finally(() => setDataLoading(false))
  }

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    setIsLoading(true)
    axios
      .put<BaseResponse>(`/api/v1/domain/update/${id}`, {
        domain: values.name,
        dns_provider: values.dnsProvider,
        cloudflare_zone_id: values.cloudflareZoneId,
        cloudflare_api_key: values.cloudflareApiKey,
      })
      .then((res) => {
        if (res.data.status === 'success') {
          toast.success('Domain updated successfully')
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

  useEffect(() => {
    fetchData()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

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
            <h2 className='text-2xl font-bold tracking-tight'>Edit Domain</h2>
            <p className='text-muted-foreground'>Edit a domain!</p>
          </div>
        </div>
        <div className='flex-1 px-4 py-1 -mx-4 overflow-auto lg:flex-row lg:space-x-12 lg:space-y-0'>
          {dataLoading ? (
            <Loading loading />
          ) : (
            <>
              {domain ? (
                <DomainsForm
                  onSubmit={onSubmit}
                  defaultValues={{
                    name: domain.domain,
                    dnsProvider: domain.dns_provider,
                    cloudflareZoneId: domain.cloudflare_zone_id,
                    cloudflareApiKey: domain.cloudflare_api_key,
                  }}
                  isLoading={isLoading}
                />
              ) : (
                <div className='mt-10 text-2xl text-center'>
                  Domain not found
                </div>
              )}
            </>
          )}
        </div>
      </Layout.Body>
    </Layout>
  )
}
