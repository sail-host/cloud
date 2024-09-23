import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { z } from 'zod'
import { useParams } from 'react-router-dom'
import { domains } from './data/domains'
import { DomainsForm, formSchema } from './components/domains-form'

export default function EditDomain() {
  const { id } = useParams()
  const domain = domains.find((domain) => domain.id === Number(id))

  if (!domain) {
    return <div>Domain not found</div>
  }

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    console.log(values)
  }

  return (
    <Layout>
      {/* ===== Top Heading ===== */}
      <Layout.Header sticky>
        <Search />
        <div className='ml-auto flex items-center space-x-4'>
          <ThemeSwitch />
          <UserNav />
        </div>
      </Layout.Header>

      <Layout.Body>
        <div className='mb-2 flex items-center justify-between space-y-2'>
          <div>
            <h2 className='text-2xl font-bold tracking-tight'>Edit Domain</h2>
            <p className='text-muted-foreground'>Edit a domain!</p>
          </div>
        </div>
        <div className='-mx-4 flex-1 overflow-auto px-4 py-1 lg:flex-row lg:space-x-12 lg:space-y-0'>
          <DomainsForm
            onSubmit={onSubmit}
            defaultValues={{
              name: domain.name,
              dnsProvider: domain.dnsProvider,
              cloudflareAccountId: domain.cloudflareAccountId,
              cloudflareApiToken: domain.cloudflareApiToken,
            }}
          />
        </div>
      </Layout.Body>
    </Layout>
  )
}
