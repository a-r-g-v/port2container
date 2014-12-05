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
        fmt.Println(err)
        return
    }

    fmt.Println(docker.Path)
    output, err := docker.ExecCommand("ps","-a");
    output, err = docker.JsonDecode();
    if err == nil {
        fmt.Printf("%s",output)
    }
        fmt.Printf("%s",err)
    if env.DEBUG {
    output,err := iptables.Exec("-L")
    fmt.Printf("Exec : %t",iptables.Exists(""))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("YoYo %s",output)
    }
}
