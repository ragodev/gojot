package jot

import (
	"math/big"
	"strings"
)

// THE FOLLOWING TAKEN FROM
// https://github.com/jbenet/go-base58/blob/d31c55ce19dab47a8c3d1020775441b1e78d595a/base58.go

// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
// Modified by Juan Benet (juan@benet.ai)

// LICENSE: https://raw.githubusercontent.com/jbenet/go-base58/d31c55ce19dab47a8c3d1020775441b1e78d595a/LICENSE
// Copyright (c) 2013 Conformal Systems LLC.
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

// alphabet is the modified base58 alphabet used by Bitcoin.
const BTCAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const FlickrAlphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Decode decodes a modified base58 string to a byte slice, using BTCAlphabet
func DecodeString(b string) []byte {
	return DecodeAlphabet(b, BTCAlphabet)
}

// Encode encodes a byte slice to a modified base58 string, using BTCAlphabet
func EncodeToString(b []byte) string {
	return EncodeAlphabet(b, BTCAlphabet)
}

// EncodeBase58 encodes a byte slice to a modified base58 string, using BTCAlphabet
func EncodeBase58(s string) string {
	return EncodeAlphabet([]byte(s), BTCAlphabet)
}

// DecodeBase58 decodes a modified base58 string to a byte slice, using BTCAlphabet
func DecodeBase58(b string) string {
	return string(DecodeAlphabet(b, BTCAlphabet))
}

// DecodeAlphabet decodes a modified base58 string to a byte slice, using alphabet.
func DecodeAlphabet(b, alphabet string) []byte {
	answer := big.NewInt(0)
	j := big.NewInt(1)

	for i := len(b) - 1; i >= 0; i-- {
		tmp := strings.IndexAny(alphabet, string(b[i]))
		if tmp == -1 {
			return []byte("")
		}
		idx := big.NewInt(int64(tmp))
		tmp1 := big.NewInt(0)
		tmp1.Mul(j, idx)

		answer.Add(answer, tmp1)
		j.Mul(j, bigRadix)
	}

	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		if b[numZeros] != alphabet[0] {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen, flen)
	copy(val[numZeros:], tmpval)

	return val
}

// Encode encodes a byte slice to a modified base58 string, using alphabet
func EncodeAlphabet(b []byte, alphabet string) string {
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0, len(b)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabet[0])
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}
