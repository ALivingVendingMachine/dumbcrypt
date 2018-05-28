// This package implements RSA and some utility functions.
package dumbcrypt

import (
  "math/big"
  "math/rand"
  "time"
  "fmt"
)

type publickey struct {
  e int
  N int
}

type secretkey struct {
  d int
  N int
}

type KeyPair struct {
  Pub *publickey
  Sec *secretkey
}

// fastModExp is a fast exponent mod some value.  It computes base^exp (mod mod)
// in a fast manner.
func fastModExp(base int, exp int, mod int) int {
  if (base < 0 || exp < 0 || mod < 0) {
    panic("fastModExp got negative values!")
  }
  var bigBase *big.Int = new(big.Int).SetInt64(int64(base))
  var bigExp *big.Int = new(big.Int).SetInt64(int64(exp))
  var bigMod *big.Int = new(big.Int).SetInt64(int64(mod))
  var ret *big.Int = new(big.Int)

  ret.Exp(bigBase, bigExp, bigMod)

  if (!ret.IsInt64() || int(ret.Int64()) < 0) {
    panic("fastModExp failed!")
  }

  return int(ret.Int64())

  // So this is bad form, but I spent time writing this so you get to spend time
  // reading it.  This is the logic for doing fast mod exponentiation in Go.
  // I'm not going to use this, because even for pretty small numbers, it will
  // overflow. (The biggest value that it could possibly run is as long as two
  // ints multiplied, which is HUGE, so I had to switch to math/big)

  // var ret int

  // for exp != 0 {
  //   if (exp % 2 == 1) {
  //     ret *= bigBase
  //     ret = ret % mod
  //   }
  //   exp = exp / 2
  //   bigBase = bigBase * bigBase % mod
  // }

  // return int(ret)
}

// xgcd is the extended euclidean algorithm, which, given 2 values, returns the
// gcd, then two values such that a(oldS) + b(oldT) = oldR.  In this case, oldR
// is the gcd.
func xgcd(a int, b int) (int, int, int) {
  var s int = 0
  var t int = 1
  var r int = b
  var oldS int = 1
  var oldT int = 0
  var oldR int = a
  var tmp int

  for r != 0 {
    var quot int =  oldR / r
    tmp = r
    r = oldR - quot * tmp
    oldR = tmp

    tmp = s
    s = oldS - quot * tmp
    oldS = tmp

    tmp = t
    t = oldT - quot * tmp
    oldT = tmp
  }

  return oldR, oldS, oldT
}

// ModInvers: given a value, a, and a modulo, mod, returns the multiplicative
// inverse of a.  What that means is that the result times a should equal 1 modulo
// the mod value.
func modInverse(a int, mod int) int {
  gcd, s, _ := xgcd(a, mod)

  if (gcd != 1) {
    return -1
  }

  if (s < 0) {
    s = s + mod
  }

  return s
}

// The Rabin-Miller Primality test.  Given some value, n, it checks if n is
// prime with a as a witness. (Note that this is a probablistic test, and so
// when we use it we have to call it a great number of times)
func rabinMillerPrimality(n int, a int) bool {
  gcd, _, _ := xgcd(n, a)
  if (n % 2 == 0 || gcd != 1) {
    return false
  }

  var r, d, holdN int = 0, 0, 0
  holdN = n - 1

  for (holdN % 2 == 0) {
    r += 1
    d = holdN / 2
    holdN = d
  }

  var x int = fastModExp(a, d, n)

  if (x == 1 || x == n - 1) {
    return true
  }

  for (r != -1) {
    x = (x * x) % n

    if (x == 1) {
      return false
    }
    if (x == -1) {
      return true
    }

    r -= 1
  }

  return false
}

// Runs the Rabin-Miller primality test times times, with random values each time.
// Assumes that the prng is already seeded.
func rabinMillerNTimes(n int, times int) bool {
  var ret bool = false

  for i := 0; i < times; i++ {
    ret = rabinMillerPrimality(n, rand.Intn(n))
  }

  return ret
}

func pAndQ() (int, int) {
  var p, q int = rand.Int(), rand.Int()
  var Nminus int = -1

  // If a value passes the Rabin-Miller primality test, it has a 1/4^n chance
  // of NOT being a true prime.  So if it passes 7 times, we are (1 - (1/4^7) ==)
  // 99.99389...% sure its really prime

  for Nminus < 0 {
    for !rabinMillerNTimes(p, 7) {
      p = p - 1
      if p <= 0 {
        p = rand.Int()
      }
    }

    for !rabinMillerNTimes(q, 7) {
      q = q - 1
      if q <= 0 {
        q = rand.Int()
      }
    }

    Nminus = (p-1) * (q-1)
  }

  return p, q
}

// GenerateRSAKeyPair does exactly what it says on the tin: it generates a RSA
// keypair and returns two type structs: PublicKeyPair and SecretKeyPair.  If
// you are sending a message to Bob, send him PublicKeyPair and the result of
// running encrypt(...) on your message.  Note that this could take A WHILE to run
// because of the limits on some values that are needed.
func GenerateRSAKeyPair() *KeyPair {
  // Seed the RNG like its your first programming class with the current time
  rand.Seed(time.Now().UTC().UnixNano())

  fmt.Println("FYI: this is randomly hanging, probably due to entropy sources")
  fmt.Println("If it looks like its hung, it probably is.  Restart the program if needed")

  var p int
  var q int
  var e int = 65535 // AFAIK this is pretty standard
  var d int

  var pubKey = new(publickey)
  var secKey = new(secretkey)

  pubKey.e = e

  // first we need to generate N.  So we need 2 primes, p and q, such that
  // gcd(e, (p-1)(q-1)) == 1 (which allows for an inverse, which we need to
  // decrypt)

  p, q = pAndQ()
  var gcd int

  gcd, _, _ = xgcd(e, ((p - 1) * (q - 1)))
  if gcd != 1 {
    p, q = pAndQ()
    gcd, _, _ = xgcd(e, ((p - 1) * (q - 1)))
  }

  pubKey.N = (p*q)

  d = modInverse(e, ((p-1) * (q-1)))

  secKey.d = d
  secKey.N = (p*q)

  var keys *KeyPair = new(KeyPair)
  keys.Pub = pubKey
  keys.Sec = secKey

  return keys
}

func (pk *publickey) Encrypt(m int) int {
  return fastModExp(m, pk.e, pk.N)
}

func (sk *secretkey) Decrypt(m int) int {
  return fastModExp(m, sk.d, sk.N)
}

// func Encrypt(m int, pubKey publickey) int {
//   return fastModExp(m, pubKey.e, pubKey.N)
// }
//
// func Decrypt(c int, secKey secretkey) int {
//   return fastModExp(c, secKey.d, secKey.N)
// }
