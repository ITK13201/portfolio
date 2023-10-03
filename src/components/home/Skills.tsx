import React from 'react';
import { Box, Typography } from '@mui/material';
import styles from 'assets/styles/components/home/Skills.module.scss';
import { SkillBarSkill } from 'react-skillbars';
import { NoSsr } from '@mui/base';
import dynamic from 'next/dynamic';

const SkillBar = dynamic(() => import('react-skillbars'), { ssr: false });

const languageSkills: SkillBarSkill[] = [
  {
    type: 'Python',
    level: 70,
    color: {
      bar: '#85C1E9',
      title: {
        background: '#5DADE2',
      },
    },
  },
  {
    type: 'Go',
    level: 60,
    color: {
      bar: '#3498DB',
      title: {
        background: '#2E86C1',
      },
    },
  },
  {
    type: 'Ruby',
    level: 50,
    color: {
      bar: '#F266AB',
      title: {
        background: '#CC0000'
      }
    }
  },
  {
    type: 'Typescript',
    level: 50,
    color: {
      bar: '#17A589',
      title: {
        background: '#148F77',
      },
    },
  },
  {
    type: 'Java',
    level: 40,
    color: {
      bar: '#D35400',
      title: {
        background: '#BA4A00',
      },
    },
  },
  {
    type: 'Javascript',
    level: 30,
    color: {
      bar: '#F7DC6F',
      title: {
        background: '#F4D03F',
      },
    },
  },
  {
    type: 'C/C++',
    level: 20,
    color: {
      bar: '#EC7063',
      title: {
        background: '#E74C3C',
      },
    },
  },
];

const techSkills: SkillBarSkill[] = [
  {
    type: 'Django',
    level: 70,
  },
  {
    type: 'Rails',
    level: 50
  },
  {
    type: 'Next.js',
    level: 30,
  },
  {
    type: 'React',
    level: 40,
  },
  {
    type: 'Nginx',
    level: 50,
  },
  {
    type: 'Git',
    level: 90,
  },
  {
    type: 'Docker',
    level: 70,
  },
  {
    type: 'Kubernetes',
    level: 30,
  },
  {
    type: 'Terraform',
    level: 50,
  },
  {
    type: 'AWS',
    level: 60,
  },
  {
    type: 'GCP',
    level: 70
  },
  {
    type: 'Heroku',
    level: 70,
  },
];

const techColors = {
  bar: '#566573',
  title: {
    background: '#2C3E50',
  },
};

const dbSkills = [
  {
    type: 'MySQL',
    level: 60,
  },
  {
    type: 'MariaDB',
    level: 60,
  },
  {
    type: 'PostgreSQL',
    level: 50,
  },
];

const dbColors = {
  bar: '#A569BD',
  title: {
    background: '#8E44AD',
  },
};

const Skills = () => {
  return (
    <NoSsr>
      <Box className={styles.root}>
        <Box className={styles.caption}>
          <Typography variant="h2" color="initial">
            <strong>Skills</strong>
          </Typography>
        </Box>
        <Box className={styles.wrapper}>
          <Box className={styles.container}>
            <Typography variant="h3" className={styles.secondCaption}>
              Languages
            </Typography>
            <SkillBar skills={languageSkills}/>
          </Box>
          <Box className={styles.container}>
            <Typography variant="h3" className={styles.secondCaption}>
              Frameworks and Technologies
            </Typography>
            <SkillBar skills={techSkills} colors={techColors} />
          </Box>
          <Box className={styles.container}>
            <Typography variant="h3" className={styles.secondCaption}>
              Databases
            </Typography>
            <SkillBar skills={dbSkills} colors={dbColors} />
          </Box>
        </Box>
      </Box>
    </NoSsr>
  );
};

export default Skills;
