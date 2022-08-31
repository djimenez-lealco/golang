package main

import "fmt"

// -------------------------------------------------------
// enviarEmails sin OpenClose
// -------------------------------------------------------

func enviarEmailSinOpenClase(comercioId int) {
	body := ""
	title := ""
	if comercioId == 1 {
		body = "<h1> Hola BIenvenido al comercio</h1>"
		title = "Coorbanca te Invita"
	} else if comercioId == 2 {
		body = "Nuevo mensaje de correo electronico"
		title = "Arturo Calle Gran descuento"
	} else {
		body = "Nuevo mensaje de correo electronico"
		title = "Arturo Calle Gran descuento"
	}
	fmt.Println(title, body)
}

// -------------------------------------------------------
// enviar aplicando el Principio Open Close
// -------------------------------------------------------
type Email struct {
	ComercioId int
	Title      string
	Body       string
}

func enviarEmailOpenClose(comercioId int, emails []Email) (bool, error) {

	var email *Email
	for index, _ := range emails {
		if emails[index].ComercioId == comercioId {
			email = &emails[index]
		}
	}
	if email == nil {
		panic("No Existe el tipo de email para el comercio")
	}
	fmt.Println(email.Title, email.Body)
	return true, nil
}

func enviador() func(int) (bool, error) {
	// OPEN CLOSE
	emails := []Email{
		{
			ComercioId: 1,
			Body:       "<h1> Hola BIenvenido al comercio</h1>",
			Title:      "Coorbanca te Invita",
		},
		{
			ComercioId: 2,
			Body:       "Nuevo mensaje de correo electronico",
			Title:      "Arturo Calle Gran descuento",
		},
		{
			ComercioId: 3,
			Body:       "Nuevo mensaje de correo electronico",
			Title:      "Arturo Calle Gran descuento 2",
		},
	}
	return func(comerdioId int) (bool, error) {
		return enviarEmailOpenClose(comerdioId, emails)
	}
}
func main() {
	// Sin OpenClose
	enviarEmailSinOpenClase(1)
	//Con OpenClose
	enviador()(1)
}
