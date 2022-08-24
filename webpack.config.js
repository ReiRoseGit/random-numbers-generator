const path = require('path')
const htmlWebpackPlugin = require('html-webpack-plugin')
const PrettierPlugin = require('prettier-webpack-plugin')
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin')
const UglifyJsPlugin = require('uglifyjs-webpack-plugin')

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
                filename: (env === 'dev') ? 'name.js' : '[hash].js',
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
            optimization: {
                minimize: true,
                minimizer: [
                    new CssMinimizerPlugin(),
                    new UglifyJsPlugin({
                        uglifyOptions: {
                            compress: {
                                unsafe: true,
                                inline: true,
                                passes: 2,
                                keep_fargs: false,
                            },
                            output: {
                                beautify: false,
                            },
                            mangle: true,
                        },
                    }),
                ],
            },
            plugins: [new htmlWebpackPlugin({ template: './src/public/index.html' }), new MiniCssExtractPlugin()],
            mode: 'production',
            devtool: 'source-map',
        }
    }
}
