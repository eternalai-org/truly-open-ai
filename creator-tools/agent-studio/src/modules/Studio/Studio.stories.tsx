import type { Meta, StoryObj } from '@storybook/react';

import { CATEGORIES } from './__mocks__/categories';
import { DATA_SOURCE } from './__mocks__/data-source';
import { Studio, StudioProps } from './Studio';

type Story = StoryObj<typeof Studio>;

const args = {
  categories: CATEGORIES,
  dataSource: DATA_SOURCE,
  data: [],
} satisfies StudioProps;

const meta: Meta<typeof Studio> = {
  title: 'Studio',
  component: Studio,
  args,
};

export const Default: Story = {
  render: function useTabs(args) {
    return (
      <div style={{ width: 'calc(100vw - 3rem)', height: 'calc(100vh - 3rem)' }}>
        <Studio
          {...args}
          onChange={(data) => {
            console.log('[Studio] onChange', data);
          }}
        />
      </div>
    );
  },
};

export default meta;
