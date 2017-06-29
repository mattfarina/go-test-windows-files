# Windows (and Linux Subsystem) filesystem bug example

In Golang, using `os.RemoveAll` and some other filesystem (e.g., `os.Rename`)
command on Windows 10 (including the Linux subsystem) causes errors where not
all files are affected.

## How to reproduce

1. Start on Windows 10 (tested in Windows 10 Pro)
2. Clone this repo into your `GOPATH`
3. Change directory to the root
4. Run `go run main.go`

This will attempt to delete all the files in the `files` directory. With the number
of files present it has been seen to fail.

The error seen is
```sh
remove files: directory not empty
```

In this case `files` is the name of the thing on the filesystem attempting to be deleted.