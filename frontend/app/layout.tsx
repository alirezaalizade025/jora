import { AppProvider } from 'context/AppContext';
import { Metadata } from 'next';
import PageLayout from 'src/Components/PageLayout';
import 'static/css/common.css';

export const metadata: Metadata = {
  viewport: { width: 1440 },
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const lang = 'fa';
  const direction = 'rtl';

  return (
    <html lang={lang} dir={direction}>
      <body>
        <AppProvider>
          <PageLayout>{children}</PageLayout>
        </AppProvider>
      </body>
    </html>
  );
}
