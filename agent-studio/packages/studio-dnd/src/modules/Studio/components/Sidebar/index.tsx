import cx from 'clsx';
import { memo, useMemo } from 'react';

import './Sidebar.scss';
import CategoryGroup from './CategoryGroup';
import SidebarOverlay from './SidebarOverlay';
import SidebarTab from './SidebarTab';
import { SidebarSide } from '../../enums/side';
import { TabBehavior } from '../../enums/tab';
import useStudioCategoryStore from '../../stores/useStudioCategoryStore';
import useStudioConfigStore from '../../stores/useStudioConfigStore';
import Factory from '../DnD/Factory';

const Sidebar = () => {
  const tabBehavior = useStudioConfigStore((state) => state.config.tab.behavior);
  const sidebarSide = useStudioConfigStore((state) => state.config.sidebar.side);

  const categories = useStudioCategoryStore((state) => state.categories);
  const filters = useStudioCategoryStore((state) => state.filters);

  const processedCategoryTabs = useMemo(() => {
    return categories.filter((item) => !item.hidden);
  }, [categories]);

  const processCategoryGroups = useMemo(() => {
    if (!filters.length || tabBehavior !== TabBehavior.FILTER) {
      return processedCategoryTabs;
    }

    return processedCategoryTabs.filter((item) => filters.includes(item.idx));
  }, [processedCategoryTabs, filters, tabBehavior]);

  return (
    <Factory className={cx('sidebar', sidebarSide === SidebarSide.LEFT ? 'sidebar--left' : 'sidebar--right')}>
      <div className="sidebar__tabs">
        <div className="sidebar__tabs__inner">
          {processedCategoryTabs.map((category) => (
            <SidebarTab {...category} key={`sidebar-tab-${category.idx}`} />
          ))}
        </div>
      </div>

      <div className="sidebar__groups">
        <SidebarOverlay />

        <div className="sidebar__groups__inner" id="sidebar-groups">
          {processCategoryGroups.map((category) => (
            <CategoryGroup {...category} key={`sidebar-group-${category.idx}`} />
          ))}
        </div>
      </div>
    </Factory>
  );
};

export default memo(Sidebar);
