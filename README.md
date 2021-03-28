# Security
Simple (unsecure) encryption and decryption in Go. No external libraries.
> Disclaimer: encryption uses the bitwise operator on bytes. NOT SECURE. This is not meant to be very secure, but to simply teach and show the idea of encryption, while using a simple encription technique. See the unchecked [Features](#Features).

## Features
- [x] Encypt/decrypt text
- [x] Encrypt/decrypt single files
- [x] Encrypt/decrypt folders
- [x] Encrypt/decrypt subfolders in folder
- [x] Overwrite files
- [x] Encrypt/decrypt to new folder
- [x] Save new files with prefix
- [x] Save new files with original name in new folder
- [ ] New, more secure cipher for encyption

## Getting Started
To start, clone the repo. Now, you can run the program directly or build it into an executable.
>- **Option 1:** Running the program directly
```shell
$ go run main.go
```
>- **Option 2:** Building the program
```shell
$ go build main.go
$ main.exe (Windows)
$ ./main (Linux and MacOS)
```
Once the program is running, everything is self-explanatory and intuitive.

## Author
- **Hussein Elguindi** - *all the work*

## License 
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
- [GPL-3.0 License](https://www.gnu.org/licenses/gpl-3.0)
- Copyright 2020 © Hussein Elguindi

