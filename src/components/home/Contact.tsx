import React from 'react';
import { Box, Typography, Link, Avatar } from '@mui/material';
import {
  Mail as MailIcon,
  Twitter as TwitterIcon,
  GitHub as GithubIcon,
} from '@mui/icons-material';
import styles from 'assets/styles/components/home/Contact.module.scss';

const Contact: React.FC = () => {
  return (
    <Box className={styles.wrapper}>
      <Box className={styles.caption}>
        <Typography variant="h5" color="initial">
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
  );
};

export default Contact;
