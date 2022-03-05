import type { NextPage } from 'next';
import Head from 'next/head';
import styles from 'assets/styles/pages/Home.module.scss';

// component imports
import Header from 'components/home/Header';
import EyeCatch from 'components/home/EyeCatch';
import About from 'components/home/About';
import Contact from 'components/home/Contact';
import Skills from 'components/home/Skills';

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

        <p className={styles.description}>
          Get started by editing{' '}
          <code className={styles.code}>pages/index.tsx</code>
        </p>

        <div className={styles.grid}>
          <a href="https://nextjs.org/docs" className={styles.card}>
            <h2>Documentation &rarr;</h2>
            <p>Find in-depth information about Next.js features and API.</p>
          </a>

          <a href="https://nextjs.org/learn" className={styles.card}>
            <h2>Learn &rarr;</h2>
            <p>Learn about Next.js in an interactive course with quizzes!</p>
          </a>

          <a
            href="https://github.com/vercel/next.js/tree/canary/examples"
            className={styles.card}
          >
            <h2>Examples &rarr;</h2>
            <p>Discover and deploy boilerplate example Next.js projects.</p>
          </a>

          <a
            href="https://vercel.com/new?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
            className={styles.card}
          >
            <h2>Deploy &rarr;</h2>
            <p>
              Instantly deploy your Next.js site to a public URL with Vercel.
            </p>
          </a>
        </div>
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
