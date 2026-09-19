package main

import "fmt"

//func printUser(name string, age int) {
//	fmt.Println("Name:", name, "Age:", age)
//}

//func calculateAge(birthYear, currentYear int) int {
//	return currentYear - birthYear
//}

//func minMax(a, b int) (int, int) {
//	if a < b {
//		return a, b
//	}
//	return b, a
//}

//func average(numbers ...float64) float64 {
//	sum := 0.0
//	for _, number := range numbers {
//		sum += number
//	}
//	return sum / float64(len(numbers))
//}

//	var sum = func(a, b int) int {
//		return a + b
//	}(10, 30)
//
//	func calculate(a, b int) (sum, multiply int) {
//		sum = a + b
//		multiply = a * b
//		return
//	}

//func increase(score *int) {
//	*score += 10
//}

//	type Product struct {
//		Name  string
//		Price int
//	}
//
//	func applyDiscount(product *Product) {
//		product.Price -= 100
//	}
//
// Method
//type Product struct {
//	Name  string
//	Price int
//}
//
//func (p Product) printInfo() {
//	fmt.Println(p.Name, p.Price)
//}

//type User struct {
//	Name string
//	Age  int
//}
//
//func (u *User) birthday() {
//	u.Age++
//}
//
//func (u *User) rename(name string) {
//	u.Name = name
//}

// Printer Interfaces
//type Printer interface {
//	Print() string
//}
//
//type Book struct {
//	Title string
//}

//	func (b Book) Print() string {
//		return b.Title
//	}
//
//	func show(p Printer) {
//		fmt.Println(p.Print())
//	}
//type Player interface {
//	Play() string
//	Stop()
//}
//
//type musicPlayer struct{}
//
//func (m musicPlayer) Play() string {
//	return "Play"
//}
//func (m musicPlayer) Stop() {
//	fmt.Println("Stop")
//}

// Generic
//func show[T any](a T) {
//	fmt.Println(a)
//}

//func printPair[A any, B any](a A, b B) {
//	fmt.Println(a, b)
//}

//type TextOrNumber interface {
//	string | int
//}

type pair[A any, B any] struct {
	First  A
	Second B
}

func main() {
	//if
	//age := 20
	//hasTicket := false
	//if age >= 18 && hasTicket {
	//	fmt.Println("You can enter")
	//} else {
	//	fmt.Println("You cannot enter")
	//}
	//if age := 17; age >= 18 {
	//	fmt.Println("Adult")
	//} else {
	//	fmt.Println("Minor")
	//}

	//switch
	//month := "March"
	//
	//switch month {
	//case "December", "January", "February":
	//	fmt.Println("Winter")
	//case "March", "April", "May":
	//	fmt.Println("Spring")
	//case "June", "July", "August":
	//	fmt.Println("Summer")
	//case "September", "October", "November":
	//	fmt.Println("Autumn")
	//default:
	//}

	//for
	//for i := 0; i < 5000; i++ {
	//	fmt.Println(i)
	//}
	//for initialization; condition; post {
	// code
	//}
	//for i := 1; i <= 5; i++ {
	//	fmt.Println(i)
	//}
	//
	//for i := 10; i >= 0; i -= 2 {
	//	fmt.Println(i)
	//}
	//
	//for i := 20; i > 0; i -= 5 {
	//	fmt.Println(i)
	//}
	//
	//for i := 1; i <= 10; i++ {
	//	if i == 5 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	//languages := []string{"Go", "Python", "Rust"}
	//
	//for index, value := range languages {
	//	fmt.Println(index, ":", value)
	//}
	//languages := []string{"Go", "Python", "Rust"}
	//
	//for _, language := range languages {
	//	fmt.Println(language)
	//}

	//languages := map[string]int{
	//	"Go":     10,
	//	"Python": 8,
	//	"Rust":   9,
	//}
	//
	//for lang, score := range languages {
	//	fmt.Println(lang, score)
	//}
	//
	//for lang := range languages {
	//	fmt.Println(lang)
	//}
	//
	//for _, score := range languages {
	//	fmt.Println(score)
	//}
	//text := "A🔥Go"
	//
	//for index, r := range text {
	//	fmt.Printf("%d -> %c\n", index, r)
	//}

	//functions
	//printUser("Ali", 24)

	//age := calculateAge(2002, 2026)
	//fmt.Println(age)

	//minimum, maximum := minMax(10, 22)
	//fmt.Println("Min:", minimum)
	//fmt.Println("Max:", maximum)

	//result := average(10, 20, 30)
	//fmt.Println(result)

	//scores := []float64{15, 18, 20, 17}
	//
	//result := average(scores...)
	//fmt.Println(result)

	//score := 10
	//up5 := func() {
	//	score += 5
	//}
	//up5()
	//fmt.Println(score)

	//fmt.Println(calculate(3, 4))
	//Pointer
	//age := 20
	//p := &age
	//*p = 24
	//fmt.Println(age)

	//score := 50
	//increase(&score)
	//increase(&score)
	//increase(&score)
	//fmt.Println(score)

	//p := Product{
	//	Name:  "Keyboard",
	//	Price: 1000,
	//}
	//
	//applyDiscount(&p)
	//fmt.Println(p)

	//p := Product{
	//	Name:  "Keyboard",
	//	Price: 1000,
	//}
	//p.printInfo()

	//u := User{
	//	Name: "Ali",
	//	Age:  24,
	//}
	//
	//u.birthday()
	//u.rename("Amir")
	//fmt.Println(u)

	//Interfaces
	//book := Book{
	//	Title: "The little Prince",
	//}

	//show(book)

	//show(5)
	//show("hello")
	//show(3.14)
	//show(true)

	//printPair(1, 2)
	//printPair("Ali", 24)
	//printPair(3.1415, 16)

	p := pair[string, bool]{First: "Go", Second: true}

	fmt.Println(p)
}
