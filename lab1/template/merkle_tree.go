package main

import (
	 "crypto/sha256"
	 "bytes"
)

// MerkleTree represent a Merkle tree
type MerkleTree struct {
	RootNode *MerkleNode
	Leafs    []*MerkleNode
}

// MerkleNode represent a Merkle tree node
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
    Parent *MerkleNode
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
// implement
func NewMerkleNode(left,right *MerkleNode, data []byte) *MerkleNode {
    mNode := MerkleNode{}

    if left == nil && right == nil {
        hash := sha256.Sum256(data)
        mNode.Data = hash[:]
    }else {
        prevHashes := append(left.Data,right.Data...)
        hash := sha256.Sum256(prevHashes)
        mNode.Data = hash[:]
    }

    mNode.Left = left
    mNode.Right = right
	mnode.Parent = nil

    return &mNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

    if len(data) % 2 != 0 {
        data = append(data, data[len(data) - 1])
    }

    for _, dataitem := range data {
        node := NewMerkleNode(nil, nil, dataitem)
        nodes = append(nodes, *node)
    }
	leafnodes := nodes
    for i := 0; i<len(data)/2; i++ {
        var newNodes []MerkleNode

        for j := 0; j < len(nodes); j += 2 {
            node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			len := len(newNodes)
			nodes[j].Parent = node
			nodes[j+1].Parent = node
            newNodes = append(newNodes , *node)
        }

        nodes = newNodes
    }

    mTree := MerkleTree{&nodes[0],leafnodes}

	return &mTree
}

func (m *MerkleTree) GetMerklePath(data []byte) ([]MerkleNode, []int64) {
    //找到要验证的节点
	for _, current := range m.Leafs {
		if bytes.Equal(data, current.Data){
			currentParent := current.Parent
			var merklePath []MerkleNode
			var index []int64
			for currentParent != nil {
				if bytes.Equal(currentParent.Left.Data, current.Data) {
					merklePath = append(merklePath, currentParent.Right)
					index = append(index, 1) // add right leaf
				} else {
					merklePath = append(merklePath, currentParent.Left)
					index = append(index, 0) // add left leaf
				}
				current = currentParent
				currentParent = currentParent.Parent
			}
			return merklePath, index
		}else{
			return nil, nil
		}
	}
	return nil, nil
}