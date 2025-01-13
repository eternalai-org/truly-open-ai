import cx from 'clsx';
import { HTMLAttributes } from 'react';
import './LegoContent.scss';

type Props = HTMLAttributes<HTMLDivElement>;

const LegoContent = ({ children, className, ...props }: Props) => {
  return (
    <div className={cx('lego-content', className)} {...props}>
      {children}
    </div>
  );
};

export default LegoContent;
