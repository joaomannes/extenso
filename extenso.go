package extenso_go

import (
	"log"
	"math"
	"strconv"
)

const numeroMaximo int64 = 999999999999999999
const tamanhoMilhar int = 3

//Extenso escreve um número por extenso
func Extenso(numero int64) string {

	verificaEstouroNumeroMaximo(numero)

	if numero == 0 {
		return "zero"
	}

	numeroExtenso := ""
	var numeroString string = strconv.FormatInt(numero, 10)

	tamanho := len(numeroString)

	var ternarios = int(math.Ceil(float64(tamanho) / float64(tamanhoMilhar)))

	numeroString = LeftPad(numeroString, "0", ternarios*3)

	for i := 0; i < ternarios; i++ {
		var parteNumero, _ = strconv.ParseInt(numeroString[i*3:3+i*3], 10, 0)

		if parteNumero == 0 {
			continue
		}

		centena := int(math.Floor(float64(parteNumero / 100)))

		dezena := int(math.Floor(float64((parteNumero - int64((centena * 100))) / 10)))

		unidade := int(parteNumero - int64(centena*100) - int64(dezena*10))

		if centena > 0 {
			if dezena == 0 && unidade == 0 && numeroExtenso != "" {
				numeroExtenso += " e "
			} else if numeroExtenso != "" {
				numeroExtenso += ", "
			}
			if centena == 1 {
				if dezena > 0 || unidade > 0 {
					numeroExtenso += centenas[centena].plural
				} else {
					numeroExtenso += centenas[centena].singular
				}
			} else {
				numeroExtenso += centenas[centena].singular
			}
		}

		if dezena > 0 {
			if numeroExtenso != "" {
				numeroExtenso += " e "
			}

			if dezena == 1 {
				dezena = 10 + unidade
				unidade = 0
				numeroExtenso += unidades[dezena]
			} else {
				numeroExtenso += dezenas[dezena]
			}
		}

		if unidade > 0 {
			if numeroExtenso != "" {
				numeroExtenso += " e "
			}
			numeroExtenso += unidades[unidade]
		}

		if parteNumero > 1 {
			numeroExtenso += " " + milhares[ternarios-i-1].plural
		} else {
			numeroExtenso += " " + milhares[ternarios-i-1].singular
		}

	}

	return numeroExtenso
}

func verificaEstouroNumeroMaximo(number int64) {
	if number > numeroMaximo {
		log.Fatal("Número informado maior que o máximo permitido")
	}
}
