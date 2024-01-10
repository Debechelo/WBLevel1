//package main
//
//import "fmt"
//
////Дана структура Human (с произвольным набором полей и методов).
////Реализовать встраивание методов в структуре Action от родительской
////структуры Human (аналог наследования).
//
//// Human Структура
//type Human struct {
//	Name string
//	Age  int
//}
//
//// SayName Методы для структуры Human
//func (h *Human) SayName() {
//	fmt.Println("Hello, my name is", h.Name)
//}
//
//// SayAge Методы для структуры Human
//func (h *Human) SayAge() {
//	fmt.Printf("I'm %d years old\n", h.Age)
//}
//
//// Action Структура с встраиванием методов от структуры Human
//type Action struct {
//	Human // Встраивание структуры Human
//}
//
//// Walk Дополнительный метод для структуры Action
//func (a *Action) Walk() {
//	fmt.Println(a.Name, "is walking")
//}
//
//func main() {
//	person := Action{
//		Human: Human{
//			Name: "Sam",
//			Age:  20,
//		},
//	}
//
//	person.SayName() // Метод SayName унаследован от Human
//	person.SayAge()  // Метод SayAge унаследован от Human
//	person.Walk()    // Метод Walk принадлежит структуре Action
//
//}

package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
