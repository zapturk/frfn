package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    argsWithoutProg := os.Args[1:]
    oldText := ""
    newText := ""

    for index, arg := range argsWithoutProg {
        switch arg {
        case "--OldText":
        case "-ot":
            oldText = argsWithoutProg[index+1]
            break
        case "--NewText":
        case "-nt":
            newText = argsWithoutProg[index+1]
        default:
            break
        }
    }



    // gets the files in this directory
    dir, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }

    for _, entry := range dir {
        oldFileName := entry.Name()
        if strings.Contains(oldFileName, oldText){
            newFileName := strings.ReplaceAll(oldFileName, oldText, newText)
            oldPath := "./" + oldFileName
            newPath := "./" + newFileName
            fmt.Println(oldFileName + " changed to " + newFileName)

            // try to rename the file
            err := os.Rename(oldPath, newPath)
            if err != nil {
                fmt.Println("Failed to rename file" + oldFileName)
                return
            }
        }
    }
}
