package pkg

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/kardianos/ftps"
)

// Client estructura que representa el cliente FTPS
type Client struct {
	ftpsClient *ftps.Client
}

func Kardianos() error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Crea una nueva instancia del cliente FTPS
	client, err := NewClient()
	if err != nil {
		fmt.Println("Error al crear el cliente FTPS:", err)
		return err
	}

	defer func(ftpsClient *ftps.Client) {
		err := ftpsClient.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(client.ftpsClient)

	fmt.Println("He podido conectar con el servidor")
	/*
	   f1Buff := &bytes.Buffer{}
	   	err = client.ftpsClient.Download(ctx, "/pub/example/readme.txt", f1Buff)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	*/

	list, err := client.ftpsClient.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range list {
		fmt.Println(item)
	}
	if g, w := len(list), 1; g != w {
		log.Fatalf("got %d items, want %d", g, w)
	}

	err = client.ftpsClient.Close()
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// NewClient crea una nueva instancia del cliente FTPS
func NewClient() (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Crear una configuración TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Configura los detalles de conexión FTPS
	opt := ftps.DialOptions{
		Host:        "test.rebex.net",
		Port:        21,
		Username:    "demo",
		Passowrd:    "password",
		TLSConfig:   tlsConfig,
		ExplicitTLS: true,
	}

	// Dial a FTPS server and return a Client
	ftpsClient, err := ftps.Dial(ctx, opt)
	if err != nil {
		return nil, fmt.Errorf("error al crear el cliente FTPS: %w", err)
	}

	return &Client{ftpsClient: ftpsClient}, nil
}
