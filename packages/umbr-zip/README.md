# Umbr-zip

**Umbr-zip** is a Node.js wrapper around the **Go Zip CLI** (`go-zip`) for quickly compressing and extracting files and folders. It is designed for scripting and automation in Node.js environments.  

> ⚠️ **Not for browser use** – this library relies on platform-specific Go binaries.

---

## Features

- Compress files and folders into zip archives.
- Extract zip archives with optional file/folder filtering.
- Flatten single-directory contents when extracting.
- Supports Windows, macOS, and Linux.

---

## Installation

```bash
npm install umbr-zip --save-dev
````

---

## Usage (Programmatic)

```javascript
const { zip, unzip } = require("umbr-zip");

// Zip a folder
await zip("./myFolder", "./archive.zip");

// Unzip an archive
await unzip("./archive.zip", "./outputFolder");

// Unzip only specific files and flatten directories
await unzip("./archive.zip", "./outputFolder", {
  includeFiles: ".*\\.txt$",
  flatten: true
});
```

---

### Commands

* `zip` – Compress files or folders.
* `unzip` – Extract files or folders.

### Flags (Optional)

* `--include-files` – Regex to include specific files (unzip only).
* `--include-folders` – Regex to include specific folders (unzip only).
* `--flatten` – Flatten single-directory contents when unzipping.
* `--timeout` – Timeout in milliseconds for the operation.

### Examples

```bash
# Zip a folder
umbr-zip zip ./myFolder ./archive.zip

# Unzip an archive
umbr-zip unzip ./archive.zip ./outputFolder

# Unzip only txt files
umbr-zip unzip ./archive.zip ./outputFolder --include-files=".*\.txt$"

# Unzip and flatten contents
umbr-zip unzip ./archive.zip ./outputFolder --flatten
```

---

## Notes

* This library spawns **platform-specific Go binaries**, so it cannot run in browsers.
* Works on **Windows, macOS, and Linux**.