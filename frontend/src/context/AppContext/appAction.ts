import { redirect } from 'next/navigation';
import AuthApi from 'src/Api/Auth';
import clearCookiesOnLogout from 'utils/clearCookiesOnLogout';
import { cookieGetter } from 'utils/cookieUtils';
import { getInitialState } from './helper';

const appActionsReducer = (dispatch: React.Dispatch<AppReducerType>) => {
  const appActionInitialData = (payload: AppAction['InitialData']) => {
    console.log("initial data action need api")
  };

  const appActionLogout = (payload: AppAction['logout']) => {
    const jwt = cookieGetter({ name: 'jwt' });
    if (!payload.channel?.isClosed) {
      payload.channel?.postMessage?.('userLogout');
    }
    const logoutRemover = () => {
      clearCookiesOnLogout();
      dispatch({ type: 'logout' });
    };
    if (payload.withoutCallApi) {
      logoutRemover();
    } else {
      if (jwt) {
        AuthApi.logout()
          .then((res) => {
            if (res.status == 200) {
              logoutRemover();
              setTimeout(() => {
                redirect('/');
              }, 3000);
            }
          })
          .catch((error) => console.log(error));
      }
    }
  };

  const appActionLogin = (payload: AppAction['login']) => {
    const { jwt, userInfo } = getInitialState();
    if (!payload.channel?.isClosed) {
      payload.channel?.postMessage?.('userLogin');
    }
    dispatch({ type: 'login', jwt, userInfo });
  };

  const actionsReducer: ActionsReducer<AppAction> = (action) => {
    console.log(`app action :${action.type}`, action);
    switch (action.type) {
      case 'InitialData': {
        appActionInitialData(action.payload);
        break;
      }
      case 'logout': {
        appActionLogout(action.payload);
        break;
      }
      case 'login': {
        appActionLogin(action.payload);
        break;
      }
      default: {
        throw new Error(`Unexpected type: ${action}`);
      }
    }
  };
  return actionsReducer;
};

export default appActionsReducer;
