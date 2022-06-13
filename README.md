# integrador-avatar-me

Esto es el trabajo integrador de conceptos básicos del curso de Golang.
El mismo es un módulo que genera un identicon único para cada usuario en base a una información del mismo.

# Explicación de su uso

Para su uso debe crearse una estructura tipo Info, la cual contendrá la información de la cual se obtendrá un hash y la dirección donde se guardará la imagen generada.

A continuación se presenta un ejemplo.
-----
```go
package main
import (
        "log"
        "github.com/tarm/serial"
)
func main() {
        c := &serial.Config{Name: "COM45", Baud: 115200}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }
        
        n, err := s.Write([]byte("test"))
        if err != nil {
                log.Fatal(err)
        }
        
        buf := make([]byte, 128)
        n, err = s.Read(buf)
        if err != nil {
                log.Fatal(err)
        }
        log.Printf("%q", buf[:n])
}
```
