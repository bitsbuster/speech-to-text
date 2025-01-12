package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"github.com/gordonklaus/portaudio"
)

func RecordAudio() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	// Configurar la grabación
	buffer := make([]int16, 64)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(buffer), func(in []int16) {
		copy(buffer, in)
	})
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	// Crear archivo de salida
	outFile, err := os.Create("output.wav")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	enc := wav.NewEncoder(outFile, 44100, 16, 1, 1)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Press 'r' to start recording. Press 'space' to stop.")

	recording := false
	stopChan := make(chan bool)

	go func() {
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		defer keyboard.Close()
		for {
			if v, key, err := keyboard.GetKey(); err == nil {
				fmt.Printf("Key: %c\n", v)
				switch {
				case v == 'r':
					if !recording {
						fmt.Println("Recording... Press 'space' to stop.")
						recording = true
						go func() {
							err := stream.Start()
							if err != nil {
								panic(err)
							}
							for recording {
								stream.Read()
								audioBuffer := &audio.IntBuffer{
									Data:   make([]int, len(buffer)),
									Format: &audio.Format{SampleRate: 44100, NumChannels: 1},
								}
								for i, sample := range buffer {
									audioBuffer.Data[i] = int(sample)
								}
								err = enc.Write(audioBuffer)
								if err != nil {
									panic(err)
								}
							}
							stream.Stop()
							stopChan <- true
						}()
					}
				case key == keyboard.KeySpace:
					if recording {
						fmt.Println("Recording stopped.")
						recording = false
						<-stopChan
					}
				case key == keyboard.KeyEsc:
					return
				}
			}
		}
	}()

	go func() {
		for recording {
			audioBuffer := &audio.IntBuffer{
				Data:   make([]int, len(buffer)),
				Format: &audio.Format{SampleRate: 44100, NumChannels: 1},
			}
			for i, sample := range buffer {
				audioBuffer.Data[i] = int(sample)
			}
			err := enc.Write(audioBuffer)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	<-c
	enc.Close()
	fmt.Println("\nApplication finished.")
}
