package snowflake

import (
	"fmt"
)

func ExampleGenerateID() {
	fmt.Println(GenerateID())
}

func ExampleResetNode() {
	fmt.Println(ResetNode())
}
