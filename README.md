# Simple Subnetting
Licencia: MIT.

Lenguaje: Go.

Versión: 2.2.0

## Descripción
Herramienta de línea de comandos que permite subnetear una red,
además de sumar y restar hosts a una dirección dada.

### Subneteo

Dada una dirección IP, su máscara de subred y los hosts para cada subred,
calcula:

* ID de red.
* Broadcast de red.
* Primera dirección usable.
* Última dirección usable.
* Máscara de subred (Decimal)
* Máscara de subred (DDN).
* Direcciones disponibles.

para cada subred requerida.

### Sumar y restar hosts a una dirección

Dada una dirección ip y una cantidad de hosts, se puede sumar o restar esa
catidad de hosts a la dirección proporcionada.

Ejemplo: 192.168.1.0 - 16 = 192.168.0.240

Ejemplo: 192.168.1.0 + 16 = 192.168.1.16

## Uso

### Subneteo

Al invocar el programa, para hacer un subneteo,
se le deben pasar obligatoriamente los flags:

* -ip
* -mask
* -req

donde -ip es la dirección base, -mask es su máscara de subred y -req son la cantidad de
hosts mínimos para cada subred (separados por espacios).

Entonces, tenemos que el formato de uso es:

```
ssbnt -ip [dirección en formato DDN] -mask [máscara en decimal]
-req [host mínimos en cada subred separados por espacios]
```

Si no es posible hacer el subneteo con la configuración inicial dada,
se deplegará por la salida de error estándar un mensaje indicándolo.

### Sumar y restar hosts a una dirección.

Al invocar el programa, para hacer una suma o resta a una dirección,
se le deben pasar obligatoriamente los flags:

* -ip
* -subtr | -add

Se puede usar `-subtr` para restar y `-add` para sumar. No se deben usar
juntos en una misma ejecución o se ignorarán.

```
ssbnt -ip [dirección en formato DDN] -subtr [cantidad de hosts a restar]

ssbnt -ip [dirección en formato DDN] -add [cantidad de hosts a sumar]
```

## Opciones
Con respecto al subnetero, se pueden usar los flags:

* -sort [desc|asc]
* -lo

`sort`: Ordena las redes antes de subnetear. Si se usa este flag y no se le
pasa argumentos, se hará el subneteo desde la red más grande hasta la más
pequeña (desc). Si se usa (asc) será al revés.

`lo`: Imprime al final los siguientes datos:
* Dirección de inicio del bloque sobrante (si los requerimientos no llenaron toda la red).
* Cantidad de direcciones sobrantes.

## Ejemplo

### Subneteo

Al escribir:
```
ssbnt -ip 192.168.23.0 -mask 24 -req 60 28 12 6 2 2 -lo
```

podrás obtener por la salida estándar:

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
Direcciones disponibles: 62
-----------------------------------------
Subred [1]:
ID de red: 192.168.23.64
Broadcast de red: 192.168.23.95
Primera dirección usable: 192.168.23.65
Última dirección usable: 192.168.23.94
Máscara de subred (Decimal): 27
Máscara de subred (DDN): 255.255.255.224
Direcciones disponibles: 30
-----------------------------------------
Subred [2]:
ID de red: 192.168.23.96
Broadcast de red: 192.168.23.111
Primera dirección usable: 192.168.23.97
Última dirección usable: 192.168.23.110
Máscara de subred (Decimal): 28
Máscara de subred (DDN): 255.255.255.240
Direcciones disponibles: 14
-----------------------------------------
Subred [3]:
ID de red: 192.168.23.112
Broadcast de red: 192.168.23.119
Primera dirección usable: 192.168.23.113
Última dirección usable: 192.168.23.118
Máscara de subred (Decimal): 29
Máscara de subred (DDN): 255.255.255.248
Direcciones disponibles: 6
-----------------------------------------
Subred [4]:
ID de red: 192.168.23.120
Broadcast de red: 192.168.23.123
Primera dirección usable: 192.168.23.121
Última dirección usable: 192.168.23.122
Máscara de subred (Decimal): 30
Máscara de subred (DDN): 255.255.255.252
Direcciones disponibles: 2
-----------------------------------------
Subred [5]:
ID de red: 192.168.23.124
Broadcast de red: 192.168.23.127
Primera dirección usable: 192.168.23.125
Última dirección usable: 192.168.23.126
Máscara de subred (Decimal): 30
Máscara de subred (DDN): 255.255.255.252
Direcciones disponibles: 2
-----------------------------------------
Dirección de inicio del bloque sobrante: 192.168.23.128
Direcciones sobrantes: 128
-----------------------------------------
```

### Sumar o restar hosts a una dirección

Al escribir

```
ssbnt -ip 192.168.1.0 -subtr 16
```

Podrás obtener

```
Cantidad de hosts restados: 16
Dirección: 192.168.0.240
```

Al escribir

```
ssbnt -ip 192.168.1.0 -add 16
```

```
Cantidad de hosts sumados: 16
Dirección: 192.168.1.16
```

## Compilación
En la raíz del proyecto:

```
go build -o ssbnt ./ssubnetting.go
```

¡Qué les sea útil!
