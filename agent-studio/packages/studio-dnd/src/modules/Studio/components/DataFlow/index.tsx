import { useEffect, useRef } from 'react';

import useStudioCategoryStore from '../../stores/useStudioCategoryStore';
import useStudioDataStore from '../../stores/useStudioDataStore';
import useStudioDndStore from '../../stores/useStudioDndStore';
import useStudioFlowStore from '../../stores/useStudioFlowStore';
import useStudioFormStore from '../../stores/useStudioFormStore';
import { GraphData, StudioCategoryOptionMapValue, StudioDataNode, StudioNode } from '../../types';
import { createNodeData } from '../../utils/data';

import { useThrottleValue } from '@/hooks/useThrottleValue';
import useStudioFlowViewStore from '@/modules/Studio/stores/useStudioFlowViewStore';

type Props = {
  throttleNodesDelay: number;
  throttleDataDelay: number;
  throttleViewDelay: number;
  onChange?: (graphData: GraphData) => void;
};

function Listen({ throttleNodesDelay, throttleDataDelay, throttleViewDelay }: Props) {
  const nodes = useStudioFlowStore((state) => state.nodes);
  const formMap = useStudioFormStore((state) => state.formMap);
  const draggingData = useStudioDndStore((state) => state.draggingData);
  const view = useStudioFlowViewStore((state) => state.view);

  const isDragging = !!draggingData;

  const throttleNodes = useThrottleValue(nodes, throttleNodesDelay);
  const throttleDataForms = useThrottleValue(formMap, throttleDataDelay);
  const throttleView = useThrottleValue(view, throttleViewDelay);

  useEffect(() => {
    // sync nodes with data

    if (!isDragging) {
      const linkedNodes = useStudioFlowStore.getState().linkedNodes;
      const usedKeyCollection: Record<string, string> = {};
      const categoryOptionMap = useStudioCategoryStore.getState().categoryOptionMap;

      const getChildrenDataFromChildren = (children: StudioNode[]): StudioDataNode[] => {
        return children
          .map((child) => {
            const id = child.data.id;
            const metadata = child.data.metadata;
            const idx = child.data.metadata.idx;
            const option: StudioCategoryOptionMapValue | undefined = categoryOptionMap[idx];

            if (metadata && option) {
              const directlyChildren = getChildrenDataFromChildren(metadata?.children);
              const formValue = throttleDataForms[id] || {};

              if (option?.parent?.idx) {
                usedKeyCollection[option.parent.idx] = option.parent.idx;
              }

              usedKeyCollection[option.idx] = option.idx;

              return createNodeData(id, option, directlyChildren, formValue, child.position, option?.parent?.idx);
            }

            return null;
          })
          .filter((item) => !!item) as StudioDataNode[];
      };

      const getDataFromLinkedNodes = (node: StudioNode): StudioDataNode[] => {
        let linkedNodeChildren: StudioNode[] = [];
        const linkedChildrenIds = linkedNodes[node.id] || [];
        if (linkedChildrenIds.length) {
          linkedNodeChildren = linkedChildrenIds
            .map((linkedId) => throttleNodes.find((node) => node.id === linkedId))
            .filter((item) => !!item);
          if (linkedNodeChildren.length) {
            const dataFromLinkedChildren = getChildrenDataFromChildren(linkedNodeChildren);
            dataFromLinkedChildren.forEach((item) => {
              if (linkedNodes[item.id]?.length) {
                const linkedNode = linkedNodeChildren.find((linkedNode) => item.id === linkedNode.id);
                if (linkedNode) {
                  item.children = getDataFromLinkedNodes(linkedNode);
                }
              }
            });

            return dataFromLinkedChildren;
          }
        }

        return [];
      };

      const newData: StudioDataNode[] = [];

      const nodeIdsShouldIgnore = Object.values(linkedNodes)
        .map((ids) => ids)
        .flat();

      const filterNodes = throttleNodes.filter((node) => !nodeIdsShouldIgnore.includes(node.id));

      filterNodes.forEach((node) => {
        const metadata = node.data.metadata;
        const id = node.data.id;
        const idx = node.data.metadata.idx;
        const option: StudioCategoryOptionMapValue | undefined = categoryOptionMap[idx];

        if (metadata && option) {
          const directlyChildren = getChildrenDataFromChildren(metadata?.children);
          const inDirectlyChildren = getDataFromLinkedNodes(node);
          const formValue = throttleDataForms[id] || {};

          if (option?.parent?.idx) {
            usedKeyCollection[option.parent.idx] = option.parent.idx;
          }

          usedKeyCollection[option.idx] = option.idx;
          newData.push(
            createNodeData(
              id,
              option,
              [...directlyChildren, ...inDirectlyChildren],
              formValue,
              node.position,
              option?.parent?.idx,
            ),
          );
        }
      });

      useStudioDataStore.getState().setData(newData);
      useStudioCategoryStore.getState().setUsedKeyCollection(usedKeyCollection);
    }
  }, [throttleNodes, throttleDataForms, isDragging]);

  useEffect(() => {
    if (!isDragging) {
      useStudioDataStore.getState().setViewport(throttleView);
    }
  }, [throttleView, isDragging]);

  return <></>;
}

function Publish({ onChange }: { onChange?: (graphData: GraphData) => void }) {
  const data = useStudioDataStore((state) => state.data);
  const viewport = useStudioDataStore((state) => state.viewport);

  const onChangeRef = useRef(onChange);
  const draggingData = useStudioDndStore((state) => state.draggingData);

  const isDragging = !!draggingData;

  useEffect(() => {
    onChangeRef.current = onChange;
  }, [onChange]);

  const viewportDelay = useThrottleValue(viewport, 500);

  useEffect(() => {
    if (!isDragging) {
      if (onChangeRef.current) {
        onChangeRef.current({
          data,
          viewport: viewportDelay,
        } satisfies GraphData);
      }
    }
  }, [data, viewportDelay, isDragging]);

  return <></>;
}

function DataSync() {
  const entry = useStudioDataStore((state) => state.entry);
  const data = useStudioDataStore((state) => state.data);
  const rootCategory = useStudioCategoryStore((state) => state.rootCategory);
  const draggingData = useStudioDndStore((state) => state.draggingData);

  const isDragging = !!draggingData;

  useEffect(() => {
    if (!isDragging) {
      if (!entry) {
        if (rootCategory) {
          const rootOptions = rootCategory.options as StudioCategoryOptionMapValue[];
          const rootOptionsKey = rootOptions.map((item) => item.idx);
          const newEntry = data?.find((item) => item.idx === rootCategory.idx || rootOptionsKey.includes(item.idx));

          if (newEntry) {
            // set entry
            useStudioDataStore.getState().setEntry(newEntry);
            useStudioCategoryStore.getState().updateCategoriesForEntry(newEntry);
          } else {
            useStudioCategoryStore.getState().updateCategoriesForEntry(null);
          }
        }
      } else {
        const existEntry = data.find((item) => item.id === entry.id);
        if (!existEntry) {
          // remove entry
          useStudioDataStore.getState().setEntry(null);
          useStudioCategoryStore.getState().updateCategoriesForEntry(null);
        }
      }
    }
    // useStudioCategoryStore.getState().updateCategoriesForEntry(entry);
  }, [entry, data, rootCategory, isDragging]);

  return <></>;
}

function DataFlow({ throttleNodesDelay, throttleDataDelay, throttleViewDelay, onChange }: Props) {
  return (
    <>
      <Listen
        throttleNodesDelay={throttleNodesDelay}
        throttleDataDelay={throttleDataDelay}
        throttleViewDelay={throttleViewDelay}
      />
      <Publish onChange={onChange} />
      <DataSync />
    </>
  );
}

export default DataFlow;
