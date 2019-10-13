package sounds

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func init() {
	speaker.Play(&q)
	sr = beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
}

var (
	sr beep.SampleRate
)

func Consume(ch <-chan string) {
	fmt.Println("waiting for a sample....")
	for {
		fmt.Println("Got one!")
		fn := <-ch

		f, err := os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}

		s, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		r := beep.Resample(4, format.SampleRate, sr, s)

		speaker.Lock()
		q.Add(r)
		speaker.Unlock()
	}
}
