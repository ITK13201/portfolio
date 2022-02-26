import React, { useEffect } from 'react';
import AnchorLink from 'react-anchor-link-smooth-scroll';
import {
  AppBar,
  Toolbar,
  Button,
  Typography,
  Container,
  List,
  Box,
  Link,
  Menu,
  MenuItem,
} from '@mui/material';
import { Menu as MenuIcon } from '@mui/icons-material';
import styles from 'styles/components/home/Navbar.module.scss';

interface navbarElement {
  href: string;
  className: string;
  title: string;
}

const navbarElements: navbarElement[] = [
  {
    href: '#about',
    className: styles.link,
    title: 'ABOUT',
  },
  {
    href: '#skills',
    className: styles.link,
    title: 'SKILLS',
  },
  {
    href: '#works',
    className: styles.link,
    title: 'WORKS',
  },
  {
    href: '#contact',
    className: styles.link,
    title: 'CONTACT',
  },
];

const SpNavbarList: React.FC = () => {
  const [anchorEl, setAnchorEl] = React.useState(null);

  const handleClick = (event: any) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div className={styles.spList}>
      <Button
        aria-controls="simple-menu"
        aria-haspopup="true"
        onClick={handleClick}
      >
        <MenuIcon />
      </Button>
      <Menu
        id="simple-menu"
        anchorEl={anchorEl}
        keepMounted
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        {navbarElements.map((element: navbarElement, index) => (
          <MenuItem key={'sp-' + element.title} onClick={handleClose}>
            <AnchorLink href={element.href} className={element.className}>
              {element.title}
            </AnchorLink>
          </MenuItem>
        ))}
        <MenuItem onClick={handleClose}>
          <Link
            href="https://itk13201.hatenablog.com/archive"
            target="_blank"
            rel="noopener noreferrer"
            className={styles.link}
            underline="none"
          >
            BLOG
          </Link>
        </MenuItem>
      </Menu>
    </div>
  );
};

const PcNavbarList: React.FC = () => {
  return (
    <div className={styles.pcList}>
      <List>
        {navbarElements.map((element: navbarElement) => (
          <AnchorLink
            key={'pc-' + element.title}
            href={element.href}
            className={element.className}
          >
            <Button>{element.title}</Button>
          </AnchorLink>
        ))}
        <Link
          href="https://itk13201.hatenablog.com/archive"
          target="_blank"
          rel="noopener noreferrer"
          className={styles.link}
        >
          <Button>BLOG</Button>
        </Link>
      </List>
    </div>
  );
};

const Navbar: React.FC = () => {
  return (
    <AppBar color="default" position="static">
      <Toolbar>
        <Container maxWidth="md" className={styles.container}>
          <Box>
            <Typography variant="h5" color="initial" className={styles.title}>
              <Link href="/" color="inherit" underline="none">
                Portfolio
              </Link>
            </Typography>
          </Box>
          {/*display if sp screen width*/}
          <SpNavbarList />
          {/*display if pc screen width*/}
          <PcNavbarList />
        </Container>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
