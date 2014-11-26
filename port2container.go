package main
import(
    "fmt"
    "iptables"
    "env"
)

func main() {
    if env.DEBUG {
    output,err := iptables.Exec("-L")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("YoYo %s",output)
    }
}
