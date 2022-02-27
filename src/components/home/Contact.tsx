import React from 'react';
import { Box, Typography, Link, Avatar } from '@mui/material';
import {
  Mail as MailIcon,
  Twitter as TwitterIcon,
  GitHub as GithubIcon,
} from '@mui/icons-material';
import styles from 'styles/components/home/Contact.module.scss';

const Contact: React.FC = () => {
  return (
    <Box p={2} className={styles.container}>
      <Box display="flex" justifyContent="center" p={1}>
        <Typography variant="h5" color="initial">
          Contact
        </Typography>
      </Box>
      <Box className={styles.root} display="flex" justifyContent="center" p={1}>
        <Link
          href="mailto:tkik2236@gmail.com?subject=Hello"
          color="inherit"
          title="Email"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Avatar className={styles.green}>
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
          <Avatar className={styles.blue}>
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
          <Avatar className={styles.purple}>
            <GithubIcon />
          </Avatar>
        </Link>
      </Box>
    </Box>
  );
};

export default Contact;
