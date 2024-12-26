/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react'
import { Card } from '@/components/ui/card'
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
  TooltipProps,
} from 'recharts'
import { Switch } from '@/components/ui/switch'
import { Label } from '@/components/ui/label'
import axios from 'axios'

interface SystemMetrics {
  cpu: number
  ram: {
    total: number
    used: number
    available: number
    usage: number
  }
  disk: {
    total: number
    used: number
    available: number
    usage: number
  }
  time: string
}

const formatValue = (value: number) => `${value.toFixed(2)}%`
const formatTime = (time: string) => new Date(time).toLocaleTimeString()

export function Analytics() {
  const [metrics, setMetrics] = useState<SystemMetrics[]>([])
  const [autoUpdate, setAutoUpdate] = useState(true)

  const fetchMetrics = () => {
    axios
      .get('/api/v1/system/metrics')
      .then((res) => {
        const data = res.data
        setMetrics((prev) => [...prev, data].slice(-20)) // Keep last 20 points
      })
      .catch((err) => {
        console.error('Error fetching metrics:', err)
      })
  }

  useEffect(() => {
    fetchMetrics() // Initial fetch

    let interval: NodeJS.Timeout | null = null
    if (autoUpdate) {
      interval = setInterval(fetchMetrics, 5000) // Update every 5 seconds
    }

    return () => {
      if (interval) clearInterval(interval)
    }
  }, [autoUpdate])

  const CustomTooltip = ({
    active,
    payload,
    label,
  }: TooltipProps<number, string>) => {
    if (active && payload && payload.length) {
      return (
        <div className='p-2 rounded-lg shadow-md bg-background/95 ring-1 ring-black/5 backdrop-blur'>
          <p className='font-medium'>{formatTime(label)}</p>
          {payload.map((entry: any) => (
            <p key={entry.name} style={{ color: entry.color }}>
              {entry.name}: {formatValue(entry.value)}
            </p>
          ))}
        </div>
      )
    }
    return null
  }

  return (
    <div className='space-y-4'>
      <div className='flex items-center space-x-2'>
        <Switch
          id='auto-update'
          checked={autoUpdate}
          onCheckedChange={setAutoUpdate}
        />
        <Label htmlFor='auto-update'>Auto Update (every 5 seconds)</Label>
      </div>

      <div className='grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3'>
        <Card className='p-4'>
          <h3 className='mb-4 text-lg font-semibold'>CPU Usage</h3>
          <div className='h-[250px] w-full'>
            <ResponsiveContainer width='100%' height='100%'>
              <LineChart data={metrics}>
                <CartesianGrid strokeDasharray='3 3' opacity={0.2} />
                <XAxis
                  dataKey='time'
                  tickFormatter={formatTime}
                  fontSize={12}
                  tickMargin={8}
                />
                <YAxis
                  domain={[0, 100]}
                  tickFormatter={(value) => `${value}%`}
                  fontSize={12}
                  tickMargin={8}
                />
                <Tooltip content={<CustomTooltip />} />
                <Legend />
                <Line
                  type='monotone'
                  dataKey='cpu'
                  stroke='#8884d8'
                  name='CPU'
                  strokeWidth={2}
                  dot={false}
                  activeDot={{ r: 4 }}
                />
              </LineChart>
            </ResponsiveContainer>
          </div>
        </Card>

        <Card className='p-4'>
          <h3 className='mb-4 text-lg font-semibold'>RAM Usage</h3>
          <div className='h-[250px] w-full'>
            <ResponsiveContainer width='100%' height='100%'>
              <LineChart data={metrics}>
                <CartesianGrid strokeDasharray='3 3' opacity={0.2} />
                <XAxis
                  dataKey='time'
                  tickFormatter={formatTime}
                  fontSize={12}
                  tickMargin={8}
                />
                <YAxis
                  domain={[0, 100]}
                  tickFormatter={(value) => `${value}%`}
                  fontSize={12}
                  tickMargin={8}
                />
                <Tooltip content={<CustomTooltip />} />
                <Legend />
                <Line
                  type='monotone'
                  dataKey='ram.usage'
                  stroke='#82ca9d'
                  name='RAM'
                  strokeWidth={2}
                  dot={false}
                  activeDot={{ r: 4 }}
                />
              </LineChart>
            </ResponsiveContainer>
          </div>
        </Card>

        <Card className='p-4'>
          <h3 className='mb-4 text-lg font-semibold'>Disk Usage</h3>
          <div className='h-[250px] w-full'>
            <ResponsiveContainer width='100%' height='100%'>
              <LineChart data={metrics}>
                <CartesianGrid strokeDasharray='3 3' opacity={0.2} />
                <XAxis
                  dataKey='time'
                  tickFormatter={formatTime}
                  fontSize={12}
                  tickMargin={8}
                />
                <YAxis
                  domain={[0, 100]}
                  tickFormatter={(value) => `${value}%`}
                  fontSize={12}
                  tickMargin={8}
                />
                <Tooltip content={<CustomTooltip />} />
                <Legend />
                <Line
                  type='monotone'
                  dataKey='disk.usage'
                  stroke='#ffc658'
                  name='Disk'
                  strokeWidth={2}
                  dot={false}
                  activeDot={{ r: 4 }}
                />
              </LineChart>
            </ResponsiveContainer>
          </div>
        </Card>
      </div>
    </div>
  )
}
