import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { Button } from '@/components/custom/button'
import { IconPlus } from '@tabler/icons-react'
import { Link } from 'react-router-dom'
import { ConfirmationModal } from '@/components/custom/confirmation-modal'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'
import { useDeleteModalStore } from '@/store/delete-modal-store'
import { DomainsTable } from './components/domains-table'
import axios from 'axios'
import { BaseResponse } from '@/types/base'
import { Domain } from '@/types/model'

export default function Domains() {
  const { open, setOpen, deleteID } = useDeleteModalStore()
  const [loading, setLoading] = useState(false)
  const [data, setData] = useState<Domain[]>([])
  const [dataLoading, setDataLoading] = useState(true)

  const fetchData = () => {
    setDataLoading(true)
    axios
      .get<BaseResponse<Domain[]>>('/api/v1/domain/list')
      .then((res) => {
        setData(res.data.data ?? [])
      })
      .finally(() => setDataLoading(false))
  }

  const handleDelete = () => {
    setLoading(true)
    axios
      .delete(`/api/v1/domain/delete/${deleteID}`)
      .then(() => {
        toast.success('Domain deleted successfully')
        setOpen(false)
        fetchData()
      })
      .finally(() => setLoading(false))
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
            <h2 className='text-2xl font-bold tracking-tight'>Domains</h2>
            <p className='text-muted-foreground'>
              Here&apos;s a list of your domains!
            </p>
          </div>
          <Button asChild>
            <Link to='/domains/create'>
              <IconPlus className='mr-2' />
              Add Domain
            </Link>
          </Button>
        </div>
        <div className='flex-1 px-4 py-1 -mx-4 overflow-auto lg:flex-row lg:space-x-12 lg:space-y-0'>
          <DomainsTable domains={data} loading={dataLoading} />
        </div>
        <ConfirmationModal
          open={open}
          setOpen={setOpen}
          confirmFunction={handleDelete}
          title='Delete Domain'
          description='Are you sure you want to delete this domain?'
          loading={loading}
        />
      </Layout.Body>
    </Layout>
  )
}
