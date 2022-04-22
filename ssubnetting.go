package main

import (
  "fmt"
  "os"
  ssbnt "github.com/alemontejolp/ssubnetting/packages/ssubnetting"
)

func main() {
  //Obtiene las entradas desde la línea de comandos.
  ip, mask, hostsReq, sort, flo, subtr, add, fok := ssbnt.CaptureData()

  if !fok {
    fmt.Fprintln(os.Stderr, "No es posible hacer el subneteo con esa configuración.")
    return
  }

  //Revisar si la accion es restar hosts a una dirección.
  if subtr != 0 {
    fmt.Printf("Cantidad de hosts restados: %d\n", subtr)
    //Realizar resta.
    ssbnt.SubAddr(&ip, subtr)
    fmt.Print("Dirección: ")
    ssbnt.PrintDDN(ip)
    fmt.Println()
    return
  }

  //Revisar si la accion es sumar hosts a una dirección.
  if add != 0 {
    fmt.Printf("Cantidad de hosts sumados: %d\n", add)
    //Realizar resta.
    ssbnt.AddAddr(&ip, add)
    fmt.Print("Dirección: ")
    ssbnt.PrintDDN(ip)
    fmt.Println()
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
  subnets, leftoverAddr, leftoverHosts := ssbnt.Subnetting(ip, mask, masks)
  //Despliega los resultados con formato.
  ssbnt.PrintSubnetting(subnets, flo, leftoverAddr, leftoverHosts)
}
