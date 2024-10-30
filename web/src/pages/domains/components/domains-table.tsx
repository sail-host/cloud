import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Button } from '@/components/custom/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'

import { DotsHorizontalIcon, ExternalLinkIcon } from '@radix-ui/react-icons'
import { Link } from 'react-router-dom'
import { useDeleteModalStore } from '@/store/delete-modal-store'
import {
  IconAlertCircle,
  IconBrandCloudflare,
  IconWorld,
} from '@tabler/icons-react'
import { Domain } from '@/types/model'
import { Loading } from '@/components/custom/loading'

interface DomainsTableProps {
  domains: Domain[]
  loading: boolean
}

export function DomainsTable({ domains, loading }: DomainsTableProps) {
  const { setOpen, setDeleteID } = useDeleteModalStore()

  const handleDelete = (id: string | number) => {
    setDeleteID(id)
    setOpen(true)
  }

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead className='w-[100px]'>ID</TableHead>
          <TableHead>Name</TableHead>
          <TableHead>Type</TableHead>
          <TableHead>Git URL</TableHead>
          <TableHead>Created At</TableHead>
          <TableHead className='text-right'>Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {loading ? (
          <TableRow>
            <TableCell colSpan={6} className='text-center'>
              <Loading loading />
            </TableCell>
          </TableRow>
        ) : domains?.length === 0 ? (
          <TableRow>
            <TableCell colSpan={6} className='text-center'>
              <div className='flex items-center justify-center'>
                <IconAlertCircle className='mr-2 h-4 w-4' />
                <span>No domains found</span>
              </div>
            </TableCell>
          </TableRow>
        ) : (
          domains.map((domain) => (
            <TableRow key={domain.id}>
              <TableCell className='font-medium'>{domain.id}</TableCell>
              <TableCell className='font-medium'>{domain.domain}</TableCell>
              <TableCell className='flex items-center gap-x-2 capitalize'>
                {domain.dns_provider === 'cloudflare' ? (
                  <IconBrandCloudflare />
                ) : (
                  <IconWorld />
                )}
                {domain.dns_provider}
              </TableCell>
              <TableCell>
                <Link
                  to={`https://${domain.domain}`}
                  target='_blank'
                  className='flex items-center gap-x-1 text-blue-500 hover:text-blue-600'
                >
                  {domain.domain}
                  <ExternalLinkIcon className='h-4 w-4' />
                </Link>
              </TableCell>
              <TableCell>
                {new Date(domain.createdAt).toLocaleString()}
              </TableCell>
              <TableCell className='flex justify-end'>
                <DropdownMenu>
                  <DropdownMenuTrigger asChild>
                    <Button
                      variant='ghost'
                      className='flex h-8 w-8 p-0 text-right data-[state=open]:bg-muted'
                    >
                      <DotsHorizontalIcon className='h-4 w-4' />
                      <span className='sr-only'>Open menu</span>
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align='end' className='w-[160px]'>
                    <DropdownMenuItem asChild>
                      <Link to={`/domains/edit/${domain.id}`}>Edit</Link>
                    </DropdownMenuItem>
                    <DropdownMenuItem>Make a copy</DropdownMenuItem>
                    <DropdownMenuItem>Favorite</DropdownMenuItem>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem onClick={() => handleDelete(domain.id)}>
                      Delete
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </TableCell>
            </TableRow>
          ))
        )}
      </TableBody>
    </Table>
  )
}
