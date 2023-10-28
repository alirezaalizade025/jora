const path = require('path');

// @ts-check
/**
 * @type {import('next/dist/server/config').NextConfig}
 **/
const nextConfig = {
  reactStrictMode: false,
  eslint: { ignoreDuringBuilds: true },
  compress: false,

  webpack: (config) => {
    config.resolve.modules.push(
      path.resolve('./src/Components'),
      path.resolve('./src'),
      path.resolve('./public'),
      path.resolve('./'),
    );

    return config;
  },
  crossOrigin: 'anonymous',
  generateBuildId: async () => {
    return process.env.NEXT_PUBLIC_APP_VERSION || `${Date.now()}`;
  },
};

module.exports = nextConfig;
