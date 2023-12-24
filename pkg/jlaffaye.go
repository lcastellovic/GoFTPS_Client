package pkg

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"net/textproto"
	"os"
	"os/exec"
	"strings"
)

// Servidor FTP gratuito para pruebas experimentales
const (
	HOST     = "test.rebex.net"
	PORT     = "21"
	USER     = "demo"
	PASSWORD = "password"
)

type lecturaStruct struct {
	Comando string
	Valor   string
}

func ClienteFTPS() error {

	fmt.Println("************************************************************")
	fmt.Println("*       FTPS-CLI - Cliente de FTPS en Consola              *")
	fmt.Println("*                 BY lcastellovic                          *")
	fmt.Println("************************************************************")
	fmt.Println()
	fmt.Println("                     ¡Bienvenido!                           ")
	fmt.Println("Este es un cliente FTPS por consola que acepta TLSv1.3")
	fmt.Println("Escribe 'ayuda' para obtener la lista de comandos disponibles.")
	fmt.Println()
	fmt.Println("************************************************************")

	err := lectorPrincipal()
	if err != nil {
		return err
	}

	return nil
}

/*
********************   Empiezan las funciones del programa     ********************
 */

func lectorPrincipal() error {
	// Crear el objeto os.Stdin para leer del canal estándar de entrada
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresa número:")
	fmt.Printf("1: %-20s 2: %-20s\n", "Start Conn", "LIST")
	fmt.Printf("3: %-20s 4: %-20s\n", "CWD", "Retr")
	fmt.Printf("5: %-20s\n", "CLEAR")
	var (
		c   *ftp.ServerConn
		err error
	)

	// Usar el scanner para leer una línea del canal estándar de entrada
	for scanner.Scan() {
		// El texto ingresado se encuentra en scanner.Text()
		texto := scanner.Text()
		lineaCompleta := strings.Split(texto, " ")
		textoStruct := lecturaStruct{}
		if len(lineaCompleta) < 2 {
			textoStruct = lecturaStruct{lineaCompleta[0], ""}
		} else {
			textoStruct = lecturaStruct{lineaCompleta[0], lineaCompleta[1]}
		}

		switch textoStruct.Comando {
		case "1":
			c, err = conexionFtp()
			fmt.Println(" [SUCCSESS] - Sesión iniciada correctamente [Conexión - Autenticación]")
			if err != nil {
				return err
			}

		case "ayuda":
			listarComandos()

		case "2":
			err = listarDirectorio(textoStruct.Valor, c)
			if err != nil {
				var tpErr *textproto.Error
				if errors.As(err, &tpErr) {
					switch tpErr.Code {
					case 550:
						fmt.Printf("[ERROR] - El directorio %s no existe\n", textoStruct.Valor)
					}
				}
			}

		case "3":
			err := cambiarDirectorio(textoStruct.Valor, c)
			if err != nil {
				var tpErr *textproto.Error
				if errors.As(err, &tpErr) {
					switch tpErr.Code {
					case 550:
						fmt.Printf("[ERROR] - El directorio %s no existe\n", textoStruct.Valor)
					}
				}
			}

		case "4":
			switch textoStruct.Valor {
			case "":
				fmt.Println("[WARNING] - Insertar el nombre del fichero")
			}
			err = descargaFichero(textoStruct.Valor, c)
			if err != nil {
				var tpErr *textproto.Error
				if errors.As(err, &tpErr) {
					switch tpErr.Code {
					case 550:
						fmt.Printf("[ERROR] - El fichero %s no existe\n", textoStruct.Valor)
					}
				}
			}

		case "5":
			err := limpiarConsola()
			if err != nil {
				fmt.Println("[ERROR] - No se pudo limpiar la consola")
			}

		default:
			fmt.Println("[WARNING] - Porfavor, ingrese un código correcto")
		}

		fmt.Print("Ingresa el comando (o presiona Ctrl+C para salir o 'ayuda' para mostrar los comandos): ")
	}

	err = c.Quit()
	if err != nil {
		return err
	}

	return err
}

// Operación #1 que inicia la conexión al servidor
func conexionFtp() (conn *ftp.ServerConn, err error) {
	// Configuración del struct tls para la encriptación
	tlsConf := tls.Config{
		InsecureSkipVerify: true,
		ClientSessionCache: tls.NewLRUClientSessionCache(0),
		ServerName:         "Server FTPS",
	}
	// Generar la conexión del servdidor especificando tipo de conexión
	// en este caso conexión FTP coon TLS explícito (FTPES).
	c, err := ftp.Dial(HOST+":"+PORT, ftp.DialWithExplicitTLS(&tlsConf))
	if err != nil {
		return nil, err
	}

	// Iniciar la sesión
	err = c.Login(USER, PASSWORD)
	if err != nil {
		return nil, err
	}

	return c, err
}

// Operación #2 para listar ficheros y directorios usando el comando LIST de FTP
func listarDirectorio(directorio string, conn *ftp.ServerConn) error {
	fmt.Printf("[SUCCSESS] - Empiezo a listar del directorio: %s\n", directorio)
	entries, err := conn.List(directorio)
	if err != nil {
		return err
	}
	printHeaders()
	for _, entry := range entries {
		printEntry(*entry)
	}
	fmt.Println(strings.Repeat("-", 98)) // Separador
	return nil
}

// Operación #3 para cambiar el Work Directory usando el comando CWD de FTP
func cambiarDirectorio(directorio string, conn *ftp.ServerConn) error {
	err := conn.ChangeDir(directorio)
	if err != nil {
		return err
	}
	return nil
}

// Operación #4 para descargar ficheros usando el comando RETR de FTP
func descargaFichero(fichero string, conn *ftp.ServerConn) error {

	response, err := conn.Retr(fichero)
	if err != nil {
		return err
	}
	defer func(response *ftp.Response) {
		err := response.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response)

	bytes, err := io.ReadAll(response)
	if err != nil {
		return err
	}

	err = os.WriteFile(fichero, bytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Operación #5 para limpiar la consola con el comando correspondiente del SO
func limpiarConsola() error {

	// Determinar el comando según el sistema operativo
	var cmd *exec.Cmd

	cmd = exec.Command("cmd", "/c", "cls")

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}
	return err
}

func listarComandos() {
	fmt.Println("Lista de comandos actuales: ")
	fmt.Println("1: Para empezar la conexión")
	fmt.Println("2: Para listar el directorio")
	fmt.Println("3: Para cambiar el directorio de trabajo")
	fmt.Println("4: Para descargar un fichero")
	fmt.Println("5: Para limpiar la consola")
}

// Función para formatear y mostrar un Entry en columnas
func printEntry(entry ftp.Entry) {
	const entryFormat = "%-20s %-30s %-15s %-10d %s\n"
	fmt.Printf(entryFormat, entry.Name, entry.Target, entry.Type.String(), entry.Size, entry.Time.Format("2006-01-02 15:04:05"))
}

// Función para imprimir los encabezados de las columnas
func printHeaders() {
	const headerFormat = "%-20s %-30s %-15s %-10s %s\n"
	fmt.Printf(headerFormat, "Name", "Target", "Type", "Size", "Time")
	fmt.Println(strings.Repeat("-", 98)) // Separador
}
