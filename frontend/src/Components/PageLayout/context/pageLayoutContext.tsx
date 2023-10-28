import React, { useState } from 'react';

const PageLayoutStateContext = React.createContext<PageLayoutContextState>(null);

const initialState: PageLayoutContextState = {
  refs: {
    pageLayoutWrapperRef: React.createRef<HTMLDivElement>(),
  },
};

const PageLayoutProvider = (props: { children: React.ReactNode }) => {
  const [state, dispatch] = useState(initialState);

  return (
    <PageLayoutStateContext.Provider value={state}>
      {props.children}
    </PageLayoutStateContext.Provider>
  );
};

const usePageLayoutState = () => {
  const context = React.useContext(PageLayoutStateContext);
  if (context === undefined) {
    throw new Error('context must use in provider');
  }
  return context;
};

export { PageLayoutProvider, usePageLayoutState };
