import {
  Flex,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalOverlay,
  Text,
} from '@chakra-ui/react';
import cs from 'classnames';
import s from './styles.module.scss';

type Props = {
  isOpen: boolean;
  onClose: () => void;

  title?: string;
  description?: string;

  dataListView?: React.ReactElement;
};

export default function CollectionModal({
  isOpen,
  onClose,
  title,
  description,
  dataListView,
}: Props) {
  return (
    <Modal
      isOpen={isOpen}
      onClose={onClose}
      isCentered
      onOverlayClick={() => {}}
    >
      <ModalOverlay />
      <ModalContent
        maxW={'max-content'}
        maxH={'max-content'}
        textAlign={'center'}
        className={cs(s.modalContent, 'white-modal')}
      >
        <ModalCloseButton
          className="white-modal__close-btn"
          onClick={onClose}
        />
        <ModalBody
          className={cs(s.modalBody, 'white-modal__body')}
          maxHeight={'500px'}
          overflowY={'auto'}
          display={'flex'}
          flexDir={'column'}
          alignItems={'center'}
        >
          <Flex
            width={'100%'}
            flexDir={'column'}
            gap={'15px'}
            align={'flex-start'}
          >
            <Text fontSize={'20px'} fontWeight={500} color={'#000'}>
              {title}
            </Text>
            <Text fontSize={'16px'} fontWeight={400} opacity={0.7}>
              {description}
            </Text>
            {dataListView}
          </Flex>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
}
