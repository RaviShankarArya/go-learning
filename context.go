package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context :", ctx)
	fmt.Println("Context Error :", ctx.Err())
	fmt.Printf("Context type : %T\n", ctx)

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context :", ctx)
	fmt.Println("Context Error :", ctx.Err())
	fmt.Printf("Context type : %T\n", ctx)
	fmt.Println("Cancel :", &cancel)
	fmt.Printf("Cancel Type : %T\n", cancel)

	cancel()

	fmt.Println("Context :", ctx)
	fmt.Println("Context Error :", ctx.Err())
	fmt.Printf("Context type : %T\n", ctx)
	fmt.Println("Cancel :", &cancel)
	fmt.Printf("Cancel Type : %T\n", cancel)

}
