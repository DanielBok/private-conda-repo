const CircularDependencyPlugin = require("circular-dependency-plugin");
const CracoAntDesignPlugin = require("craco-antd");
const CracoRawLoaderPlugin = require("@baristalabs/craco-raw-loader");
const WebpackBar = require("webpackbar");
const getCSSModuleLocalIdent = require("react-dev-utils/getCSSModuleLocalIdent");

const path = require("path");

const isDev = process.env.NODE_ENV === "development";

const extraWebpackPlugins = isDev
  ? [
      new CircularDependencyPlugin({
        exclude: /node_modules/,
        failOnError: true,
        allowAsyncCycles: false,
        cwd: process.cwd(),
      }),
    ]
  : []; // prod plugins

module.exports = {
  webpack: {
    performance: {
      hints: true,
    },
    alias: {
      "@": path.resolve(__dirname, "src/"),
    },
    devServer: {
      historyApiFallback: true, // * to handle react-router-dom browserHistory
      inline: true,
      compress: true,
      open: false,
      port: 3000,
    },
    plugins: [new WebpackBar({ profile: true }), ...extraWebpackPlugins],
  },
  plugins: [
    {
      plugin: CracoAntDesignPlugin,
      options: {
        customizeTheme: {
          "@primary-color": "#43b02a",
          "@primary-color-light": "#46d42a",
          "@primary-color-dark": "#025c02",
        },
        cssLoaderOptions: {
          modules: {
            localIdentName: isDev ? "[path][name]_[local]" : "[hash:base64]",
            getLocalIdent: (context, localIdentName, localName, options) =>
              context.resourcePath.includes("node_modules")
                ? localName
                : getCSSModuleLocalIdent(
                    context,
                    localIdentName,
                    localName,
                    options
                  ),
          },
          localsConvention: "camelCase",
        },
      },
    },
    {
      plugin: CracoRawLoaderPlugin,
      options: {
        test: /\.md$/,
      },
    },
  ],
};
