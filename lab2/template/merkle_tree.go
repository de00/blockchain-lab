<<<<<<< HEAD
package main

import (
	"crypto/sha256"
	"bytes"
	"gopkg.in/eapache/queue.v1"

	// "encoding/hex"
	// "fmt"
)

// MerkleTree represent a Merkle tree
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode represent a Merkle tree node
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode{
	
	var node MerkleNode
	if(left == nil && right == nil){
		// fmt.Printf("NewMerkleNode left = nil, right = nil\n")
		hash := sha256.Sum256(data)
		node = MerkleNode{nil, nil, hash[:]}
	} else {
		// fmt.Printf("NewMerkleNode left = %s, right = %s\n", hex.EncodeToString(left.Data), hex.EncodeToString(right.Data))
		var buf bytes.Buffer
		buf.Write(left.Data)
		buf.Write(right.Data)
		data := buf.Bytes()
		hash := sha256.Sum256(data)
		node = MerkleNode{left, right, hash[:]}
	}
	
	return &node
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
// implement
func NewMerkleTree(data [][]byte) *MerkleTree {
	// var node = MerkleNode{nil,nil,data[0]}
	// var mTree = MerkleTree{&node}
	que := queue.New()
	var lev_size1, lev_size2 int
	var tmp *MerkleNode
	for i:=0; i<len(data); i++{
		node := NewMerkleNode(nil, nil, data[i])
		que.Add(node)
		tmp = node
	}
	// fmt.Printf("out of first loop\n")
	lev_size1 = 0
	if(len(data) % 2 != 0){
		que.Add(tmp)
		lev_size1 += 1
	}
	lev_size1 += len(data)
	// var lev int // debug
	for que.Length() > 0{
		lev_size2 = 0
		// lev ++
		// fmt.Printf("lev = %d, lev size = %d\n", lev, lev_size1)
		for i:=0; i<lev_size1/2 && que.Length()>1; i++{
			
			left_ptr := que.Peek().(*MerkleNode)
			que.Remove()
			right_ptr := que.Peek().(*MerkleNode)
			que.Remove()
			node := NewMerkleNode(left_ptr, right_ptr, nil)
			que.Add(node)
			lev_size2++
			tmp = node
		}
		if(lev_size2 == 1){
			break;
		}
		if(lev_size2%2 != 0){
			que.Add(tmp)
			lev_size2++
		}
		lev_size1 = lev_size2
	
	}
	mTree := MerkleTree{tmp}
	return &mTree
}

=======
package main

import (
	// "crypto/sha256"
)

// MerkleTree represent a Merkle tree
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode represent a Merkle tree node
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
// implement
func NewMerkleTree(data [][]byte) *MerkleTree {
	var node = MerkleNode{nil,nil,data[0]}
	var mTree = MerkleTree{&node}

	return &mTree
}
>>>>>>> a8f7e8aacf32b00972f74cc3119f9dd72652936f
