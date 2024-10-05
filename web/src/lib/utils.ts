import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

dayjs.extend(relativeTime)
export function formatDistanceToNow(date: Date) {
  return dayjs(date).fromNow(true)
}

export function randomString(length: number) {
  return Math.random()
    .toString(36)
    .substring(2, length + 2)
}

export function slugify(text: string, separator = '-') {
  return text
    .toString()
    .toLowerCase()
    .replace(/\s+/g, separator)
    .replace(/[^\w-]+/g, '')
    .replace(/^-+/, '')
    .replace(/-+$/, '')
}
