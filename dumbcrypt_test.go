package dumbcrypt_test

import (
  "dumbcrypt"
  "testing"
  "strconv"
)

var fastModTests = []struct {
  base int
  exp int
  mod int
  out int
} {
  {3, 218, 1000, 489}, // 1
  {39, 22, 100, 21}, // 2
  {2435, 2355, 100, 75}, // 3
  {3215031751, 11, 2345, 841}, // 4
}

func TestFastMod(t *testing.T)  {
  t.Parallel()
  for i, tc := range fastModTests {
    t.Run(strconv.Itoa(i), func(t *testing.T) {
      out := dumbcrypt.FastModExp(tc.base, tc.exp, tc.mod)

      if out != tc.out {
        t.Errorf("Got %d, want %d", out, tc.out)
      }
    })
  }
}

var xgcdTests = []struct {
  a int
  b int
  gcd int
  s int
  t int
} {
  {17, 237, 1, 14, -1}, // 1
  {45, 135, 45, 1, 0}, // 2
  {1, 1, 1, 0, 1}, // 3
  {44, 34, 2, 7, -9}, // 4
}

func TestXGCD(t *testing.T)  {
  t.Parallel()

  for i, tc := range xgcdTests {
    t.Run(strconv.Itoa(i), func(t *testing.T) {
      outGcd, outS, outT := dumbcrypt.XGCD(tc.a, tc.b)

      if (outGcd != tc.gcd || outS != tc.s || outT != tc.t) {
        t.Errorf("Got %d %d %d, wanted %d %d %d", outGcd, outS, outT, tc.gcd, tc.s, tc.t)
      }
    })
  }

}

var modInverseTests = []struct {
  a int
  mod int
  out int
} {
  {17, 34, -1}, // 1
  {134, 227, 144}, // 2
  {449, 557, 459}, // 3
}

func TestModInverse(t *testing.T)  {
  t.Parallel()

  for i, tc := range modInverseTests {
    t.Run(strconv.Itoa(i), func(t *testing.T) {
      out := dumbcrypt.ModInverse(tc.a, tc.mod)

      if out != tc.out {
        t.Errorf("Got %d, wanted %d", out, tc.out)
      }
    })
  }
}

var rabinMillerTests = []struct {
  n int
  witness int
  out bool
} {
  {6, 2, false}, // 1
  {7, 2, true}, // 2
  {3215031751, 2, true}, // 3
  {3215031751, 3, true}, // 4
  {3215031751, 5, true}, // 5
  {3215031751, 11, false}, // 6
}

func TestRabinMiller(t *testing.T)  {
  t.Parallel()

  for i, tc := range rabinMillerTests {
    t.Run(strconv.Itoa(i), func(t *testing.T) {
      out := dumbcrypt.RabinMillerPrimality(tc.n, tc.witness)

      if out != tc.out {
        t.Errorf("Got %t, wanted %t", out, tc.out)
      }
    })
  }
}
