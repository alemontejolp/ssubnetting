package ssubnetting

type Subnet struct {
  DecMask int
  DDNMask [4]int
  Id [4]int
  Broadcast [4]int
  FirstU [4]int
  LastU [4]int
  HostsU int
}
