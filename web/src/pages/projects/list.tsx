import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { ConfirmationModal } from '@/components/custom/confirmation-modal'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'
import { useDeleteModalStore } from '@/store/delete-modal-store'
import { ProjectsFilter } from './components/projects-filter'
import { ProjectsGridTable } from './components/projects-grid-table'
import axios from 'axios'
import { Loading } from '@/components/custom/loading'
import { BaseResponse } from '@/types/base'

export interface Project {
  id: number
  name: string
  domain: string
  git_hash: string
  git_date: string
  git_branch: string
  git_commit: string
}

export default function Projects() {
  const { open, setOpen, deleteID } = useDeleteModalStore()
  const [loading, setLoading] = useState(false)
  const [projects, setProjects] = useState<Project[]>([])
  const [dataLoading, setDataLoading] = useState(true)

  const fetchProjects = () => {
    setDataLoading(true)
    axios
      .get<BaseResponse<Project[]>>('/api/v1/project/list')
      .then((res) => {
        setProjects(res.data.data)
      })
      .catch(() => {
        toast.error('Error fetching projects')
      })
      .finally(() => {
        setDataLoading(false)
      })
  }

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

  useEffect(() => {
    fetchProjects()
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
            <h2 className='text-2xl font-bold tracking-tight'>Projects</h2>
            <p className='text-muted-foreground'>
              Here&apos;s a list of your projects!
            </p>
          </div>
        </div>
        <ProjectsFilter />
        {dataLoading ? (
          <div className='flex items-center justify-center my-10'>
            <Loading loading />
          </div>
        ) : (
          <ProjectsGridTable projects={projects} />
        )}
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
