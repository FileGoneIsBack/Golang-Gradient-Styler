package main

import (
    "fmt"
    "io/ioutil"
	"styler/style"
)

func main() {
	styler := style.NewStyler()

    inputFile := "textfile.txt"
    outputFile := "newtext.txt"

	if err := style.ApplyStylingToFile(inputFile, outputFile, styler); err != nil {
		fmt.Println("Error:", err)
		return
	}
    styledText, err := ioutil.ReadFile(outputFile)
    if err != nil {
        fmt.Println("Error reading output file:", err)
        return
    }

    fmt.Println("Styled text from", outputFile+":")
    fmt.Println(string(styledText))
}  

//made by filegone (for once just me!)