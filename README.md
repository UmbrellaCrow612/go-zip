# Go Zip

A wrapper around a Go CLI binary to **zip and unzip files and folders quickly**, intended for scripting in Node.js scripts.

---

## CLI Usage

```bash
<cmd> <path> <outPath> [flags]
```

- **command** – The operation to perform (`zip` or `unzip`)
- **options** – Additional flags and arguments

### Examples

```bash
# Zip a folder or file
cli zip /path/to/input /path/to/output.zip

# Zip a folder excluding files/folders matching a regex
cli zip /path/to/input /path/to/output.zip --exclude="\.git|node_modules"

# Unzip an archive
cli unzip /path/to/archive.zip /path/to/extracted-folder --name=new_folder_name

# Unzip an archive excluding files/folders matching a regex
cli unzip /path/to/archive.zip /path/to/extracted-folder --exclude="file\.txt|skip_folder"
```

---

## CLI Arguments Table

| Argument / Flag | Type    | Required | Description                                                                   | Example                                                                     |     |     |
| --------------- | ------- | -------- | ----------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --- | --- |
| `zip`           | command | Yes      | Compress files or folders into a zip archive                                  | `cli zip input_folder output.zip`                                           |     |     |
| `unzip`         | command | Yes      | Extract files or folders from a zip archive                                   | `cli unzip archive.zip extracted_folder`                                    |     |     |
| `/in-path`      | string  | Yes      | Input file or folder to zip                                                   | `cli zip ./folder ./archive.zip`                                            |     |     |
| `/out-path`     | string  | Yes      | Output zip file or extraction folder                                          | `cli unzip archive.zip ./output`                                            |     |     |
| `--recursive`   | bool    | No       | Include all files recursively when zipping                                    | `cli zip -r folder archive.zip`                                             |     |     |
| `--name`        | string  | No       | Rename extracted folder when unzipping                                        | `cli unzip archive.zip ./output --name=myfolder`                            |     |     |
| `--exclude`     | string  | No       | **Regular expression** for files/folders to exclude when zipping or unzipping | `cli unzip archive.zip ./output --name=myfolder --exlude=regularExpression` |     |     |
