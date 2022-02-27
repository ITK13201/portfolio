import { createTheme } from '@mui/material';

export const CustomMaterialTheme = createTheme({
  typography: {
    h1: {
      fontSize: '4rem',
      fontWeight: 500,
      '@media screen and (max-width: 600px)': {
        fontSize: '3rem',
      },
    },
    h2: {
      fontSize: '2rem',
      fontWeight: 400,
      '@media screen and (max-width: 600px)': {
        fontSize: '1.5rem',
      },
    },
  },
});
