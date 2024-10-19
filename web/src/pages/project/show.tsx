import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { ProjectTab } from './components/project-tab'
import { DeploymentsTab } from './components/depoloyments-tab'
import { LogsTab } from './components/logs-tab'
import { StorageTab } from './components/storage-tab'
import { SettingsTab } from './settings/settings-tab'
import { useEffect, useState } from 'react'
import axios from 'axios'
import { useParams } from 'react-router-dom'
import { Loading } from '@/components/custom/loading'
import NotFoundError from '../errors/not-found-error'
import { BaseResponse } from '@/types/base'
import { Project, useProjectStore } from '@/store/project-store'

export default function ProjectShow() {
  const { uuid } = useParams()
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(false)
  const { project, setProject } = useProjectStore()

  const fetchProject = () => {
    setLoading(true)
    axios
      .get<BaseResponse<Project>>(`/api/v1/project/show/${uuid}`)
      .then((res) => {
        if (res.data.status === 'success') {
          setProject(res.data.data)
        } else {
          setError(true)
        }
      })
      .catch((err) => {
        setError(true)
        console.error(err)
      })
      .finally(() => {
        setLoading(false)
      })
  }

  useEffect(() => {
    fetchProject()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return (
    <Layout>
      {!error && (
        <Layout.Header sticky>
          <Search />
          <div className='flex items-center ml-auto space-x-4'>
            <ThemeSwitch />
            <UserNav />
          </div>
        </Layout.Header>
      )}

      <Layout.Body>
        {loading ? (
          <div className='flex h-[80svh] items-center justify-center'>
            <Loading loading />
          </div>
        ) : error ? (
          <NotFoundError />
        ) : (
          <>
            <div className='flex items-center justify-between mb-2 space-y-2'>
              <h1 className='text-2xl font-bold tracking-tight'>
                {project?.name}
              </h1>
            </div>
            <div className='mt-4'>
              <Tabs
                orientation='vertical'
                defaultValue='project'
                className='space-y-4'
              >
                <div className='w-full pb-2 overflow-x-auto'>
                  <TabsList>
                    <TabsTrigger value='project'>Project</TabsTrigger>
                    <TabsTrigger value='deployments'>Deployments</TabsTrigger>
                    <TabsTrigger value='logs'>Logs</TabsTrigger>
                    <TabsTrigger value='storage'>Storage</TabsTrigger>
                    <TabsTrigger value='settings'>Settings</TabsTrigger>
                  </TabsList>
                </div>

                <ProjectTab />
                <TabsContent value='deployments' className='space-y-4'>
                  <DeploymentsTab uuid={uuid} />
                </TabsContent>
                <LogsTab />
                <StorageTab />
                <SettingsTab />
              </Tabs>
            </div>
          </>
        )}
      </Layout.Body>
    </Layout>
  )
}
