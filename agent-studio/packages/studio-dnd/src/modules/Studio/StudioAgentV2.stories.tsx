import type { Meta, StoryObj } from '@storybook/react';

import AGENT_MODEL_CATEGORIES from './__mocks__/agent-categories';
import { AGENT_DATA_SOURCE } from './__mocks__/agent-data-source';
import { Studio, StudioProps } from './Studio';

type Story = StoryObj<typeof Studio>;

const args = {
  categories: AGENT_MODEL_CATEGORIES,
  dataSource: AGENT_DATA_SOURCE,
  graphData: {
    data: [],
    viewport: {
      x: 0,
      y: 0,
      zoom: 1,
    },
  },
} satisfies StudioProps;

const meta: Meta<typeof Studio> = {
  title: 'Studio',
  component: Studio,
  args,
};

export const AgentStudioV2: Story = {
  render: function useTabs(args) {
    return (
      <div style={{ width: 'calc(100vw - 3rem)', height: 'calc(100vh - 3rem)' }}>
        <Studio
          {...args}
          // ref={ref}
          onChange={(data) => {
            console.log('[Studio] onChange', data);
          }}
        />
      </div>
    );
  },
};

export default meta;
