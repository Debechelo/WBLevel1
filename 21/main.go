package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

func (e *Employee) GetFullName() string {
	return e.FirstName + " " + e.LastName
}

// Интерфейс, ожидаемый клиентским кодом
type Printer interface {
	PrintName() string
}

// Адаптер для адаптации интерфейса работника к интерфейсу клиентского кода
type EmployeeAdapter struct {
	Employee *Employee
}

func (ea *EmployeeAdapter) PrintName() string {
	return ea.Employee.GetFullName()
}

// Функция, принимающая интерфейс Printer
func PrintEmployeeName(printer Printer) {
	fmt.Println("Employee's full name:", printer.PrintName())
}

func main() {
	// Создаем объект работника
	employee := &Employee{FirstName: "John", LastName: "Doe"}

	// Создаем адаптер
	adapter := &EmployeeAdapter{Employee: employee}

	// Передаем адаптер в функцию, ожидающую интерфейс Printer
	PrintEmployeeName(adapter)
}
