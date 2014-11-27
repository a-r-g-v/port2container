package main
import(
    "fmt"
    "iptables"
    "env"
)

func main()  {

    // Load Env Value
    r, err := env.Load()    
    if err != nil {
        fmt.Print(err)
        return
    }
    
    fmt.Print(r);
    
    if env.DEBUG {
    output,err := iptables.Exec("-L")
    fmt.Printf("Exec : %t",iptables.Exists(""))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("YoYo %s",output)
    }
}
