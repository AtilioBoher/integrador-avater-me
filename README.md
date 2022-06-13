# integrador-avatar-me

Esto es el trabajo integrador de conceptos básicos del curso de Golang.
El mismo es un módulo que genera un identicon único para cada usuario en base a una información del mismo.

# Explicación de su uso

Para su uso debe crearse una estructura tipo Info, la cual contendrá la información de la cual se obtendrá un hash y la dirección donde se guardará la imagen generada.

info := avatar.Info{
		StrInfo:  "atilio",
		FilePath: "identicon.png",
	}

Luego debe crearese un generador de avatar

a := avatar.GimmeAnAvatarGenerator()

Este generador de avatar genera la imagen en base a la información guardada en la estructura info

err := a.GenerateAndSaveAvatar(info)
	if err != nil {
		fmt.Println(err)
	}
