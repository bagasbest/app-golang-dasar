package main

import (
	_ "app-golang-dasar/database"
	_ "app-golang-dasar/helper" /// BLANK IDENTIFIER
	_ "container/list"
	_ "container/ring"
	"errors"
	"fmt"
	_ "math"
	"reflect"
	_ "reflect"
	_ "strconv"
	_ "strings"
	_ "time"
)

func main() {
	// NUMBER
	// fmt.Println("Satu = ", 1)
	// fmt.Println("Tiga Setengah = ", 3.5)

	// BOOLEAN
	// fmt.Println("Benar = ", true)
	// fmt.Println("Salah = ", false)

	// STRING
	// fmt.Println("Bagas Pangestu Ganteng")
	// fmt.Println("Bagas Pangestu")
	// fmt.Println("Bagas ")

	// panjang string
	// fmt.Println(len("Bagas Pangestu Ganteng"))
	// fmt.Println("Bagas Pangestu Ganteng"[0]) /// byte B = 66
	// fmt.Println("Bagas Pangestu Ganteng"[1])

	// VARIABLE (using var)
	// var name string
	// name = "Bagas Pangestu Ganteng"
	// fmt.Println(name)
	// name = "Bagas Pangestu"
	// fmt.Println(name)

	// var friendName = "Budi"
	// fmt.Println(friendName)

	// var age int8 = 30
	// fmt.Println(age)

	// VARIABLE (not using var, :=)
	// country := "Indonesia"
	// fmt.Println(country)

	// MULTIPLE VARIABLE DECLARATION
	// var (
	// 	firstName  = "Bagas"
	// 	middleName = "Ganteng"
	// 	lastName   = "Pangestu"
	// )

	// fmt.Println(firstName + " " + middleName + " " + lastName)

	/// CONSTANT
	// const (
	// 	firstName string = "Bagas"
	// 	lastName         = "Pangestu"
	// 	value            = 150000
	// )

	/// KONVERSI TIPE DATA
	// nilai32 := int32(12lai32)
	// nilai8 := int8(nilai32) // hati hati ketika konversi tipe data

	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai8)

	// name := "Bagas"
	// e := byte(name[0])
	// eString := string(e)
	// fmt.Println(eString) // ngambil karakter "B"

	// TYPE DECLARATION
	// type (
	// 	NoKTP   string
	// 	Married bool
	// 	Age     uint8
	// )

	// ktpNumber := NoKTP("123543423324324")
	// isMarried := Married(false)
	// userAge := Age(24)

	// fmt.Println(ktpNumber)
	// fmt.Println(isMarried)
	// fmt.Println(userAge)

	// OPERASI MATEMATIKA
	// var result = 10 + 10
	// fmt.Println(result)

	// a := 10
	// b := 10
	// c := a * b
	// fmt.Println(c)

	// i := 10
	// i += 10
	// fmt.Println(i)

	// // UNARY OPERATOR
	// i++
	// fmt.Println(i)

	// OPERASI PERBANDINGAN
	// name1 := "Bagas"
	// name2 := "Pangestu"

	// result := name1 == name2
	// fmt.Println(result)

	// OPERASI BOOLEAN
	// ujian := 80
	// absensi := 80

	// lulusUjian := ujian > 80
	// lulusAbsensi := absensi > 80
	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsensi)

	// lulus := lulusUjian && lulusAbsensi
	// fmt.Println(lulus)

	// fmt.Println(ujian >= 80 && absensi >= 80)

	// ARRAY
	// var names [3]string

	// names[0] = "Bagas"
	// names[1] = "Pangestu"
	// names[2] = "Ganteng"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	// var values = [3]int{
	// 	90, 95, 80,
	// }

	// fmt.Println(values)

	// // MENDAPATKAN PANJANG ARRAY
	// fmt.Println(len(values))

	// // MENGUBAH VALUE DARI INDEX ARRAY
	// values[2] = 100
	// fmt.Println(values)

	// TIPE DATA SLICE (pointer, length, capacity)
	// var months = [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	// slice1 := months[4:7]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// var slice2 = months[10:]
	// fmt.Println(slice2)

	// var slice3 = append(slice2, "Bagas")
	// fmt.Println(slice3)

	// slice3[1] = "Bukan Desember"
	// fmt.Println(slice3)
	// fmt.Println(slice2)
	// fmt.Println(months)

	// MAKE
	// newSlice := make([]string, 2, 5)
	// newSlice[0] = "Bagas"
	// newSlice[1] = "Pangestu"

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// // COPY
	// copySlice := make([]string, len(newSlice), cap(newSlice))
	// copy(copySlice, newSlice)
	// fmt.Println(copySlice)

	// // ARRAY VS SLICE
	// iniArray := [...]int{0, 1, 2, 3, 4}
	// iniSlice := []int{0, 1, 2, 3, 4}

	// fmt.Println(iniArray)
	// fmt.Println(iniSlice)

	// MAP
	// person := map[string]string{
	// 	"name":    "Bagas Pangestu",
	// 	"address": "Kebon Kacang",
	// }

	// fmt.Println(person)
	// fmt.Println(len(person))
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])

	// MAKE
	// book := make(map[string]string)

	// book["title"] = "Belajar GOlang"
	// book["Author"] = "Bagas"
	// book["ups"] = "Salah"

	// fmt.Println(len(book))
	// delete(book, "ups")
	// fmt.Println(len(book))

	// fmt.Println(book)

	// IF EXPRESSION
	// name := "Joko"

	// if name == "Bagas" {
	// 	fmt.Println("Hello, Bagas!")
	// } else if name == "Joklo" {
	// 	fmt.Println("Hello, Joklo!")
	// } else {
	// 	fmt.Println("Hello," + name)
	// }

	// // SHORT STATEMENT
	//   if length := len(name); length > 5 {
	// 	fmt.Println("nama sudha benar")
	// }

	// SWITCH
	// name := "Bagas"

	// switch name {
	// case "Bagas":
	// 	fmt.Println("Hello Bagas")
	// case "Joko":
	// 	fmt.Println("Hello Joko")
	// default:
	// 	fmt.Println("Kenalan Donk")
	// }

	// // SHORT STATEMENT
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// SWITCH TANPA KONSIDI
	// length := len(name)
	// switch {
	// case length > 10:
	// 	fmt.Println("nama terlalu panjang")
	// case length > 5:
	// 	fmt.Println("nama lumayan panjang")
	// case length > 10:
	// 	fmt.Println("nama sudah benar")
	// }

	/// FOR LOOPS BIASA
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke : ", counter)
	// 	counter++
	// }

	/// FOR LOOPS DENGAN STATEMENT
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke: ", counter)
	// }

	/// FOR LOOPS UNTUK AKSES SLICE
	// slice := []string{"Bagas", "Pangestu", "Eko", "Kurniawan", "Khannedy"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	/// FOR UNTUK MENGAKSES DENGAN INDEX DARI SLICE (RANGE)
	// for index, value := range slice {
	// 	fmt.Println("Index: ", index, " = ", value)
	// }

	// for _, value := range slice {
	// 	fmt.Println("value: ", value)
	// }

	/// FOR LOOPS DENGAN MAP
	// person := make(map[string]string)
	// person["name"] = "Eko"
	// person["title"] = "Programmer"

	// for key, value := range person {
	// 	fmt.Println("Key : ", key, " Value : ", value)
	// }

	/// BREAK DAN CONTINUE
	// for i := 0; i < 10; i++ {
	// if i == 5 {
	// 	break
	// }

	// if i % 2 == 0 {
	// 	continue
	// }
	// fmt.Println("Perulangan ke : ", i)
	//}

	/// FUNCTION
	// sayHello()

	/// FUNCTION PARAMETER
	// sayHelloName("Bagas", "Pangestu")

	/// FUNCTION RETURN VALUE
	// name := "Bagas Pangestu"
	// fmt.Println(getHello(name))

	/// FUNCTION RETURN MULTIPLE VALUE
	// name, _ := getUserNameAndAge()
	// fmt.Println(name)
	// fmt.Println(age)

	/// NAMED RETURN VALUE
	// firstName, middleName, lastName := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)

	/// VARIADIC FUNCTION
	// total := sumAll() // gak wajid masukin parameter
	// total := sumAll(10, 5, 6, 4, 5)
	/// DENGAN SLICE
	// slice := []int{10, 20, 30, 40, 50, 60, 80}
	// total := sumAll(slice...)
	// fmt.Println(total)

	/// FUNCTION VALUE
	// sayGoodBye := getGoodBye

	// result := sayGoodBye("Bagas")
	// fmt.Println(result)

	/// FUNCTION AS PARAMETER
	//sayHelloWithFilter("Anjay", spamFilter)
	/// FUNCTION TYPE DECLARATION

	/// ANONYMOUS FUNCTION
	// blacklist := func(name string) bool {
	// 	return name == "admin"
	// }
	// registerUser("admin", blacklist)
	// registerUser("user", blacklist)

	/// RECURSIVE FUNCTION
	// fmt.Println(factorialLoop(5))
	// fmt.Println(factiorialRecursive(5))

	/// CLOSURE FUNCTION
	// counter := 0
	// name := "Eko"

	// increment := func() {
	// 	name := "Bagas"
	// 	fmt.Println("Increment")
	// 	counter++
	// 	fmt.Println(name)
	// }

	// increment()
	// increment()

	// fmt.Println(counter)
	// fmt.Println(name)

	/// DEFER, PANIC, RECOVER
	// runApplication(0) // defer
	// runApp(true) // panic, recover

	/// KOMENTAR
	/**
	Komentar
	Multiline
	tidak terbatas
	**/

	/// STRUCT

	// var bagas Customer
	// bagas.Name = "Bagas"
	// bagas.Address = "Jakarta"
	// bagas.Age = 24

	// fmt.Println(bagas.Name)
	// fmt.Println(bagas.Address)
	// fmt.Println(bagas.Age)

	/// STRUCT LITERAL
	// joko := Customer{
	// 	Name:    "joko",
	// 	Address: "Indonesia",
	// 	Age:     30,
	// }
	// fmt.Println(joko)

	// budi := Customer{"Budi", "Jakarta", 35}
	// fmt.Println(budi)

	/// STRUCT METHOD
	// bagas.sayHelloWithStruct("Ucok")

	/// INTERFACE
	// var eko Person
	// eko.Name = "Eko"

	// sayHelloWithInterface(eko)

	/// INTERFACE 2
	// cat := Animal{
	// 	Name: "cat",
	// }

	// sayHelloWithInterface(cat)

	/// INTERFACE KOSONG
	// emptyInterface := Ups(3)
	// fmt.Println(emptyInterface)

	/// NIL
	// data := NewMap("")
	// if data != nil {
	// 	fmt.Println(data)
	// } else {
	// 	fmt.Println("Data Kosong")
	// }

	/// ERROR INTERFACE
	// contohError := errors.New("Ups, error")
	// fmt.Println(contohError.Error())

	// hasil, err := Pembagi(100, 0)
	// if err == nil {
	// 	fmt.Println("Hasil", hasil)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	/// TYPE ASSERTION
	// result := random()
	// resultToString := result.(string)
	// fmt.Println(resultToString)

	// switch value := result.(type) {
	// case string:
	// 	fmt.Println("Value", value, "is string")
	// case int:
	// 	fmt.Println("Value", value, "is int")
	// default:
	// 	fmt.Println("Unknown type")
	// }

	/// POINTER (PASS BY REFERENCE)
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // passing by value
	//address2 := &address1 // passing by reference

	//address2.City = "Bandung"

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	//*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	// fmt.Println(address1)
	// fmt.Println(address2)

	// address4 := new(Address)
	// address4.City = "Jakarta"
	// fmt.Println(address4)

	/// POINTER DI FUNCTION
	// alamat := Address {
	// 	City: "Subang",
	// 	Provence: "Jawa Barat",
	// 	Country: "",
	// }

	// ChangeCountryToIndonesia(&alamat)
	// fmt.Println(alamat)

	/// POINTER DI METHOD
	// bagas := Man{"Bagas"}
	// bagas.Married()

	// fmt.Println(bagas.Name)

	/// GEPATH (Sekarang lebih di rekomendasikan gomodule)

	/// PACKAGE AND IMPORT
	// helper.SayHello("haloooo Bagas Ganteng")

	/// ACCESS MODIFIER (variable diawali huruf besar itu public & variable diawali huruf kecil itu private)

	/// PACKAGE INITIALIZATION
	// result := database.GetDatabase()
	// fmt.Println(result)

	/// BLANK IDENTIFIER
	/// DI PACKAGE BISA DI TAMBAHKAN _ SEBELUM NAMA PACKAGE

	/// PACKAGE OS
	// args := os.Args
	// fmt.Println(args)

	// name, err := os.Hostname() /// UNTUK MENGAMBIL NAMA KOMPUTER / HOST
	// if err == nil {
	// 	fmt.Println("Hostname: ", name)
	// } else {
	// 	fmt.Println("Error: ", err.Error())
	// }

	// username := os.Getenv("USERNAME")
	// password := os.Getenv("PASSWORD")

	// fmt.Println(username)
	// fmt.Println(password)

	/// PACKAGE FLAG
	// host := flag.String("host", "localhost", "Put your host")
	// user := flag.String("user", "root", "Put your database user")
	// password := flag.String("password", "root", "Put your database password")
	// number := flag.Int("number", 100, "Put your number")

	// // untuk melakukan proses parsing kita panggil method dibawah
	// flag.Parse()

	// fmt.Println("Host:", *host)
	// fmt.Println("User:", *user)
	// fmt.Println("Password:", *password)
	// fmt.Println("Number:", *number)

	/// PACKAGE STRINGS
	// fmt.Println(strings.Contains("Bagas Pangestu", "Bagas"))
	// fmt.Println(strings.Split("Bagas Pangestu", " "))
	// fmt.Println(strings.ToLower("Bagas Pangestu"))
	// fmt.Println(strings.ToUpper("Bagas Pangestu"))
	// fmt.Println(strings.ReplaceAll("Bagas Pangestu", "a", "i"))
	// fmt.Println(strings.Trim("              Bagas Pangestu        ", " "))

	/// PACKAGE MATH
	// fmt.Println(math.Round(0.7))
	// fmt.Println(math.Round(1.3))
	// fmt.Println(math.Floor(1.5))
	// fmt.Println(math.Ceil(1.5))
	// fmt.Println(math.Max(1.5, 1.4))
	// fmt.Println(math.Min(29, 28))
	// fmt.Println(math.Abs(-10 / 2))
	// fmt.Println(math.Pow(10, 2))

	/// PACKAGE CONTAINER LIST (implementasi double linked list)
	// data := list.New()
	// data.PushBack("Bagas")
	// data.PushBack("Ganteng")
	// data.PushBack("Pangestu")
	// data.PushFront("Awal")

	// for e := data.Front(); e != nil; e = e.Next() {
	// 	fmt.Println("Get Isi: ", e.Value)
	// }

	// data := ring.New(5)
	// for i := 0; i < data.Len(); i++ {
	// 	data.Value = "Data " + strconv.FormatInt(int64(i), 10)
	// 	data = data.Next()
	// }

	// data.Do(func(value interface{}) {
	// 	fmt.Println(value)
	// })

	/// PACKAGE SORT
	// users := []User{
	// 	{"Bagas", 24},
	// 	{"Joko", 30},
	// 	{"Rudi", 27},
	// 	{"Andi", 40},
	// 	{"Sugeng", 23},
	// }

	// SORT
	// sort.Sort(UserSlice(users))
	// fmt.Println(users)

	/// PACKAGE TIME
	// now := time.Now()
	// fmt.Println(now)
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Day())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// fmt.Println(now.Nanosecond())

	// utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	// fmt.Println(utc)

	// layout := "2006-01-02"
	// parse, _ := time.Parse(layout, "2020-10-01")
	// fmt.Println(parse)

	/// PACKAGE REFLECT
	// sample := Sample{"Bagas"}

	// sampleType := reflect.TypeOf(sample)
	// fmt.Println(sampleType.NumField())
	// fmt.Println(sampleType.Field(0).Name)

	// fmt.Println(sampleType.Field(0).Tag.Get("required"))
	// fmt.Println(sampleType.Field(0).Tag.Get("max"))
	// fmt.Println(sampleType.Field(0).Tag.Get("min"))

	// fmt.Println(IsValid(sample))

	// contoh := ContohLagi{"", ""}
	// contoh := ContohLagi{"Bagas", "Oke"}
	// fmt.Println(IsValid(contoh))

	/// PACKAGE REGEX
	// regex := regexp.MustCompile(`e([a-z])o`)
	// fmt.Println(regex.MatchString("eko"))
	// fmt.Println(regex.MatchString("eto"))
	// fmt.Println(regex.MatchString("eDo"))

	// search := regex.FindAllString("eko eka eyo eki", -1)
	// fmt.Println(search)

	/// MATERI SELANJUTNYA
	// GOLANG UNIT TEST
	// GOLANG MODULE
	// GOLANG CUNCURRENCY (MULTICORE PROGRAMMING)
	// GOLANG WEB

}

func sayHello() {
	fmt.Println("Hello World!")
}

func sayHelloName(firstName string, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}

func getHello(name string) string {
	if name != "" {
		return "Hello, " + name
	} else {
		return "Hello, User"
	}
}

func getUserNameAndAge() (string, uint8) {
	return "Bagas", uint8(24)
}

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Bagas"
	middleName = "Ganteng"
	lastName = "Pangestu"

	return
}

// / VAR ARG HANYA BISA DI LETAKKAN DI PALING KANAN
func sumAll(number ...int) int {
	total := 0

	for _, value := range number {
		total += value
	}
	return total
}

func getGoodBye(name string) string {
	return "Good Bye, " + name
}

type Filter func(string) string ///// FUNCTION TYPE DECLARATION
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello, " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjay" {
		return "..."
	} else {
		return name
	}
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked,", name)
	} else {
		fmt.Println("Welcome,", name)
	}
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factiorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factiorialRecursive(value-1)
	}
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() /// biasakan gunakan defer diatas
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println(result)
}

// / recover di simpan pada defer function
func endApp() {
	message := recover()
	if message == nil {
		fmt.Println("Error dengan message,", message)
	}
	fmt.Println("Aplikasi selesai")
}

// / defer diletakkan di bagian atas
func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

// / awalnya uppercase untuk struct
// / nama field nya awalnya kapital
type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHelloWithStruct(name string) {
	fmt.Println("Hello,", name, "My Name is", customer.Name)
}

type HasName interface {
	GetName() string
}

func sayHelloWithInterface(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func random() interface{} {
	return "Ups"
}

type Address struct {
	City, Provence, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}
