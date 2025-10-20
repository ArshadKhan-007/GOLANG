package main

import (
	"fmt"
	"os"
)

func main() {
	// Get file information
	// Here, os.Open ask for error handling because the file may not exist
	// f, is a pointer to os.File means file object
	// err, is an error type
	f, err := os.Open("file.txt") // Ensure the file exists in the same directory
	// Panic if there is an error
	if err != nil {
		panic(err)
	}
	// Here stat() method also ask for error handling
	// fileInfo is of type FileInfo
	// err is of type error
	fileInfo, err := f.Stat() //stat() method returns a FileInfo type
	if err != nil {
		panic(err)
	}
	fmt.Println("File Name:", fileInfo.Name())          // Name() method returns the name of the file
	fmt.Println("File or Folder:", fileInfo.IsDir())    // IsDir() method returns true if it is a directory
	fmt.Println("File Size:", fileInfo.Size(), "bytes") // Size() method returns the size in bytes
	fmt.Println("File Permission:", fileInfo.Mode())    // Mode() method returns the file permission
	fmt.Println("Last Modified:", fileInfo.ModTime())   // ModTime() method returns the last modified time

	fmt.Println("---------")

	// Reading Contents of the file
	f, err = os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close() // Close the file when the function ends

	data := make([]byte, fileInfo.Size())
	d, err := f.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bytes Read:", d)
	fmt.Println("Data Read:", string(data))

	fmt.Println("---------")

	// Read Folders
	folder, err := os.Open("../") // Open the parent directory,If we pass the current directory that will return false for IsDir()
	// Bcz we have files here not folders
	if err != nil {
		panic(err)
	}
	defer folder.Close()

	folderInfo, err := folder.ReadDir(-1) // ReadDir(-1) reads all the directory entries

	for _, fi := range folderInfo {
		fmt.Println(fi.Name(), fi.IsDir())
	}

	// Note: os.Open can open both files and folders

	fmt.Println("---------")

	// Create a file
	file, err := os.Create("file2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write data to the file
	file.WriteString("Hii This is file2")
	fmt.Println("file2.txt created and data written successfully")

}
