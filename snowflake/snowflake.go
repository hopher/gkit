package snowflake

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

var _node, _ = snowflake.NewNode(createNodeID())

// GenerateID 雪花算法
// mysql 用 bigint(20)
func GenerateID() int64 {
	return _node.Generate().Int64()
}

// ResetNode 节点ID重置
func ResetNode() int64 {
	nodeID := createNodeID()
	_node, _ = snowflake.NewNode(nodeID)
	return nodeID
}

// createNodeID 随机生成 [0-1024) 区间数
func createNodeID() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(1024)
}
