import cx from 'clsx';
import { CSSProperties, useMemo } from 'react';

import ImageRender from '../../Render/ImageRender';
import TextRender from '../../Render/TextRender';
import Tooltip from '../../ui/Tooltip';

import { TabBehavior } from '@/modules/Studio/enums/tab';
import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import useStudioConfigStore from '@/modules/Studio/stores/useStudioConfigStore';
import { StudioCategory } from '@/modules/Studio/types/category';
import { adjustColorShade } from '@/modules/Studio/utils/ui';

import './SidebarTab.scss';

type Props = StudioCategory;

const SidebarTab = (props: Props) => {
  const { idx, icon, title, color = '#CC6234', customizeRenderOnTab, tooltip } = props;

  const tabBehavior = useStudioConfigStore((state) => state.config.tab.behavior);

  const { filters, setFilters } = useStudioCategoryStore();

  const isActive = useMemo(() => {
    return filters.includes(idx) || filters.length === 0;
  }, [filters, idx]);

  const handleClick = () => {
    if (tabBehavior === TabBehavior.FILTER) {
      setFilters(idx);
    } else if (tabBehavior === TabBehavior.SCROLL) {
      const groups = document.getElementById('sidebar-groups');
      const groupElement = document.getElementById(`category-group-${idx}`);
      const groupElementTop = groupElement?.offsetTop;

      if (groups && groupElementTop !== undefined) {
        groups.scroll({ behavior: 'smooth', top: groupElementTop });
      }
    }
  };

  if (customizeRenderOnTab && typeof customizeRenderOnTab === 'function') {
    return customizeRenderOnTab(props);
  }

  return (
    <Tooltip label={tooltip}>
      <div
        className={cx('sidebar-tab', { 'sidebar-tab--active': isActive })}
        onClick={handleClick}
        style={
          {
            '--color': color,
            '--border-color': adjustColorShade(color, -20),
          } as CSSProperties
        }
      >
        <ImageRender data={icon} />

        <span className="sidebar-tab__title">
          <TextRender data={title} />
        </span>
      </div>
    </Tooltip>
  );
};

export default SidebarTab;
