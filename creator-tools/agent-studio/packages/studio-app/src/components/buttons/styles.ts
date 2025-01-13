import { SystemStyleObject } from '@chakra-ui/react';

export const ButtonStyleMap: Record<string, SystemStyleObject> = {
  DEFAULT_STYLE: {
    borderRadius: '999px',
    fontSize: '12px',
    lineHeight: 'calc(16 / 12)',
    fontWeight: '700',
    fontFamily: 'var(--font-SFProDisplay)',
    textAlign: 'center',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    gap: '6px',
    px: '8px',
    py: '2px',
    height: '24px',
    _hover: {
      opacity: 0.7,
      cursor: 'pointer',
    },
  },
};
