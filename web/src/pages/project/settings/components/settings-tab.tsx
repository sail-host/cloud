import { CrontabTab } from './crontab-tab'
import { DomainsTab } from './domains-tab'
import { EnvironmentsTab } from './environments-tab'
import { GeneralTab } from './general-tab'
import { GitTab } from './git-tab'
import { NginxConfigTab } from './nginx-config-tab'
import { SslTab } from './ssl-tab'

interface SettingsTabProps {
  activeTab: string
}

export function SettingsTab({ activeTab }: SettingsTabProps) {
  return (
    <>
      {activeTab === 'general' && <GeneralTab />}
      {activeTab === 'crontab' && <CrontabTab />}
      {activeTab === 'nginx-config' && <NginxConfigTab />}
      {activeTab === 'environments' && <EnvironmentsTab />}
      {activeTab === 'domains' && <DomainsTab />}
      {activeTab === 'git' && <GitTab />}
      {activeTab === 'ssl' && <SslTab />}
    </>
  )
}
