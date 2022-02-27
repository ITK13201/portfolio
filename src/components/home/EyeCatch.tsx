import React from 'react';
import { Box, Avatar, Typography, ThemeProvider } from '@mui/material';

import ImageFile from 'assets/images/components/home/avatar.jpg';
import styles from 'assets/styles/components/home/EyeCatch.module.scss';
import { CustomMaterialTheme } from 'hooks/theme';

const EyeCatch: React.FC = () => {
  return (
    <ThemeProvider theme={CustomMaterialTheme}>
      <Box className={styles.root}>
        <Box className={styles.avatarContainer}>
          <Avatar
            alt="Takumi Ikeda"
            src={ImageFile.src}
            className={styles.avatar}
          />
        </Box>
        <Box className={styles.whoami}>
          <Typography variant="h1" color="initial">
            <strong>Takumi Ikeda</strong>
          </Typography>
          <Typography variant="h2" color="initial">
            (handle name: I.TK, itk13201)
          </Typography>
        </Box>
      </Box>
    </ThemeProvider>
  );
};

export default EyeCatch;
