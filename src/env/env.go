package env

import(
   "code.google.com/p/gcfg"
   "fmt"
)

const(
    DEBUG = true
    FILE_NAME = "config.gcfg"
)

var cfg Config

type Config struct {
    General struct {
        Httpproxy bool
        Sqlite3file string
    }
    Nginx struct {
        Httpport int
        Httpsport int
        Confdir    string
    }

}
func Load() (bool,error){
   err := gcfg.ReadFileInto(&cfg, FILE_NAME) 
   if err != nil {
        return false,fmt.Errorf("Error pkg/env : %s", err)
   }
   if DEBUG == true {
    fmt.Printf("Load Config : %s\n",FILE_NAME)
    }
   return true,nil
}

