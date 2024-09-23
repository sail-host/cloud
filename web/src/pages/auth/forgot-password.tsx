import { Card } from '@/components/ui/card'
import { ForgotForm } from './components/forgot-form'
import { Link } from 'react-router-dom'

export default function ForgotPassword() {
  // TODO: Implement register button
  const isRegister = true

  return (
    <>
      <div className='container grid h-svh flex-col items-center justify-center bg-background lg:max-w-none lg:px-0'>
        <div className='mx-auto flex w-full flex-col justify-center space-y-2 sm:w-[480px] lg:p-8'>
          <div className='mb-4 flex items-center justify-center'>
            <img src='/images/logo.svg' alt='SailHost' className='h-6' />
          </div>
          <Card className='p-6'>
            <div className='mb-2 flex flex-col space-y-2 text-left'>
              <h1 className='text-md font-semibold tracking-tight'>
                Forgot Password
              </h1>
              <p className='text-sm text-muted-foreground'>
                Enter your registered email and <br /> we will send you a link
                to reset your password.
              </p>
            </div>
            <ForgotForm />
            {isRegister && (
              <p className='mt-4 px-8 text-center text-sm text-muted-foreground'>
                Don't have an account?{' '}
                <Link
                  to='/register'
                  className='underline underline-offset-4 hover:text-primary'
                >
                  Register
                </Link>
                .
              </p>
            )}
          </Card>
        </div>
      </div>
    </>
  )
}
