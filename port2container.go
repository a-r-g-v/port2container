package main

import(
    "fmt"
    "iptables"
    "env"
    "docker"
    "os"
)


func main()  {

    if os.Geteuid() != 0 {
        fmt.Println("must be superuser")
        os.Exit(-1)
    }

    // Load Env Value
    _, err := env.Load()    
    if err != nil {
        fmt.Print(err)
        return
    }

    // Check Docker
    fmt.Println(docker.Path)
    
    if env.DEBUG {
    output,err := iptables.Exec("-L")
    fmt.Printf("Exec : %t",iptables.Exists(""))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("YoYo %s",output)
    }
}
