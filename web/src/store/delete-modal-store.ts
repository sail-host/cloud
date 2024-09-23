import { create } from 'zustand'

interface DeleteModalStore {
  open: boolean
  setOpen: (open: boolean) => void
  deleteID: string | number
  setDeleteID: (deleteID: string | number) => void
}

export const useDeleteModalStore = create<DeleteModalStore>((set) => ({
  open: false,
  setOpen: (open) => set({ open }),
  deleteID: '',
  setDeleteID: (deleteID) => set({ deleteID }),
}))
