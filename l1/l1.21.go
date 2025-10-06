package l1

import "fmt"

// Реализовать паттерн проектирования «Адаптер» на любом примере.
// Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого, который ожидает клиент.
// Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс (или структура) и другой, несовместимый по интерфейсу потребитель — напишите адаптер, который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.
// Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.

// интерфейс, который ожидает клиент
type EuropeanSocket interface {
	ChargeRoundPlug()
}

// класс, не совместимый с клиентским интерфейсом
type AmericanSocket struct{}

func (as *AmericanSocket) ChargeFlatPlug() {
	fmt.Println("american socket")
}

// адаптер
type Middleware struct {
	americanSocket *AmericanSocket
}

func (m *Middleware) ChargeRoundPlug() {
	fmt.Println("adapt EuropeanSocket to AmericanSocket")
	m.americanSocket.ChargeFlatPlug()
}
