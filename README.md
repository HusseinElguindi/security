# Security
Simple encryption and decryption in Go. No external libraries.
> Small disclaimer: encryption uses the bitwise operator on bytes. Not extremely secure. See [TODO](#TODO).

## Features
- [x] Encrypt/decrypt single files
- [x] Encrypt/decrypt folders
- [x] Encrypt/decrypt subfolders in folder
- [x] Overwrite files
- [x] Encrypt/decrypt to new folder
- [x] Save new files with prefix
- [x] Save new files with original name in new folder

## Getting Started
To start, you can run the program directly or build it into an executable.
>- Run the program
```shell
$ go run main.go
$ brew install fvcproductions
```
>- Build the program
```shell
$ go build main.go
$ main.exe (Windows) or ./main (Linux and MacOS)
```

## TODO
- [ ] New, more secure cipher for encyption
- [ ] Option to encypt/decrypt text

## Author
- **Hussein Elguindi** - *all the work*
