import path from 'path'
import webpack from 'webpack'

export default {
    entry:  {
        chat_room: './static_source/chat_room.js'
    },
    output:  {
        path: path.resolve(__dirname, 'public/js'),
        filename: '[name].entry.js'
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        ignore: './node_modules/',
                        presets: [
                            ['es2015', { modules: false }]
                        ]
                    }
                }
            },
        ]
    }
}