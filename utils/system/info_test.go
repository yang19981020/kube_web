package	system

import (
	"fmt"
	"testing"
)

func TestGetCpuPercent(t *testing.T) {
	percent := GetCpuPercent()
	fmt.Print(percent)
}
func TestGetMemPercent(t *testing.T) {
	percent := GetMemPercent()
	fmt.Print(percent)
}