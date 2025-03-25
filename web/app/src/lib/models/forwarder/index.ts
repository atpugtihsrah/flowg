import { DatadogForwarderModel } from '@/lib/models/forwarder/datadog'
import { HttpForwarderModel } from '@/lib/models/forwarder/http'
import { SyslogForwarderModel } from '@/lib/models/forwarder/syslog'

export type ForwarderModel = {
  config: ForwarderConfigModel
}

export const ForwarderTypeValues = [
  { key: 'http', label: 'Webhook' },
  { key: 'syslog', label: 'Syslog Server' },
  { key: 'datadog', label: 'Datadog' },
] as const

export const ForwarderTypeLabelMap = ForwarderTypeValues.reduce(
  (acc, cur) => {
    acc[cur.key] = cur.label
    return acc
  },
  {} as Record<ForwarderTypes, string>,
)

export type ForwarderConfigModel =
  | HttpForwarderModel
  | SyslogForwarderModel
  | DatadogForwarderModel

export type ForwarderTypes = ForwarderConfigModel['type']
