package module1

//func DecToBase(n, convBase int) string {
//	var base []string
//	var curr int
//	for n >= convBase {
//		curr, n = n % convBase, n / convBase
//		base = append([]string{strconv.Itoa(curr)}, base...)
//	}
//	base = append([]string{strconv.Itoa(n)}, base...)
//	return strings.Join(base, "")
//}

//func DecToBase(n, convBase int) string {
//	var res string
//	var rem int
//	for n >0 {
//		rem, n = n % convBase, n / convBase
//		res = fmt.Sprintf("%X%s", rem, res)
//	}
//	return res
//}

func DecToBase(n, convBase int) string {
	const charset = "0123456789ABCDEF"
	var res string
	var rem int
	for n > 0 {
		rem, n = n%convBase, n/convBase
		res = string(charset[rem]) + res
	}
	return res
}
