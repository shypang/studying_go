// 문제 설명
// 소수점 아래 숫자가 계속되지 않고 유한개인 소수를 유한소수라고 합니다. 분수를 소수로 고칠 때 유한소수로 나타낼 수 있는 분수인지 판별하려고 합니다. 유한소수가 되기 위한 분수의 조건은 다음과 같습니다.

// 기약분수로 나타내었을 때, 분모의 소인수가 2와 5만 존재해야 합니다.
// 두 정수 a와 b가 매개변수로 주어질 때, a/b가 유한소수이면 1을, 무한소수라면 2를 return하도록 solution 함수를 완성해주세요.

// 제한사항
// a, b는 정수
// 0 < a ≤ 1,000
// 0 < b ≤ 1,000
// 입출력 예
// a	b	result
// 7	20	1
// 11	22	1
// 12	21	2
// 입출력 예 설명
// 입출력 예 #1

// 분수 7/20은 기약분수 입니다. 분모 20의 소인수가 2, 5 이기 때문에 유한소수입니다. 따라서 1을 return합니다.
// 입출력 예 #2

// 분수 11/22는 기약분수로 나타내면 1/2 입니다. 분모 2는 소인수가 2 뿐이기 때문에 유한소수 입니다. 따라서 1을 return합니다.
// 입출력 예 #3

// 분수 12/21는 기약분수로 나타내면 4/7 입니다. 분모 7은 소인수가 7 이므로 무한소수입니다. 따라서 2를 return합니다.
// Hint
// 분자와 분모의 최대공약수로 약분하면 기약분수를 만들 수 있습니다.
// 정수도 유한소수로 분류합니다.

// 유한소수, 무한소수, 기약분수
// 유리수 : 분수꼴로 나타낼 수 있는 수고 정수와 정수가 아닌 수로 나뉜다. 정수에는 부호에 따라 양의 유리수, 0, 음의 유리가 있다. 정수가 아닌 수로는 분수가 있다.
// 정수(양의정수, 0, 음의 정수) / 정수 아닌 유리수 (+1/2. -2/3, 0.4, -1.5)
// 유한소수 : 1/2, 3/4 처럼 소수점 아래 숫자가 어느정도 이어지다가 멈추는 소수지만
// 무한소수 : 2/3은 0.66666....처럼 소수점 아래에서 숫자가 멈추지 않고 계속된다. 이렇게 분수를 소수로 바꿨을 때 숫자가 계속되지 않고 멈추는 소수를 유한소수, 멈추지 않고 계속 이어지는 소수를 무한 소수라고 한다.
// 기약분수는 분자와 분모의 공약수가 1뿐이어서 더 이상 약분되지 않는 분수다.

package main

import "fmt"

// func solution(a int, b int) int {
// 	check := false

// 	if b%a == 0 {
// 		b /= a
// 		a /= a
// 	}

// 	if b%a == 0 {
// 		check = true
// 	}
// 	if b%2 == 0 || b%5 == 0 {
// 		check = true
// 	} else {
// 		check = false
// 	}

// 	if check == true {
// 		return 1
// 	} else {
// 		return 2
// 	}
// }

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solution(a, b int) int {
	g := gcd(a, b)
	b /= g
	num := []int{}
	i := 2
	for i <= b {
		if b%i == 0 {
			b /= i
			num = append(num, i)
		} else {
			i++
		}
	}

	for _, v := range num {
		if v != 2 && v != 5 {
			fmt.Println(2)
			return 2
		}
	}
	fmt.Println(1)
	return 1
}

func main() {
	a := 7
	b := 20
	// a := 11
	// b := 22
	// a := 12
	// b := 21
	solution(a, b)
}
