# Shadcn Admin Dashboard

Admin Dashboard UI crafted with Shadcn and Vite. Built with responsiveness and accessibility in mind.

![alt text](public/images/shadcn-admin.png)

I've been creating dashboard UIs at work and for my personal projects. I always wanted to make a reusable collection of dashboard UI for future projects; and here it is now. While I've created a few custom components, some of the code is directly adapted from ShadcnUI examples.

> This is not a starter project (template) though. I'll probably make one in the future.

## Features

- Light/dark mode
- Responsive
- Accessible
- Sidebar and header layouts
- 10+ pages
- Extra custom components

## Tech Stack

**UI:** [ShadcnUI](https://ui.shadcn.com) (TailwindCSS + RadixUI)

**Build Tool:** [Vite](https://vitejs.dev/)

**Routing:** [React Router](https://reactrouter.com/en/main)

**Type Checking:** [TypeScript](https://www.typescriptlang.org/)

**Linting/Formatting:** [Eslint](https://eslint.org/) & [Prettier](https://prettier.io/)

**Icons:** [Tabler Icons](https://tabler.io/icons)

## Run Locally

Clone the project

```bash
  git clone https://github.com/satnaing/shadcn-admin.git
```

Go to the project directory

```bash
  cd shadcn-admin
```

##### Run using bun

For better performance, consider using [Bun](https://bun.sh/).

Install `bun` for **macOS** and **Linux**.

```bash
# for macOS, Linux, and WSL
curl -fsSL https://bun.sh/install | bash
# to install a specific version
curl -fsSL https://bun.sh/install | bash -s "bun-v1.0.0"
```

Install `bun` for **Windows**.

```bash
powershell -c "irm bun.sh/install.ps1|iex"
```

Install dependencies

```bash
  bun install
```

Start the server

```bash
  bun run dev
```

##### Run using npm

Install dependencies

```bash
  npm install
```

Start the server

```bash
  npm run dev
```

## Author

Crafted with 🤍 by [@satnaing](https://github.com/satnaing)

## License

Licensed under the [MIT License](https://choosealicense.com/licenses/mit/)
