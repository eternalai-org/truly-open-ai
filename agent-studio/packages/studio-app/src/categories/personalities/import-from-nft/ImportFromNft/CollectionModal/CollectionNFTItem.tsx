import { Flex, Image, Text } from '@chakra-ui/react';

type Props = {
  item: any;
  index?: number;
  itemOnClick?: (item: any) => void;
};

export const CollectionNFTItem = (props: Props) => {
  const { item, index, itemOnClick } = props;
  return (
    <Flex
      key={`${item}-${index}`}
      flexDir={'row'}
      align={'center'}
      flexDirection={'row'}
      gap="8px"
      p={'5px 20px'}
      _hover={{
        cursor: 'pointer',
        opacity: 0.7,
        bgColor: '#E5E7EB',
      }}
      onClick={() => {
        itemOnClick && itemOnClick(item);
      }}
    >
      <Image
        src={item?.image_url}
        w={'30px'}
        h={'30px'}
        borderRadius={'100%'}
        overflow={'hidden'}
      />
      <Text fontSize={'14px'} fontWeight={500} color={'#000'}>
        {item?.name}
      </Text>
    </Flex>
  );
};
