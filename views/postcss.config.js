const fs = require('fs');
const path = require('path');

let combinedJson = {};

module.exports = {
  plugins: [
    require('postcss-modules')({
      getJSON: function(cssFileName, json, outputFileName) {
        // Get the relative path from the base CSS directory to the file
        const relativePath = path.relative("./", cssFileName);
        const cleanKey = relativePath.replace(path.extname(relativePath), '').replace(new RegExp(path.sep, 'g'), '.');
        combinedJson[cleanKey] = json;
      },
    }),
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

// Hook into the build process to write the combined JSON file once all CSS files are processed
process.on('exit', () => {
  const jsonFileName = path.resolve(__dirname, '../public/css', 'styles.map.json');
  fs.mkdirSync(path.dirname(jsonFileName), { recursive: true });
  fs.writeFileSync(jsonFileName, JSON.stringify(flattenObject(combinedJson), null, 2)); // Added indentation for better readability
});

// Flattening function that excludes empty objects
function flattenObject(obj, parentKey = '', res = {}) {
  for (let key in obj) {
    const newKey = parentKey ? `${parentKey}.${key}` : key;
    if (typeof obj[key] === 'object' && !Array.isArray(obj[key]) && Object.keys(obj[key]).length > 0) {
      flattenObject(obj[key], newKey, res);
    } else if (typeof obj[key] !== 'object') {
      res[newKey] = obj[key];
    }
  }
  return res;
}
