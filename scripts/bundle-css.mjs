import { execSync } from 'child_process';
import path from 'path';
import fs from 'fs';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

// Get the directory of the current script
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const cssDirectory = path.join(__dirname, '../frontend/css');
const outputDirectory = path.join(__dirname, '../public/css');
const frontendDir = path.join(__dirname, '../frontend');

function processAllCSSFiles(directory) {
  // Remove the output directory if it exists
  if (fs.existsSync(outputDirectory)) {
    fs.rmSync(outputDirectory, { recursive: true, force: true });
  }

  // Recreate the output directory
  fs.mkdirSync(outputDirectory, { recursive: true });
  // Change the working directory to frontend
  process.chdir(frontendDir);

  // Process all CSS files initially using PostCSS with glob pattern
  execSync(`bunx postcss "${directory}/**/*.css" --config ./postcss.config.js --base css --dir ${outputDirectory} --ext .min.css`);
  console.log(`Processed all initial CSS files in ${directory}`);

  // Change back to the original working directory
  process.chdir(__dirname);
}

// Process all CSS files initially
processAllCSSFiles(cssDirectory);
