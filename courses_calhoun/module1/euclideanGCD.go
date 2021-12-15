package module1

func EuclideanGCD(a, b int) int {
	//if b < 1 {
	//	return a
	//}
	//a,b = b, a%b
	//return EuclideanGCD(a, b)

	//Step 2
	//for b != 0 {
	//	a,b = b, a%b
	//}
	//return a

	//Step 3

	for {
		if b < 1 {
			return a
		}
		a, b = b, a%b
	}
}
