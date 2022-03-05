import React from 'react';
import { Box, Avatar, Typography } from '@mui/material';

import ImageFile from 'assets/images/avatar.jpg';
import styles from 'assets/styles/components/home/EyeCatch.module.scss';

const EyeCatch = (): JSX.Element => {
  return (
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
  );
};

export default EyeCatch;
