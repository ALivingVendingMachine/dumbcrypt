# dumbcrypt
Implements RSA in golang for kicks and to play around with TDD in go.

# Warnings
Don't use this.  I'm going to set this up to use very small primes (small enough
  to fit into a `int`), which are going to be way, way, way too small to do anything
  close to actual cryptography that can't just be broken with a modernish machine
  brute forcing all the possible exponents. \

Also, guess what?  I'm not using a cryptographically secure PRNG.  This baby will
  grab the same N every single time if its ran with the same seed.  I might change
  this in the future, but I'm not making you a cryptotool, I'm making me an excuse
  to work with Go.

# Usage
TODO
