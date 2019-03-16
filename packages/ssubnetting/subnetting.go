package ssubnetting

import (
  "math"
)

func HostsByMask(m int) int {
  return int(math.Pow(2, float64(32 - m)))
}

func DDNMask(mask *[4]int, m int) {
  h := HostsByMask(m)
  i := 3
  for ; h > 256; i-- {
    h = h / 256;
  }
  mask[i] = 256 - h
  FillArr(mask, 255, 0, i)
  FillArr(mask, 0, i + 1, 4)
}

func CalcMinMask(h int) (m int) {
  for m = 30; HostsByMask(m) - 2 < h; m-- {}
  return
}

func AddAddr(addr *[4]int, h int)  {
  addr[3] += h
  for i := 3; addr[i] > 255; i-- {
    addr[i - 1] = addr[i - 1] + (addr[i] / 256)
    addr[i] = addr[i] % 256
  }
}

func SubAddr(addr *[4]int, h int)  {
  addr[3] -= h
  for i := 3; i >= 0; i-- {
    for addr[i] < 0 {
      r := addr[i] % 256
      if -addr[i] >= 256 {
        addr[i - 1] += addr[i] / 256
        if r != 0 {
          if r < 0 {
            addr[i] = r
          }
        } else {
          addr[i] = 0
        }
      } else {
        addr[i - 1]--
        if r != 0 {
          addr[i] = 256 + r
        } else {
          addr[i] = 0
        }
      }
    }
  }
}

func CalcSubnet(addr [4]int, mask int, last, fusb, lusb *[4]int)  {
  CopyAddr(addr, last)
  AddAddr(last, HostsByMask(mask) - 1)
  CopyAddr(addr, fusb)
  AddAddr(fusb, 1)
  CopyAddr(*last, lusb)
  SubAddr(lusb, 1)
}

//hr: hosts required, hn: hosts number.
func MinMaskByHosts(hr int) (mask int)  {
  mask = 30;
  for hn := 0; hn - 2 < hr; mask-- {
    hn = HostsByMask(mask)
  }
  mask++
  return
}

// bm: base mask, masks[]: las máscaras de cada subred.
func ValidSubnetting(bm int, masks []int) bool {
  l := len(masks)
  ac := 0
  for i := 0; i < l; i++ {
    ac += HostsByMask(masks[i])
  }
  return HostsByMask(bm) >= ac
}

func GetMaskByHostReq(hosts []int) (masks []int)  {
  l := len(hosts)
  masks = make([]int, l)
  for i := 0; i < l; i++ {
    masks[i] = MinMaskByHosts(hosts[i])
  }
  return
}

func Subnetting(ip [4]int, masks []int) []Subnet {
  var currAddr[4]int
  l := len(masks)
  sn := make([]Subnet, l)
  CopyAddr(ip, &currAddr)
  for i := 0; i < l; i++ {
    CopyAddr(currAddr, &sn[i].Id)
    CalcSubnet(sn[i].Id, masks[i], &sn[i].Broadcast, &sn[i].FirstU, &sn[i].LastU)
    CopyAddr(sn[i].Broadcast, &currAddr)
    AddAddr(&currAddr, 1)
    sn[i].DecMask = masks[i]
    DDNMask(&sn[i].DDNMask, masks[i])
  }
  return sn
}
