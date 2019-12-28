const CircularDependencyPlugin = require("circular-dependency-plugin");
const CracoAntDesignPlugin = require("craco-antd");
const WebpackBar = require("webpackbar");

const path = require("path");

const extraWebpackPlugins =
  process.env.NODE_ENV === "development"
    ? [
        new CircularDependencyPlugin({
          exclude: /node_modules/,
          failOnError: true,
          allowAsyncCycles: false,
          cwd: process.cwd()
        })
      ]
    : []; // prod plugins

module.exports = {
  webpack: {
    performance: {
      hints: true
    },
    alias: {
      "@": path.resolve(__dirname, "src/")
    },
    devServer: {
      historyApiFallback: true, // * to handle react-router-dom browserHistory
      inline: true,
      compress: true,
      open: false,
      port: 3000
    },
    plugins: [new WebpackBar({ profile: true }), ...extraWebpackPlugins]
  },
  plugins: [
    {
      plugin: CracoAntDesignPlugin,
      options: {
        lessLoaderOptions: {
          javascriptEnabled: true,
          noIeCompat: true
        },
        cssLoaderOptions: {
          modules: {
            localIdentName: "[name]_[local]_[hash:base64:5]"
          },
          localsConvention: "camelCase"
        }
      }
    }
  ]
};
