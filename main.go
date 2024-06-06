package main

import (
    "fmt"
    "github.com/casbin/casbin/v2"
)

func main() {
    // Initialize the enforcer
    e, _ := casbin.NewEnforcer("./model.conf", "./policy.csv")

    // Check if a user has a permission
    hasPerm, _ := e.Enforce("alice", "data1", "read")
    fmt.Println(hasPerm) // should print: true

    hasPerm, _ = e.Enforce("bob", "data2", "write")
    fmt.Println(hasPerm) // should print: true

	hasPerm, _ = e.Enforce("alice", "data2", "read")
    fmt.Println(hasPerm) // should print: false

    hasPerm, _ = e.Enforce("bob", "data1", "read")
    fmt.Println(hasPerm) // should print: false

	hasPerm, _ = e.Enforce("bob", "data2", "read")
    fmt.Println(hasPerm) // should print: true
}