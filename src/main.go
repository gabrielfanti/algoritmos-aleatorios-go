package main

import (
	"algoritmos-aleatorios/src/exercicio11"
	"algoritmos-aleatorios/src/exercicio12"
	"algoritmos-aleatorios/src/exercicio13"
	"algoritmos-aleatorios/src/exercicio14"
	"algoritmos-aleatorios/src/exercicio15"
	"algoritmos-aleatorios/src/exercicio16"
	"fmt"
)

func main() {
	var exercicio int

	fmt.Println("----------Listagem de exercícios----------")
	fmt.Println("A1. Two Fer (2-Fer)")
	fmt.Println("A2. Troca")
	fmt.Println("A3. Ajuste Salarial")
	fmt.Println("A4. FizzBuzz")
	fmt.Println("A5. Área de uma circunferência")
	fmt.Println("A6. Soma dos pares)")
	fmt.Println("A7. Salário Líquido Professor")
	fmt.Println("A8. Volume da caixa")
	fmt.Println("A9. Conversão em Real")
	fmt.Println("A10. Conversão em Dólar")
	fmt.Println("A11. Soma dos Quadrados")
	fmt.Println("A12. Quadrado da Soma")
	fmt.Println("A13. Velocidade do Projétil")
	fmt.Println("A14. Eleição Sindical")
	fmt.Println("A15. Média do Aluno")
	fmt.Println("A16. Somando Naturais")
	fmt.Println("A17. Ordenando Inteiros")
	fmt.Println("A18. Palíndromo")
	fmt.Println("A19. Scrabble Score")
	fmt.Println("A20. Fatorial")
	fmt.Println("A21. Algarismos Romanos")
	fmt.Println("A22. Primo")
	fmt.Println("A23. JoKenPo ")
	fmt.Println("A24. Pedra-papel-tesoura-lagarto-Spock")
	fmt.Println("A25. Soma dos pares")
	fmt.Println("A26. Anagramas")
	fmt.Println("A27. Letras repetidas")
	fmt.Println("A28. Fibonacci")
	fmt.Println("A29. Pangrama")
	fmt.Println("A30. Troco")
	fmt.Println("A31. Quebra de Linha")
	fmt.Println("A32. Sudoku")
	fmt.Println("----------------------------------------")

	fmt.Print("Digite o número do exercício que deseja executar): ")
	fmt.Scanln(&exercicio)

	switch exercicio {
	case 11:
		exercicio11.SomaQuadrados()
	case 12:
		exercicio12.QuadradoSoma()
	case 13:
		exercicio13.VelocidadeProjetil()
	case 14:
		exercicio14.EleicaoSindical()
	case 15:
		exercicio15.MediaAluno()
	case 16:
		exercicio16.SomandoNaturais()
	}

}
