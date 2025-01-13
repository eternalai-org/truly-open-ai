import React, { FunctionComponent } from 'react';

type Props = {
  data: React.ReactNode | FunctionComponent | string;
  size?: number;
};

function ImageRender({ data, size = 24 }: Props) {
  if (!data) {
    return <></>;
  }

  if (typeof data === 'string') {
    return (
      <img
        src={data}
        alt="Image"
        style={{
          width: `${size}px`,
          height: `${size}px`,
        }}
      />
    );
  }

  if (typeof data === 'function') {
    return data({});
  }

  return data;
}

export default ImageRender;
