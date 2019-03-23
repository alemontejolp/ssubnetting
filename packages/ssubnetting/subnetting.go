package ssubnetting

import (
  "math"
)

// Calcula cuántos hosts pueden haber con la máscara dada.
func HostsByMask(m int) int {
  return int(math.Pow(2, float64(32 - m)))
}

// Convierte la máscara de notación decimal
// a notación decimal por puntos.
// mask: Array donde se guardará la máscara.
// m: máscara en notacón decimal.
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

// Calcula la máscara mínima que debería tener una red para
// soportar los hosts indicados.
func CalcMinMask(h int) (m int) {
  for m = 30; HostsByMask(m) - 2 < h; m-- {}
  return
}

// Suma a la dirección proporcionada una cantidad 'h' de hosts.
// Retorna cómo quedaría la dirección después de haberse
// recorrido los 'h' hosts hacia delante.
// Ejemplo: 192.168.1.0 + 16 = 192.168.1.16
// Ejemplo: 192.168.1.254 + 20 = 192.168.2.18
func AddAddr(addr *[4]int, h int)  {
  addr[3] += h
  for i := 3; addr[i] > 255; i-- {
    addr[i - 1] = addr[i - 1] + (addr[i] / 256)
    addr[i] = addr[i] % 256
  }
}

// Resta una cantidad 'h' de hosts a la dirección proporcionada.
// Retorna cómo quedaría la dirección después de haberse
// recorrido los 'h' hosts hacia atrás.
// Ejemplo: 192.168.1.0 - 16 = 192.168.0.240
// Ejemplo: 192.168.1.254 - 20 = 192.168.1.234
func SubAddr(addr *[4]int, h int)  {
  //Restar todo al último octeto.
  addr[3] -= h
  //Pasar por los 4 octetos para asegurarnos de hacerle
  //las adecuaciones necesarias a la dirección.
  for i := 3; i >= 0; i-- {
    //Mientras el octeto actual sea negativo, debemos hacer
    //las siguientes adecuaciones.
    for addr[i] < 0 {
      //Calcular cuanto sobraría de repartir el número actual en
      //bloques de 256.
      r := addr[i] % 256
      //Si es posible hacer al menos un bloque de 265 con el valor del
      //octeto actual, restar ese número de bloques al octeto anterior
      //y asignar el residuo al octeto actual.
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
        //Si no se pueden hacer bloques de 256, 1 al octeto anterior y
        //al octeto actual asignar el valor de 256 - |addr[i]|.
        addr[i - 1]--
        if r != 0 {
          addr[i] = 256 + r //recordar que 'r' aquí es negativo...
        } else {
          addr[i] = 0 //Este podría no ser necesario...
        }
      }
    }
  }
  //Al llegar aquí, ya se habŕa valanceado/adecuado correctamente la dirección.
}

// Calcula los siguientes datos de una subred, dada el ID y la máscara:
// addr: ID de red. Debe ser proporcionado correctamente.
// mask: Máscara de red. Debe ser proporcionado correctamente.
// last: Dirección de broadcast. Calculada por la función.
// fusb: primera dirección usable. Calculada por la función.
// lusb: última dirección usable. Calculada por la función.
func CalcSubnet(addr [4]int, mask int, last, fusb, lusb *[4]int, hostsav *int)  {
  CopyAddr(addr, last)
  AddAddr(last, HostsByMask(mask) - 1)
  CopyAddr(addr, fusb)
  AddAddr(fusb, 1)
  CopyAddr(*last, lusb)
  SubAddr(lusb, 1)
  *hostsav = HostsByMask(mask) - 2
}

//Suma los hosts de todas las máscaras proveidas.
func SumOfMasks(masks []int) int {
  l := len(masks)
  ac := 0
  for i := 0; i < l; i++ {
    ac += HostsByMask(masks[i])
  }
  return ac;
}

// Revisa si las subredes requeridas no desbordan la red base.
// Es decir, si caben en la subred base.
// bm: base mask, masks[]: las máscaras de cada subred.
func ValidSubnetting(bm int, masks []int) bool {
  ac := SumOfMasks(masks)
  return HostsByMask(bm) >= ac
}

// Calcula la máscara mínima para satisfacer cada requerimiento
// de host en el slide.
func GetMaskByHostReq(hosts []int) (masks []int)  {
  l := len(hosts)
  masks = make([]int, l)
  for i := 0; i < l; i++ {
    masks[i] = CalcMinMask(hosts[i])
  }
  return
}

//Dada una ip base y la lista de máscaras, realiza el subneteo de la red.
//Retorna el arreglo de subredes, la dirección de inicio del bloque
//restante y los hosts sobrantes. So no hay sobante, la dirección de
//inicio del bloque restante resán todos ceros.
func Subnetting(ip [4]int, mask int, masks []int) ([]Subnet, [4]int, int) {
  var currAddr[4]int
  l := len(masks)
  sn := make([]Subnet, l)
  CopyAddr(ip, &currAddr)
  for i := 0; i < l; i++ {
    CopyAddr(currAddr, &sn[i].Id)
    CalcSubnet(sn[i].Id, masks[i], &sn[i].Broadcast, &sn[i].FirstU, &sn[i].LastU, &sn[i].HostsAvailable)
    CopyAddr(sn[i].Broadcast, &currAddr)
    AddAddr(&currAddr, 1)
    sn[i].DecMask = masks[i]
    DDNMask(&sn[i].DDNMask, masks[i])
  }
  // leftoverHosts = totalhosts - usedhosts
  leftoverHosts := HostsByMask(mask) - SumOfMasks(masks)
  if leftoverHosts == 0 {
    for i := 0; i < 4; i++ {
      currAddr[i] = 0
    }
  }
  return sn, currAddr, leftoverHosts
}
