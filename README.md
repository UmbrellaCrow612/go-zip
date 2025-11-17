# Go Zip

A wrapper around a Go CLI binary to **zip and unzip files and folders quickly**, designed for scripting in Node.js or other environments.

---

## CLI Usage

```bash
<cmd> <command> <path> <outPath> [flags]
```

- **command** – The operation to perform (`zip` or `unzip`)
- **path** – Input file or folder
- **outPath** – Output file (for zip) or folder (for unzip)
- **flags** – Optional arguments

### Examples

---

## CLI Arguments Table

| Argument / Flag     | Type    | Required | Description                                                    | Example                                                          |
| ------------------- | ------- | -------- | -------------------------------------------------------------- | ---------------------------------------------------------------- |
| `zip`               | command | Yes      | Compress files or folders into a zip archive                   | `cli zip input_folder output.zip`                                |
| `unzip`             | command | Yes      | Extract files or folders from a zip archive                    | `cli unzip archive.zip extracted_folder`                         |
| `<path>`            | string  | Yes      | Input file or folder to zip or unzip                           | `cli zip ./folder ./archive.zip`                                 |
| `<outPath>`         | string  | Yes      | Output zip file or extraction folder                           | `cli unzip archive.zip ./output`                                 |
| `--include-files`   | string  | No       | **Regex** for files to include (only during unzip)             | `cli unzip archive.zip ./output --include-files="file\.txt"`     |
| `--include-folders` | string  | No       | **Regex** for folders to include (only during unzip)           | `cli unzip archive.zip ./output --include-folders="folder_name"` |
| `--flatten`         | bool    | No       | If true, flattens single-directory zip contents when unzipping | `cli unzip archive.zip ./output --flatten`                       |
