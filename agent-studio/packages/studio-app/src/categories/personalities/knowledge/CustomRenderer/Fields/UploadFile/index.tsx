import { Flex, Image, ListItem, UnorderedList } from '@chakra-ui/react';
import cs from 'classnames';
import { useState } from 'react';
import { FileUploader } from 'react-drag-drop-files';

import s from './styles.module.scss';

export interface IProps {
  acceptedFileType?: Array<string>;
  fileOrFiles?: File[] | null;
  labelText: string;
  maxSize: number;
  onChange: (files: FileList) => void;
  multiple?: boolean;
  disabled?: boolean;
}

const UploadFile: React.FC<IProps> = ({
  acceptedFileType,
  fileOrFiles,
  labelText,
  maxSize,
  onChange,
  multiple = false,
  disabled = false,
}: IProps) => {
  const [file, setFile] = useState<FileList | null>(null);
  const [error, setError] = useState<string | null>(null);

  const onChangeFile = (file: FileList): void => {
    setFile(file);
    setError('');
    onChange(file);
  };

  const onSizeError = (): void => {
    setError(`File size error.`);
  };

  const onTypeError = (): void => {
    setError('Invalid file extension.');
  };

  return (
    <div
      className={cs(s.dropFile, {
        [s.dropFile__drag]: false,
        [s.dropFile__error]: !!error,
      })}
    >
      <FileUploader
        handleChange={onChangeFile}
        name={'zipFileUploader'}
        maxSize={maxSize}
        minSize={0}
        types={acceptedFileType}
        onTypeError={onTypeError}
        onSizeError={onSizeError}
        multiple={multiple}
        fileOrFiles={fileOrFiles}
        classes={s.dropZone}
        disabled={disabled}
      >
        <Flex direction="column" gap="6px">
          <Flex alignItems="center" gap="8px" flex="1">
            <Image
              className={s.dropZoneThumbnail}
              src="https://storage.googleapis.com/eternal-ai/images/docs.svg"
              w="60px"
              h="60px"
            />
            <p className={cs(s.dropZoneDescription, 'description')}>
              {labelText}
            </p>
          </Flex>
          <UnorderedList className={s.uploadFile_description} flex="1">
            {acceptedFileType && (
              <ListItem>
                Supported file extensions are {acceptedFileType.join(', ')}.
              </ListItem>
            )}
          </UnorderedList>
        </Flex>
      </FileUploader>
    </div>
  );
};

export default UploadFile;
