import React from 'react';
import ScrollToTop from 'react-scroll-to-top';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';

const ScrollUp = (): JSX.Element => {
  return (
    <ScrollToTop
      smooth
      style={{
        borderRadius: '50%',
      }}
      component={<KeyboardArrowUpIcon />}
    />
  );
};

export default ScrollUp;
