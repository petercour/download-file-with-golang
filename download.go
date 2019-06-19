// Download file from URL
// https://golangr.com
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("usage: go run download.go https://golang.org/doc/gopher/pkg.png ")
        os.Exit(1)
    }
    url := os.Args[1]
    filename := filepath.Base(url)
    
    err := downloadFile(url, filename)
    if err != nil {
        panic(err)
    }

}

// The function downloadFile downloads the file and saves it to
// the destination. It does not load the entire file into memory,
// so you can download large files.
//
func downloadFile(url string, filepath string) error {
    // Create file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Get data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Write body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}
