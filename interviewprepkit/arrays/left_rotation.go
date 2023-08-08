package arrays

func LeftRotation(arr []int32, d int32) []int32 {
	// Tomando com base o seguinte array: [1, 2, 3, 4, 5], e rotacionaremos ele 4 vezes

	/*
		Podemos dividir em duas partes:
		- A primeira que seria os elementos que efetivamente irão ser "rotacionados", ou seja, eles sairão da parte
		esquerda e serão colocados na última parte do array.

		arr[:(int(d)%len(arr))]
		[1, 2, 3, 4, 5] -> [:(4%5)] -> [:4] = [1, 2, 3, 4]

		- A segunda seria os elementos restantes, deslocados internamente pela saída dos elementos
		a partir da esquerda.

		arr[(int(d)%len(arr)):]
		[1, 2, 3, 4, 5] -> [(4%5):] -> [4:] = [5]

		- Por fim, apenas agrupamos a segunda parte com a primeira.

		[5] + [1, 2, 3, 4] = [5, 1, 2, 3, 4]
	*/

	rotatedArr := append([]int32{}, arr[(int(d)%len(arr)):]...)
	rotatedArr = append(rotatedArr, arr[:(int(d)%len(arr))]...)

	return rotatedArr
}
