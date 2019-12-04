package gopherconuk

type Animal struct{}

type Dog struct {
	Animal
}

func main() {
	var a Animal
	a = Dog{}
}
