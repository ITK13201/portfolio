import { CustomMaterialTheme } from 'hooks/theme';
import {
  Avatar,
  Box,
  ThemeProvider,
  Typography,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  ListItemProps,
  Card,
} from '@mui/material';
import { Folder as FolderIcon } from '@mui/icons-material';
import styles from 'assets/styles/components/home/Works.module.scss';
import React from 'react';

// Programming Language image files imports
import JavaImage from 'assets/images/languages/java.png';
import PythonImage from 'assets/images/languages/python.png';
import TypescriptImage from 'assets/images/languages/typescript.png';
import GolangImage from 'assets/images/languages/golang.png';

function ListItemLink(props: ListItemProps<'a', { button?: true }>) {
  return <ListItem button component="a" {...props} />;
}

interface worksElement {
  href: string;
  title: string;
  description: string;
  img: StaticImageData;
}

const worksElements: worksElement[] = [
  {
    href: 'https://github.com/MISW/Museum',
    title: 'MISW Museum',
    description:
      'サークル内成果物（ゲーム，動画，イラスト，アプリケーションなど）を投稿するためのWebアプリケーションです．Djangoを用いて5名の共同開発で作成しました．',
    img: PythonImage,
  },
  {
    href: 'https://github.com/ITK13201/private_diary',
    title: 'Private Diary',
    description:
      'Web上で管理できる個人日記サービスです．Djangoを用いて個人開発で作成しました．',
    img: PythonImage,
  },
  {
    href: 'https://github.com/ITK13201/ControlPDF',
    title: 'Control PDF',
    description:
      'PDFファイルを分割/結合できるツールです．Pythonを用いて個人開発で作成しました．',
    img: PythonImage,
  },
  {
    href: '',
    title: 'Portfolio',
    description:
      '現在ご覧のこのポートフォリオです．Next.jsを使って作成しました．',
    img: TypescriptImage,
  },
  {
    href: 'https://github.com/ITK13201/waseda-moodle-scheduler',
    title: 'Waseda Moodle Scheduler',
    description:
      '早稲田大学ではMoodleというWebサービスを授業支援として用いいており，その中に課題の締切などが登録されるカレンダー機能があります．そのカレンダー機能を利用してDiscordサーバに通知を送ることで締め切り釣果を防ぐWebAPIアプリケーションです．Djangoで作成しました．',
    img: PythonImage,
  },
  {
    href: 'https://github.com/ITK13201/TodoManagerProject',
    title: 'TODO Manager Project',
    description:
      'Socket通信を用いてTODO管理ができるアプリケーションです．Javaで作成しました．',
    img: JavaImage,
  },
  {
    href: 'https://github.com/ITK13201/smart-watch-iot-server',
    title: 'Smart Watch IoT Project',
    description:
      'Smartwatchで測定した心拍数に応じで最適な音楽をGoogle Homeから流すIoTアプリケーションです．Django, Flaskを用いて個人開発で作成しました．このプロジェクトは3つのレポジトリに分かれております．このリンク先(GitHub)のREADMEにそれぞれのレポジトリURLが記載されているのでご参照ください．',
    img: PythonImage,
  },
  {
    href: 'https://github.com/ITK13201/holodule-bot',
    title: 'Holodule Bot',
    description:
      'ホロライブというVirtual YouTuber事務所の配信スケジュールをWebサイトからスクレイピングにより取得し，Dsicordサーバーに通知する趣味用アプリケーションです．Go言語で作成しました．',
    img: GolangImage,
  },
];

const Works = (): JSX.Element => {
  return (
    <ThemeProvider theme={CustomMaterialTheme}>
      <Box className={styles.root}>
        <Box className={styles.caption}>
          <Typography variant="h2" color="initial">
            <strong>Works</strong>
          </Typography>
        </Box>
        <Box className={styles.container}>
          <List className={styles.list}>
            {worksElements.map((element: worksElement) => (
              <ListItem key={element.title}>
                <Card className={styles.contentCard}>
                  <ListItemLink
                    href={element.href}
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    <ListItemAvatar>
                      <Avatar src={element.img.src} />
                    </ListItemAvatar>
                    <ListItemText
                      primary={element.title}
                      secondary={element.description}
                    />
                  </ListItemLink>
                </Card>
              </ListItem>
            ))}
          </List>
        </Box>
      </Box>
    </ThemeProvider>
  );
};

export default Works;
