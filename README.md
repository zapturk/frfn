# Find and Replace File Name
This is a simple CLI tool to find and replace text in the file name in the current directry.

Install
==================
- Install Go
- Add the Go bin folder to your path
- Run this in your terminal
```
go install github.com/zapturk/frfn@latest
```
```
go install github.com/zapturk/frfn@v1.0.0
```

Usage
==================
This tool has three main commands: `replace`, `prepend`, and `append`.

### Replace

The `replace` command (aliased as `r`) finds and replaces text within file names in the current directory.

**Options:**

*   `--OldText`, `-o`: (Required) The text to be replaced in the file name.
*   `--NewText`, `-n`: (Required) The new text to replace the old text.
*   `--Force`, `-f`: (Optional) By default, the command runs in a "dry run" mode, showing the changes without applying them. Use this flag to execute the file renaming.

**Example:**

To replace "old" with "new" in file names:

```bash
frfn replace -o "old" -n "new"
```

To execute the replacement:

```bash
frfn replace -o "old" -n "new" -f
```

### Prepend

The `prepend` command (aliased as `p`) adds text to the beginning of file names in the current directory.

**Options:**

*   `--Text`, `-t`: (Required) The text to prepend to the file names.
*   `--Force`, `-f`: (Optional) By default, the command runs in a "dry run" mode. Use this flag to execute the file renaming.

**Example:**

To prepend "prefix-" to all file names:

```bash
frfn prepend -t "prefix-"
```

To execute the prepending:

```bash
frfn prepend -t "prefix-" -f
```

### Append

The `append` command (aliased as `a`) adds text to the end of file names in the current directory.

**Options:**

*   `--Text`, `-t`: (Required) The text to append to the file names.
*   `--Force`, `-f`: (Optional) By default, the command runs in a "dry run" mode. Use this flag to execute the file renaming.

**Example:**

To append "-suffix" to all file names:

```bash
frfn append -t "-suffix"
```

To execute the appending:

```bash
frfn append -t "-suffix" -f
```


Contributors
==================
<a href="https://github.com/zapturk/frfn/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=zapturk/frfn&max=500&columns=20&anon=1" />
</a>
