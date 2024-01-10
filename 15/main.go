package main

import (
	"bytes"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string
//
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }
//
// func main() {
// 	someFunc()
// }

// Проблема в утечке памяти так как мы создаем строку размера 2^10 байт
// после создаем срез длиной 100, но это срез продолжает ссылаться на массив размера 2^10
// и появляется утечка пямяти в размере 2^10 - 100

// Самым адекватную решением было бы сначало создать строку размера 100 байт

var justString string

func someFunc() {
	v := createHugeString(100)
	justString = v[:100]
}

func createHugeString(cnt int) string {
	return strings.Repeat("a", cnt)
}

// но если нам прям очень хочется создать большую строку
// можно изпользовать буффер прочитав только первые 100 байт

func someFunc2() {
	var buffer bytes.Buffer
	createHugeString2(&buffer, 1<<10)
	first100Bytes := make([]byte, 100)
	// Читаем первые 100 байт
	buffer.Read(first100Bytes)

	justString = string(first100Bytes)
}

func createHugeString2(buffer *bytes.Buffer, cnt int) {
	buffer.WriteString(strings.Repeat("a", cnt))
}

func main() {
	someFunc()
	someFunc2()
}
