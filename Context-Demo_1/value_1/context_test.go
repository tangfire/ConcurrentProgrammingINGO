package value_1

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateCtx(t *testing.T) {
	ctx := context.Background()
	//ctx1 := context.TODO()
	ctx = context.WithValue(ctx, "userId", 1024)

	func1(ctx)
}

func func1(ctx context.Context) {
	userId := ctx.Value("userId")

	ctx = context.WithValue(ctx, "username", "tangfire")

	fmt.Println("userId:", userId)

	func2(ctx)
}

func func2(ctx context.Context) {
	username := ctx.Value("username")
	userId := ctx.Value("userId")

	fmt.Println("username:", username)
	fmt.Println("userId:", userId)

}
