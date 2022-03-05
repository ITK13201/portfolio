import React from 'react';
import { Box, Typography, Link, Avatar, ThemeProvider } from '@mui/material';
import {
  Mail as MailIcon,
  Twitter as TwitterIcon,
  GitHub as GithubIcon,
} from '@mui/icons-material';
import styles from 'assets/styles/components/home/Contact.module.scss';
import { CustomMaterialTheme } from '../../hooks/theme';

const Contact = (): JSX.Element => {
  return (
    <ThemeProvider theme={CustomMaterialTheme}>
      <Box className={styles.root}>
        <Box className={styles.caption}>
          <Typography variant="h2" color="initial">
            Contact
          </Typography>
        </Box>
        <Box className={styles.container}>
          <Link
            href="mailto:tkik2236@gmail.com?subject=Hello"
            color="inherit"
            title="Email"
            target="_blank"
            rel="noopener noreferrer"
          >
            <Avatar className={styles.email}>
              <MailIcon />
            </Avatar>
          </Link>
          <Link
            href="https://twitter.com/itk13201"
            color="inherit"
            title="Twitter"
            target="_blank"
            rel="noopener noreferrer"
          >
            <Avatar className={styles.twitter}>
              <TwitterIcon />
            </Avatar>
          </Link>
          <Link
            href="https://github.com/ITK13201"
            color="inherit"
            title="Github"
            target="_blank"
            rel="noopener noreferrer"
          >
            <Avatar className={styles.github}>
              <GithubIcon />
            </Avatar>
          </Link>
        </Box>
      </Box>
    </ThemeProvider>
  );
};

export default Contact;
