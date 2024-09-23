import { Card } from '@/components/ui/card'
import { OtpForm } from './components/otp-form'

export default function Otp() {
  const handleResendCode = () => {
    console.log('Resend code')
  }

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
                Two-factor Authentication
              </h1>
              <p className='text-sm text-muted-foreground'>
                Please enter the authentication code. <br /> We have sent the
                authentication code to your email.
              </p>
            </div>
            <OtpForm />
            <p className='mt-4 px-8 text-center text-sm text-muted-foreground'>
              Haven't received it?{' '}
              <button
                type='button'
                className='underline underline-offset-4 hover:text-primary'
                onClick={handleResendCode}
              >
                Resend a new code.
              </button>
              .
            </p>
          </Card>
        </div>
      </div>
    </>
  )
}
