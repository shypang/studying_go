package main

import (
	"fmt"
	"math"
	"sort"
)

func Solution(a int, b int, c int, d int) int {

	//문제 조건
	//네 주사위 값이 모두 p로 같다면 1111 곱하기 p
	//세 주사위가 p로 같고 나머지 주사위에서 나온 숫자가 q라면 (10 X p + q)^2
	//주사위가 두 개씩 같은 값이 나오고, 나온 숫자를 각각 p,q(p!=q)면 (p+q) * 절대값 p-q
	//두 주사위는 p로 같고, 나머지 두 주사위에서 나온 숫자가 p와 다른 q, r이라면 (q !=r) q * r 점을 얻는다.
	//네 주사위에 적힌 숫자가 모두 다르면 나온 숫자 중 가장 작은 숫자만큼의 점수를 얻는다.

	//네개가 같은 경우 다음과 같이 구현
	//네개가 다를 경우 오름차순으로 정렬한 다음 가장 앞의 값을 반환하기 이것도 슬라이스를 사용해야할 것 같음
	//두개씩 같은 경우 연산해서 p+q * p-q 계산후 리턴 이것도 슬라이스 사용해야할 것 같음
	//두 주사위는 같고 각각 두개가 다른 경우에는 q * r 반환
	//세 개는 p로 같고 나머지주사위에서 나온 숫자가 q면 10 x p + q의 제곱, 슬라이스로 앞의3개, 뒤의3개 계산하면 될 것 같음

	arr := []int{}
	arr = append(arr, a)
	arr = append(arr, b)
	arr = append(arr, c)
	arr = append(arr, d)
	answer := 0
	sort.Sort(sort.IntSlice(arr))
	// fmt.Println("arr", arr)

	//4개가 같은 경우 1111 * p 반환
	//2222
	if arr[0] == arr[1] && arr[2] == arr[3] && arr[1] == arr[2] {
		answer = 1111 * a
		//4개가 다를 경우 가장 작은 수 반환
		//1,2,3,4
	} else if arr[0] != arr[1] && arr[2] != arr[3] && arr[1] != arr[2] {
		answer = arr[0]
		//2개씩 같을 경우
		//3,3,6,6 
	} else if arr[0] == arr[1] && arr[2] == arr[3] && arr[1] != arr[2] {
		answer = (arr[0] + arr[2]) * (int(math.Abs(float64(arr[0]) - float64(arr[2]))))

		//2개는 같고 나머지 2개는 각각 p와 다를 경우
		//2,2,5,6
	} else if arr[0] == arr[1] && arr[2] != arr[3] && arr[1] != arr[2] {
		answer = arr[2] * arr[3]

		//2개는 같고 나머지 2개는 각각 p와 다를 경우
		/// 1,3,3,5
	} else if arr[1] == arr[2] && arr[0] != arr[3] && arr[2] != arr[3] {
		answer = arr[0] * arr[3]

		//2개는 같고 나머지 2개는 다를 경우
		//0,1,2,2
	} else if arr[2] == arr[3] && arr[0] != arr[1] && arr[1] != arr[2] {
		answer = arr[0] * arr[1]

		//3개는 같고 나머지 하나는 다를 경우
		//1,1,1,4
	} else if arr[0] == arr[1] && arr[1] == arr[2] && arr[0] != arr[3] {

		answer = (10*arr[0] + arr[3]) * (10*arr[0] + arr[3])
		//3개는 같고 나머지 하나는 다를 경우
		//1,2,2,2
	} else if arr[1] == arr[2] && arr[2] == arr[3] && arr[1] != arr[0] {

		answer = (10*arr[1] + arr[0]) * (10*arr[1] + arr[0])
	}
	fmt.Println("answer", answer)
	return answer

}

// func main() {
// 	a1 := 2
// 	b1 := 2
// 	c1 := 2
// 	d1 := 2

// 	// a2 := 4
// 	// b2 := 1
// 	// c2 := 4
// 	// d2 := 4

// 	// a3 := 6
// 	// b3 := 3
// 	// c3 := 3
// 	// d3 := 6

// 	// a4 := 2
// 	// b4 := 5
// 	// c4 := 2
// 	// d4 := 6

// 	Solution(a1, b1, c1, d1)

// }
