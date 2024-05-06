import {vitePlugin as remix} from '@remix-run/dev'
import {defineConfig} from 'vite'

export default defineConfig({
  plugins: [remix()],
  server: {
    host: true,
    port: 3000,
    // hmr: {
    //   port: 3010,
    // },
  },
  ignoredRouteFiles: ['**/*.css'],
})
