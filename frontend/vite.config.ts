import { defineConfig, splitVendorChunkPlugin } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react({
      babel: {
        code: true,
        compact: true,
        comments: false,
      },
    }),
    splitVendorChunkPlugin(),
  ],
  build: {
    copyPublicDir: true,
    rollupOptions: {
      output: {
        manualChunks(id) {
          const packageName = id.match(/node_modules\/([^/]+)/)
          if (packageName) {
            return packageName[1]
          }
        },
      },
    },
  },
})
