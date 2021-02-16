package mp

import (
	"fmt"
)

type WAVPlayer struct {
	stat     int
	progress int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music", source)

	fmt.Println("\nFinished playing", source)
}
