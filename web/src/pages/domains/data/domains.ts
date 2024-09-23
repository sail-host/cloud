export interface Domain {
  id: number
  name: string
  dnsProvider: 'cloudflare' | 'manual'
  cloudflareAccountId?: string
  cloudflareApiToken?: string
  createdAt: string
}

export const domains: Domain[] = [
  {
    id: 1,
    name: 'techsphere.io',
    dnsProvider: 'cloudflare',
    cloudflareAccountId: '9876543210',
    cloudflareApiToken: 'abcdefghij1234567890',
    createdAt: '2023-05-15',
  },
  {
    id: 2,
    name: 'digitalwave.com',
    dnsProvider: 'manual',
    createdAt: '2023-06-22',
  },
  {
    id: 3,
    name: 'innovatelab.net',
    dnsProvider: 'cloudflare',
    cloudflareAccountId: '1357924680',
    cloudflareApiToken: 'qwertyuiop0987654321',
    createdAt: '2023-07-30',
  },
  {
    id: 4,
    name: 'cyberbridge.org',
    dnsProvider: 'manual',
    createdAt: '2023-08-12',
  },
  {
    id: 5,
    name: 'quantumleap.tech',
    dnsProvider: 'cloudflare',
    cloudflareAccountId: '2468135790',
    cloudflareApiToken: 'zxcvbnmasdfghjkl1234',
    createdAt: '2023-09-05',
  },
  {
    id: 6,
    name: 'nexusportal.co',
    dnsProvider: 'manual',
    createdAt: '2023-10-18',
  },
  {
    id: 7,
    name: 'futurehub.dev',
    dnsProvider: 'cloudflare',
    cloudflareAccountId: '3692581470',
    cloudflareApiToken: 'poiuytrewqlkjhgfdsa',
    createdAt: '2023-11-27',
  },
  {
    id: 8,
    name: 'codeforge.app',
    dnsProvider: 'manual',
    createdAt: '2023-12-09',
  },
  {
    id: 9,
    name: 'datastream.cloud',
    dnsProvider: 'cloudflare',
    cloudflareAccountId: '7531902468',
    cloudflareApiToken: 'mnbvcxzlkjhgfdsa9876',
    createdAt: '2024-01-14',
  },
  {
    id: 10,
    name: 'aiplatform.io',
    dnsProvider: 'manual',
    createdAt: '2024-02-28',
  },
]
