import { useEffect, useState } from 'react'
import ContentSection from '../components/content-section'
import { Loading } from '@/components/custom/loading'
import { Button } from '@/components/ui/button'
import axios from 'axios'
import { BaseResponse } from '@/types/base'
import { toast } from 'sonner'

export default function SettingsAccount() {
  const [isLoading, setIsLoading] = useState(false)
  const [appVersion, setAppVersion] = useState('')
  const [lastVersion, setLastVersion] = useState('')
  const [isUpdating, setIsUpdating] = useState(false)

  const handleUpdate = () => {
    setIsUpdating(true)

    axios
      .post<BaseResponse>('/api/v1/upgrade/update')
      .then((res) => {
        toast.success(
          res.data?.message || 'Update success, please reload the page'
        )
      })
      .catch((err) => {
        toast.error(err.response?.data?.message || 'Failed to update')
      })
      .finally(() => {
        setIsUpdating(false)
      })
  }

  const checkVersion = () => {
    setIsLoading(true)

    axios
      .get<
        BaseResponse<{
          current_version: string
          last_version: string
        }>
      >('/api/v1/upgrade/check')
      .then((res) => {
        setAppVersion(res.data.data.current_version)
        setLastVersion(res.data.data.last_version)
      })
      .catch((err) => {
        console.error(err)
      })
      .finally(() => {
        setIsLoading(false)
      })
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
                  {appVersion !== lastVersion && (
                    <Button
                      type='button'
                      onClick={handleUpdate}
                      loading={isUpdating}
                    >
                      Upgrade Now
                    </Button>
                  )}
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </ContentSection>
  )
}
