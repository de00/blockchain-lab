func (m *MerkleTree)Verify(path []MerkleNode, index []int64, data []byte) bool{
	prenodehash := data[:]
	for i := 0;i<len(path);i++{
		if index[i] == 0{
			var buf bytes.Buffer
			buf.Write(path[i].Data)
			buf.Write(prenodehash)
			data := buf.Bytes()
			prenodehash := sha256.Sum256(data)
		}else{
			var buf bytes.Buffer
			buf.Write(prenodehash)
			buf.Write(path[i].Data)
			data := buf.Bytes()
			prenodehash := sha256.Sum256(data)
		}
	}
	if m.RootNode.Data = prenodehash{
		return true
	}
	return false

}