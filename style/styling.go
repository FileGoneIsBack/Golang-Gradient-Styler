package style

import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"regexp"
)

func ApplyStylingToFile(inputFile, outputFile string, styler *Styler) error {
    input, err := os.Open(inputFile)
    if err != nil {
        return fmt.Errorf("error opening input file: %w", err)
    }
    defer input.Close()

    output, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("error creating output file: %w", err)
    }
    defer output.Close()

    scanner := bufio.NewScanner(input)
    writer := bufio.NewWriter(output)
    defer writer.Flush()

    for scanner.Scan() {
        line := scanner.Text()
        styledLine := ApplyStylingToLine(line, styler)
        _, err := writer.WriteString(styledLine + "\n")
        if err != nil {
            return fmt.Errorf("error writing to output file: %w", err)
        }
    }

    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error scanning input file: %w", err)
    }

    return nil
}

func ApplyStylingToLine(line string, styler *Styler) string {
    reColor := regexp.MustCompile(`<<Start:\s*#?([#a-fA-F0-9]+)\s*>>\s*(.*?)\s*<<End:\s*#?([#a-fA-F0-9]+)\s*>>`)
    matchesColor := reColor.FindAllStringSubmatch(line, -1)

    for _, match := range matchesColor {
        startColor := match[1]
        endColor := match[3]

        startColorWithHash := "#" + startColor
        endColorWithHash := "#" + endColor

        text := match[2]

        styledText := styler.Style(text, "", []string{startColorWithHash, endColorWithHash}, "")

        line = strings.ReplaceAll(line, match[0], styledText)
    }


    reSpecial := regexp.MustCompile(`<<Start:\s*(.*?)>>\s*(.*?)\s*<<End:\s*(.*?)>>`)
    matchesSpecial := reSpecial.FindAllStringSubmatch(line, -1)

    for _, match := range matchesSpecial {
        tag := match[1]
        text := match[2]

        var styledText string
        switch tag {
        case "Bold":
            styledText = styler.applyStyle(text, "bold")
        case "Strike":
            styledText = styler.applyStyle(text, "strikethrough")
		case "Italic":
			styledText = styler.applyStyle(text, "italic")
		case "Underline":
			styledText = styler.applyStyle(text, "underline")
        default:
            styledText = text
        }

        line = strings.ReplaceAll(line, match[0], styledText)
    }

    return line
}
