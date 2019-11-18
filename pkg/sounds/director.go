package sounds

import (
	"errors"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
	"sync"
	"time"
)

func init() {
	speaker.Play(&q)
	sr = beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
}

type AudioFormat int

const (
	MP3 AudioFormat = iota
	WAV
)

func (af AudioFormat) String() string {
	return [...]string{"mp3", "wav"}[af]
}

var audioFormats = []string{"mp3", "wav"}

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

	var (
		s      beep.StreamSeekCloser
		format beep.Format
		err    error
		f      *os.File
	)

	for {
		fn := <-ch

		f, err = os.Open(clipPath + fn)
		if err != nil {
			log.Fatal(err)
		}

		re := regexp.MustCompile(`[^.]+$`)
		ext := re.Find([]byte(f.Name()))

		if ext == nil {
			log.Fatalln(errors.New("can not infer audio encoding without extension for file " + f.Name()))
		}

		switch string(ext) {
		case MP3.String():
			s, format, err = mp3.Decode(f)
		case WAV.String():
			s, format, err = wav.Decode(f)
		default:
			log.Fatalln(errors.New("unrecognized audio encoding format found: " + string(ext)))
		}

		if err != nil {
			log.Fatal(err)
		}

		r := beep.Resample(4, format.SampleRate, sr, s)

		speaker.Lock()
		q.Add(r)
		speaker.Unlock()
	}
}
