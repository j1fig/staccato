package midi

import (
	// "gitlab.com/gomidi/midi/reader"
	"fmt"
	"gitlab.com/gomidi/rtmididrv"
	"log"
)

func List() {
	drv, err := rtmididrv.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	inputs, err := drv.Ins()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, i := range inputs {
		fmt.Printf("port: %d\tname: %s", i.Number(), i.String())
	}
}
