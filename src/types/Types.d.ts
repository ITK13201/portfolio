// Type define
declare module 'react-anchor-link-smooth-scroll' {
  interface Props {
    href: string;
    offset?: function | number;
    onClink?: (e: Event) => void;
    [key: string]: any;
  }

  export default class AnchorLink extends React.Component<Props> {}
}

declare module '*.css' {
  const styles: { [className: string]: string };
  export default styles;
}

declare module '*.scss' {
  const styles: { [className: string]: string };
  export default styles;
}
