const path = require('path')
const htmlWebpackPlugin = require('html-webpack-plugin')
const PrettierPlugin = require('prettier-webpack-plugin')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')

module.exports = (env) => {
    if (env.development) {
        console.log('Development mode!')
        return {
            entry: './src/scripts/script.ts',
            output: {
                path: path.resolve(__dirname, 'dist'),
                filename: 'bundle.js',
            },
            resolve: {
                extensions: ['.tsx', '.ts', '.js'],
            },
            module: {
                rules: [
                    {
                        test: /\.tsx?$/,
                        loader: 'ts-loader',
                    },
                    {
                        enforce: 'pre',
                        test: /\.ts?$/,
                        use: 'tslint-loader',
                    },
                    { test: /\.css$/, use: ['style-loader', 'css-loader'] },
                ],
            },
            plugins: [
                new htmlWebpackPlugin({ template: './src/public/index.html' }),
                new PrettierPlugin({
                    printWidth: 120,
                    tabWidth: 4,
                    useTabs: false,
                    semi: false,
                    encoding: 'utf-8',
                    extensions: ['.ts', 'js', '.css', '.less', '.html', '.xml'],
                    singleQuote: true,
                    arrowParens: 'always',
                    trailingComma: 'es5',
                    'end-of-line': 'lf',
                }),
            ],
            mode: 'development',
        }
    } else if (env.production) {
        console.log('Production mode!')
        return {
            entry: './src/scripts/script.ts',
            output: {
                path: path.resolve(__dirname, 'dist'),
                filename: '[name].js',
            },
            resolve: {
                extensions: ['.tsx', '.ts', '.js'],
            },
            module: {
                rules: [
                    {
                        test: /\.tsx?$/,
                        loader: 'ts-loader',
                    },
                    { test: /\.css$/, use: [MiniCssExtractPlugin.loader, 'css-loader'] },
                ],
            },
            plugins: [
                new htmlWebpackPlugin({ template: './src/public/index.html', filename: '[name].html', hash: true }),
                new MiniCssExtractPlugin(),
            ],
            mode: 'production',
            devtool: 'source-map',
        }
    }
}
