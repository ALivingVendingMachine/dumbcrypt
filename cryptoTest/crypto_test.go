package main

import (
  "testing"
  "cryptoFuncs"
)

func TestFastMod(t *testing.T)  {
  t.Parallel()
  total := cryptoFuncs.FastModExp(3, 218, 1000)

  t.Log("First pass")
  t.Log("3 ** 218 (mod 1000) is ", total)
  if (total != 489) {
    t.Errorf("Total should be 489")
  }

  t.Log("Second pass")
  total = cryptoFuncs.FastModExp(39, 22, 100)

  t.Log("39 ** 22 (mod 100) is ", total)
  if (total != 21) {
    t.Errorf("Total should be 21")
  }

  t.Log("Third pass")
  total = cryptoFuncs.FastModExp(2435, 2355, 100)

  t.Log("2435 ** 2355 (mod 100) is ", total)
  if (total != 75) {
    t.Errorf("Total should be 75")
  }
}

func TestXGCD(t *testing.T)  {
  t.Parallel()
  t.Log("First pass")
  a, b, c := cryptoFuncs.XGCD(17, 237)

  t.Log("a is ", a)
  t.Log("b is ", b)
  t.Log("c is ", c)
  if (a != 1) {
    t.Errorf("a should be 1")
  }

  if (b != 14) {
    t.Errorf("b should be 14")
  }

  if (c != -1) {
    t.Errorf("c should be -1")
  }

  t.Log("Second pass")
  a, b, c = cryptoFuncs.XGCD(45, 135)

  t.Log("a is ", a)
  t.Log("b is ", b)
  t.Log("c is ", c)
  if (a != 45) {
    t.Errorf("a should be 45")
  }

  if (b != 1) {
    t.Errorf("b should be 1")
  }

  if (c != 0) {
    t.Errorf("c should be 0")
  }

  t.Log("Third pass")
  a, b, c = cryptoFuncs.XGCD(1, 1)

  t.Log("a is ", a)
  t.Log("b is ", b)
  t.Log("c is ", c)
  if (a != 1) {
    t.Errorf("a should be 45")
  }

  if (b != 0) {
    t.Errorf("b should be 1")
  }

  if (c != 1) {
    t.Errorf("c should be 0")
  }
}

func TestModInverse(t *testing.T)  {
  t.Parallel()
  var a int = cryptoFuncs.ModInverse(17, 34)

  t.Log("a is ", a)
  if (a != -1) {
    t.Errorf("a should be -1")
  }

  a = cryptoFuncs.ModInverse(134, 227)

  t.Log("a is ", a)
  if (a != 144) {
    t.Errorf("a should be 144")
  }

  a = cryptoFuncs.ModInverse(449, 557)

  t.Log("a is ", a)
  if (a != 459) {
    t.Errorf("a should be 459")
  }
}

func TestRabinMiller(t *testing.T)  {
  t.Parallel()

  var a bool = cryptoFuncs.RabinMillerPrimality(6, 2)

  t.Log("Is 6 prime with 2 as a witness")
  if (a) {
    t.Errorf("a should be false")
  }

  a = cryptoFuncs.RabinMillerPrimality(7, 2)

  t.Log("Is 7 prime with 2 as a witness")
  if (!a) {
    t.Errorf("a should be true")
  }

  a = cryptoFuncs.RabinMillerPrimality(3215031751, 2)

  t.Log("Is 3215031751 prime with 2 as a witness")
  if (a) {
    t.Errorf("a should be false")
  }
}
