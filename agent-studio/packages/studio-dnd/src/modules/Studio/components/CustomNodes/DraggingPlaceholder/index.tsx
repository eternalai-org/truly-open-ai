import LegoRender from '../NodeBase/LegoRender';

import useStudioDndStore from '@/modules/Studio/stores/useStudioDndStore';

function DraggingPlaceholder({ showDraggingFloating = false }: { showDraggingFloating?: boolean }) {
  const draggingElement = useStudioDndStore((state) => state.draggingElement);

  if (draggingElement) {
    return (
      <div
        style={{
          opacity: 0.2,
        }}
      >
        {showDraggingFloating ? (
          draggingElement
        ) : (
          <LegoRender
            background="#00000055"
            icon={null}
            title="Drop item at here"
            id="dragging-placeholder"
            schemaData={{}}
            idx=""
          />
        )}
      </div>
    );
  }

  return <></>;
}

export default DraggingPlaceholder;
