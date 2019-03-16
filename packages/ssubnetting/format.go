package ssubnetting

import (
  "fmt"
  "os"
  "strconv"
)

// Lee y transforma la configuración del subneteo desde la línea de comandos.
func CaptureData() ([4]int, int, []int, bool) {
  var (
    ip [4]int
    hostsReq []int
  )

  fIp := GetFlagValue("-ip")
  fMask, err := strconv.Atoi(GetFlagValue("-mask"))
  fReq := GetFlagValue("-req")
  //fmt.Fprintln(os.Stderr, fReq)

  if err != nil {
    fMask = 32
    fmt.Fprintln(os.Stderr, "Falló al convertir la máscara a entero.")
    return ip, fMask, hostsReq, false
  }

  _ip, fok := StrToSeqOfInt(fIp, ".")
  if(fok) {
    for i := 0; i < 4; i++ {
      ip[i] = _ip[i]
    }
  } else {
    fmt.Fprintln(os.Stderr, "Falló al parsear la IP.")
    return ip, fMask, hostsReq, false
  }
  hostsReq, fok = StrToSeqOfInt(fReq, " ")
  if !fok {
    fmt.Fprintln(os.Stderr, "Falló al parsear los requerimeintos.")
    return ip, fMask, hostsReq, false
  }
  return ip, fMask, hostsReq, true
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
func PrintSubnetting(sn []Subnet)  {
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
    fmt.Println("-----------------------------------------")
  }
}
