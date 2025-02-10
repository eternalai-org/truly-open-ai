import CREATE_FLOW_FARCASTER_ON_X_VALIDATORS from './create';
import UPDATE_FLOW_REPLY_ON_FARCASTER_VALIDATORS from './update';

const REPLY_ON_FARCASTER_VALIDATES = {
  create: CREATE_FLOW_FARCASTER_ON_X_VALIDATORS,
  update: UPDATE_FLOW_REPLY_ON_FARCASTER_VALIDATORS,
};

export default REPLY_ON_FARCASTER_VALIDATES;
