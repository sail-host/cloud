export interface User {
  id: string
  email: string
  name: string
  role: 'admin' | 'user'
}

export interface GitAccount {
  id: string
  name: string
  url: string
  type: 'github' | 'gitlab' | 'bitbucket' | 'gitea'
  token: string
  createdAt: string
  updatedAt: string
}

export interface Domain {
  id: string
  domain: string
  dns_provider: 'cloudflare' | 'manual'
  cloudflare_zone_id: string
  cloudflare_api_key: string
  createdAt: string
  updatedAt: string
}
