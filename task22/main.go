package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Используем пакет math/big, так как int64 не удовлетворяет условиям
	bigInteger := NewBigInteger(3512030103010230012, 959593945994499393)
	fmt.Println(bigInteger.Add())
	fmt.Println(bigInteger.Sub())
	fmt.Println(bigInteger.Mul())
	fmt.Println(bigInteger.Div())
}

// Структура для работы с двумя числами типа big.Int
type BigInteger struct {
	a, b *big.Int
}

func (b *BigInteger) Add() *big.Int {
	return new(big.Int).Add(b.a, b.b)
}

func (b *BigInteger) Sub() *big.Int {
	return new(big.Int).Sub(b.a, b.b)
}

func (b *BigInteger) Mul() *big.Int {
	return new(big.Int).Mul(b.a, b.b)
}

func (b *BigInteger) Div() *big.Int {
	return new(big.Int).Quo(b.a, b.b)
}

func NewBigInteger(a, b int64) *BigInteger {
	floats := &BigInteger{new(big.Int), new(big.Int)}
	floats.a.SetInt64(a)
	floats.b.SetInt64(b)
	return floats
}