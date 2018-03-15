// De um a um trilhão
func intPorExtenso(n int64) string {
	negativo := n < 0
	menos := ""

	if negativo {
		n *= -1
		menos = "menos "
	}

	const (
		UM_TRILHAO = 1000000000000
		UM_BILHAO  = 1000000000
		UM_MILHAO  = 1000000
		MIL        = 1000
		CEM        = 100
		VINTE      = 20
	)

	if n < VINTE {
		a := [...]string{"", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove", "dez", "onze", "doze", "treze", "quatorze", "quinze", "dezesseis", "dezessete", "dezoito", "dezenove"}

		return fmt.Sprintf("%s%s", menos, a[n])
	}

	if n < CEM {
		quantos := n / 10
		resto := n % 10

		a := [...]string{"", "", "vinte", "trinta", "quarenta", "cinqüenta", "sessenta", "setenta", "oitenta", "noventa"}
		e := " "

		if resto > 0 {
			e = " e "
		}

		return strings.Replace(fmt.Sprintf("%s%s%s%s", menos, a[quantos], e, intPorExtenso(resto)), "  ", " ", -1)
	}

	if n < MIL {
		if n == CEM {
			return "cem"
		}

		quantos := n / CEM
		resto := n % CEM

		a := [...]string{"", "cento", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seicentos", "setecentos", "oitocentos", "novecentos"}
		e := " "

		if resto > 0 {
			e = " e "
		}

		return strings.Replace(fmt.Sprintf("%s%s%s%s", menos, a[quantos], e, intPorExtenso(resto)), "  ", " ", -1)
	}

	if n < UM_MILHAO {
		if n == MIL {
			return "mil"
		}

		quantos := n / MIL
		resto := n % MIL
		e := " "

		if ( resto > 0 ) && ( resto < CEM ) {
			e = " e "
		}

		return strings.Replace(fmt.Sprintf("%s%s mil %s%s", menos, intPorExtenso(quantos), e, intPorExtenso(resto)), "  ", " ", -1)
	}

	if n < UM_BILHAO {
		if n == UM_MILHAO  {
			return "um milhão"
		}

		quantos := n / UM_MILHAO
		resto := n % UM_MILHAO
		e := " "
		u := " milhão "

		if quantos>1 {
			u = " milhões "
		}

		if ( resto > 0 ) && ( resto < MIL ) {
			e = " e "
		}

		return strings.Replace(fmt.Sprintf("%s%s%s%s%s", menos, intPorExtenso(quantos), u, e, intPorExtenso(resto)), "  ", " ", -1)
	}

	if n < UM_TRILHAO {
		if n == UM_BILHAO {
			return "um bilhão"
		}

		quantos := n / UM_BILHAO
		resto := n % UM_BILHAO
		e := " "
		u := " bilhão "

		if quantos>1 {
			u = " bilhões "
		}

		if ( resto > 0 ) && ( resto < MIL ) {
			e = " e "
		}

		return strings.Replace(fmt.Sprintf("%s%s%s%s%s", menos, intPorExtenso(quantos), u, e, intPorExtenso(resto)), "  ", " ", -1)
	}

	if n==UM_TRILHAO {
		return "um trilhão"
	}

	return fmt.Sprintf("%s%d", menos, n)
}
