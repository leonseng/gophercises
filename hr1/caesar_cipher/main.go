package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
    runes := []rune(s)
    rotFactor := k % 26
    for i, r := range runes {
        if (r >= 65 && r <= 90) || (r >= 97 && r <= 122) {
            newRune := r + rotFactor
            if (newRune > 90 && newRune <= 90 + rotFactor) || newRune > 122 {
                newRune -= 26
            }
            runes[i] = newRune
        }
    }

    return string(runes)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    strconv.ParseInt(readLine(reader), 10, 64)

    s := readLine(reader)

    kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := caesarCipher(s, k)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
