package websocket

type Retorno struct {
	Resposta string
	Latencia float64
}

func getterRetorno() Retorno {
	var retorno Retorno
	return retorno
}
