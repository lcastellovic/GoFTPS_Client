package pkg

import (
	"fmt"
	"github.com/sacloud/ftps"
	"log"
)

func Webguerilla() error {
	ftpsCli := new(ftps.FTPS)

	ftpsCli.TLSConfig.InsecureSkipVerify = true
	err := ftpsCli.Connect("test.rebex.net", 21)
	if err != nil {
		return err
	}

	err = ftpsCli.Login("demo", "password")
	if err != nil {
		return err
	}

	err = ftpsCli.ChangeWorkingDirectory("pub")
	if err != nil {
		return err
	}

	directory, err := ftpsCli.PrintWorkingDirectory()
	if err != nil {
		return err
	}
	log.Printf("Current working directory: %s", directory)

	entries, err := ftpsCli.List()
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Println(entry)
	}

	defer func(ftpsCli *ftps.FTPS) {
		err := ftpsCli.Quit()
		if err != nil {
			log.Fatal(err)
		}
	}(ftpsCli)

	return nil
}
