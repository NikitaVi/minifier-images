package main

import (
	_ "embed"
	minifier "github.com/NikitaVi/minifier_core/pkg"
	"github.com/gen2brain/beeep"
	"path/filepath"
	"sync"
)

//go:embed resources/ok.png
var icon []byte

func main() {

	err := minifier.EmbedInit()

	//err := file_init.EmbedInit(embeddedBinaries)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	patterns := []string{"*.jpg", "*.jpeg", "*.png", "*.webp"}
	var files []string

	for _, p := range patterns {
		matches, _ := filepath.Glob(p)
		files = append(files, matches...)
	}

	for _, img := range files {
		wg.Add(1)

		go func(img string) {
			defer wg.Done()

			minifier.Minifier(img)
		}(img)
	}

	wg.Wait()

	err = beeep.Notify("Успешно!", "Все файлы были сжаты!", icon)
	if err != nil {
		panic(err)
	}
}
