# integrador-avatar-me

Esto es el trabajo integrador de conceptos básicos del curso de Golang.
El mismo es un módulo que genera un identicon único para cada usuario en base a una información del mismo.

# Explicación de su uso

Para su uso debe crearse una estructura tipo Info, la cual contendrá la información de la cual se obtendrá un hash y la dirección donde se guardará la imagen generada. A continuación se encuentra un ejemplo.

```go
package main

import (
	"fmt"

	"github.com/AtilioBoher/integrador-avater-me/avatar"
)

func main() {
	info := avatar.Info{
		StrInfo:  "atilio",
		FilePath: "identicon.png",
	}

	a := avatar.GimmeAnAvatarGenerator()
	err := a.GenerateAndSaveAvatar(info)
	if err != nil {
		fmt.Println(err)
	}

}
```
