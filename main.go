package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/NikitaVi/image_minifier/internal/format_detector"
	"github.com/gen2brain/beeep"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

//go:embed bin/*
var embeddedBinaries embed.FS

//go:embed ok.png
var icon []byte

func main() {

	err := EmbedInit()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	// Список поддерживаемых форматов
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

			format, err := format_detector.FormatDetector(img)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(format, img)

			switch format {
			case "jpeg":
				MinifierJPG(img)
			case "png":
				MinifierPNG(img)
			case "webp":
				MinifierWEBP(img)
			default:
				fmt.Println("Unsupported format:", format)
			}
		}(img)
	}

	wg.Wait()

	err = beeep.Notify("Успешно!", "Все файлы были сжаты!", icon)
	err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}

func MinifierPNG(imageName string) {
	alg := pathGenerator("pngquant")

	cmd := exec.Command(alg, "--quality=80-90", "--output", imageName, "--force", imageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("png: %s \n", imageName)

	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			println("Невозможно сжать")
		}
		fmt.Println(err.Error())
	}
}

func MinifierJPG(imageName string) {
	alg := pathGenerator("jpegoptim")

	cmd := exec.Command(alg, "--max=90", imageName)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		println(err.Error())
	}
}

func MinifierWEBP(imageName string) {
	metaDataMin := pathGenerator("webpmux")
	recodeMin := pathGenerator("cwebp")
	fmt.Println(metaDataMin)

	if err := exec.Command(metaDataMin, "-strip", "icc", imageName, "-o", imageName).Run(); err != nil {
		println(err.Error())
	}

	if err := exec.Command(metaDataMin, "-strip", "exif", imageName, "-o", imageName).Run(); err != nil {
		println(err.Error())
	}

	if err := exec.Command(metaDataMin, "-strip", "xmp", imageName, "-o", imageName).Run(); err != nil {
		println(err.Error())
	}

	if err := exec.Command(recodeMin, "-q", "75", "-m", "6", imageName, "-o", imageName).Run(); err != nil {
		println("Невозможно сжать")
	}
}

func pathGenerator(pathName string) string {
	ext := ""

	tempDir := os.TempDir()

	if runtime.GOOS == "windows" {
		ext = ".exe"
	}

	return fmt.Sprintf("%s\\%s%s", tempDir, pathName, ext)
}

func EmbedInit() error {
	sys := runtime.GOOS
	algorithms, _ := embeddedBinaries.ReadDir("bin/" + sys)

	for _, a := range algorithms {
		tmpPath := filepath.Join(os.TempDir(), a.Name())

		data, err := embeddedBinaries.ReadFile(filepath.Join("bin/"+sys, a.Name()))

		if err = os.WriteFile(tmpPath, data, 0755); err != nil {
			return err
		}
	}

	return nil
}
