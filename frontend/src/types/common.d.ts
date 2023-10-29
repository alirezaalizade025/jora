type Globals = { dataTest?: string };
type EventType<T> = (event: T) => void | Promise<any>;
type Canceler = import('axios').Canceler;
type ReducerType<K, V = void> = V extends void ? { type: K } : { type: K } & V;
type ActionMap<M extends { [index: string]: any }> = {
  [Key in keyof M]: M[Key] extends undefined
    ? {
        type: Key;
      }
    : {
        type: Key;
        payload: M[Key];
      };
};

type ActionsReducer<K> = (action: ActionMap<K>[keyof ActionMap<K>]) => void;
