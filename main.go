package main

import (
    "log"
    "github.com/tetymd/api/account"
)

func main() {
    u := account.User{
        UserId: "test",
        UserName: "testuser",
        Password:  "test",
    }

    log.Println(u.UserName)
}
