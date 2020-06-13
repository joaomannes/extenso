package extenso_go

//LeftPad Preenche uma string à esquerda
func LeftPad(str, pad string, length int) string {
	for {
		if len(str) < length {
			str = pad + str
		} else {
			return str
		}
	}
}

//RightPad Preenche uma string à direita
func RightPad(str, pad string, length int) string {
	for {
		if len(str) < length {
			str += pad
		} else {
			return str
		}
	}
}
