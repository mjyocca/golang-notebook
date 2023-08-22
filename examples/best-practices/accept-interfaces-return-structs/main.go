package main

import (
	"context"
	"fmt"

	"github.com/mjyocca/golang-notebook/best-practices/accept-interfaces-return-structs/service"

	"github.com/mjyocca/golang-notebook/best-practices/accept-interfaces-return-structs/db"
)

func main() {
	ctx := context.Background()
	// store injected into user service
	store := db.NewDB()
	// user service struct, can now use it's exposed methods
	useService := service.NewUserService(store)

	user := &service.User{}
	if err := useService.CreateUser(ctx, user); err != nil {
		fmt.Println(fmt.Errorf("error: %s", err))
	}

	fmt.Println(fmt.Printf("User created: %+v", user))
}
