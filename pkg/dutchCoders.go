package pkg

import (
	"crypto/tls"
	"fmt"

	"github.com/dutchcoders/goftp"
)

func DutchCoders() error {
	// For debug messages: goftp.ConnectDbg("ftp.server.com:21")
	ftp, err := goftp.Connect("test.rebex.net:21")
	if err != nil {
		return err
	}
	defer func(ftp *goftp.FTP) {
		err := ftp.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ftp)

	fmt.Println("Successfully connected")

	// TLS client authentication
	config := tls.Config{
		InsecureSkipVerify: true,
		ClientSessionCache: tls.NewLRUClientSessionCache(0),
		ServerName:         "Server FTP",
	}

	err = ftp.AuthTLS(&config)
	if err != nil {
		panic(err)
	}

	// Username / password authentication
	err = ftp.Login("demo", "password")
	if err != nil {
		return err
	}

	entries, err := ftp.List("")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Println(entry)
	}
	return nil
}
