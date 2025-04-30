package utils

import (
	"strconv"
	"strings"
)

func ValidateCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	repeated := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			repeated = false
			break
		}
		if repeated {
			return false
		}
	}

	for i := 9; i < 11; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			num, _ := strconv.Atoi(string(cpf[j]))
			sum += num * (i + 1 - j)
		}
		dv := (sum * 10) % 11
		if dv == 10 {
			dv = 0
		}
		if dv != int(cpf[i]-'0') {
			return false
		}
	}

	return true
}
