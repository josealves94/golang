package main

import (
        "fmt"
        "io"
        "os"
)

var path = "/home/jalves/golang/test_file.txt"


func main() {
        createFile()
        writeFile()
        readFile()
        deleteFile()
}

func createFile() {
        // detect if file exists
        var _, err= os.Stat(path)

        // Create file if not exists

        if os.IsNotExist(err) {
                var file, err = os.Create(path)
                checkError(err)
                defer file.Close()
        }

}


func writeFile() {
        // Open file using read and write permission
        var file, err = os.OpenFile(path, os.O_RDWR, 0644)
        checkError(err)
        defer file.Close()

        // Write some text to file


        _, err = file.WriteString("hello\n")
        checkError(err)
        _, err = file.WriteString("I'm learning Golang\n")
        checkError(err)

        // save changes
        err = file.Sync()
        checkError(err)

}


func readFile() {
        // re-open file
        var file, err = os.OpenFile(path, os.O_RDWR, 0644)
        checkError(err)
        defer file.Close()

        // read file
        var text = make([]byte, 1024)
        for {

                n, err := file.Read(text)
                if err != io.EOF {
                        checkError(err)
                }

                if n == 0 {
                        break
                }
        }
        fmt.Println(string(text))
        checkError(err)
}

func deleteFile() {
        // delete file
        var err = os.Remove(path)
        checkError(err)
}

func checkError(err error) {
        if err != nil {
                fmt.Println(err.Error())
                os.Exit(0)
        }
}
