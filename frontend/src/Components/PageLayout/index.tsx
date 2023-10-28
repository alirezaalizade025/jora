'use client';
import PageLayoutWrapper from './components/PageLayoutWrapper';
import { PageLayoutProvider } from './context/pageLayoutContext';

const PageLayout = ({ children }) => {
  return (
    <PageLayoutProvider>
      <PageLayoutWrapper>{children}</PageLayoutWrapper>
    </PageLayoutProvider>
  );
};

export default PageLayout;
