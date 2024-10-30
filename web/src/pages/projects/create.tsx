import { Layout } from '@/components/custom/layout'
import { Search } from '@/components/search'
import ThemeSwitch from '@/components/theme-switch'
import { UserNav } from '@/components/user-nav'
import { CreateNewProject } from './components/create-new-project'
import { ConfigureNewProject } from './components/configure-new-project'
import {
  useProjectCreateStore,
  useProjectStore,
} from '@/store/project-create-store'
import { useEffect } from 'react'

export default function CreateProject() {
  const { step, setStep, setGitAccount } = useProjectCreateStore()
  const { setProject } = useProjectStore()

  useEffect(() => {
    setStep('1')
    setGitAccount(null)
    setProject(null)
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
              Create Project
            </h2>
            <p className='text-muted-foreground'>Create a new project!</p>
          </div>
        </div>
        <div className='-mx-4 flex-1 overflow-auto px-4 py-1 lg:flex-row lg:space-x-12 lg:space-y-0'>
          {step === '1' ? <CreateNewProject /> : <ConfigureNewProject />}
        </div>
      </Layout.Body>
    </Layout>
  )
}
