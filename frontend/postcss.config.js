module.exports = {
  plugins: [
    require('postcss-preset-env')({
      stage: 0, // Enables all modern CSS features
      browsers: 'last 2 versions, not ie > 0',
      autoprefixer: { grid: true }
    }),
    require('cssnano')({
      preset: 'default',
    }),
  ],
};

