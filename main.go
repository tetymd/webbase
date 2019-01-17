package main

import (
    "log"
    "github.com/tetymd/webbase/account"
)

func main() {
    u := account.User{
        UserId:   "dev",
        UserName: "devuser",
        Password: "devpass",
    }

    log.Println(u.UserName)
}
