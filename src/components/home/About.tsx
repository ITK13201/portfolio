import React from 'react';
import { Avatar, Box, Card, CardContent, Typography } from '@mui/material';
import {
  EmojiEvents as EmojiEventsIcon,
  Engineering as EngineeringIcon,
  Interests as InterestsIcon,
  LocalActivity as LocalActivityIcon,
  School as SchoolIcon,
  Work as WorkIcon,
  SvgIconComponent,
} from '@mui/icons-material';
import styles from 'assets/styles/components/home/About.module.scss';
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
    description: `- 早稲田大学基幹理工学部情報理工学科（2019年〜2023年卒業）
        - 早稲田大学基幹理工学研究科情報理工・情報通信専攻（2023年〜2025年卒業）`,
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
    description: `- 普通自動車第一種運転免許（AT限定） 取得
        - 応用情報技術者試験 合格`,
  },
  {
    title: 'Career',
    titleIcon: WorkIcon,
    description: `2020/12 - present: エヌ次元株式会社
      2022/09, 3weeks: クックパッド株式会社`,
  },
];

const About = (): JSX.Element => {
  return (
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
  );
};

export default About;
