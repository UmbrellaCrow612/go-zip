#!/usr/bin/env node

const { spawn } = require("child_process");
const path = require("path");
const os = require("os");

/**
 * Get the platform-specific Go Zip binary path.
 *
 * Determines the correct binary based on the current OS:
 * - Windows: `go-zip-windows.exe`
 * - macOS: `go-zip-darwin`
 * - Linux: `go-zip-linux`
 *
 * @throws {Error} Throws if the platform is unsupported.
 * @returns {string} Absolute path to the Go Zip binary.
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
 * Spawn the Go Zip binary with specified arguments.
 *
 * @param {string[]} args - Arguments to pass to the binary.
 * @param {number} [timeout] - Optional timeout in milliseconds.
 * @returns {Promise<void>} Resolves when the process exits successfully, rejects on error or timeout.
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
 * Compress a file or folder into a zip archive.
 *
 * @param {string} input - Path to the file or folder to compress.
 * @param {string} output - Path where the zip archive will be created.
 * @param {Object} [options] - Optional parameters.
 * @param {number} [options.timeout] - Timeout in milliseconds for the operation.
 * @returns {Promise<void>} Resolves when compression is complete.
 *
 * @example
 * await zip("./folder", "./archive.zip", { timeout: 5000 });
 */
function zip(input, output, options = {}) {
  const args = ["zip", input, output];
  return runBinary(args, options.timeout);
}

/**
 * Extract files from a zip archive.
 *
 * Supports optional filtering and flattening of extracted files.
 *
 * @param {string} input - Path to the zip archive.
 * @param {string} output - Folder path where files will be extracted.
 * @param {Object} [options] - Optional parameters.
 * @param {string} [options.includeFiles] - Regex pattern to include specific files.
 * @param {string} [options.includeFolders] - Regex pattern to include specific folders.
 * @param {boolean} [options.flatten] - If true, flattens single-directory contents.
 * @param {number} [options.timeout] - Timeout in milliseconds for the operation.
 * @returns {Promise<void>} Resolves when extraction is complete.
 *
 * @example
 * await unzip("./archive.zip", "./output", { includeFiles: ".*\\.txt$", flatten: true });
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

module.exports = { zip, unzip };
