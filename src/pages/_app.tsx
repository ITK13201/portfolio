import * as React from 'react';
import Head from 'next/head';
import { AppProps } from 'next/app';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { CacheProvider, EmotionCache } from '@emotion/react';
import { theme } from 'components/theme/theme';
import createEmotionCache from 'hooks/createEmotionCache';

// roboto font imports
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';

// Client-side cache, shared for the whole session of the user in the browser.
const clientSideEmotionCache = createEmotionCache();

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache;
}

export default function MyApp(props: MyAppProps) {
  const { Component, emotionCache = clientSideEmotionCache, pageProps } = props;
  const faviconPath = '/images/favicon';

  return (
    <CacheProvider value={emotionCache}>
      <Head>
        {/*title information*/}
        <title>Takumi Ikeda&apos;s Portfolio</title>
        <meta name="description" content="Takumi Ikeda's Portfolio" />

        <meta charSet="UTF-8" />
        <meta
          name="viewport"
          content="width=device-width, initial-scale=1.0, minimum-scale=1.0"
        />

        {/*favicon settings*/}
        <link
          rel="apple-touch-icon"
          sizes="180x180"
          href={faviconPath + '/apple-touch-icon.png'}
        />
        <link
          rel="icon"
          type="image/png"
          sizes="32x32"
          href={faviconPath + '/favicon-32x32.png'}
        />
        <link
          rel="icon"
          type="image/png"
          sizes="16x16"
          href={faviconPath + '/favicon-16x16.png'}
        />
        <link rel="manifest" href={faviconPath + '/site.webmanifest'} />
        <link
          rel="mask-icon"
          href={faviconPath + '/safari-pinned-tab.svg'}
          color="#5bbad5"
        />
        <meta name="msapplication-TileColor" content="#da532c" />
        <meta name="theme-color" content="#ffffff" />
      </Head>
      <ThemeProvider theme={theme}>
        {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
        <CssBaseline />
        <Component {...pageProps} />
      </ThemeProvider>
    </CacheProvider>
  );
}
