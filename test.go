package main

import (
    "fmt"
    "math/big"
    "encoding/hex"
)

var Diff1 = StringToBig("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")

func StringToBig(h string) *big.Int {
    n := new(big.Int)
    n.SetString(h, 0)
    return n
}

func main() {
    fmt.Println(Diff1)
    padded := make([]byte, 32)

    //diff,_ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",16)
    //diff.Rsh(diff, uint(1))
    

    
    //diffBuff := new(big.Int).Div(Diff1, diff).Bytes()
    //fmt.Println(diffBuff)
    //copy(padded[32-len(diffBuff):], diffBuff)
    //fmt.Println(padded)
    
    //buff := padded[0:4]
    //fmt.Println(buff)
    //targetHex := hex.EncodeToString(reverse(buff))
    //fmt.Println(targetHex)
    
    targetHex := "c5a70000"
    targetHex = "ffff0f00"
    fmt.Println(targetHex)
    decoded, _ := hex.DecodeString(targetHex)
    fmt.Println(decoded)
    decoded = reverse(decoded)
    fmt.Println(decoded)
    copy(padded[:len(decoded)], decoded)
    fmt.Println(padded)
    newDiff := new(big.Int).SetBytes(padded)
    fmt.Println(newDiff)
    newDiff = new(big.Int).Div(Diff1, newDiff)

    fmt.Println(newDiff)
}

func reverse(src []byte) []byte {
    dst := make([]byte, len(src))
    for i := len(src); i > 0; i-- {
        dst[len(src)-i] = src[i-1]
    }
    return dst
}