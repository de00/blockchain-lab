package main

import (
	"encoding/binary"
	"fmt"
	"bytes"
	"log"
)

func main(){
	var str string = "abcd"
	var data []byte = []byte(str)
	fmt.Printf("%x", Sha256Compute(data))
	return
}
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
func Sha256Compute(message []byte) [32]byte {
    //初始哈希值为前8个质数(2到19)的平方根的小数部分的前32位
	h0 := uint32(0x6a09e667)
	h1 := uint32(0xbb67ae85)
	h2 := uint32(0x3c6ef372)
	h3 := uint32(0xa54ff53a)
	h4 := uint32(0x510e527f)
	h5 := uint32(0x9b05688c)
	h6 := uint32(0x1f83d9ab)
	h7 := uint32(0x5be0cd19)

    //计算过程当中用到的常数,即前64个质数(2到311)的立方根小数部分的前32位:
	k := [64]uint32{
        0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
		0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
		0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
		0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
		0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
		0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
		0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
		0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2
	}
//前期处理
	tobecompute := append(message, 0x80)
	if len(tobecompute) % 64 < 56 {
		suffix := make([]byte, 56 - (len(tobecompute) % 64))
		tobecompute = append(tobecompute, suffix...)
	} else {
		suffix := make([]byte, 64 + 56 - (len(tobecompute) % 64))
		tobecompute = append(tobecompute, suffix...)
	}
	msgLen := len(message) * 8
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(msgLen))
	tobecompute = append(tobecompute, bs...)

	slices := [][]byte{};
    
	for i := 0; i < len(tobecompute) / 64; i++ {
		slices = append(slices, tobecompute[i * 64: i * 64 + 63])
	}
    
    //主循环
	for _, chunk := range slices {
		w := []uint32{}
        
		for i := 0; i < 16; i++ {
			w = append(w, binary.BigEndian.Uint32(chunk[i * 4:i * 4 + 4]))
		}
		w = append(w, make([]uint32, 48)...)
        
        //W消息区块处理
		for i := 16; i < 64; i++ {
			s0 := Loopright(w[i - 15], 7) ^ Loopright(w[i - 15], 18) ^ (w[i - 15] >> 3)
			s1 := Loopright(w[i - 2], 17) ^ Loopright(w[i - 2], 19) ^ (w[i - 2] >> 10)
			w[i] = w[i - 16] + s0 + w[i - 7] + s1
		}

		a := h0
		b := h1
		c := h2
		d := h3
		e := h4
		f := h5
		g := h6
		h := h7
        
        //在主循环中用压缩函数处理
		for i := 0; i < 64; i++ {
			S1 := Loopright(e, 6) ^ Loopright(e, 11) ^ Loopright(e, 25)
			ch := (e & f) ^ ((^e) & g)
			temp1 := h + S1 + ch + k[i] + w[i]
			S0 := Loopright(a, 2) ^ Loopright(a, 13) ^ Loopright(a, 22)
			maj := (a & b) ^ (a & c) ^ (b & c)
			temp2 := S0 + maj

			h = g
			g = f
			f = e
			e = d + temp1
			d = c
			c = b
			b = a
			a = temp1 + temp2
		}
        //将压缩后的尾端加到现有的hash值
		h0 = h0 + a
		h1 = h1 + b
		h2 = h2 + c
		h3 = h3 + d
		h4 = h4 + e
		h5 = h5 + f
		h6 = h6 + g
		h7 = h7 + h
	}
	hashedbytes := [][]byte{IntToByte(h0), IntToByte(h1), IntToByte(h2), IntToByte(h3), IntToByte(h4), IntToByte(h5), IntToByte(h6), IntToByte(h7)}
	hash := []byte{}
	hasharr := [32]byte{}
	for i := 0; i < 8; i ++ {
		hash = append(hash, hashedbytes[i]...)
	}
	copy(hasharr[:], hash[0:32])
	return hasharr
}

func IntToByte(i uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return bs
}

//循环右移函数
func Loopright(n uint32, d uint) uint32 {
	return (n >> d) | (n << (32 - d))
}