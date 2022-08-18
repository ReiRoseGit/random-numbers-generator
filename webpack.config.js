const path = require('path')
const htmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    entry: './src/scripts/script.ts',
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: 'bundle.js'
    },
    resolve: {
        extensions: [".tsx", ".ts", ".js"]
    },
    module:{
        rules: [
            {
                test: /\.tsx?$/,
                loader: 'ts-loader',
            },
            {
                enforce: 'pre',
                test: /\.ts?$/,
                use: "tslint-loader"
            }
        ]
    },
    plugins: [new htmlWebpackPlugin({template:'./src/public/index.html'})],
    mode:'production'
}