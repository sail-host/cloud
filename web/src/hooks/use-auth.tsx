import { useUserStore } from '@/store/user-store'
import axios from 'axios'
import { useState } from 'react'
export default function useAuth() {
  const [loading, setLoading] = useState(false)
  const { user, setUser } = useUserStore()

  const login = () => {
    setLoading(true)
    const token = localStorage.getItem('auth_token')

    if (!token) {
      setLoading(false)
      return
    }

    axios
      .get('/api/v1/user', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        setUser(res.data.data)
      })
      .catch((err) => {
        console.error(err)
      })
      .finally(() => {
        setLoading(false)
      })
  }

  const logout = () => {
    // TODO: Logout user
    setUser(null)
    localStorage.removeItem('auth_token')
  }

  return { user, login, logout, loading }
}
