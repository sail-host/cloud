import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { ConfirmationModal } from '@/components/custom/confirmation-modal'
import { useState } from 'react'
import { toast } from 'sonner'
import { useDeleteModalStore } from '@/store/delete-modal-store'
import { ProjectsFilter } from './components/projects-filter'
import { ProjectsGridTable } from './components/projects-grid-table'
import { projects } from './data/projects'

export default function Projects() {
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
            <h2 className='text-2xl font-bold tracking-tight'>Projects</h2>
            <p className='text-muted-foreground'>
              Here&apos;s a list of your projects!
            </p>
          </div>
        </div>
        <ProjectsFilter />
        <ProjectsGridTable projects={projects} />
        <ConfirmationModal
          open={open}
          setOpen={setOpen}
          confirmFunction={handleDelete}
          title='Delete Project'
          description='Are you sure you want to delete this project?'
          loading={loading}
        />
      </Layout.Body>
    </Layout>
  )
}
