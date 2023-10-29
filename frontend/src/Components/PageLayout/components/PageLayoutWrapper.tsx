import { usePageLayoutState } from '../context/pageLayoutContext';

const PageLayoutWrapper = ({ children }) => {
  const pageLayoutState = usePageLayoutState();
  return <div ref={pageLayoutState.refs.pageLayoutWrapperRef}>{children}</div>;
};

export default PageLayoutWrapper;
