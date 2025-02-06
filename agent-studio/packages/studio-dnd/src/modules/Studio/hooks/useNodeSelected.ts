import { useOnSelectionChange } from '@xyflow/react';
import { useState } from 'react';

type Props = {
  id: string;
};

const useNodeSelected = ({ id }: Props) => {
  const [isSelected, setIsSelected] = useState(false);

  useOnSelectionChange({
    onChange: ({ nodes }) => {
      const isSelectedThisNode = nodes.some((node) => node.id === id);

      setIsSelected(isSelectedThisNode);
    },
  });

  return { isSelected, setIsSelected };
};

export default useNodeSelected;
