/*
Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?


var justString string

func someFunc() {
	v := createHugeString(1 &lt;&lt; 10)
	justString = v[:100]
}

func main() {
	someFunc()
}


Ответ:
Под капотом, строки, это массивы байт, если создаем срез - создаем указатель на этот массив.
Значит строка будет жить в памяти

*/

// Исправленная реализация
package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	var builder strings.Builder
	builder.WriteString(v[:100])
	justString = builder.String()
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}
