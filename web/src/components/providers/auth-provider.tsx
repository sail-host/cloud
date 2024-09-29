import { useNavigate } from 'react-router-dom'
import axios from 'axios'
import { useEffect, useState } from 'react'
import { Loading } from '../custom/loading'
import { useUserStore } from '@/store/user-store'

axios.interceptors.request.use((config) => {
  const token = localStorage.getItem('auth_token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode
}) {
  const [loading, setLoading] = useState(true)
  const { setUser } = useUserStore()
  const token = localStorage.getItem('auth_token')
  const navigate = useNavigate()

  useEffect(() => {
    if (token) {
      axios
        .get('/api/v1/user')
        .then((res) => {
          setUser(res.data.data)
        })
        .catch(() => {
          localStorage.removeItem('auth_token')
          navigate('/login')
        })
        .finally(() => {
          setLoading(false)
        })
    } else {
      navigate('/login')
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [token, navigate])

  if (loading) {
    return (
      <div className='z-50 flex items-center justify-center h-screen'>
        <Loading loading={true} />
      </div>
    )
  }

  return children
}
