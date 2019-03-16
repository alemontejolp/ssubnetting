# Simple Subnetting.

## Descripción.
Herramienta de línea de comandos que permite subnetear una red.

Dada una dirección IP, su máscara de subred y los hosts para cada subred,
calcula:

* ID de red.
* Broadcast de red.
* Primera dirección usable.
* Última dirección usable.
* Máscara de subred (Decimal)
* Máscara de subred (DDN).

para cada subred requerida.

## Uso.
Al invocar el programa, se le deben pasar obligatoriamente los flags:

* -ip
* -mask
* -req

donde -ip es la dirección base, -mask es su máscara y -req son la cantidad de
hosts mínimos para cada subred (separados por espacios).

Entonces, tenemos que el formato de uso es:

"ssubnetting -ip [dirección en formato DDN] -mask [máscara en decimal]
-req [host mínimis en cada subred separados por espacios]"

Si el subneteo no es posible hacerse con la configuración inicial dada,
se deplegará por la salida de error estándar un mensaje indicándolo.

## Ejemplo.

Al escribir: "ssubnetting -ip 192.168.23.0 -mask 24 -req 60 28 12 6 2 2".

podrás obtener:

```
Subneteo:
-----------------------------------------
Subred [0]:
ID de red: 192.168.23.0
Broadcast de red: 192.168.23.63
Primera dirección usable: 192.168.23.1
Última dirección usable: 192.168.23.62
Máscara de subred (Decimal): 26
Máscara de subred (DDN): 255.255.255.192
-----------------------------------------
Subred [1]:
ID de red: 192.168.23.64
Broadcast de red: 192.168.23.95
Primera dirección usable: 192.168.23.65
Última dirección usable: 192.168.23.94
Máscara de subred (Decimal): 27
Máscara de subred (DDN): 255.255.255.224
-----------------------------------------
Subred [2]:
ID de red: 192.168.23.96
Broadcast de red: 192.168.23.111
Primera dirección usable: 192.168.23.97
Última dirección usable: 192.168.23.110
Máscara de subred (Decimal): 28
Máscara de subred (DDN): 255.255.255.240
-----------------------------------------
Subred [3]:
ID de red: 192.168.23.112
Broadcast de red: 192.168.23.119
Primera dirección usable: 192.168.23.113
Última dirección usable: 192.168.23.118
Máscara de subred (Decimal): 29
Máscara de subred (DDN): 255.255.255.248
-----------------------------------------
Subred [4]:
ID de red: 192.168.23.120
Broadcast de red: 192.168.23.123
Primera dirección usable: 192.168.23.121
Última dirección usable: 192.168.23.122
Máscara de subred (Decimal): 30
Máscara de subred (DDN): 255.255.255.252
-----------------------------------------
Subred [5]:
ID de red: 192.168.23.124
Broadcast de red: 192.168.23.127
Primera dirección usable: 192.168.23.125
Última dirección usable: 192.168.23.126
Máscara de subred (Decimal): 30
Máscara de subred (DDN): 255.255.255.252
-----------------------------------------
```

¡Qué les sea útil!
