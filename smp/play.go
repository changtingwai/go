// play.go
package mp

import (
	"fmt"
)

type Player interface {
	Play(source string)
}

func play(source string, mtype string) {
	var p Player
	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("unsupported music type", mtype)
		return
	}
	p.Play(source)
}
