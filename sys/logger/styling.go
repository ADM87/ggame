package logger

import "fmt"

// Note: This isn't cross-platform safe. Windows terminals may not support ANSI escape codes.
const (
	StyleRed     = "\033[31m%s\033[0m"
	StyleGreen   = "\033[32m%s\033[0m"
	StyleYellow  = "\033[33m%s\033[0m"
	StyleBlue    = "\033[34m%s\033[0m"
	StyleMagenta = "\033[35m%s\033[0m"
	StyleCyan    = "\033[36m%s\033[0m"
	StyleWhite   = "\033[37m%s\033[0m"
	StyleGray    = "\033[90m%s\033[0m"
)

func Colorize(color string, text string) string {
	return fmt.Sprintf(color, text)
}

func Red(args ...any) string {
	return fmt.Sprintf(StyleRed, fmt.Sprint(args...))
}

func Redf(format string, args ...any) string {
	return fmt.Sprintf(StyleRed, fmt.Sprintf(format, args...))
}

func Green(args ...any) string {
	return fmt.Sprintf(StyleGreen, fmt.Sprint(args...))
}

func Greenf(format string, args ...any) string {
	return fmt.Sprintf(StyleGreen, fmt.Sprintf(format, args...))
}

func Yellow(args ...any) string {
	return fmt.Sprintf(StyleYellow, fmt.Sprint(args...))
}

func Yellowf(format string, args ...any) string {
	return fmt.Sprintf(StyleYellow, fmt.Sprintf(format, args...))
}

func Blue(args ...any) string {
	return fmt.Sprintf(StyleBlue, fmt.Sprint(args...))
}

func Bluef(format string, args ...any) string {
	return fmt.Sprintf(StyleBlue, fmt.Sprintf(format, args...))
}

func Magenta(args ...any) string {
	return fmt.Sprintf(StyleMagenta, fmt.Sprint(args...))
}

func Magentaf(format string, args ...any) string {
	return fmt.Sprintf(StyleMagenta, fmt.Sprintf(format, args...))
}

func Cyan(args ...any) string {
	return fmt.Sprintf(StyleCyan, fmt.Sprint(args...))
}

func Cyanf(format string, args ...any) string {
	return fmt.Sprintf(StyleCyan, fmt.Sprintf(format, args...))
}

func White(args ...any) string {
	return fmt.Sprintf(StyleWhite, fmt.Sprint(args...))
}

func Whitef(format string, args ...any) string {
	return fmt.Sprintf(StyleWhite, fmt.Sprintf(format, args...))
}

func Gray(args ...any) string {
	return fmt.Sprintf(StyleGray, fmt.Sprint(args...))
}

func Grayf(format string, args ...any) string {
	return fmt.Sprintf(StyleGray, fmt.Sprintf(format, args...))
}
