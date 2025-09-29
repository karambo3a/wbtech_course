package l1

// Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?
// Приведите корректный пример реализации.

// var justString string

// func someFunc() {
//   v := createHugeString(1 &lt;&lt; 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// Вопрос: что происходит с переменной justString?

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	if len(runes) > 100 {
		runes = runes[:100]
	}
	justString = string(runes)
}
