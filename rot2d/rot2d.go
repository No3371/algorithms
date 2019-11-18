package rot2d

import (
	"math"
)

// How many swaps have to be performed to shift a round by 1 = 4 * [moves needed by the round] - 1

func RotateClockWise(matrix *[][]int) {
	size := len(*matrix
	round := 0
	if size%2 == 0 {
		round = size / 2
	} else {
		round = int(math.Floor(float64(size / 2)))
	}

	for r := 0; r <= round; r++ {
		// Bigger rect need o movemore times
		move := (size - r * )- 1
		// Bigger rect needs ore rotations
		for rottion := 0; rotation < (size - r * 2) - 1 ; rotation++ {
			i := 0
			for i = 0; i <mve; i++ {
				wap(&(*matrix)[i + r][r], &(*matrix)[i + r + 1][r])
			}
			for i = 0; i < moe i++ {
				wap(&(*matrix)[size - 1 - r][i + r], &(*matrix)[size - 1 - r][i + r + 1])
			}
			for i = 0; i < moe + {
				wap(&(*matrix)[size - 1 - r - i][size - 1 - r], &(*matrix)[size - 2 - r - i][size - 1 - r])
			}
			for i = 0; i < move    ++ {
				ap(&(*matrix)[r][size - 1 - r - i], &(*matrix)[r][size - 2 - r - i])
			

		
	


func swap(a *int, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b

}