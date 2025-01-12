import "fmt"

func main() {
    // Using the `var` keyword with explicit type
    var name string = "Payce"
    var age int = 16

    // Type inference
    var score = 100
    var isStudent = false

    // Short variable declaration
    greeting := "Hello, Go!"
    piValue := 3.14159

    // Declaring multiple variables
    var x, y, z int = 10, 20, 30
    a, b, c := "GoLang", 2.718, true

    // Constants
    const Pi = 3.14
    const Company = "The Boring Company"

    // Default zero values
    var uninitializedInt int
    var uninitializedString string
    var uninitializedBool bool

    // ----- Other Data Types -----
    // Floating point types
    var bigFloat float64 = 123456.789
    var smallFloat float32 = 3.14

    // Complex numbers
    var complexNum complex64 = 1 + 2i
    var complexNum64 complex128 = 3 + 4i

    // Arrays
    var arr = [3]int{1, 2, 3}

    // Slices (dynamic arrays)
    var slice = []int{10, 20, 30}

    // Maps (key-value pairs)
    var personInfo = map[string]string{
        "name":    "Payce",
        "country": "USA",
    }

    // Struct (custom type)
    type Person struct {
        Name    string
        Age     int
        Country string
    }
    var person = Person{
        Name:    "Payce",
        Age:     16,
        Country: "USA",
    }

    // Pointers
    var ptr *int
    ptr = &age

    // Print all values
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Score:", score)
    fmt.Println("Is Student:", isStudent)
    fmt.Println("Greeting:", greeting)
    fmt.Println("Value of Pi (piValue):", piValue)

    fmt.Println("Multiple Declarations:")
    fmt.Println("x:", x, "y:", y, "z:", z)
    fmt.Println("a:", a, "b:", b, "c:", c)

    fmt.Println("Constants:")
    fmt.Println("Pi:", Pi)
    fmt.Println("Company:", Company)

    fmt.Println("Default Zero Values:")
    fmt.Println("Uninitialized Int:", uninitializedInt)
    fmt.Println("Uninitialized String:", uninitializedString)
    fmt.Println("Uninitialized Bool:", uninitializedBool)

    // ----- Other Data Types -----
    fmt.Println("\nOther Data Types:")
    fmt.Println("Big Float (float64):", bigFloat)
    fmt.Println("Small Float (float32):", smallFloat)
    fmt.Println("Complex Number (complex64):", complexNum)
    fmt.Println("Complex Number (complex128):", complexNum64)

    fmt.Println("\nArray:", arr)
    fmt.Println("Slice:", slice)

    fmt.Println("Map (personInfo):", personInfo)
    fmt.Println("Struct (person):", person)

    fmt.Println("Pointer to Age:", *ptr)
}