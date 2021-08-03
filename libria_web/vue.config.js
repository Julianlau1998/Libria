module.exports = {
    pwa: {
      name: 'Libria',
      themeColor: '#000000',
      msTileColor: '#000000',
      backgroundColor: '#ffffff',
      appleMobileWebAppCapable: 'yes',
      appleMobileWebAppStatusBarStyle: 'default',
      workboxPluginMode: 'InjectManifest',
      workboxOptions: {
        swSrc: 'src/service-worker.js'
      }
    },
    devServer: {
      host: '0.0.0.0',
      disableHostCheck: true,
      proxy: {
        '^/api/*': {
          target: 'http://192.168.178.58:1324',
          changeOrigin: true
        }
      }
    }
  }
  