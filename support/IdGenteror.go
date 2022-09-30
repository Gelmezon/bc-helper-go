package support

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

var node *snowflake.Node

func init() {
	err := InitSnowflake()
	if err != nil {
		log.Println(err.Error())
	}
}

// InitSnowflake initiate Snowflake node singleton.
func InitSnowflake() error {
	// Get node number from env TIX_NODE_NO
	// Create snowflake node
	n, err := snowflake.NewNode(12)
	if err != nil {
		return err
	}
	// Set node
	node = n
	return nil
}

// GenerateSnowflake generate Twitter Snowflake ID
func GenerateSnowflake() uint {
	iInt := node.Generate().Int64()
	return uint(iInt)

}
