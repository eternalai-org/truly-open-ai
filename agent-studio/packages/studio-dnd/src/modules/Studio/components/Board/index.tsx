import { Background, ConnectionMode, Controls, MiniMap, ReactFlow, Viewport } from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import { useEffect, useMemo, useState } from 'react';

import BoardOverlay from './BoardOverlay';
import { DEFAULT_EDGE_TYPES, DEFAULT_NODE_TYPES } from '../../constants/key-map';
import useStudioConfigStore from '../../stores/useStudioConfigStore';
import useStudioDataStore from '../../stores/useStudioDataStore';
import useStudioFlowStore from '../../stores/useStudioFlowStore';
import useStudioFlowViewStore from '../../stores/useStudioFlowViewStore';
import Distribution from '../DnD/Distribution';

import './Board.scss';

function Board() {
  const [isRendered, setIsRendered] = useState(false);

  const initialViewport = useStudioDataStore((state) => state.viewport);
  const boardConfig = useStudioConfigStore((state) => state.config.board);

  const disabledConnection = useStudioDataStore((state) => state.disabledConnection);
  const nodes = useStudioFlowStore((state) => state.nodes);
  const edges = useStudioFlowStore((state) => state.edges);

  const onNodesChange = useStudioFlowStore((state) => state.onNodesChange);
  const onEdgesChange = useStudioFlowStore((state) => state.onEdgesChange);

  const reloadFlowCounter = useStudioFlowStore((state) => state.reloadFlowCounter);
  const setView = useStudioFlowViewStore((state) => state.setView);

  const [currentView, setCurrentView] = useState<Viewport>({
    x: 0,
    y: 0,
    zoom: 1,
  });

  const renderEdges = useMemo(() => {
    if (!disabledConnection) {
      return edges;
    }

    return [];
  }, [edges, disabledConnection]);

  useEffect(() => {
    setCurrentView(useStudioFlowViewStore.getState().view);
  }, [reloadFlowCounter]);

  useEffect(() => {
    if (isRendered) return;

    setIsRendered(true);
    setCurrentView(initialViewport);

    useStudioFlowStore.getState().reloadFlow();
  }, [isRendered, initialViewport]);

  return (
    <Distribution className="board">
      <BoardOverlay />
      <ReactFlow
        key={JSON.stringify(currentView)}
        nodes={nodes}
        nodeTypes={DEFAULT_NODE_TYPES}
        onNodesChange={(changes) => {
          if (changes.some((change) => change.type === 'remove')) {
            return;
          }

          onNodesChange(changes);
        }}
        edges={renderEdges}
        edgeTypes={DEFAULT_EDGE_TYPES}
        onEdgesChange={onEdgesChange}
        onViewportChange={setView}
        defaultViewport={currentView}
        edgesFocusable={false}
        deleteKeyCode="Backspace"
        fitViewOptions={{ padding: 1 }}
        connectionMode={ConnectionMode.Loose}
        zoomOnDoubleClick={false}
        selectNodesOnDrag={false}
        disableKeyboardA11y
        {...boardConfig}
      >
        <Controls fitViewOptions={{ padding: 1 }} />
        <Background />
        <MiniMap />
      </ReactFlow>
    </Distribution>
  );
}

export default Board;
