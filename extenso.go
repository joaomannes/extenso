package extenso

import (
	"errors"
	"math"
	"strconv"
)

const nMax int64 = 999999999999999999
const tamanhoMilhar int = 3

//From escreve um número por extenso
func From(n int64) (string, error) {

	var err = verificaEstouroNumeroMaximo(n)
	if err != "" {
		return "", errors.New(err)
	}

	err = verificaNumeroNegativo(n)
	if err != "" {
		return "", errors.New(err)
	}

	if n == 0 {
		return "zero", nil
	}

	extenso := ""
	var numeroString string = strconv.FormatInt(n, 10)

	tamanho := len(numeroString)

	var ternarios = int(math.Ceil(float64(tamanho) / float64(tamanhoMilhar)))

	numeroString = leftPad(numeroString, "0", ternarios*3)

	for i := 0; i < ternarios; i++ {
		var parteNumero, _ = strconv.ParseInt(numeroString[i*3:3+i*3], 10, 0)

		if parteNumero == 0 {
			continue
		}

		centena := int(math.Floor(float64(parteNumero / 100)))

		dezena := int(math.Floor(float64((parteNumero - int64((centena * 100))) / 10)))

		unidade := int(parteNumero - int64(centena*100) - int64(dezena*10))

		if centena > 0 {
			if dezena == 0 && unidade == 0 && extenso != "" {
				extenso += " e "
			} else if extenso != "" {
				extenso += ", "
			}
			if centena == 1 {
				if dezena > 0 || unidade > 0 {
					extenso += centenas[centena].plural
				} else {
					extenso += centenas[centena].singular
				}
			} else {
				extenso += centenas[centena].singular
			}
		}

		if dezena > 0 {
			if extenso != "" {
				extenso += " e "
			}

			if dezena == 1 {
				dezena = 10 + unidade
				unidade = 0
				extenso += unidades[dezena]
			} else {
				extenso += dezenas[dezena]
			}
		}

		if unidade > 0 {
			if extenso != "" {
				extenso += " e "
			}
			extenso += unidades[unidade]
		}

		if parteNumero > 1 {
			extenso += " " + milhares[ternarios-i-1].plural
		} else {
			extenso += " " + milhares[ternarios-i-1].singular
		}

	}

	return extenso, nil
}

func verificaEstouroNumeroMaximo(n int64) string {
	if n > nMax {
		return "Número informado maior que o máximo permitido"
	} else {
		return ""
	}
}

func verificaNumeroNegativo(n int64) string {
	if n < 0 {
		return "Não é possível escrever números negativos"
	} else {
		return ""
	}
}
