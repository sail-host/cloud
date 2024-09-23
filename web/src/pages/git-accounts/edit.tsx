import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { formSchema, GitAccountForm } from './components/git-account-form'
import { z } from 'zod'
import { useParams } from 'react-router-dom'
import { gitAccounts } from './data/gitAccounts'

export default function EditGitAccount() {
  const { id } = useParams()
  const gitAccount = gitAccounts.find((account) => account.id === id)

  if (!gitAccount) {
    return <div>Git account not found</div>
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
            <h2 className='text-2xl font-bold tracking-tight'>
              Edit Git Account
            </h2>
            <p className='text-muted-foreground'>Edit a git account!</p>
          </div>
        </div>
        <div className='-mx-4 flex-1 overflow-auto px-4 py-1 lg:flex-row lg:space-x-12 lg:space-y-0'>
          <GitAccountForm
            onSubmit={onSubmit}
            defaultValues={{
              name: gitAccount.name,
              gitUrl: gitAccount.gitUrl,
              type: gitAccount.type,
            }}
          />
        </div>
      </Layout.Body>
    </Layout>
  )
}
