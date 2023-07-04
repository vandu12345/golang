// Types Of Package
// Executable -> Generate a file that we can Execute
// Reusable -> Code used as helpers ,Good Place to put reusable logic
package Packages

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
