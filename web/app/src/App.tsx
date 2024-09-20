import { RouterProvider } from 'react-router-dom'

import { createTheme, ThemeProvider } from '@mui/material/styles'
import * as colors from '@mui/material/colors'

import { DialogsProvider } from '@toolpad/core/useDialogs'
import { NotificationsProvider } from '@toolpad/core/useNotifications'

import router from '@/router'

const theme = createTheme({
  shape: {
    borderRadius: 0,
  },
  palette: {
    primary: {
      main: colors.blue[800],
    },
    secondary: {
      main: colors.teal[400],

    }
  }
})

export default function App() {
  return (
    <ThemeProvider theme={theme}>
      <DialogsProvider>
        <NotificationsProvider>
          <RouterProvider router={router} />
        </NotificationsProvider>
      </DialogsProvider>
    </ThemeProvider>
  )
}
