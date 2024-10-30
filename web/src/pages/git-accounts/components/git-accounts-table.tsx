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
import {
  IconBrandGithub,
  IconBrandGitlab,
  IconBrandBitbucket,
  IconAlertCircle,
} from '@tabler/icons-react'
import { useDeleteModalStore } from '@/store/delete-modal-store'
import { GitAccount } from '@/types/model'
import { Loading } from '@/components/custom/loading'

interface GitAccountsTableProps {
  gitAccounts: GitAccount[]
  loading: boolean
}

export function GitAccountsTable({
  gitAccounts,
  loading,
}: GitAccountsTableProps) {
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
              <Loading loading={loading} />
            </TableCell>
          </TableRow>
        ) : gitAccounts?.length === 0 ? (
          <TableRow>
            <TableCell colSpan={6} className='text-center'>
              <div className='flex items-center justify-center'>
                <IconAlertCircle className='mr-2 h-4 w-4' />
                <span>No git accounts found</span>
              </div>
            </TableCell>
          </TableRow>
        ) : (
          gitAccounts.map((gitAccount) => (
            <TableRow key={gitAccount.id}>
              <TableCell className='font-medium'>{gitAccount.id}</TableCell>
              <TableCell className='font-medium'>{gitAccount.name}</TableCell>
              <TableCell className='flex items-center gap-x-2 capitalize'>
                {gitAccount.type === 'github' ? (
                  <IconBrandGithub />
                ) : gitAccount.type === 'gitlab' ? (
                  <IconBrandGitlab />
                ) : (
                  <IconBrandBitbucket />
                )}
                {gitAccount.type}
              </TableCell>
              <TableCell>
                <Link
                  to={gitAccount.url}
                  target='_blank'
                  className='flex items-center gap-x-1 text-blue-500 hover:text-blue-600'
                >
                  {gitAccount.url}
                  <ExternalLinkIcon className='h-4 w-4' />
                </Link>
              </TableCell>
              <TableCell>
                {new Date(gitAccount.createdAt).toLocaleDateString()}
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
                      <Link to={`/git-accounts/edit/${gitAccount.id}`}>
                        Edit
                      </Link>
                    </DropdownMenuItem>
                    <DropdownMenuItem>Make a copy</DropdownMenuItem>
                    <DropdownMenuItem>Favorite</DropdownMenuItem>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem
                      onClick={() => handleDelete(gitAccount.id)}
                    >
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
