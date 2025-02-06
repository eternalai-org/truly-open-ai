import cx from 'clsx';
import { CSSProperties, FunctionComponent, HTMLAttributes, memo, useMemo } from 'react';

import { adjustColorShade } from '../../utils/ui';
import './Lego.scss';
import { TooltipIcon } from '../icons/common';
import { StudIcon } from '../icons/lego';
import ImageRender from '../Render/ImageRender';
import Tooltip from '../ui/Tooltip';

type Props = HTMLAttributes<HTMLDivElement> & {
  background?: string; // HEX color
  disabled?: boolean;

  icon?: React.ReactNode | FunctionComponent;
  actions?: React.ReactNode;

  fixedHeight?: boolean;

  title?: React.ReactNode | FunctionComponent;
  tooltip?: string;
};

const Lego = ({
  fixedHeight = true,
  background = '#CC6234',
  disabled,
  className,
  style,
  icon,
  children,
  tooltip,
  ...props
}: Props) => {
  const borderColor = useMemo(() => adjustColorShade(background, -20), [background]);

  return (
    <div
      {...props}
      className={cx('lego', className, {
        'lego--disabled': disabled,
        'lego--dynamic': !fixedHeight,
      })}
      style={
        {
          '--border-color': borderColor,
          '--background-color': background,
          ...style,
        } as CSSProperties
      }
    >
      <div className="lego__stud">
        <StudIcon />
      </div>
      <div className="lego__stud lego__stud--bottom">
        <StudIcon />

        <div className="lego__stud__mask" />
        <div className="lego__stud__mask" />
      </div>

      <div className="lego__body">
        {icon && (
          <div className="lego__icon">
            <ImageRender data={icon} size={20} />
          </div>
        )}

        <div className="lego__content">{children}</div>
        {tooltip && (
          <Tooltip label={tooltip}>
            <TooltipIcon />
          </Tooltip>
        )}
      </div>
    </div>
  );
};

export default memo(Lego);
