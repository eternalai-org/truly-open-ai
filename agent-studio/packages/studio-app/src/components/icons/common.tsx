import { SVGProps } from "react";

export const InfoIcon = (props: SVGProps<SVGSVGElement>) => (
  <svg width={20} height={20} fill="none" {...props}>
    <path
      fill="#fff"
      d="M10 1.25A8.745 8.745 0 0 0 1.25 10 8.745 8.745 0 0 0 10 18.75 8.745 8.745 0 0 0 18.75 10 8.745 8.745 0 0 0 10 1.25Zm0 16.28c-4.152 0-7.53-3.378-7.53-7.53S5.849 2.47 10 2.47s7.53 3.378 7.53 7.53-3.379 7.53-7.53 7.53Z"
    />
    <path
      fill="#fff"
      d="M10.002 8.344c-.593 0-1.014.25-1.014.618v5.015c0 .316.421.632 1.014.632.566 0 1.026-.316 1.026-.632V8.962c0-.368-.46-.618-1.026-.618ZM10.001 5.242c-.605 0-1.08.435-1.08.935s.475.948 1.08.948c.593 0 1.067-.448 1.067-.948 0-.5-.474-.935-1.067-.935Z"
    />
  </svg>
);

export const categoryImageIcon = (src: string) => {
  return () => {
    return (
      <img
        src={src}
        alt="Lego Component Icon"
        style={{
          width: 24,
          height: 24,
          borderRadius: 4,
        }}
      />
    );
  };
};
