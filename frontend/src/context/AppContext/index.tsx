'use client';

import { usePathname, useRouter } from 'next/navigation';
import React, { useEffect, useMemo, useReducer, useState } from 'react';
import { cookieGetter } from 'utils/cookieUtils';
import appActionsReducer from './appAction';
import appReducer from './appReducer';
import { getInitialState, initialState } from './helper';

const AppStateContext = React.createContext<AppContextState | undefined>(undefined);
const AppActionsContext = React.createContext<ActionsReducer<AppAction> | undefined>(undefined);

const AppProvider = (props: { children: React.ReactNode }) => {
  const [state, dispatch] = useReducer(appReducer, initialState, getInitialState);
  const [isMount, setIsMount] = useState(false);

  const actions = React.useMemo(() => appActionsReducer(dispatch), [dispatch]);
  const pathname = usePathname();
  const router = useRouter();

  const needSeeAuth = useMemo(
    () => !state.jwt && pathname != '/auth' && pathname != '/register',
    [pathname, state.jwt],
  );

  useEffect(() => {
    setIsMount(true);
  }, []);

  useEffect(() => {
    if (needSeeAuth) {
      router.replace('/auth');
    }
  }, [needSeeAuth, router]);

  useEffect(() => {
    if (state.jwt) {
      actions({ type: 'InitialData', payload: { clearCacheEntry: false } });
    }
  }, [actions, state.jwt]);

  useEffect(() => {
    const onMessage = (msg) => {
      if (msg == 'userLogout') {
        dispatch({ type: 'logout' });
      }
      if (msg == 'userLogin') {
        const jwt = cookieGetter({ name: 'jwt' }) || '';
        if (jwt) {
          actions({ type: 'InitialData', payload: { clearCacheEntry: false } });
        }
      }
    };
    const initChannel = () => {
      state.channel?.addEventListener?.('message', onMessage);
    };
    const removeChannel = () => {
      state.channel?.removeEventListener?.('message', onMessage);
      state.channel?.close?.();
    };

    window.addEventListener('pageshow', initChannel);
    window.addEventListener('pagehide', removeChannel);

    return () => {
      removeChannel();
      window.removeEventListener('pageshow', initChannel);
      window.removeEventListener('pagehide', removeChannel);
    };
  }, [actions, state.channel]);

  return (
    <AppStateContext.Provider value={state}>
      <AppActionsContext.Provider value={actions}>
        {!needSeeAuth && isMount ? props.children : 'is loading ...'}
      </AppActionsContext.Provider>
    </AppStateContext.Provider>
  );
};

const useAppState = () => {
  const context = React.useContext(AppStateContext);
  if (context === undefined) {
    console.log('AppStateContext provider undefined');
  }
  return context;
};

const useUserInfo = () => {
  const context = React.useContext(AppStateContext);
  if (context === undefined) {
    throw new Error('AppStateContext operator provider undefined');
  }
  const userInfo = useMemo(() => context.userInfo, [context.userInfo]);
  return userInfo;
};

const useAppDispatch = () => {
  const context = React.useContext(AppActionsContext);
  if (context === undefined) {
    console.log('appActionsContext provider undefined');
  }
  return context;
};

export { AppProvider, useAppDispatch, useAppState, useUserInfo };
