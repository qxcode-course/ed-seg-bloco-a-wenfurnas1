package main

import "fmt"
//esta atividade eu fiz no dia 25/03/2026 no fup is fun e copiei para cá, não usei IA, apenas como tenho tempo limitado fiz 
// no moodl2 pois estou sem maquina;

func main() {
	quantidade := 0
	fmt.Scanf("%d", &quantidade)
	ganhador := 1000
	posicao := 0
	a_resultado := 0
	auxiliar := 0
	for i := 0; i < quantidade; i++ {
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)
		if a >= 10 && b >= 10 {
			if (a-b == a_resultado) || (b-a == a_resultado) {
				if a < auxiliar {
					posicao = i + 1
				}
				if b < auxiliar {
					posicao = i
				}

			}
			if a > b {
				auxiliar = a
				a_resultado = a - b
			} else {
				auxiliar = b
				a_resultado = b - a
			}
			if a_resultado < ganhador {
				ganhador = a_resultado
				posicao = i
			}

		}

	}

	if ganhador == 1000 {
		fmt.Println("sem ganhador")
	} else {

		fmt.Println(posicao)
	}

}
