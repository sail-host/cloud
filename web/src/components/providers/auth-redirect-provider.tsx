import { Navigate } from "react-router-dom"

export default function AuthRedirectProvider({ children }: { children: React.ReactNode }) {
    // Check if user is authenticated redirect to home

    // TODO: Implement auth check
    const isAuthenticated = localStorage.getItem('auth_token')
    if (isAuthenticated) {
        return <Navigate to='/' />
    }

    return children
}