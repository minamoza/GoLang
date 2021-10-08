package main

import (
	"fmt"
	"math"
	"net/http"
)

var a, b, c = 1.0, 3.0, 1.0
var discriminant, root1, root2 float64

func welcome_page(page http.ResponseWriter, r *http.Request){
	fmt.Fprintf(page, "Welcome!!! That page is calculate Discriminant\n")
	// fmt.Scan(&a, &b, &c)
	r1, r2 := Discriminant(a, b, c)
	
	fmt.Fprintf(page, "For value of a = %v, b= %v, c = %v the first root is %v and second root is %v", a, b, c, r1, r2)
}

func hundleRequest(){
	http.HandleFunc("/", welcome_page)
	http.ListenAndServe(":8080", nil)
}

func Discriminant(a float64, b float64, c float64)(float64, float64) { //public function
    discriminant = math.Pow(b, 2) - 4 * a * c
    switch {
    case discriminant > 0:
        root1 = (-b + math.Sqrt(discriminant)/(2*a))
        root2 = (-b - math.Sqrt(discriminant)/(2*a))
    case discriminant == 0:
        root1 = -b / (2 * a)
        root2 = -b / (2 * a)
    case discriminant < 0:
		return 0, 0
        // fmt.Println("Warning, Discriminant is less than 0")
    }
	return root1, root2
}

func main(){
	hundleRequest()
}