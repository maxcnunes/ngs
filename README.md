# n git status (ngs)

Git status command in **n** directories. This is a pretty simple command line to show the status from multiple git repositories.
Currently it can only search in the current directory (or the one specified on `dir`) and it's first children directories.

### Download

[Releases](https://github.com/maxcnunes/ngs/releases)

### Installing

Download the bin file for your specif OS and copy into `/usr/local/bin` or include the binary on `$PATH`.

```bash
sudo mv ngs /usr/local/bin/
sudo chmod +x /usr/local/bin/ngs
```

### Options

- **-h**: Help message
- **-dir=DIRNAME**: The base directory to start searching

### Example

```bash
cd my-dev-folder
$ ngs
»  my-first-project/
?? .editorconfig
?? README.md
?? bin/
?? fig.yml
?? main.go
»  my-second-project/
M app/index.html
M app/styles/main.scss
```


## Development

Running with `Docker` and `Fig`:

```bash
fig run --rm local
```
