import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { Button } from '@/components/custom/button'
import { IconPlus } from '@tabler/icons-react'
import { Link } from 'react-router-dom'
import { ConfirmationModal } from '@/components/custom/confirmation-modal'
import { useState } from 'react'
import { toast } from 'sonner'
import { useDeleteModalStore } from '@/store/delete-modal-store'
import { DomainsTable } from './components/domains-table'

export default function Domains() {
  const { open, setOpen, deleteID } = useDeleteModalStore()
  const [loading, setLoading] = useState(false)

  const handleDelete = () => {
    setLoading(true)
    console.log('delete', deleteID)

    // TODO: Delete domain

    setTimeout(() => {
      setLoading(false)
      setOpen(false)

      toast.success('Domain deleted successfully')
    }, 2000)
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
        <div className='-mx-4 flex-1 overflow-auto px-4 py-1 lg:flex-row lg:space-x-12 lg:space-y-0'>
          <DomainsTable />
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
