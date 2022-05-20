import React from 'react';
import AnchorLink from 'react-anchor-link-smooth-scroll';
import { AppBar, Box, Button, Container, Link, List, Menu, MenuItem, Toolbar, Typography } from '@mui/material';
import { Menu as MenuIcon } from '@mui/icons-material';
import styles from 'assets/styles/components/home/Header.module.scss';

interface navbarElement {
  href: string;
  title: string;
  isExternal: boolean;
}

const navbarElements: navbarElement[] = [
  {
    href: '#about',
    title: 'ABOUT',
    isExternal: false,
  },
  {
    href: '#skills',
    title: 'SKILLS',
    isExternal: false,
  },
  {
    href: '#works',
    title: 'WORKS',
    isExternal: false,
  },
  {
    href: '#contact',
    title: 'CONTACT',
    isExternal: false,
  },
  {
    href: 'https://itk13201.hatenablog.com/archive',
    title: 'BLOG',
    isExternal: true,
  },
];

const SpNavbar = (): JSX.Element => {
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
        {navbarElements.map((element: navbarElement) =>
          !element.isExternal ? (
            <MenuItem key={'sp-' + element.title} onClick={handleClose}>
              <AnchorLink href={element.href} className={styles.link}>
                {element.title}
              </AnchorLink>
            </MenuItem>
          ) : (
            <MenuItem key={'sp-' + element.title} onClick={handleClose}>
              <Link
                href={element.href}
                target="_blank"
                rel="noopener noreferrer"
                className={styles.link}
                underline="none"
              >
                {element.title}
              </Link>
            </MenuItem>
          )
        )}
      </Menu>
    </div>
  );
};

const PcNavbar = (): JSX.Element => {
  return (
    <div className={styles.pcList}>
      <List>
        {navbarElements.map((element: navbarElement) =>
          !element.isExternal ? (
            <AnchorLink
              key={'pc-' + element.title}
              href={element.href}
              className={styles.link}
            >
              <Button>{element.title}</Button>
            </AnchorLink>
          ) : (
            <Link
              key={'pc-' + element.title}
              href={element.href}
              target="_blank"
              rel="noopener noreferrer"
              className={styles.link}
            >
              <Button>{element.title}</Button>
            </Link>
          )
        )}
      </List>
    </div>
  );
};

const Header = (): JSX.Element => {
  return (
    <AppBar color="default" position="static">
      <Toolbar>
        <Container maxWidth="md" className={styles.container}>
          <Box>
            <Typography variant="h2" color="initial" className={styles.title}>
              <Link href="/" color="inherit" underline="none">
                Portfolio
              </Link>
            </Typography>
          </Box>
          {/*display if sp screen width*/}
          <SpNavbar />
          {/*display if pc screen width*/}
          <PcNavbar />
        </Container>
      </Toolbar>
    </AppBar>
  );
};

export default Header;
