package main
import(
    "fmt"
    "iptables"
    "env"
)

func main() {
    if env.DEBUG {
    iptables.Opps()
    fmt.Printf("Hi opps")
    }
}
