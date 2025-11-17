#!/usr/bin/env node

const { spawn } = require("child_process");
const path = require("path");
const os = require("os");

/**
 * Get the path to the platform-specific Go Zip binary
 * @returns {string} Binary path
 */
function getBinaryPath() {
  const platform = os.platform();
  let binaryName;

  switch (platform) {
    case "win32":
      binaryName = "go-zip-windows.exe";
      break;
    case "darwin":
      binaryName = "go-zip-darwin";
      break;
    case "linux":
      binaryName = "go-zip-linux";
      break;
    default:
      throw new Error(`Unsupported platform: ${platform}`);
  }

  return path.join(__dirname, "bin", binaryName);
}

/**
 * Spawn the Go Zip binary with the given arguments
 * @param {string[]} args Arguments to pass to the binary
 * @param {number} [timeout] Optional timeout in milliseconds
 * @returns {Promise<void>} Resolves when process exits successfully
 */
function runBinary(args, timeout) {
  return new Promise((resolve, reject) => {
    const child = spawn(getBinaryPath(), args, { stdio: "inherit" });
    let killed = false;

    // Handle timeout
    let timer;
    if (timeout) {
      timer = setTimeout(() => {
        killed = true;
        child.kill();
        reject(new Error(`Process timed out after ${timeout}ms`));
      }, timeout);
    }

    child.on("close", (code) => {
      if (timer) clearTimeout(timer);
      if (killed) return;
      if (code === 0) resolve();
      else reject(new Error(`Process exited with code ${code}`));
    });
  });
}

/**
 * Zip a file or folder
 * @param {string} input Input file or folder path
 * @param {string} output Output zip file path
 * @param {Object} [options] Optional arguments
 * @param {number} [options.timeout] Timeout in milliseconds
 * @returns {Promise<void>}
 */
function zip(input, output, options = {}) {
  const args = ["zip", input, output];
  return runBinary(args, options.timeout);
}

/**
 * Unzip an archive
 * @param {string} input Input zip file path
 * @param {string} output Output folder path
 * @param {Object} [options] Optional arguments
 * @param {string} [options.includeFiles] Regex to include files
 * @param {string} [options.includeFolders] Regex to include folders
 * @param {boolean} [options.flatten] Flatten single-directory contents
 * @param {number} [options.timeout] Timeout in milliseconds
 * @returns {Promise<void>}
 */
function unzip(input, output, options = {}) {
  const args = ["unzip", input, output];

  if (options.includeFiles)
    args.push(`--include-files=${options.includeFiles}`);
  if (options.includeFolders)
    args.push(`--include-folders=${options.includeFolders}`);
  if (options.flatten) args.push("--flatten");

  return runBinary(args, options.timeout);
}

// Export helpers for programmatic use
module.exports = { zip, unzip };
