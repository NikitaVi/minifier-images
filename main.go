package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/NikitaVi/image_minifier/internal/file_init"
	"github.com/NikitaVi/image_minifier/internal/minifiers"
	"github.com/NikitaVi/image_minifier/internal/utils"
	"github.com/gen2brain/beeep"
	"path/filepath"
	"sync"
)

//go:embed bin/*
var embeddedBinaries embed.FS

//go:embed resources/ok.png
var icon []byte

func main() {

	err := file_init.EmbedInit(embeddedBinaries)
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

			format, err := utils.FormatDetector(img)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(format, img)

			switch format {
			case "jpeg":
				minifiers.MinifierJPG(img)
			case "png":
				minifiers.MinifierPNG(img)
			case "webp":
				minifiers.MinifierWEBP(img)
			default:
				fmt.Println("Unsupported format:", format)
			}
		}(img)
	}

	wg.Wait()

	err = beeep.Notify("Успешно!", "Все файлы были сжаты!", icon)
	if err != nil {
		panic(err)
	}
}
