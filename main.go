package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// Task - encrypt/decrypt struct
type Task struct {
	File    *os.File
	newFile *os.File
	wg      *sync.WaitGroup
}

func (t *Task) encrypt(key int) {
	if t.wg != nil {
		defer t.wg.Done()
	}

	buf := make([]byte, 8*1024) // buffer size here

	for {
		n, err := t.File.Read(buf)

		if n > 0 {
			// fmt.Print(buf[:n]) // read buffer
			for i := 0; i < n; i++ {
				buf[i] ^= byte(key)
				// fmt.Println(buf[i])
			}
			t.newFile.Write(buf[:n])
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
	}

	fmt.Printf("Completed %s -> %s.\n", filepath.Base(t.File.Name()), filepath.Base(t.newFile.Name()))
}

func readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

func getKey() (int, error) {
	fmt.Print("Encryption Key: ")
	key, err := strconv.Atoi(readline())
	if err != nil || key == 0 {
		fmt.Println("Encryption Key must be an integer value and not 0.")
		return 0, err
	}
	return key, nil
}

func main() {
	// var err error

	fmt.Println("Security - Hussein Elguindi")
	fmt.Println("1. Encrypt/Decrypt File")
	fmt.Println("2. Encrypt/Decrypt Files In Folder")
	fmt.Print("\nSelection: ")
	choice := readline()

	if choice == "1" {
		key, err := getKey()
		if err != nil {
			return
		}

		t := Task{}

		fmt.Print("Filepath: ")
		fPath := readline()
		t.File, err = os.Open(fPath)
		if err != nil {
			fmt.Println("File does not exist.")
			return
		}
		defer t.File.Close()

		fmt.Print("New Filename: ")
		t.newFile, err = os.OpenFile(filepath.Join(filepath.Dir(fPath), readline()+filepath.Ext(fPath)), os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("Error creating new file.")
			return
		}
		defer t.newFile.Close()

		fmt.Println("\nStarting...")
		t.encrypt(key)

	} else if choice == "2" {
		key, err := getKey()
		if err != nil {
			return
		}

		fmt.Print("Folder Path: ")
		dirPath := readline()

		var fileCount int
		var filePaths []string
		var files []os.FileInfo
		fmt.Print("Every file in sub folders? [y/n] ")

		if strings.EqualFold(readline(), "y") {
			err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
				// fmt.Println(path)
				if !info.IsDir() {
					filePaths = append(filePaths, path)
					fileCount++
				}
				return nil
			})
			if err != nil {
				fmt.Println("Directory does not exist.")
				return
			}
		} else {
			files, err = ioutil.ReadDir(dirPath)
			fileCount = len(files)
			if err != nil {
				fmt.Println("Directory does not exist.")
				return
			}
		}

		fmt.Print("New Files Prefix: ")
		prefix := readline()

		newPath := filepath.Join(dirPath, "Security")
		_, err = os.Stat(newPath)
		if os.IsNotExist(err) {
			err := os.MkdirAll(newPath, 0755)
			if err != nil {
				fmt.Println("Failed to create directory.")
				return
			}
		}

		fmt.Println("\nStarting...")

		i := 0
		wg := &sync.WaitGroup{}
		// for _, f := range files {
		for x := 0; x < fileCount; x++ {
			var fPath string
			var newWritePath string

			if files != nil {
				f := files[x]

				if f.IsDir() {
					continue
				}

				fPath = filepath.Join(dirPath, f.Name())
				newWritePath = filepath.Join(newPath, fmt.Sprintf("%s_%d%s", prefix, i, filepath.Ext(fPath)))

			} else {
				fPath = filePaths[x]
				newWritePath = filepath.Join(newPath, fmt.Sprintf("%s_%d%s", prefix, i, filepath.Ext(fPath)))
			}

			t := Task{}
			t.wg = wg
			wg.Add(1)

			i++
			// // fPath := filepath.Join(dirPath, f.Name()) // f.Name() could be full path so line is useless
			// fPath := f.Name() // f.Name() could be full path so line is useless

			// fmt.Println(fPath)
			t.File, err = os.Open(fPath)
			if err != nil {
				fmt.Println("Error opening file.")
				return
			}
			defer t.File.Close()

			t.newFile, err = os.OpenFile(newWritePath, os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				fmt.Println("Error opening new file.")
				return
			}
			defer t.newFile.Close()

			go t.encrypt(key)
		}

		wg.Wait()

	} else {
		fmt.Println("Invalid selection.")
		return
	}
}
