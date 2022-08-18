const path = require('path')
const htmlWebpackPlugin = require('html-webpack-plugin');
const PrettierPlugin = require("prettier-webpack-plugin");

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
    plugins: [
        new htmlWebpackPlugin({template:'./src/public/index.html'}),
        new PrettierPlugin({
            "printWidth": 120,
            "tabWidth": 4,
            "useTabs": false,
            "semi": false,
            "encoding": "utf-8",
            "extensions": [".ts", "js", ".css", ".less", ".html", ".xml"],
            "singleQuote": true,
            "arrowParens": "always",
            "trailingComma": "es5",
            "end-of-line": "lf"
        }
        )
    ],
    mode:'production'
}