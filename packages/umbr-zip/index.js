#!/usr/bin/env node

const { spawn } = require("child_process");
const path = require("path");
const os = require("os");

// Determine platform
let platform = os.platform();
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
    console.error(`Unsupported platform: ${platform}`);
    process.exit(1);
}

// Path to the binary inside this package
const binaryPath = path.join(__dirname, "bin", binaryName);

// Pass all CLI arguments to the binary
const args = process.argv.slice(2);

// Spawn the platform-specific binary
const child = spawn(binaryPath, args, { stdio: "inherit" });

child.on("close", (code) => {
  process.exit(code);
});
