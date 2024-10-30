import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { formSchema, GitAccountForm } from './components/git-account-form'
import { z } from 'zod'
import { useNavigate, useParams } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { GitAccount } from '@/types/model'
import axios from 'axios'
import { BaseResponse } from '@/types/base'
import { Loading } from '@/components/custom/loading'
import { toast } from 'sonner'

export default function EditGitAccount() {
  const { id } = useParams()
  const [gitAccount, setGitAccount] = useState<GitAccount | null>(null)
  const [loading, setLoading] = useState(false)
  const [loadingUpdate, setLoadingUpdate] = useState(false)
  const navigate = useNavigate()

  const fetchData = () => {
    setLoading(true)
    axios
      .get<BaseResponse<GitAccount>>(`/api/v1/git/show/${id}`)
      .then((res) => {
        setGitAccount(res.data.data)
      })
      .finally(() => setLoading(false))
  }

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    setLoadingUpdate(true)
    axios
      .put(`/api/v1/git/update/${id}`, {
        name: values.name,
        url: values.gitUrl,
        type: values.type,
        token: values.token,
      })
      .then(() => {
        toast.success('Git account updated successfully')
        navigate('/git-accounts')
      })
      .finally(() => setLoadingUpdate(false))
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
          {loading ? (
            <Loading loading />
          ) : (
            <>
              {gitAccount ? (
                <GitAccountForm
                  onSubmit={onSubmit}
                  isLoading={loadingUpdate}
                  defaultValues={{
                    name: gitAccount.name,
                    gitUrl: gitAccount.url,
                    type: gitAccount.type,
                    token: gitAccount.token,
                  }}
                />
              ) : (
                <div className='mt-10 text-center text-2xl'>
                  Git Account not found
                </div>
              )}
            </>
          )}
        </div>
      </Layout.Body>
    </Layout>
  )
}
