// This package implements RSA and some utility functions.
package cryptoFuncs

// FastModExp is a fast exponent mod some value.  It computes base^exp (mod mod)
// in a fast manner.
func FastModExp(base int, exp int, mod int) int {
  var ret int = 1

  for exp != 0 {
    if (exp % 2 == 1) {
      ret *= base
      ret = ret % mod
    }
    exp = exp / 2
    base = (base * base) % mod
  }
  return ret
}

// XGCD is the extended euclidean algorithm, which, given 2 values, returns the
// gcd, then two values such that a(oldS) + b(oldT) = oldR.  In this case, oldR
// is the gcd.
func XGCD(a int, b int) (int, int, int) {
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
func ModInverse(a int, mod int) int {
  gcd, s, _ := XGCD(a, mod)

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
func RabinMillerPrimality(n int, a int) bool {
  gcd, _, _ := XGCD(n, a)
  if (n % 2 == 0 || gcd != 1) {
    return false
  }

  var r, d, holdN int = 0, 0, 0
  holdN = n - 1

  var x int = FastModExp(a, n, n)

  if (x == 1 || x == n - 1) {
    return true
  }

  for (holdN % 2 == 0) {
    r += 1
    d = holdN / 2
    holdN = d
  }

  for (r != 0) {
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
