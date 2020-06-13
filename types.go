package extenso

type palavra struct {
	singular string
	plural   string
}

var unidades = map[int]string{
	1:  "um",
	2:  "dois",
	3:  "três",
	4:  "quatro",
	5:  "cinco",
	6:  "seis",
	7:  "sete",
	8:  "oito",
	9:  "nove",
	10: "dez",
	11: "onze",
	12: "doze",
	13: "treze",
	14: "quatorze",
	15: "quinze",
	16: "dezesseis",
	17: "dezessete",
	18: "dezoito",
	19: "dezenove"}

var dezenas = map[int]string{
	2: "vinte",
	3: "trinta",
	4: "quarenta",
	5: "cinquenta",
	6: "sessenta",
	7: "setenta",
	8: "oitenta",
	9: "noventa"}

var centenas = map[int]palavra{
	1: palavra{"cem", "cento"},
	2: palavra{"duzentos", "duzentos"},
	3: palavra{"trezentos", "trezentos"},
	4: palavra{"quatrocentos", "quatrocentos"},
	5: palavra{"quinhentos", "quinhentos"},
	6: palavra{"seiscentos", "seiscentos"},
	7: palavra{"setecentos", "setecentos"},
	8: palavra{"oitocentos", "oitocentos"},
	9: palavra{"novecentos", "novecentos"}}

var milhares = []palavra{
	palavra{"", ""},
	palavra{"mil", "mil"},
	palavra{"milhão", "milhões"},
	palavra{"bilhão", "bilhões"},
	palavra{"trilhão", "trilhões"},
	palavra{"quatrilhão", "quatrilhões"}}
