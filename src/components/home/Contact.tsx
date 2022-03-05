import React from 'react';
import { Avatar, Box, Link, Typography } from '@mui/material';
import { GitHub as GithubIcon, Mail as MailIcon, Twitter as TwitterIcon } from '@mui/icons-material';
import styles from 'assets/styles/components/home/Contact.module.scss';

const Contact = (): JSX.Element => {
  return (
    <Box className={styles.root}>
      <Box className={styles.caption}>
        <Typography variant="h2" color="initial">
          <strong>Contact</strong>
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
  );
};

export default Contact;
