const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin')
const webpack = require('webpack')

module.exports = {
    entry:'./src/script.js',
    output:{
        path:path.resolve(__dirname, 'dist'),
        filename:'bundle.js'
    },
    plugins : [
        new HtmlWebpackPlugin({template: './public/index.html'})
    ],
    mode : 'development'
};