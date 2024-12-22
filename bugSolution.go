```go
func main() {
    var m map[string]int
    val, ok := m["0"] //Safe way to access map elements. 
    if ok {
        fmt.Println(val)
    } else {
        fmt.Println("Key not found")
    }
}
```