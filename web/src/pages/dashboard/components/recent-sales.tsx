import { Link } from 'react-router-dom'

interface Data {
  name: string
  domain: string
  date: string
  status: string
  slug: string
}

const data: Data[] = [
  {
    date: '2024-01-01',
    status: 'success',
    slug: 'ui',
    domain: 'ui.sailhost.local',
    name: 'UI',
  },
  {
    date: '2024-01-02',
    status: 'error',
    slug: 'api',
    domain: 'api.sailhost.local',
    name: 'API Service',
  },
  {
    date: '2024-01-03',
    status: 'success',
    slug: 'docs',
    domain: 'docs.sailhost.local',
    name: 'Documentation',
  },
  {
    date: '2024-01-03',
    status: 'success',
    slug: 'blog',
    domain: 'blog.sailhost.local',
    name: 'Company Blog',
  },
  {
    date: '2024-01-04',
    status: 'pending',
    slug: 'dashboard',
    domain: 'dashboard.sailhost.local',
    name: 'Admin Dashboard',
  },
]

export function RecentSales() {
  return (
    <div className='space-y-8'>
      {data.map((item, index) => (
        <div className='flex items-center' key={index}>
          <div>
            <p className='flex items-center gap-2 text-sm font-medium leading-none'>
              <Link to={`/projects/${item.slug}`} className='hover:underline'>
                {item.name}
              </Link>
              <div className='flex flex-col'>
                <div className='flex items-center'>
                  {item.status === 'success' && (
                    <>
                      <div className='w-2 h-2 bg-green-500 rounded-full' />
                      <span className='ml-1'>Active</span>
                    </>
                  )}
                  {item.status === 'error' && (
                    <>
                      <div className='w-2 h-2 bg-red-500 rounded-full' />
                      <span className='ml-1'>Error</span>
                    </>
                  )}
                  {item.status === 'pending' && (
                    <>
                      <div className='w-2 h-2 bg-yellow-500 rounded-full' />
                      <span className='ml-1'>Pending</span>
                    </>
                  )}
                  {item.status === 'building' && (
                    <>
                      <div className='w-2 h-2 bg-blue-500 rounded-full' />
                      <span className='ml-1'>Building</span>
                    </>
                  )}
                  {item.status === 'deploying' && (
                    <>
                      <div className='w-2 h-2 bg-purple-500 rounded-full' />
                      <span className='ml-1'>Deploying</span>
                    </>
                  )}
                </div>
              </div>
            </p>
            <p className='text-sm text-muted-foreground'>
              <a href={`https://${item.domain}`} className='underline'>
                {item.domain}
              </a>
            </p>
          </div>
          <div className='ml-auto text-sm font-medium'>{item.date}</div>
        </div>
      ))}
    </div>
  )
}
