package extenso

//leftPad Preenche uma string à esquerda
func leftPad(str, pad string, length int) string {
	for {
		if len(str) < length {
			str = pad + str
		} else {
			return str
		}
	}
}

//rightPad Preenche uma string à direita
func rightPad(str, pad string, length int) string {
	for {
		if len(str) < length {
			str += pad
		} else {
			return str
		}
	}
}
