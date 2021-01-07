package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type Apple struct {
	X int
	Y int
	Z int
	R int
}

func main() {

	var N, index int

	fmt.Printf("請輸入蘋果的數量(最小值1，最大值99)：")
	for {
		fmt.Scan(&N)
		if N >= 1 && N < 100 {
			break
		} else {
			fmt.Printf("輸入錯誤，請重新輸入(最小值1，最大值99)：")
		}
	}

	fmt.Printf(fmt.Sprintf("請輸入要從「第i個」蘋果開始掉落(最小值1，最大值%d)：", N))
	for {
		fmt.Scan(&index)
		if index >= 1 && index <= N {
			break
		} else {
			fmt.Println(fmt.Sprintf("輸入錯誤，請重新輸入(最小值1，最大值%d)：", N))
		}
	}

	// 這裡代表一團蘋果
	apples := make([]Apple, N)

	for i := 0; i < N; i++ {

		fmt.Scan(&apples[i].X, &apples[i].Y, &apples[i].Z, &apples[i].R)
	}

	fmt.Println(apples)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("0") // Write answer to stdout
}
