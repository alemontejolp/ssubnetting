package main

import (
  "fmt"
  ssbnt "./packages/ssubnetting"
  "os"
)

func main() {
  ip, mask, hostsReq, fok := ssbnt.CaptureData()
  if !fok {
    fmt.Fprintln(os.Stderr, "No es posible hacer el subneteo con esa configuración.")
    return
  }
  masks := ssbnt.GetMaskByHostReq(hostsReq)
  ssbnt.SortMasks(masks)
  if !ssbnt.ValidSubnetting(mask, masks) {
    fmt.Fprintln(os.Stderr, "No se puede hacer el subneteo. Los requerimientos desbordan la red base.")
    return
  }

  subnets := ssbnt.Subnetting(ip, masks)
  ssbnt.PrintSubnetting(subnets)
}
