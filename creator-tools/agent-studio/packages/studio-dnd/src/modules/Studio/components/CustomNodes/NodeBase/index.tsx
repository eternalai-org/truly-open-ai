import NodeSingle from './NodeSingle';
import NodeStacks from './NodeStacks';
import { NodeBaseProps } from './types';
import './NodeBase.scss';

type Props = NodeBaseProps;

export default function NodeBase(props: Props) {
  const { data } = props;
  const children = data?.metadata?.children;

  if (children.length) {
    return <NodeStacks {...props} />;
  }

  return <NodeSingle {...props} />;
}
