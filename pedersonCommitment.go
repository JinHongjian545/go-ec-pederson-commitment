package go_ec_pederson_commitment

import (
	"encoding/hex"
	"github.com/bwesterb/go-ristretto"
	"math/big"
)

/**
基于EC的 pederson commitment
原论文：Non-Interactive and Information-Theoretic Secure Verifiable Secret Sharing
在原论文里是基于循环群的实现，说的是g,h是群里的元素，没有人知道log_g(h)
对应到椭圆曲线的话，就是对G、H的要求是G和H是椭圆曲线上的点，没有人知道 H=aG 中a的值，那么G和H都采用随机选取的方式，应该是满足要求的？！ 或者更简单的方式就是G使用基点，H随机选，这里采用都随机生成的方式
秘密x，随机数r， G和H是椭圆曲线上的随机的点, --> c = xG + rH
G，H 由被承诺方选定，r由承诺方选定
*/

//生成两个曲线上的点，用于计算承诺
func ParamsGen() (G, H ristretto.Point) {
	G.Rand()
	H.Rand()
	return G, H
}

func ParamsGenToString() (GString, HString string) {
	var G, H ristretto.Point
	G.Rand()
	H.Rand()
	GBytes, _ := G.MarshalText()
	HBytes, _ := H.MarshalText()
	GString = hex.EncodeToString(GBytes)
	HString = hex.EncodeToString(HBytes)
	return
}

//生成一个随机阶数
func RandomGen() (r ristretto.Scalar) {
	r.Rand()
	return r
}

func RandomGenToNumberString() string {
	var r ristretto.Scalar
	r.Rand()
	return r.BigInt().String()
}

//计算 秘密 的 Pederson 承诺， 返回曲线上的点
func Commit(G, H ristretto.Point, secret []byte, r ristretto.Scalar) (commit ristretto.Point) {
	var x ristretto.Scalar
	x.Derive(secret)
	//c = xG + rH
	var comm ristretto.Point
	comm.Add(G.ScalarMult(&G, &x), H.ScalarMult(&H, &r))
	return comm
}

//计算 秘密 的 Pederson 承诺， 返回字符串
func CommitToString(GString, HString, rString string, secret []byte) (commitString string, err error) {
	var G, H ristretto.Point
	GBytes, _ := hex.DecodeString(GString)
	HBytes, _ := hex.DecodeString(HString)
	err = G.UnmarshalText(GBytes)
	if err != nil {
		return "", err
	}
	err = H.UnmarshalText(HBytes)
	if err != nil {
		return "", err
	}
	var r ristretto.Scalar
	var bigInt big.Int
	bigInt.SetString(rString, 10)
	r.SetBigInt(&bigInt)

	comm := Commit(G, H, secret, r)
	bytes, err := comm.MarshalText()
	if err != nil {
		return "", err
	} else {
		commitString = hex.EncodeToString(bytes)
		return commitString, err
	}
}

//通过原始参数验证承诺
func Open(comm, G, H ristretto.Point, secret []byte, r ristretto.Scalar) bool {
	var x ristretto.Scalar
	x.Derive(secret)
	var calculateComm ristretto.Point
	calculateComm.Add(G.ScalarMult(&G, &x), H.ScalarMult(&H, &r))
	return calculateComm.Equals(&comm)
}

//通过Marshal和hex编码的字符串参数 以及big.Int转字符串得到的 r值 验证承诺
func OpenByString(commString, GString, HString, rString string, secret []byte) bool {
	var G, H, verifyComm ristretto.Point
	verifyCommBytes, _ := hex.DecodeString(commString)
	e := verifyComm.UnmarshalText(verifyCommBytes)
	if e != nil {
		return false
	}
	GBytes, _ := hex.DecodeString(GString)
	e = G.UnmarshalText(GBytes)
	if e != nil {
		return false
	}
	HBytes, _ := hex.DecodeString(HString)
	e = H.UnmarshalText(HBytes)
	if e != nil {
		return false
	}
	var r ristretto.Scalar
	var bigInt big.Int
	bigInt.SetString(rString, 10)
	r.SetBigInt(&bigInt)

	comm := Commit(G, H, secret, r)
	return verifyComm.Equals(&comm)
}
