package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func fib1(n int) int {
	if n < 0 {
		//fmt.Println(n, " idx")
		return -1
	}
	if n == 1 || n == 2 {
		fmt.Println(n, " idx")
		fmt.Println()
		return 1
	}
	fmt.Println(n, " idx----")
	return fib1(n-1) + fib1(n-2)
}

func fib(n int) int {
	dp := map[int]int{1: 1, 2: 1}
	if n < 0 {
		return -1
	}
	if n < 2 {
		return 1 // base class
	}
	if _, ok := dp[n-1]; !ok {
		dp[n-1] = fib(n - 1)
	}
	if _, ok := dp[n-2]; !ok {
		dp[n-2] = fib(n - 2)
	}
	return dp[n-1] + dp[n-2]
}

func solve(idx, n int, dp map[int]int) int {

	if idx > n {
		fmt.Println(idx, " idx")
		return idx
	}

	if value, ok := dp[idx]; ok {
		return value
	}
	return solve(idx+1, n, dp) + solve(idx+2, n, dp)
}

func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum(n-1)
}
func getFloorNumber(n int) int {
	return int(n / 2)
}

// for sorted arrays
func binarySearch(listOfInt []int, searchItem int) int {
	copyList := make([]int, len(listOfInt))
	copy(copyList, listOfInt)
	middleIndexItem := getFloorNumber(len(copyList))
	middleItem := copyList[middleIndexItem]
	if middleItem == searchItem {
		return middleItem
	}
	if len(copyList) < 2 {
		return -1
	}
	if middleItem > searchItem {
		copyList = copyList[:middleIndexItem]
	}
	if middleItem < searchItem {
		copyList = copyList[middleIndexItem+1:]
	}
	return binarySearch(copyList, searchItem)
}

func main() {
	//dp := map[int]int{1: 1, 2: 1}
	//fmt.Println(solve(3, 4, dp))
	//fmt.Println(sum(4))
	//a := []int{1, 34, 12, 3, 12, 7, 9, 10, 12, 5, 6, 13}
	//b := []int{34, 23, 42, 5, 10, 23, 4, 22}
	//c := make([]int, len(b))
	//copy(c, b)
	//c[0] = 100
	//fmt.Println(c, b)
	//fmt.Println()
	//a = append(a[:3], a[7:]...) // cut index 3
	//a = append(a[:1], a[1+1:]...)
	//fmt.Println(a)
	//fmt.Println()
	//sortedArray := []int{1, 5, 8, 9, 11, 13, 15, 19, 21}
	//fmt.Println(binarySearch(sortedArray, 1))
	//fmt.Println(sortedArray)
	//for _, r := range "hello" {
	//	fmt.Println(r)
	//}
	//ls := make([]int, 5)
	//fmt.Println(ls)
	//
	//countryPopulation := make(map[string]float64)
	//fmt.Println()
	//fmt.Println(countryPopulation["name"])
	//getListOfUrl("../whatsJoblist.txt")
	var (
		i int
		c int
	)
	fmt.Println(c, i)
}

func ReadMyFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file, nil
}

func WriteMyFile(filePath string, file string) error {
	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	w := bufio.NewWriter(newFile)
	_, werr := w.WriteString(file)
	w.Flush()
	return werr
}

func getListOfUrl(filePath string) {
	const XMLPattern = `(?P<head>http[s]?:\/\/)(?P<path>[^\s]*)`
	linkRegex := regexp.MustCompile(XMLPattern)
	newLinks := `%d)  %s`
	reader, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	links := linkRegex.FindAllString(string(reader), -1)
	var savedLinks []string

	for idx, link := range links {
		savedLinks = append(savedLinks, fmt.Sprintf(newLinks, idx, link))
	}

	error := WriteMyFile("myNewList.txt", strings.Join(savedLinks, "\n"))
	if error != nil {
		panic(error)
	}

}
