package ssubnetting

import (
  "fmt"
  "os"
  "strconv"
)

// Lee y transforma la configuración del subneteo desde la línea de comandos.
// @return (ip, mask, host requirements, sort, flo, fok)
func CaptureData() ([4]int, int, []int, string, bool, bool) {
  var (
    ip [4]int
    hostsReq []int
  )

  fIp, _ := GetFlagValue("-ip")
  strmasks, _ := GetFlagValue("-mask")
  fMask, err := strconv.Atoi(strmasks)
  fReq, _ := GetFlagValue("-req")
  fSort, fse := GetFlagValue("-sort") //fse : Flag Sort Exists.
  _, flo := GetFlagValue("-lo") //flo : Flag leftover

  if fse && fSort == ""{
    fSort = "desc"
  }

  if err != nil {
    fMask = 32
    fmt.Fprintln(os.Stderr, "Falló al convertir la máscara a entero.")
    return ip, fMask, hostsReq, fSort, flo, false
  }

  _ip, fok := StrToSeqOfInt(fIp, ".")
  if(fok) {
    for i := 0; i < 4; i++ {
      ip[i] = _ip[i]
    }
  } else {
    fmt.Fprintln(os.Stderr, "Falló al parsear la IP.")
    return ip, fMask, hostsReq, fSort, flo, false
  }
  hostsReq, fok = StrToSeqOfInt(fReq, " ")
  if !fok {
    fmt.Fprintln(os.Stderr, "Falló al parsear los requerimeintos.")
    return ip, fMask, hostsReq, fSort, flo, false
  }
  return ip, fMask, hostsReq, fSort, flo, true
}

// Imprime una dirección en formato Dot Decimal Nonation.
func PrintDDN(a [4]int) {
  l := len(a)
  if(0 < l) {
    fmt.Print(a[0])
  }
  for i := 1; i < l; i++ {
    fmt.Printf(".%d", a[i])
  }
}

// Despliega una red con formato:
// [nombre red]: [red]
func DisplayNet(sn [4]int, message string)  {
  fmt.Printf("%s: ", message)
  PrintDDN(sn)
  fmt.Println()
}

//Imprie los detalles del subneteo.
func PrintSubnetting(sn []Subnet, flo bool, leftoverAddr [4]int, leftoverHosts int)  {
  fmt.Println("Subneteo:")
  fmt.Println("-----------------------------------------")
  l := len(sn)
  for i := 0; i < l; i++ {
    fmt.Printf("Subred [%d]:\n", i)
    DisplayNet(sn[i].Id, "ID de red")
    DisplayNet(sn[i].Broadcast, "Broadcast de red")
    DisplayNet(sn[i].FirstU, "Primera dirección usable")
    DisplayNet(sn[i].LastU, "Última dirección usable")
    fmt.Printf("Máscara de subred (Decimal): %d\n", sn[i].DecMask)
    DisplayNet(sn[i].DDNMask, "Máscara de subred (DDN)")
    fmt.Printf("Direcciones disponibles: %d\n", sn[i].HostsAvailable)
    fmt.Println("-----------------------------------------")
  }
  if flo {
    if leftoverHosts != 0 {
      DisplayNet(leftoverAddr, "Dirección de inicio del bloque sobrante")
    }
    fmt.Printf("Direcciones sobrantes: %d\n", leftoverHosts)
    fmt.Println("-----------------------------------------")
  }
}
