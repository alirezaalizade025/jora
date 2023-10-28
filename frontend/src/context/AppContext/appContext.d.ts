type BroadcastChannelMessages = any;

type UserInfo = any
type channelBroadCastType = import('broadcast-channel').BroadcastChannel<BroadcastChannelMessages>;
type AppContextState = {
  jwt: string;
  userInfo: UserInfo;
  channel: channelBroadCastType;
};

type AppReducerType =
  | ReducerType<'logout'>
  | ReducerType<
    'login',
    {
      jwt: string;
      userInfo: UserInfo;
    }
  >
  | ReducerType<'setUserInfo', { userInfo: UserInfo }>;

type AppAction = {
  InitialData: { clearCacheEntry?: boolean };
  login: { channel: channelBroadCastType };
  logout: {
    shallowLogout?: boolean;
    withoutCallApi?: boolean;
    channel: channelBroadCastType;
    router: import('next/dist/shared/lib/app-router-context').AppRouterInstance;
  };
};
