package context

import (
	"context"
	"fmt"
)

func contextBackground() {
	ctx := context.Background();

	fmt.Println("Context: ", ctx);
	fmt.Println("Context ERROR: ", ctx.Err());
	fmt.Printf("Context Type: %T\n", ctx);
	fmt.Println("----------");
}

