package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}

/*
	1、确定状态"
	2、转义方程
	3、初始条件和边界条件
	4、计算顺序
*/
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	n := len(coins)
	last := n - 1

	//确定状态
	//1、最后一步 ：求f(x-ak)=y,其中y枚硬币，可以拼成x-ak的金额
	//2、子问题：f(x)=min{f(x-coins[i])+1}

	//转移方程：f(x)=y.x表示总金额，y表示总金额为x时的最少硬币数

	//初始条件和边界情况
	/*
		f(0)=0
		f(-)=∞
	*/

	//计算顺序，从小到大
	var min int
	for i, v := range coins {

	}

}

func f(sum int) int {
	if f < 0 {
		return int(math.MaxInt64)
	}

	if f == 0 {
		return 0
	}
}
