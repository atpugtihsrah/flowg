import { useRef, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useNotifications } from '@toolpad/core/useNotifications'
import { useConfig } from '@/lib/context/config'

import KeyIcon from '@mui/icons-material/Key'

import Card from '@mui/material/Card'
import CardHeader from '@mui/material/CardHeader'
import CardContent from '@mui/material/CardContent'

import { AgGridReact } from 'ag-grid-react'
import { ColDef } from 'ag-grid-community'

import { Actions } from '@/components/table/actions'

import { CreateTokenButton } from './create-btn'
import { RowType } from './types'
import { TokenCell } from './cell'

import { UnauthenticatedError } from '@/lib/api/errors'
import * as tokenApi from '@/lib/api/operations/token'

type TokenListProps = {
  tokens: string[]
}

export const TokenList = ({ tokens }: TokenListProps) => {
  const navigate = useNavigate()
  const notifications = useNotifications()
  const config = useConfig()

  const [loading, setLoading] = useState(false)

  const gridRef = useRef<AgGridReact<RowType>>(null!)
  const [rowData] = useState<RowType[]>(
    tokens.map(token => ({ token })),
  )
  const [columnDefs] = useState<ColDef<RowType>[]>([
    {
      headerName: 'Token',
      field: 'token',
      cellRenderer: TokenCell,
      suppressMovable: true,
      sortable: false,
      flex: 1,
    },
    {
      headerName: 'Actions',
      headerClass: 'flowg-actions-header',
      cellRenderer: Actions,
      cellRendererParams: {
        onDelete: async (data: RowType) => {
          setLoading(true)

          try {
            await tokenApi.deleteToken(data.token)

            const rowNode = gridRef.current.api.getRowNode(data.token)
            if (rowNode !== undefined) {
              gridRef.current.api.applyTransaction({
                remove: [rowNode.data],
              })
            }

            notifications.show('Token deleted', {
              severity: 'success',
              autoHideDuration: config.notifications?.autoHideDuration,
            })
          }
          catch (error) {
            if (error instanceof UnauthenticatedError) {
              notifications.show('Session expired', {
                severity: 'error',
                autoHideDuration: config.notifications?.autoHideDuration,
              })
              navigate('/web/login')
            }
            else {
              notifications.show('Unknown error', {
                severity: 'error',
                autoHideDuration: config.notifications?.autoHideDuration,
              })
            }

            console.error(error)
          }

          setLoading(false)
        },
      },
      suppressMovable: true,
      sortable: false,
    },
  ])

  const onNewToken = (token: string) => {
    gridRef.current.api.applyTransaction({
      add: [{ token }],
    })
  }

  return (
    <>
      <Card className="max-lg:min-h-96 lg:h-full flex flex-col items-stretch">
        <CardHeader
          title={
            <div className="flex items-center gap-3">
              <KeyIcon />
              <span className="flex-grow">API Tokens</span>
              <CreateTokenButton onTokenCreated={onNewToken} />
            </div>
          }
          className="bg-blue-400 text-white shadow-lg z-20"
        />
        <CardContent className="!p-0 flex-grow flex-shrink h-0 ag-theme-material flowg-table">
          <AgGridReact<RowType>
            ref={gridRef}
            loading={loading}
            rowData={rowData}
            columnDefs={columnDefs}
            enableCellTextSelection
            getRowId={({ data }) => data.token}
          />
        </CardContent>
      </Card>
    </>
  )
}
