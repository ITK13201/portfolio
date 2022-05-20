import type { NextPage } from 'next';
import styles from 'assets/styles/pages/Home.module.scss';

// component imports
import Header from 'components/home/Header';
import EyeCatch from 'components/home/EyeCatch';
import About from 'components/home/About';
import Skills from 'components/home/Skills';
import Works from 'components/home/Works';
import Contact from 'components/home/Contact';
import ScrollUp from 'components/design/ScrollUp';

const Home: NextPage = () => {
  const faviconPath = '/images/favicon';

  return (
    <div className={styles.container}>
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

      <ScrollUp />
    </div>
  );
};

export default Home;
