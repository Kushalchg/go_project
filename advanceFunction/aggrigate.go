package advancefunction

func Add(a, b int) int {
	return a + b

}
func Mul(a, b int) int {
	return a * b

}
func Avg(a, b int) int {
	return (a + b) / 2

}

func Aggrigate(a, b, c int, arthmetic func(int, int) int) int {
	return arthmetic(arthmetic(a, b), c)
}
