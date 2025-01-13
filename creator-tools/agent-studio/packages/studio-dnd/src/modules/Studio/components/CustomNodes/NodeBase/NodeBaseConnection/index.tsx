import { Handle, Position } from '@xyflow/react';
import './NodeBaseConnection.scss';

const NodeBaseConnection = () => {
  return (
    <>
      <Handle
        id="a"
        type="source"
        position={Position.Top}
        className="node-base__handles__handle"
        isConnectable={false}
      />
      <Handle
        id="b"
        type="source"
        position={Position.Right}
        className="node-base__handles__handle"
        isConnectable={false}
      />
      <Handle
        id="c"
        type="source"
        position={Position.Bottom}
        className="node-base__handles__handle"
        isConnectable={false}
      />
      <Handle
        id="d"
        type="source"
        position={Position.Left}
        className="node-base__handles__handle"
        isConnectable={false}
      />
    </>
  );
};

export default NodeBaseConnection;
