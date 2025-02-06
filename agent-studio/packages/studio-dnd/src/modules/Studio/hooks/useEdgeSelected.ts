import { useOnSelectionChange } from '@xyflow/react';
import { useState } from 'react';

type Props = {
  id: string;
};

const useEdgeSelected = ({ id }: Props) => {
  const [isSelected, setIsSelected] = useState(false);

  useOnSelectionChange({
    onChange: ({ edges }) => {
      const isSelectedThisEdge = edges.some((edge) => edge.id === id);

      setIsSelected(isSelectedThisEdge);
    },
  });

  return { isSelected, setIsSelected };
};

export default useEdgeSelected;
