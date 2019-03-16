package ssubnetting

import (
  "sort"
  "strconv"
  "strings"
  "os"
)

func FillArr(arr *[4]int, v, begin, end int) {
  for i := begin; i < end; i++ {
    arr[i] = v
  }
}

func SortMasks(masks []int) {
  sort.Sort(sort.Reverse(sort.IntSlice(masks)))
}

func CopyAddr(soucre [4]int, dest *[4]int)  {
  for i := 0; i < 4; i++ {
    dest[i] = soucre[i]
  }
}

// Toma un string de la forma x1.x2.x3.x4 y devuelve un
// slice de enteros con cada xi.
func ParseAddr(addr string) ([4]int, bool) {
  var intAddr [4]int
  var fok error
  sepAddr := strings.Split(addr, ".")
  l := len(sepAddr)
  for i := 0; i < l; i++ {
    intAddr[i], fok = strconv.Atoi(sepAddr[i])
    if fok != nil {
      return intAddr, false
    }
  }
  return intAddr, true
}

//String to sequence of integers.
func StrToSeqOfInt(req string, sep string) ([]int, bool){
  var err error
  strReq := strings.Split(req, sep)
  l := len(strReq)
  intReq := make([]int, l)
  for i := 0; i < l; i++ {
    intReq[i], err = strconv.Atoi(strReq[i])
    if(err != nil) {
      return intReq, false
    }
  }
  return intReq, true
}

//Obtene lo que haya después de un argumento de línea de comandos.
func GetFlagValue(f string) string {
  var r string
  flg := false
  first := true
  l := len(os.Args)
  for i := 0; i < l; i++ {
    if os.Args[i] == f {
      flg = true
    } else if flg && os.Args[i][0] != '-' {
      if first {
        r += os.Args[i]
        first = false
      } else {
        r += " " + os.Args[i]
      }
    } else if(flg && os.Args[i][0] == '-') {
      break
    }
  }
  return r
}
