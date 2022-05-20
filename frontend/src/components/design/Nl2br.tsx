import React from 'react';

interface Nl2brProps {
  children: string;
}

const Nl2br = (props: Nl2brProps): JSX.Element => {
  const { children } = props;
  return (
    <>
      {children.split('\n').map((text: string, index: number) => (
        <React.Fragment key={index}>
          {text}
          <br />
        </React.Fragment>
      ))}
    </>
  );
};

export default Nl2br;
