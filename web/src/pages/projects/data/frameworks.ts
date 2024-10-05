export interface Framework {
  name: string
  value: string
  image: string
}

export const frameworks: Framework[] = [
  {
    name: 'Next.js',
    value: 'nextjs',
    image: '/images/frameworks/next.png',
  },
  {
    name: 'React',
    value: 'react',
    image: '/images/frameworks/react.png',
  },
  {
    name: 'Vue',
    value: 'vue',
    image: '/images/frameworks/vue.png',
  },
  {
    name: 'Vite',
    value: 'vite',
    image: '/images/frameworks/vite.png',
  },
  {
    name: 'Svelte',
    value: 'svelte',
    image: '/images/frameworks/svelte.png',
  },
  {
    name: 'Astro',
    value: 'astro',
    image: '/images/frameworks/astro.png',
  },
  {
    name: 'Nuxt',
    value: 'nuxt',
    image: '/images/frameworks/nuxt.png',
  },
  {
    name: 'Remix',
    value: 'remix',
    image: '/images/frameworks/remix.png',
  },
]
