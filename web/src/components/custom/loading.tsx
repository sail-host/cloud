import { cn } from '@/lib/utils'

interface LoadingProps extends React.HTMLAttributes<HTMLDivElement> {
  loading: boolean
  className?: string
}

export const Loading: React.FC<LoadingProps> = ({
  loading,
  className,
  ...props
}) => {
  if (!loading) return null

  return (
    <div {...props} className={cn('spinner', className)}>
      {[...Array(12)].map((_, index) => (
        <div key={index} className='spinner-blade' />
      ))}
    </div>
  )
}
