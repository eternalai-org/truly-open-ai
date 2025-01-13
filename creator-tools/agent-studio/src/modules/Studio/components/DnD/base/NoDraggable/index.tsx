import cx from 'clsx';

import './NoDraggable.scss';

function NoDraggable({ children, className }: { children: React.ReactNode; className?: string }) {
  return (
    <div
      className={cx('studio-no-draggable', className)}
      onPointerDown={(e) => e.stopPropagation()}
      onMouseDown={(e) => e.stopPropagation()}
      onMouseMove={(e) => e.stopPropagation()}
      onClick={(e) => e.stopPropagation()}
    >
      {children}
    </div>
  );
}

export default NoDraggable;
