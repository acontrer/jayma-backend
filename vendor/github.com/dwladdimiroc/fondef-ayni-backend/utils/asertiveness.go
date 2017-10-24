package utils

import (
	"os/exec"
	"strconv"
	"strings"
)

func Asertiveness(text string) float64 {
	output, err := exec.Command("java", "-cp", Config.Asertividad.Jar, Config.Asertividad.Main, text).Output()
	Check(err)

	strings.TrimSpace(string(output))
	valor, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	Check(err)
	
	return valor
}
