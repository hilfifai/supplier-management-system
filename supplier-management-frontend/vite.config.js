import { defineConfig } from 'vite';
import laravel from 'laravel-vite-plugin';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
    plugins: [
        laravel({
            input: ['resources/css/app.css', 'resources/js/app.js'],
            refresh: true,
        }),
        vue({
            template: {
                transformAssetUrls: {
                    base: null,
                    includeAbsolute: false,
                },
				compilerOptions: {
				  
				  isCustomElement: tag => tag.startsWith('ion-') 
				}
            },
			
        }),
    ],
    resolve: {
        alias: {
            vue: 'vue/dist/vue.esm-bundler.js',
			crypto: 'crypto-browserify',
        },
    },
	 server: {
    host: '127.0.0.1', // Ganti dari localhost ke 127.0.0.1
    port: 8000,
    strictPort: true,
    https: false, // Nonaktifkan HTTPS
    hmr: {
      host: '127.0.0.1', // Ganti ke 127.0.0.1
    },
    proxy: {
      '/api': 'http://127.0.0.1:8000', // Proxy ke Laravel
      '/build': 'http://127.0.0.1:8000', // Untuk assets
    },
  },
});

