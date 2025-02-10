import { Flex, SimpleGrid } from '@chakra-ui/react';

type Props = {
  dataList: any[];
  renderItem: (item: any, index: number) => React.ReactElement | undefined;
};

export const CollectionListView = (props: Props) => {
  const { dataList, renderItem } = props;

  return (
    <Flex
      flexDir={'column'}
      maxH={'500px'}
      p="20px 0px"
      w={'100%'}
      border={'1px solid #E5E7EB'}
      borderRadius={'4px'}
      bgColor={'#fff'}
      overflow={'hidden'}
    >
      <SimpleGrid
        flex={1}
        columns={1}
        spacingX="10px"
        spacingY="10px"
        overflow={'auto'}
      >
        {dataList.map((item, index) => renderItem(item, index))}
      </SimpleGrid>
    </Flex>
  );
};
