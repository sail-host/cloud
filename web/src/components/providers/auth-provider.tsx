import { Navigate } from "react-router-dom"

export default function AuthProvider({ children }: { children: React.ReactNode }) {
    // TODO: Fake auth provider
    const isAuthenticated = localStorage.getItem('auth_token') !== null

    if (!isAuthenticated) {
        return <Navigate to='/login' />
    }

    return children
}