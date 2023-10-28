import { BroadcastChannel } from 'broadcast-channel';
import { cookieGetter } from 'utils/cookieUtils';

export const initialState: AppContextState = {
  jwt: '',
  userInfo: { permissions: null, id: null, username: undefined },
  channel: null,
};

export const getInitialState = (): AppContextState => {
  const jwt = cookieGetter({ name: 'jwt' }) || '';
  let channel = null;
  if (typeof window != 'undefined') {
    try {
      channel = new BroadcastChannel<BroadcastChannelMessages>('kutam');
    } catch (error) {
      channel = null;
      console.log('BroadcastChannel is not supported', { e: error });
    }
  }

  return {
    ...initialState,
    jwt,
    channel,
  };
};
