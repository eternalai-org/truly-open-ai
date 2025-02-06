import { useMemo } from 'react';

import Source from '../../DnD/Source';
import Lego from '../../Lego';
import LegoContent from '../../LegoContent';
import TextRender from '../../Render/TextRender';
import Tooltip from '../../ui/Tooltip';

import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import { StudioCategory, StudioCategoryOption } from '@/modules/Studio/types';

import './CategoryGroup.scss';

type Props = StudioCategory;

const CategoryOption = ({
  categoryKey,
  isRoot,
  color,
  option,
  disabled,
  multipleOption,
  tooltip,
}: {
  categoryKey: string;
  isRoot?: boolean;
  color?: string;
  option: StudioCategoryOption;
  disabled?: boolean;
  multipleOption?: boolean;
  tooltip?: string;
}) => {
  const usedKeyCollection = useStudioCategoryStore((state) => state.usedKeyCollection);
  const usedCategoryKey = usedKeyCollection[categoryKey];
  const usedOptionKey = usedKeyCollection[option.idx];

  const isDisabled = useMemo(() => {
    let returnVal = disabled || option.disabled;

    // check parent first
    if (!returnVal) {
      if (!multipleOption) {
        if (usedCategoryKey) {
          returnVal = true;
        }
      }
    }

    if (!returnVal) {
      if (!option.multipleChoice) {
        if (usedOptionKey) {
          returnVal = true;
        }
      }
    }

    return returnVal;
  }, [disabled, multipleOption, option.disabled, option.multipleChoice, usedCategoryKey, usedOptionKey]);

  return (
    <Source
      id={option.idx}
      key={`sidebar-source-${categoryKey}-${option.idx}`}
      data={{ categoryKey, optionKey: option.idx, isRoot }}
      disabled={isDisabled}
    >
      <Lego background={color} icon={option.icon} disabled={isDisabled} tooltip={tooltip}>
        <LegoContent>
          <TextRender data={option.title} />
        </LegoContent>
      </Lego>
    </Source>
  );
};

const CustomCategoryOption = ({
  categoryKey,
  isRoot,
  color,
  option,
  disabled,
  multipleOption,
  children,
  tooltip,
}: {
  categoryKey: string;
  isRoot?: boolean;
  color?: string;
  option: StudioCategoryOption;
  disabled?: boolean;
  multipleOption?: boolean;
  children: React.ReactNode;
  tooltip?: string;
}) => {
  const usedKeyCollection = useStudioCategoryStore((state) => state.usedKeyCollection);
  const usedCategoryKey = usedKeyCollection[categoryKey];
  const usedOptionKey = usedKeyCollection[option.idx];

  const isDisabled = useMemo(() => {
    let returnVal = disabled || option.disabled;

    // check parent first
    if (!returnVal) {
      if (!multipleOption) {
        if (usedCategoryKey) {
          returnVal = true;
        }
      }
    }

    if (!returnVal) {
      if (!option.multipleChoice) {
        if (usedOptionKey) {
          returnVal = true;
        }
      }
    }

    return returnVal;
  }, [disabled, multipleOption, option.disabled, option.multipleChoice, usedCategoryKey, usedOptionKey]);

  return (
    <Source
      id={option.idx}
      key={`sidebar-source-${categoryKey}-${option.idx}`}
      data={{ categoryKey, optionKey: option.idx, isRoot }}
      disabled={isDisabled}
    >
      <Lego background={color} disabled={isDisabled} tooltip={tooltip}>
        {children}
      </Lego>
    </Source>
  );
};

const CategoryGroup = (props: Props) => {
  const {
    idx: categoryKey,
    title,
    color,
    options,
    required,
    disabled,
    isRoot,
    multipleOption,
    tooltip,
    customizeRenderOnSidebar,
  } = props;

  const filteredOptions = useMemo(() => {
    return options.filter((item) => !item.hidden);
  }, [options]);

  if (customizeRenderOnSidebar && typeof customizeRenderOnSidebar === 'function') {
    return customizeRenderOnSidebar(props);
  }

  return (
    <div className="category-group" id={`category-group-${categoryKey}`}>
      {tooltip ? (
        <Tooltip label={tooltip}>
          <h5 className="category-group__title">
            <TextRender data={title} /> {!!required && <span className="sidebar-tab__required">*</span>}
          </h5>
        </Tooltip>
      ) : (
        <h5 className="category-group__title">
          <TextRender data={title} /> {!!required && <span className="sidebar-tab__required">*</span>}
        </h5>
      )}

      <div className="category-group__options">
        {filteredOptions.map((option) => {
          if (option.customizeRenderOnSideBar && typeof option.customizeRenderOnSideBar === 'function') {
            return (
              <CustomCategoryOption
                key={`sidebar-source-${categoryKey}-${option.idx}`}
                categoryKey={categoryKey}
                isRoot={isRoot}
                color={color}
                option={option}
                disabled={disabled}
                multipleOption={multipleOption}
                tooltip={option.tooltip}
              >
                {option.customizeRenderOnSideBar(option)}
              </CustomCategoryOption>
            );
          }

          return (
            <CategoryOption
              key={`sidebar-source-${categoryKey}-${option.idx}`}
              categoryKey={categoryKey}
              isRoot={isRoot}
              color={color}
              option={option}
              disabled={disabled}
              multipleOption={multipleOption}
              tooltip={option.tooltip}
            />
          );
        })}
      </div>
    </div>
  );
};

export default CategoryGroup;
