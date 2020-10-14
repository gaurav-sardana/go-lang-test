package main

import (
	"net/http"

	"github.com/halodoc/webservice/controllers"
)

func main() {
	// var i int
	// for i < 5 {
	// 	fmt.Println(i)
	// }
	// fmt.Println("Hello My first program, Hahah, Bye")
	// var i int
	// // i = 10
	// fmt.Println(i)
	// var d float32
	// // d = 3.14
	// fmt.Println(d)

	// firstName := "Gaurav"
	// fmt.Println("First Name : " + firstName)

	// var isTrue bool
	// fmt.Println(isTrue)

	// var j uint
	// // j = 10
	// fmt.Println(j)

	// b := true
	// fmt.Println(b)

	// c := complex(3, 4)
	// fmt.Println(c)

	// r, img := real(c), imag(c)
	// fmt.Println(r, img)

	// var lastName *string = new(string)
	// *lastName = "Sardana"
	// fmt.Println(lastName, *lastName)
	// *lastName = "Jain"
	// ptr := &*lastName
	// fmt.Println(ptr, *ptr)

	// friendName := "Harsh Jain"
	// fPtr := &friendName
	// fmt.Println(fPtr, friendName, *fPtr)
	// friendName = "Divam"
	// fmt.Println(fPtr, friendName, *fPtr)

	// const value = 10
	// fmt.Println(value + 2)
	// fmt.Println(value + 1.2)

	// var value1 int = 10
	// fmt.Println(value1 + 5)
	// fmt.Println(float32(value1) + 1.2)

	// var x float32 = 34.7
	// fmt.Println(int(x))

	// const (
	// 	first = iota
	// 	second
	// 	third
	// )

	// fmt.Println(first, second, third)

	/* var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[2])

	arr1 := [3]int{3, 4, 5}
	fmt.Println(arr1)
	fmt.Println(arr1[2])

	slice := arr1[:]
	fmt.Println(slice)

	arr2 := []int{1, 2, 3}
	fmt.Println(arr2)
	arr3 := append(arr2, 4)
	fmt.Println(arr3)
	arr3[1] = 5
	fmt.Println(arr2, arr3)

	m := map[string]int{"foo": 23, "bar": 32}
	fmt.Println(m)
	fmt.Println(m["bar"])
	m["bar"] = 40
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m["foo"])
	fmt.Println(m)*/

	// var m map[string]float32
	// n := make(map[string]int)
	// n["cat"] = 2
	// fmt.Println(m)
	// fmt.Println(n)
	// fmt.Println(m[""])

	// type user struct {
	// 	ID        int
	// 	FirstName string
	// 	LastName  string
	// }
	// var u user
	// u.ID = 1
	// u.FirstName = "Gaurav"
	// u.LastName = "Sardana"
	// fmt.Println(u)
	// fmt.Println(u.FirstName, u.LastName)
	// u1 := user{ID: 1, FirstName: "Gaurav", LastName: "Sardana"}
	// fmt.Println(u1)
	// u := models.User{
	// 	ID:        2,
	// 	FirstName: "Divam",
	// 	LastName:  "Bharadwaj",
	// }

	// k := &u
	// *k = models.User{
	// 	ID:        2,
	// 	FirstName: "Divam",
	// 	LastName:  "Jain",
	// }
	// fmt.Println(*k)
	// port := 3000
	// _, err := startWebServer(port, 2)
	// fmt.Println(err)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	// var i int
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println("continuing!")
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println("continuing!")
	// }
	// fmt.Printf(i)

	// var i int
	// for {
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// 	i++

	// }

	// slice := []int{1, 2, 3, 4, 5}
	// // for i := 0; i < len(slice); i++ {
	// // 	fmt.Println(slice[i])
	// // }

	// for i := range slice {
	// 	// fmt.Println(i, v)
	// 	fmt.Println(slice[i])
	// }

	// users := map[int]string{1: "Gaurav", 2: "Harsh"}
	// // _ is write only variable
	// for _, v := range users {
	// 	fmt.Println(v)
	// 	panic("Something happened bad!")
	// }

	// r := HTTPRequest{Method: "GET"}
	// switch r.Method {
	// case "GET":
	// 	fmt.Println("Get Request")
	// 	fallthrough
	// 	//implicit breaking
	// 	//fallthrough?
	// case "DELETE":
	// 	fmt.Println("Delete Request")
	// 	fallthrough
	// case "POST":
	// 	fmt.Println("POST Request")
	// 	fallthrough
	// default:
	// 	fmt.Println("Default")
	// }

}

// func startWebServer(port, noOfRetries int) (int, error) {
// 	fmt.Println("Starting Web Server on port", port)
// 	fmt.Println("............")
// 	fmt.Println("Stopping Web Server")
// 	fmt.Println("Number of retries", noOfRetries)
// 	return port, errors.New("Bad happened")
// }

type HTTPRequest struct {
	Method string
}
