import { initialState } from './helper';

const appReducer = (state: AppContextState, action: AppReducerType): AppContextState => {
  console.log(`app reducer :${action.type}`, action);
  switch (action.type) {
    case 'setUserInfo': {
      return { ...state, userInfo: action.userInfo };
    }

    case 'logout': {
      return initialState;
    }

    case 'login': {
      return {
        ...state,
        jwt: action.jwt,
        userInfo: action.userInfo,
      };
    }
    default: {
      throw new Error(`Unexpected action: ${action}`);
    }
  }
};

export default appReducer;
