import type { NextPage } from 'next';
import Head from 'next/head';
import styles from 'assets/styles/pages/Home.module.scss';

// component imports
import Header from 'components/home/Header';
import EyeCatch from 'components/home/EyeCatch';
import About from 'components/home/About';
import Skills from 'components/home/Skills';
import Works from 'components/home/Works';
import Contact from 'components/home/Contact';

const Home: NextPage = () => {
  const faviconPath = '/images/favicon';

  return (
    <div className={styles.container}>
      <Head>
        <meta charSet="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        {/*title information*/}
        <title>Takumi Ikeda&apos;s Portfolio</title>
        <meta name="description" content="Takumi Ikeda's Portfolio" />
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

      <Header />

      <main className={styles.main}>
        <EyeCatch />

        <section id="about">
          <About />
        </section>

        <section id="skills">
          <Skills />
        </section>

        <section id="works">
          <Works />
        </section>
      </main>

      <footer>
        <section id="contact">
          <Contact />
        </section>
        <div className={styles.copyright}>
          <small>
            Copyright &copy; 2022 Takumi Ikeda. All Rights Reserved.
          </small>
        </div>
      </footer>
    </div>
  );
};

export default Home;
