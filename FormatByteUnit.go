package csy

import "fmt"

func FormatByteUnit(bytes float64) string {
	const (
		KB = 1 << 10
		MB = 1 << 20
		GB = 1 << 30
		TB = 1 << 40
		PB = 1 << 50
		EB = 1 << 60
		ZB = 1 << 70
		YB = 1 << 80
	)

	switch {
	case bytes >= YB:
		return fmt.Sprintf("%.2f YB", float64(bytes)/YB)
	case bytes >= ZB:
		return fmt.Sprintf("%.2f ZB", float64(bytes)/ZB)
	case bytes >= EB:
		return fmt.Sprintf("%.2f EB", float64(bytes)/EB)
	case bytes >= PB:
		return fmt.Sprintf("%.2f PB", float64(bytes)/PB)
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/TB)
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}
