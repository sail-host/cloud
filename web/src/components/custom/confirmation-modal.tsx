import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { Button } from '../ui/button'

interface ConfirmationModalProps {
  open: boolean
  setOpen: (open: boolean) => void
  confirmFunction: () => void
  title: string
  description: string
  loading?: boolean
}

export function ConfirmationModal({
  open,
  setOpen,
  confirmFunction,
  title,
  description,
  loading,
}: ConfirmationModalProps) {
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
          <DialogDescription>{description}</DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button
            onClick={confirmFunction}
            loading={loading}
            type='button'
            variant='destructive'
          >
            Confirm
          </Button>
          <Button
            type='button'
            variant='outline'
            onClick={() => setOpen(false)}
          >
            Cancel
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
