package sounds

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
	"time"
)

func init() {
	speaker.Play(&q)
	sr = beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
}

var (
	sr       beep.SampleRate
	clipPath string
	once     sync.Once
)

func Consume(ch <-chan string) {
	once.Do(func() {
		clipPath = viper.GetString("clip_path")
		if clipPath == "" {
			fmt.Println("Unable to access path to clip. Is it specified?")
			os.Exit(1)
		}
		fmt.Printf("Use path %s for finding audio clips\n", clipPath)
	})

	for {
		fn := <-ch

		f, err := os.Open(clipPath + fn)
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
