package main

import "fmt"

func main() {
	str := "hello 월드!" //한영이 섞인 문자열

	for i := 0; i < len(str); i++ {
		//문자열 크기를 얻어 순회
		//바이트 단위로 출력 13바이트 출력, 문자열 요소 하나씩 나오는게 아님
		//T%는 데이터 타입, %d는 문자코드 값 rune?값 %c는 문자값
		//영문과 공백은 정상 출력됐으나 한글은 그렇지 못함
		//str[i]처럼 인덱스로 접근하면 요소 타입은 uint8, 즉 byte 단위
		//그래서 1바이트인 영문자는 잘 표시되는데 3바이트 크기인 한글은 깨져 표시된 것
		//한 글자씩 순회하려면 []rune으로 변환해서 순회하는 방법과 range 키워드를 순회하는 방법 2가지가 있다.
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])

	}
}
