package prayers

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
)

// Color codes list : https://www.ditig.com/256-colors-cheat-sheet

func BasmalahFormatter(basmalah string) string {
	switch userOS := runtime.GOOS; userOS {

	case "darwin":
		return fmt.Sprintf("%-7s%s%-3s\n", "", basmalah, "")
	case "linux":
		return fmt.Sprintf("%-6s%s%-3s\n", "", basmalah, "")

	default:
		return ""
	}
}

func BasmalahPrinter(basmalahSprint string) {
	// Choose colors here
	switch userOS := runtime.GOOS; userOS {

	case "darwin":
		fmt.Printf(basmalahSprint)
	case "linux":
		whitePrint := color.New(color.FgHiWhite, color.Bold).PrintFunc()
		whitePrint(basmalahSprint)

	default:
		return
	}
}

func DateFormatter(todayDate string) string {
	switch userOS := runtime.GOOS; userOS {

	case "darwin":
		return fmt.Sprintf("%-1s%s\n", "", todayDate)
	case "linux":
		return fmt.Sprintf("%-1s%s\n", "", todayDate)

	default:
		return ""
	}
}

func DatePrinter(dateSprint string) {
	// Choose colors here
	switch userOS := runtime.GOOS; userOS {

	case "darwin":
		fmt.Printf(boldColor256(27, dateSprint))
	case "linux":
		fmt.Printf(boldColor256(61, dateSprint))

	default:
		return
	}
}

func LoaderFormatter(nextPrayerName string, loadingSquares string, timeRemaining string) string {
	switch userOS := runtime.GOOS; userOS {

	case "darwin":
		return fmt.Sprintf("%-1sNext: %s %s %s\n", "", nextPrayerName, loadingSquares, timeRemaining)
	case "linux":
		return fmt.Sprintf("%-1sNext: %s %s %s\n", "", nextPrayerName, loadingSquares, timeRemaining)

	default:
		return ""
	}
}

func LoaderPrinter(loaderSprint string) {
	// Choose colors here
	switch userOS := runtime.GOOS; userOS {

	case "darwin":
		fmt.Printf(boldColor256(21, loaderSprint))
	case "linux":
		fmt.Printf(boldColor256(122, loaderSprint))

	default:
		return
	}
}
