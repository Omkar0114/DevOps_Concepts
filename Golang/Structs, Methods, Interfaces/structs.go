package main
import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// structs are used to group together related data
	// structs are used to group together related functions
	// struct is a collection of fields
	// struct is a value type
	// struct is a user defined type
	// struct is a composite type
	// struct is a reference type
	// struct is a data type
	// struct is a data structure
	// struct is a logical data type
	// struct is a logical data structure
	// struct is a logical composite type
	// struct is a logical reference type
	// struct is a logical value type
	// struct is a logical user defined type

	s := person{firstName: "Omkar", lastName: "Patil", age: 23}
	fmt.Println(s)
	fmt.Println(s.firstName)
	fmt.Println(s.lastName)
	fmt.Println(s.age)
	fmt.Printf("%+v", s)

	mov := getMovie("xyz", 2.1)
	mov1 := getMovie("abc", 3.3)
	movies := make([]Movie, 5)
	movies = append(movies, mov)
	movies = append(movies, mov1)
	for _, value := range movies {
			fmt.Println(value)
	}
}

type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
			name:   s,
			rating: r,
	}
	return
}

