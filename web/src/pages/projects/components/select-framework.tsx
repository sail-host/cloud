import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import { Button } from '@/components/ui/button'
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command'
import { frameworks } from '../data/frameworks'
import { useState } from 'react'
import { IconCheck, IconChevronDown } from '@tabler/icons-react'
import { cn } from '@/lib/utils'

interface SelectFrameworkProps {
  selectedFramework: string | null
  setSelectedFramework: (framework: string | null) => void
}

export function SelectFramework({
  selectedFramework,
  setSelectedFramework,
}: SelectFrameworkProps) {
  const [open, setOpen] = useState(false)

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant='outline'
          role='combobox'
          aria-expanded={open}
          className='w-full justify-between'
        >
          {selectedFramework &&
          frameworks.find(
            (framework) => framework.value === selectedFramework
          ) ? (
            <div className='flex items-center'>
              <img
                src={
                  frameworks.find(
                    (framework) => framework.value === selectedFramework
                  )?.image
                }
                alt={
                  frameworks.find(
                    (framework) => framework.value === selectedFramework
                  )?.name
                }
                width={20}
                height={20}
                className='mr-2'
              />
              {
                frameworks.find(
                  (framework) => framework.value === selectedFramework
                )?.name
              }
            </div>
          ) : (
            'Select framework...'
          )}
          <IconChevronDown className='ml-2 h-4 w-4 shrink-0 opacity-50' />
        </Button>
      </PopoverTrigger>
      <PopoverContent className='w-[740px] p-0'>
        <Command>
          <CommandInput placeholder='Search framework...' />
          <CommandList>
            <CommandEmpty>No framework found.</CommandEmpty>
            <CommandGroup>
              {frameworks.map((framework) => (
                <CommandItem
                  key={framework.value}
                  value={framework.value}
                  onSelect={(currentValue) => {
                    setSelectedFramework(
                      currentValue === selectedFramework ? '' : currentValue
                    )
                    setOpen(false)
                  }}
                  className='relative'
                >
                  <img
                    src={framework.image}
                    alt={framework.name}
                    width={20}
                    height={20}
                    className='mr-2'
                  />
                  {framework.name}
                  <IconCheck
                    className={cn(
                      'absolute right-2 h-4 w-4',
                      selectedFramework === framework.value
                        ? 'opacity-100'
                        : 'opacity-0'
                    )}
                  />
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  )
}
