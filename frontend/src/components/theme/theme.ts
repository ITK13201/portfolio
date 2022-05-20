import { createTheme } from '@mui/material';

export const theme = createTheme({
  typography: {
    h1: {
      fontSize: '4rem',
      fontWeight: 500,
      '@media screen and (max-width: 599px)': {
        fontSize: '3rem',
      },
    },
    h2: {
      fontSize: '1.8rem',
      fontWeight: 400,
      '@media screen and (max-width: 599px)': {
        fontSize: '1.5rem',
      },
    },
    h3: {
      fontSize: '1.5rem',
      fontWeight: 400,
      '@media screen and (max-width: 599px)': {
        fontSize: '1.3rem',
      },
    },
    h4: undefined,
    h5: undefined,
  },
});
