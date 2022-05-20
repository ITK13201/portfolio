import React from 'react';
import ScrollToTop from 'react-scroll-to-top';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import styles from 'assets/styles/components/design/ScrollUp.module.scss';

const ScrollUp = (): JSX.Element => {
  return (
    <ScrollToTop
      smooth
      className={styles.root}
      component={<KeyboardArrowUpIcon />}
    />
  );
};

export default ScrollUp;
