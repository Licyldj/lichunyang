package LICY_BLC

import "crypto/sha256"

type Licy_MerkleTree struct {
	Licy_RootNode *Licy_MerkleNode
}

type Licy_MerkleNode struct {
	Licy_Left  *Licy_MerkleNode
	Licy_Right *Licy_MerkleNode
	Licy_Data  []byte
}

func Licy_NewMerkleTree(data [][]byte) *Licy_MerkleTree {

	//[tx1,tx2,tx3]

	var nodes []Licy_MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
		//[tx1,tx2,tx3,tx3]
	}

	// 创建叶子节点
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}


	//MerkleNode{nil,nil,tx1Bytes}
	//MerkleNode{nil,nil,tx2Bytes}
	//MerkleNode{nil,nil,tx3Bytes}
	//MerkleNode{nil,nil,tx3Bytes}



	// 　循环两次
	for i := 0; i < len(data)/2; i++ {

		var newLevel []Licy_MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		//MerkleNode{MerkleNode{nil,nil,tx1Bytes},MerkleNode{nil,nil,tx2Bytes},sha256(tx1Bytes,tx2Bytes)}
		//
		//MerkleNode{MerkleNode{nil,nil,tx3Bytes},MerkleNode{nil,nil,tx3Bytes},sha256(tx3Bytes,tx3Bytes)}
		//


		nodes = newLevel
	}

	//MerkleNode:
	//	left: MerkleNode{MerkleNode{nil,nil,tx1Bytes},MerkleNode{nil,nil,tx2Bytes},sha256(tx1Bytes,tx2Bytes)}
	//
	//	right: MerkleNode{MerkleNode{nil,nil,tx3Bytes},MerkleNode{nil,nil,tx3Bytes},sha256(tx3Bytes,tx3Bytes)}
	//
	//	sha256(sha256(tx1Bytes,tx2Bytes)+sha256(tx3Bytes,tx3Bytes))

	mTree := Licy_MerkleTree{&nodes[0]}

	return &mTree
}

func NewMerkleNode(left, right *Licy_MerkleNode, data []byte) *Licy_MerkleNode {
	mNode := Licy_MerkleNode{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		mNode.Licy_Data = hash[:]
	} else {
		prevHashes := append(left.Licy_Data, right.Licy_Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Licy_Data = hash[:]
	}

	mNode.Licy_Left = left
	mNode.Licy_Right = right

	return &mNode
}