# Security
Simple encryption and decryption in Go. No external libraries.
> Small disclaimer: encryption uses the bitwise operator on bytes. Not extremely secure. See [TODO](#TODO).

## Features
- [x] Encypt/decrypt text
- [x] Encrypt/decrypt single files
- [x] Encrypt/decrypt folders
- [x] Encrypt/decrypt subfolders in folder
- [x] Overwrite files
- [x] Encrypt/decrypt to new folder
- [x] Save new files with prefix
- [x] Save new files with original name in new folder

## Getting Started
To start, you can run the program directly or build it into an executable.
>- **Option 1:** Run the program
```shell
$ go run main.go
```
>- **Option 2:** Build the program
```shell
$ go build main.go
$ main.exe (Windows)
$ ./main (Linux and MacOS)
```
Once the program is running, everything is self-explanatory and intuitive.

## TODO
- [ ] New, more secure cipher for encyption

## Author
- **Hussein Elguindi** - *all the work*

## License [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
- [GPL-3.0 License](https://www.gnu.org/licenses/gpl-3.0)

