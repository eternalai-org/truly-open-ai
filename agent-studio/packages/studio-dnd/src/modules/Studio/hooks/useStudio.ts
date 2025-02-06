/**
 * The hook to export the function can integrate/manipulation to studio
 */

import { useCallback, useMemo } from 'react';

import { useMultipleStore } from './useNewStore';
import useStudioCategoryStore from '../stores/useStudioCategoryStore';
import useStudioDataStore from '../stores/useStudioDataStore';
import useStudioDndStore from '../stores/useStudioDndStore';
import useStudioFlowStore from '../stores/useStudioFlowStore';
import useStudioFlowViewStore from '../stores/useStudioFlowViewStore';
import useStudioFormStore from '../stores/useStudioFormStore';
import { GraphData } from '../types/graph';
import { getFieldDataFromRawData } from '../utils/data';
import { transformDataToNodes } from '../utils/node';

import useStudioConfigStore from '@/modules/Studio/stores/useStudioConfigStore';

export const useStudio = () => {
  const cleanup = useCallback(() => {
    useStudioFlowStore.getState().clear();
    useStudioFormStore.getState().clear();
    useStudioDataStore.getState().clear();
    useStudioFlowViewStore.getState().clear();
    useStudioDndStore.getState().clear();
    useMultipleStore.getState().clear();
    // useStudioCategoryStore.getState().clear();
  }, []);

  const redraw = useCallback((graphData: GraphData) => {
    // scan all categories and push if not exist
    useStudioCategoryStore.getState().scanFromData(graphData.data);
    useStudioDataStore.getState().setData(graphData.data);
    useStudioDataStore.getState().setViewport(graphData.viewport);
    useStudioFlowViewStore.getState().setView(graphData.viewport);

    const initNodes = transformDataToNodes(graphData.data);
    useStudioFlowStore.getState().addNodes(initNodes);

    const formData = getFieldDataFromRawData(graphData.data);
    useStudioFormStore.getState().initDataForms(formData);

    useStudioFlowStore.getState().reloadFlow();
  }, []);

  const getFormDataById = useCallback(<T>(id: string): T => {
    return (useStudioFormStore.getState().getFormById(id) || {}) as T;
  }, []);

  const setFormFields = useCallback(<T>(id: string, fields: Partial<T> & Record<string, unknown>) => {
    return useStudioFormStore.getState().setFormFields(id, fields);
  }, []);

  const draggingData = useStudioDndStore((state) => state.draggingData);

  const isDragging = !!draggingData;

  const enableZoom = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledZoom: false,
      },
    });
  }, []);

  const disableZoom = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledZoom: true,
      },
    });
  }, []);

  const disableDrag = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledDrag: true,
      },
    });
  }, []);

  const enableDrag = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledDrag: false,
      },
    });
  }, []);

  const enableBackground = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledBackground: false,
      },
    });
  }, []);

  const disableBackground = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledBackground: true,
      },
    });
  }, []);

  const enableMiniMap = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledMiniMap: false,
      },
    });
  }, []);

  const disableMiniMap = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledMiniMap: true,
      },
    });
  }, []);

  const enableControls = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledControls: false,
      },
    });
  }, []);

  const disableControls = useCallback(() => {
    return useStudioConfigStore.getState().setConfig({
      ...useStudioConfigStore.getState().config,
      board: {
        ...useStudioConfigStore.getState().config.board,
        disabledControls: true,
      },
    });
  }, []);

  const data = useStudioDataStore((state) => state.data);
  const memorizedValue = useMemo(() => {
    return {
      cleanup,
      redraw,
      getFormDataById,
      setFormFields,
      data,
      isDragging,
      enableZoom,
      disableZoom,
      enableDrag,
      disableDrag,
      enableBackground,
      disableBackground,
      enableMiniMap,
      disableMiniMap,
      enableControls,
      disableControls,
    };
  }, [
    cleanup,
    redraw,
    getFormDataById,
    setFormFields,
    data,
    isDragging,
    enableZoom,
    disableZoom,
    enableDrag,
    disableDrag,
    enableBackground,
    disableBackground,
    enableMiniMap,
    disableMiniMap,
    enableControls,
    disableControls,
  ]);

  return memorizedValue;
};
