import '../../styles/global.scss';
import './Studio.scss';

import { ReactFlowProvider } from '@xyflow/react';
import cx from 'clsx';
import React, { useEffect, useImperativeHandle } from 'react';

import Board from './components/Board';
import DataFlow from './components/DataFlow';
import DragMask from './components/DnD/base/DragMask';
import DndFlow from './components/DnD/DndFlow';
import EventHandler from './components/EventHandler';
import Sidebar from './components/Sidebar';
import { MIN_THROTTLE_DATA_DELAY, MIN_THROTTLE_NODES_DELAY } from './constants/configs';
import { DEFAULT_THROTTLE_DATA_DELAY, DEFAULT_THROTTLE_NODES_DELAY } from './constants/default-values';
import { SidebarSide } from './enums/side';
import { useStudio } from './hooks/useStudio';
import useStudioCategoryStore from './stores/useStudioCategoryStore';
import useStudioConfigStore from './stores/useStudioConfigStore';
import useStudioDataSourceStore from './stores/useStudioDataSourceStore';
import useStudioDataStore from './stores/useStudioDataStore';
import { DataSource, FormDataMap, StudioCategory, StudioDataNode } from './types';
import { StudioConfig } from './types/config';
import { min } from './utils/data';

export type StudioProps = {
  className?: string;
  sidebarWidth?: string | number;

  // Data
  categories: StudioCategory[];
  dataSource?: Record<string, DataSource[]>;
  data: StudioDataNode[];

  throttleNodesDelay?: number;
  throttleDataDelay?: number;

  // Configs
  config?: StudioConfig;

  // Events
  onChange?: (data: StudioDataNode[]) => void;
};

export type StudioRef = {
  cleanup: () => void;
  redraw: (data: StudioDataNode[]) => void;
  getOptionPlaceQuantity: (optionId: string) => number;
  checkOptionIsPlaced: (optionId: string) => boolean;
  getFormDataById: (id: string) => FormDataMap | undefined;
};

const StudioComponent = ({
  className,
  categories,
  dataSource,
  data,
  throttleNodesDelay = DEFAULT_THROTTLE_NODES_DELAY,
  throttleDataDelay = DEFAULT_THROTTLE_DATA_DELAY,
  config,
  onChange,
  sidebarWidth = 400,
  ...rest
}: StudioProps): React.ReactNode => {
  const [isFirstDraw, setIsFirstDraw] = React.useState(false);
  const { redraw, cleanup } = useStudio();

  const sidebarSide = useStudioConfigStore((state) => state.config.sidebar.side);

  useEffect(() => {
    // eslint-disable-next-line @typescript-eslint/no-floating-promises
    useStudioCategoryStore.getState().setCategories(categories);
  }, [categories]);

  useEffect(() => {
    if (dataSource) {
      useStudioDataSourceStore.getState().setDataSource(dataSource);
    }
  }, [dataSource]);

  useEffect(() => {
    useStudioConfigStore.getState().setConfig(config);

    useStudioDataStore.getState().setDisabledConnection(!!config?.board.disabledConnection);
  }, [config]);

  // for mounted
  useEffect(() => {
    redraw([]);

    return () => {
      cleanup();
    };
  }, []);

  useEffect(() => {
    if (!isFirstDraw && data.length) {
      redraw(data);
      setIsFirstDraw(true);
    }
  }, [data.length, isFirstDraw]);

  return (
    <DndFlow>
      <DataFlow
        throttleNodesDelay={min(throttleNodesDelay, MIN_THROTTLE_NODES_DELAY)}
        throttleDataDelay={min(throttleDataDelay, MIN_THROTTLE_DATA_DELAY)}
        onChange={onChange}
      />

      <div
        className={cx('studio', sidebarSide === SidebarSide.LEFT ? 'studio--left' : 'studio--right', className)}
        {...rest}
      >
        <DragMask />

        <div
          className="studio__sidebar"
          style={{
            minWidth: 300,
            width: sidebarWidth,
          }}
        >
          <Sidebar />
        </div>

        <EventHandler className="studio__board">
          <Board />
        </EventHandler>
      </div>
    </DndFlow>
  );
};

export const Studio = React.forwardRef<StudioRef, StudioProps>((props: StudioProps, ref) => {
  const { redraw, cleanup, getOptionPlaceQuantity, checkOptionIsPlaced, getFormDataById } = useStudio();

  useImperativeHandle(
    ref,
    () => ({
      cleanup: () => {
        cleanup();
      },
      redraw: (data: StudioDataNode[]) => {
        redraw(data);
      },
      getOptionPlaceQuantity: (optionId: string) => {
        return getOptionPlaceQuantity(optionId);
      },
      checkOptionIsPlaced: (optionId: string) => {
        return checkOptionIsPlaced(optionId);
      },
      getFormDataById: (id: string) => {
        return getFormDataById(id);
      },
    }),
    [cleanup, redraw, getOptionPlaceQuantity, checkOptionIsPlaced, getFormDataById],
  );

  return (
    <ReactFlowProvider>
      <StudioComponent {...props} />
    </ReactFlowProvider>
  );
});
