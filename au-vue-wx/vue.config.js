const port = process.env.port || process.env.npm_config_port || 2222 // dev port

module.exports = {
  // 选项...
  publicPath: './',
  outputDir: 'dist',
  assetsDir: 'static',
  lintOnSave: process.env.NODE_ENV === 'development',
  // devServer: {
  //   proxy: {
  //     '/wxapi': {
  //       // 需要代理的url
  //       target: 'http://localhost:9528',
  //       changeOrigin: true
  //       // pathRewrite: {
  //       //   '^/wxapi': '/'
  //       // }
  //     }
  //   }
  // }
  devServer: {
    port: port,
    open: true,
    overlay: {
      warnings: false,
      errors: true
    }
  }
  // before: require('./mock/mock-server.js')
}
