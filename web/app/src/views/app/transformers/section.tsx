import { useEffect } from 'react'
import { useLoaderData, useNavigate } from 'react-router-dom'

import Backdrop from '@mui/material/Backdrop'
import CircularProgress from '@mui/material/CircularProgress'

import { NewTransformerButton } from './new-btn'

export const TransformerView = () => {
  const { transformers } = useLoaderData() as { transformers: string[] }
  const navigate = useNavigate()

  useEffect(
    () => {
      if (transformers.length > 0) {
        navigate(`/web/transformers/${transformers[0]}`)
      }
    },
    [],
  )

  return (
    <>
      {transformers.length > 0
        ? (
          <Backdrop open={true}>
            <CircularProgress color="inherit" />
          </Backdrop>
        )
        : (
          <div className="w-full h-full flex flex-col items-center justify-center gap-5">
            <h1 className="text-3xl font-semibold">No transformer found, create one</h1>

            <NewTransformerButton
              onTransformerCreated={(name) => {
                window.location.pathname = `/web/transformers/${name}`
              }}
            />
          </div>
        )
      }
    </>
  )
}
