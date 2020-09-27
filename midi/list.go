package midi

import (
	// "gitlab.com/gomidi/midi/reader"
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
		log.Println(i.String())
	}
}
