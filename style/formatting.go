package style

import (
    "fmt"
    "math"
    "strconv"
)

type Styler struct {
    colors map[string]string
    reset  bool
}

func NewStyler() *Styler {
    return &Styler{
        colors: map[string]string{
            "reset":        "\033[0m",
        },
    }
}

func (s *Styler) hexToRGB(hexColor string) (int, int, int) {
    hexColor = hexColor[1:]
    r, _ := strconv.ParseInt(hexColor[0:2], 16, 0)
    g, _ := strconv.ParseInt(hexColor[2:4], 16, 0)
    b, _ := strconv.ParseInt(hexColor[4:6], 16, 0)
    return int(r), int(g), int(b)
}

func (s *Styler) rgbToANSI(rgbColor [3]int) string {
    r, g, b := rgbColor[0], rgbColor[1], rgbColor[2]
    return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func (s *Styler) formatColor(text string, color string) string {
    return fmt.Sprintf("%s%s%s", s.colors[color], text, s.colors["reset"])
}

func (s *Styler) applyGradient(text string, startColor string, endColor string) string {
    startR, startG, startB := s.hexToRGB(startColor)
    endR, endG, endB := s.hexToRGB(endColor)
    gradient := ""
    for i, char := range text {
        r := startR + int(math.Round(float64(endR-startR)*float64(i)/float64(len(text))))
        g := startG + int(math.Round(float64(endG-startG)*float64(i)/float64(len(text))))
        b := startB + int(math.Round(float64(endB-startB)*float64(i)/float64(len(text))))
        gradient += fmt.Sprintf("%s%c", s.rgbToANSI([3]int{r, g, b}), char)
    }
    return gradient + s.colors["reset"]
}

func (s *Styler) applyStyle(text string, style string) string {
    styles := map[string]string{
        "strikethrough": "\033[9m",
        "bold":          "\033[1m",
        "italic":        "\033[3m",
        "underline":     "\033[4m",
        "reset":         "\033[0m",
    }
    return fmt.Sprintf("%s%s%s", styles[style], text, styles["reset"])
}

func (s *Styler) Style(text string, color string, gradient []string, style string) string {
    if gradient != nil {
        startColor, endColor := gradient[0], gradient[1]
        text = s.applyGradient(text, startColor, endColor)
    }
    if color != "" {
        text = s.formatColor(text, color)
    }
    if style != "" {
        text = s.applyStyle(text, style)
    }
    return text
}
