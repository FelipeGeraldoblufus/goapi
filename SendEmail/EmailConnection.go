package SendEmail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type Dest struct {
	Name string
}

func EmailConnect(toName, toAddress, subjecto string) {

	from := mail.Address{"testingparago", "testingparago@gmail.com"}
	to := mail.Address{toName, toAddress}
	subject := subjecto
	dest := Dest{Name: to.Name} // destinatario parseado para el template

	// configuracion de headers

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = "text/html; charset='UTF-8'"

	// armar el correo
	message := ""
	// recorrer el mapa para agregar la informacion a la variable message
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	// tomar el template y pasar la informacion de la estructura
	t, err := template.ParseFiles("template.html")
	checkErr(err)
	// Ejecutar el template y guardarlo
	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkErr(err)

	message += buf.String()

	// conexion al correo para enviar el mensaje

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	// obtener auth para iniciar sesion en correo test
	auth := smtp.PlainAuth("", "testingparago@gmail.com", "atuy akve aide ifls", host)
	// configurar tls
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkErr(err)

	// crear cliente de la conexion
	client, err := smtp.NewClient(conn, host)
	checkErr(err)

	err = client.Auth(auth)
	checkErr(err)

	// quien sera el que envie el correo y el receptor
	err = client.Mail(from.Address)
	checkErr(err)

	err = client.Rcpt(to.Address)
	checkErr(err)

	// cual sera la data
	w, err := client.Data()
	checkErr(err)

	_, err = w.Write([]byte(message))
	checkErr(err)

	err = w.Close()
	checkErr(err)

	err = client.Quit()
	checkErr(err)
}
