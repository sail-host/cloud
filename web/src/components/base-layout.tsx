import { Outlet } from 'react-router-dom'
import AuthRedirectProvider from './providers/auth-redirect-provider'

export default function BaseLayout() {
  return (
    <>
      <AuthRedirectProvider>
        <Outlet />
      </AuthRedirectProvider>
    </>
  )
}
