package snowflake

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
)

func init() {
	// Create a new Node with a Node number of 1
	tmp, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	node = tmp
}

func GenerateId() snowflake.ID {
	return node.Generate()
}
