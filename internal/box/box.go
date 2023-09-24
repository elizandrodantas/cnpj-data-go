package box

import "fmt"

var (
	horizontal = "─"
	vertical   = "│"

	leftCornerTop  = "┌"
	rigthCornerTop = "┐"

	leftCornerLower  = "└"
	rigthCornerLower = "┘"

	padding = 4
)

func New(t string) {
	println()

	width := len(t) + padding

	// higher
	fmt.Printf("%s%s%s\n", leftCornerTop, repeat(horizontal, width-2), rigthCornerTop)

	// center
	fmt.Printf("%s %s %s\n", vertical, center(t, width-4), vertical)

	// lower
	fmt.Printf("%s%s%s\n", leftCornerLower, repeat(horizontal, width-2), rigthCornerLower)

	println()
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func center(s string, width int) string {
	if len(s) >= width {
		return s
	}
	padding := width - len(s)
	leftPadding := padding / 2
	rightPadding := padding - leftPadding
	return repeat(" ", leftPadding) + s + repeat(" ", rightPadding)
}
