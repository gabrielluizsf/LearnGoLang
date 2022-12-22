package context

import (
	"context"
	"fmt"
)

func withCancelContext() {
	ctx := context.Background();

	fmt.Println("Context: ", ctx);
	fmt.Println("Context [ERROR]: ", ctx.Err());
	fmt.Printf("Context type: %T\n", ctx);
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx);

	fmt.Println("context: ", ctx);
	fmt.Println("context err: ", ctx.Err());
	fmt.Printf("context type: %T\n", ctx);
	fmt.Println("cancel:  ", cancel);
	fmt.Printf("cancel type: %T\n", cancel);
	fmt.Println("----------");

	cancel();

	fmt.Println("context: ", ctx);
	fmt.Println("context err: ", ctx.Err());
	fmt.Printf("context type: %T\n", ctx);
	fmt.Println("cancel:  ", cancel);
	fmt.Printf("cancel type: %T\n", cancel);
	fmt.Println("----------");
}
