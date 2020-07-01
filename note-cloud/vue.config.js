module.exports = {
    devServer: {
        proxy : {
          '^/api' : {
              target : "http://localhost:8001",
              changeOrigin: true,//是否允许跨越
              pathRewrite: {
                  '^/api': '/',//重写,
              }
          }
        },
    }
}