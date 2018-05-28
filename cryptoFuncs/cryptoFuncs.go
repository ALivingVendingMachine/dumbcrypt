package cryptoFuncs

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

func RabinMillerPrimality(n int, a int) bool {
  gcd, _, _ := XGCD(n, a)
  if (n % 2 == 0 || gcd != 1) {
    return false
  }

  var k, q, holdN int = 0, 0, 0
  holdN = n - 1

  for (holdN % 2 == 0) {
    k += 1
    q = holdN / 2
    holdN = q
  }

  a = FastModExp(a, q, n)
  if (a == 1 || a == n - 1) {
    return true
  }

  for (a != n - 1) {
    a = FastModExp(a, 2, n)
    if (a == n-1) {
      return true
    }
    if (a == 1) {
      return false
    }
  }

  return false
}
