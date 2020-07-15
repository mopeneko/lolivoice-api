package model

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mopeneko/lolivoice"
	"io/ioutil"
	"os"
)

func GenerateLoliVoice(text string) ([]byte, error) {
	fileName := fmt.Sprintf(
		"%s.wav",
		uuid.New().String(),
	)

	dicPath := os.Getenv("LV_DICPATH")
	voicePath := os.Getenv("LV_VOICEPATH")

	generator := lolivoice.NewLoliVoiceGenerator(dicPath, voicePath)
	err := generator.Generate(text, fileName)
	if err != nil {
		return []byte{}, err
	}

	defer os.Remove(fileName)

	bin, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []byte{}, err
	}

	return bin, nil
}

