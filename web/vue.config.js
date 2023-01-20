/*
* Config for project
* Author: Kinice
*/
const fs = require('fs')
const path = require('path')
const packagejson = fs.readFileSync(path.resolve('./package.json'))
const json = JSON.parse(packagejson.toString())
const buildArgv = require('yargs-parser')(process.env.BUILD_ARGV_STR || '')

module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
    ? `${buildArgv['def_publish_env'] === 'prod' ? 'https://g.alicdn.com' : 'https://dev.g.alicdn.com'}/room-paas/${
      json.name
    }/${json.version}/`
    : '/',
  filenameHashing: false,
  // All webpack-dev-server are supported
  devServer: {
    port: 7000,
    host: '0.0.0.0',
    hot: true,
    overlay: true,
    openPage: '/',
    // new proxy tables as same as the past 'proxyTable'
    proxy: {
      '/v1': {
        target: 'http://localhost:7001/',
        changeOrigin: true,
        pathRewrite: {
          '^/v1': '/v1'
        }
      }
    }
  }
}
