package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	res := make([][]string, n) //res二维数组记录状态

	res[0] = []string{""} //初始条件
	//计算顺序为从小到大
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {

			//状态转移方程，res[i] = "("+A+")"+B ,其中A和B为res[i-1]的子集
			for _, p := range res[j] {
				for _, q := range res[i-j-1] {
					str := "(" + string(p) + ")" + string(q)
					res[i] = append(res[i], str)
				}
			}

		}
	}

	return res[n]
}
