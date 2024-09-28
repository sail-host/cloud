import { Navigate } from 'react-router-dom'
import axios from 'axios'
import useAuth from '@/hooks/use-auth'
import { useEffect } from 'react'
import { Loading } from '../custom/loading'

axios.interceptors.request.use((config) => {
  const token = localStorage.getItem('auth_token')
  config.headers.Authorization = `Bearer ${token}`

  return config
})

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode
}) {
  const { user, login, loading } = useAuth()

  useEffect(() => {
    login()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  if (loading) {
    return (
      <div className='z-50 flex items-center justify-center h-screen'>
        <Loading loading={true} />
      </div>
    )
  }

  if (!login && !user) {
    return <Navigate to='/login' />
  }

  return children
}
