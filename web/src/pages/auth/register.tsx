import { Card } from '@/components/ui/card'
import { SignUpForm } from './components/sign-up-form'
import { Link } from 'react-router-dom'
import { useState, useEffect } from 'react'
import { Loading } from '@/components/custom/loading'
import axios from 'axios'
import { BaseResponse } from '@/types/base'

export default function Register() {
  const [loading, setLoading] = useState(true)
  const [show, setShow] = useState(false)

  const handleCheckUser = () => {
    axios
      .get<BaseResponse>('/api/v1/auth/check-user-first-time')
      .then((res) => {
        setShow(res.data.status === 'success')
      })
      .finally(() => {
        setLoading(false)
      })
  }

  useEffect(() => {
    handleCheckUser()
  }, [])

  return (
    <>
      <div className='container grid flex-col items-center justify-center h-svh bg-background lg:max-w-none lg:px-0'>
        <div className='mx-auto flex w-full flex-col justify-center space-y-2 sm:w-[480px] lg:p-8'>
          <div className='flex items-center justify-center mb-4'>
            <img src='/images/logo.svg' alt='SailHost' className='h-6' />
          </div>
          <Card className='p-6'>
            <div className='flex flex-col mb-2 space-y-2 text-left'>
              <h1 className='text-lg font-semibold tracking-tight'>
                Create an account
              </h1>
              <p className='text-sm text-muted-foreground'>
                Enter your email and password to create your account. You can
                only create one account. <br />
                Already have an account?
                <Link
                  to='/login'
                  className='ml-1 underline underline-offset-4 hover:text-primary'
                >
                  Login
                </Link>
              </p>
            </div>

            {loading ? (
              <div className='flex items-center justify-center my-10'>
                <Loading loading={loading} />
              </div>
            ) : show ? (
              <SignUpForm />
            ) : (
              <div className='text-center'>Your account is already created</div>
            )}
          </Card>
        </div>
      </div>
    </>
  )
}
