package main

import (
  "fmt"
  ssbnt "./packages/ssubnetting"
  "os"
)

func main() {
  //Obtiene las entradas desde la línea de comandos.
  ip, mask, hostsReq, sort, fok := ssbnt.CaptureData()
  if !fok {
    fmt.Fprintln(os.Stderr, "No es posible hacer el subneteo con esa configuración.")
    return
  }
  //Obtiene la máscara mínima para cada requerimiento de host.
  masks := ssbnt.GetMaskByHostReq(hostsReq)
  if sort != "" {
    //Ordena las máscaras de mayor a menor.
    ssbnt.SortMasks(masks, sort)
  }
  //Revisa los requerimientos no desbordan la red base "ip/mask".
  if !ssbnt.ValidSubnetting(mask, masks) {
    fmt.Fprintln(os.Stderr, "No se puede hacer el subneteo. Los requerimientos desbordan la red base.")
    return
  }
  //Hace el cálculo de las subredes.
  subnets := ssbnt.Subnetting(ip, masks)
  //Despliega los resultados con formato.
  ssbnt.PrintSubnetting(subnets)
}
