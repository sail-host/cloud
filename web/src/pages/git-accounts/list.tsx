import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { GitAccountsTable } from './components/git-accounts-table'
import { Button } from '@/components/custom/button'
import { IconPlus } from '@tabler/icons-react'
import { Link } from 'react-router-dom'
import { ConfirmationModal } from '@/components/custom/confirmation-modal'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'
import { useDeleteModalStore } from '@/store/delete-modal-store'
import { GitAccount } from '@/types/model'
import axios from 'axios'
import { BaseResponse } from '@/types/base'

export default function GitAccounts() {
  const { open, setOpen, deleteID } = useDeleteModalStore()
  const [loading, setLoading] = useState(false)
  const [loadingDelete, setLoadingDelete] = useState(false)
  const [data, setData] = useState<GitAccount[]>([])

  const fetchData = () => {
    setLoading(true)
    axios
      .get<BaseResponse<GitAccount[]>>('/api/v1/git/list')
      .then((res) => {
        setData(res.data.data ?? [])
      })
      .finally(() => setLoading(false))
  }

  const handleDelete = () => {
    setLoadingDelete(true)

    axios
      .delete(`/api/v1/git/delete/${deleteID}`)
      .then(() => {
        toast.success('Git account deleted successfully')
        fetchData()
      })
      .finally(() => {
        setLoadingDelete(false)
        setOpen(false)
      })
  }

  useEffect(() => {
    fetchData()
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
            <h2 className='text-2xl font-bold tracking-tight'>Git Accounts</h2>
            <p className='text-muted-foreground'>
              Here&apos;s a list of your git accounts!
            </p>
          </div>
          <Button asChild>
            <Link to='/git-accounts/create'>
              <IconPlus className='mr-2' />
              Add Git Account
            </Link>
          </Button>
        </div>
        <div className='flex-1 px-4 py-1 -mx-4 overflow-auto lg:flex-row lg:space-x-12 lg:space-y-0'>
          <GitAccountsTable gitAccounts={data} loading={loading} />
        </div>
        <ConfirmationModal
          open={open}
          setOpen={setOpen}
          confirmFunction={handleDelete}
          title='Delete Git Account'
          description='Are you sure you want to delete this git account?'
          loading={loadingDelete}
        />
      </Layout.Body>
    </Layout>
  )
}
