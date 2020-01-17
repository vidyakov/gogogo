package rsa

import (
    "errors"
    "math/big"
    "math/rand"
)

type Key struct {
    key int
    n int
}

type KeyPair struct {
    Private Key
    Public Key
}

func isPrime(n int) bool {
    i := 1
    total := 0

    for (n >= i) {
        division := n % i
        if division == 0 {
            total++
        }
        i++
    }

    if total == 2 {
        return true
    } else {
        return false
    }
}


func gcd(a int, b int) int {
   i := 1
   gc := 1

   for (a >= i) || (b >= i) {
       if b % i == 0 && a % i == 0 {
           if gc < i {
               gc = i
           }
       }
       i++
   }
   return gc
}


func multiplicativeInverse(e int, phi int) int {
    var tbl [][]int

    A := phi
    B := e
    row := []int{A, B, A % B, A / B, -1, -1}
    tbl = append(tbl, row)
    for i := 0; tbl[i][2] != 0; i++ {
        A = tbl[i][1]
        B = tbl[i][2]
        row := []int{A, B, A % B, A / B, -1, -1}
        tbl = append(tbl, row)
    }

    tbl[len(tbl)-1][4] = 0
    tbl[len(tbl)-1][5] = 1
    for i := len(tbl)-2; i >= 0; i-- {
        tbl[i][4] = tbl[i+1][5]
        tbl[i][5] = tbl[i+1][4] - tbl[i+1][5]*tbl[i][3]
    }

    r := tbl[0][5] % phi
    if r < 0 {
        r = r + phi
    }
    return r
}


func GenerateKeypair(p int, q int) (KeyPair, error) {
  if !isPrime(p) || !isPrime(q) {
      return KeyPair{}, errors.New("Both numbers must be prime.")
  } else if  p == q {
      return KeyPair{}, errors.New("p and q can't be equal.")
  }
  // Проверить
  n := p * q
  // Проверить
  phi := (p-1)*(q-1)

  e := rand.Intn(phi - 1) + 1
  g := gcd(e, phi)
  for g != 1 {
      e = rand.Intn(phi - 1) + 1
      g = gcd(e, phi)
  }

  d := multiplicativeInverse(e, phi)
  return KeyPair{Key{e, n}, Key{d, n}}, nil
}


func Encrypt(pk Key, plaintext string) []int {
    cipher := []int{}
    n := new(big.Int)
    for _, ch := range plaintext {
        n = new(big.Int).Exp(
            big.NewInt(int64(ch)), big.NewInt(int64(pk.key)), nil)
        n = new(big.Int).Mod(n, big.NewInt(int64(pk.n)))
        cipher = append(cipher, int(n.Int64()))
    }
    return cipher
}


func Decrypt(pk Key, cipher []int) string {
    plaintext := ""
    n := new(big.Int)
    for _, ch := range cipher {
        n = new(big.Int).Exp(
            big.NewInt(int64(ch)), big.NewInt(int64(pk.key)), nil)
        n = new(big.Int).Mod(n, big.NewInt(int64(pk.n)))
        plaintext += string(rune(int(n.Int64())))
    }
    return plaintext
}
