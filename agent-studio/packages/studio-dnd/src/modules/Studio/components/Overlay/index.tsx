import cx from 'clsx';
import { memo, PropsWithChildren } from 'react';
import './Overlay.scss';

type Props = PropsWithChildren & {
  active: boolean;
};

const Overlay = ({ active, children }: Props) => {
  return (
    <div className={cx('overlay', { 'overlay--active': active })}>
      <div className="overlay__background" />
      <div className="overlay__content">{children}</div>
    </div>
  );
};

export default memo(Overlay);
