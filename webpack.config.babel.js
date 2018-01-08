import path from 'path'
import webpack from 'webpack'
import ExtractTextPlugin from 'extract-text-webpack-plugin'

const extractText = new ExtractTextPlugin({
  filename: "style.css",
});

var DIST_PATH = path.resolve( __dirname, 'public/js' );
var SOURCE_PATH = path.resolve( __dirname, 'static_source' );

export default {
    entry: SOURCE_PATH,
    output: {
        path: DIST_PATH,
        filename: 'app.js',
    },
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        ignore: './node_modules/',
                    }
                }
            },
            {
                test: /\.css$/,
                use: ExtractTextPlugin.extract({
                    fallback: 'style-loader',
                    use: "css-loader"
                })
            }
        ]
    },
    plugins: [
        extractText
    ]
};