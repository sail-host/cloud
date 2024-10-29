import { useEffect, useState } from 'react'
import ContentSection from '../components/content-section'
import { Loading } from '@/components/custom/loading'
import { Button } from '@/components/ui/button'

export default function SettingsAccount() {
  const [isLoading, setIsLoading] = useState(false)
  const [appVersion, setAppVersion] = useState('')
  const [lastVersion, setLastVersion] = useState('')

  const checkVersion = () => {
    setIsLoading(true)

    // Fake timeout
    setTimeout(() => {
      setAppVersion('1.0.0')
      setLastVersion('1.0.1')
      setIsLoading(false)
    }, 2000)
  }

  useEffect(() => {
    checkVersion()
  }, [])

  return (
    <ContentSection
      title='Upgrade'
      desc='Upgrade your account to get more features.'
    >
      <div>
        <div className='flex flex-col gap-4'>
          {/* Loading State */}
          {isLoading ? (
            <div className='flex items-center justify-center p-8'>
              <Loading loading />
            </div>
          ) : (
            <div className=''>
              <div className='flex flex-col gap-6'>
                <div className='flex items-center gap-10'>
                  <div className='space-y-2'>
                    <span className='text-lg font-medium text-muted-foreground'>
                      Current Version
                    </span>
                    <div className='flex items-center gap-2'>
                      <span className='text-xl font-semibold'>
                        {appVersion}
                      </span>
                      <div className='rounded-full bg-green-100 px-2 py-0.5 text-xs font-medium text-green-700'>
                        Installed
                      </div>
                    </div>
                  </div>

                  <div className='space-y-2'>
                    <span className='text-lg font-medium text-muted-foreground'>
                      Latest Version
                    </span>
                    <div className='flex items-center gap-2'>
                      <span className='text-xl font-semibold'>
                        {lastVersion}
                      </span>
                      {appVersion !== lastVersion && (
                        <div className='rounded-full bg-yellow-100 px-2 py-0.5 text-xs font-medium text-yellow-700'>
                          Update Available
                        </div>
                      )}
                    </div>
                  </div>
                </div>
                <div className='inline-block'>
                  {appVersion !== lastVersion && <Button>Upgrade Now</Button>}
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </ContentSection>
  )
}
