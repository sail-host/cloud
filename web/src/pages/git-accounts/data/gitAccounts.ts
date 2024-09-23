export interface GitAccount {
  id: string
  name: string
  gitUrl: string
  type: 'github' | 'gitlab' | 'bitbucket'
  createdAt: Date
}

export const gitAccounts: GitAccount[] = [
  {
    id: '1',
    name: 'JohnDoe123',
    gitUrl: 'https://github.com/JohnDoe123',
    type: 'github',
    createdAt: new Date('2023-01-15'),
  },
  {
    id: '2',
    name: 'AliceSmith',
    gitUrl: 'https://gitlab.com/AliceSmith',
    type: 'gitlab',
    createdAt: new Date('2023-02-28'),
  },
  {
    id: '3',
    name: 'BobJohnson',
    gitUrl: 'https://bitbucket.org/BobJohnson',
    type: 'bitbucket',
    createdAt: new Date('2023-03-10'),
  },
  {
    id: '4',
    name: 'EvaWilliams',
    gitUrl: 'https://github.com/EvaWilliams',
    type: 'github',
    createdAt: new Date('2023-04-22'),
  },
  {
    id: '5',
    name: 'MikeBrown',
    gitUrl: 'https://gitlab.com/MikeBrown',
    type: 'gitlab',
    createdAt: new Date('2023-05-05'),
  },
  {
    id: '6',
    name: 'SarahLee',
    gitUrl: 'https://bitbucket.org/SarahLee',
    type: 'bitbucket',
    createdAt: new Date('2023-06-18'),
  },
  {
    id: '7',
    name: 'ChrisGreen',
    gitUrl: 'https://github.com/ChrisGreen',
    type: 'github',
    createdAt: new Date('2023-07-30'),
  },
  {
    id: '8',
    name: 'EmmaWhite',
    gitUrl: 'https://gitlab.com/EmmaWhite',
    type: 'gitlab',
    createdAt: new Date('2023-08-12'),
  },
  {
    id: '9',
    name: 'DavidBlack',
    gitUrl: 'https://bitbucket.org/DavidBlack',
    type: 'bitbucket',
    createdAt: new Date('2023-09-25'),
  },
  {
    id: '10',
    name: 'OliviaGray',
    gitUrl: 'https://github.com/OliviaGray',
    type: 'github',
    createdAt: new Date('2023-10-08'),
  },
]
