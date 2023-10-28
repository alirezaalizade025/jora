import { Metadata } from 'next';
import { Suspense } from 'react';
import 'static/css/common.css';

export const metadata: Metadata = {
  viewport: { width: 1440 },
};

export default function RootLayout({
  // Layouts must accept a children prop.
  // This will be populated with nested layouts or pages
  children,
}: {
  children: React.ReactNode;
}) {
  const lang = 'fa';
  const direction = 'rtl';

  return (
    <html lang={lang} dir={direction}>
      <body>
          <>{children}</>
      </body>
    </html>
  );
}
