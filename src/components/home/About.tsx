import React from 'react';
import {
  Box,
  ThemeProvider,
  Typography,
  Card,
  CardContent,
  Avatar,
} from '@mui/material';
import {
  SvgIconComponent,
  School as SchoolIcon,
  Engineering as EngineeringIcon,
  LocalActivity as LocalActivityIcon,
  Interests as InterestsIcon,
  EmojiEvents as EmojiEventsIcon,
} from '@mui/icons-material';
import styles from 'assets/styles/components/home/About.module.scss';
import { CustomMaterialTheme } from '../../hooks/theme';
import Nl2br from 'components/design/Nl2br';

interface AboutElement {
  title: string;
  titleIcon: SvgIconComponent;
  description: string;
}

const AboutElements: AboutElement[] = [
  {
    title: 'Belong To',
    titleIcon: SchoolIcon,
    description: '早稲田大学基幹理工学部情報理工学科（2019〜2023卒業予定）',
  },
  {
    title: 'Wanna be',
    titleIcon: EngineeringIcon,
    description: 'Software and Web Engineer',
  },
  {
    title: 'Recent Activities',
    titleIcon: LocalActivityIcon,
    description:
      'Go言語を使ったWebAPIの開発とC++で競技プログラミングの勉強を主に行っています．',
  },
  {
    title: 'Interested in',
    titleIcon: InterestsIcon,
    description:
      'kubernetesを用いた本番環境の構築や，Go言語とNext.JSを連携させたWebアプリケーション開発に興味があります．',
  },
  {
    title: 'Awards and Certifications',
    titleIcon: EmojiEventsIcon,
    description: '普通自動車第一種運転免許（AT限定） 取得',
  },
];

const About = () => {
  return (
    <ThemeProvider theme={CustomMaterialTheme}>
      <Box className={styles.root}>
        <Box className={styles.caption}>
          <Typography variant="h2" color="initial">
            <strong>About</strong>
          </Typography>
        </Box>
        <Box className={styles.container}>
          {AboutElements.map((element: AboutElement) => (
            <Card className={styles.card} key={element.title}>
              <CardContent>
                <Box display="flex">
                  <Avatar>
                    <element.titleIcon />
                  </Avatar>
                  <Typography variant="h3" padding="5px">
                    {element.title}
                  </Typography>
                </Box>
                <p>
                  <Nl2br>{element.description}</Nl2br>
                </p>
              </CardContent>
            </Card>
          ))}
        </Box>
      </Box>
    </ThemeProvider>
  );
};

export default About;
