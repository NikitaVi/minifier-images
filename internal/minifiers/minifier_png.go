package minifiers

import (
	"fmt"
	"github.com/NikitaVi/image_minifier/internal/utils"
	"os"
	"os/exec"
)

func MinifierPNG(imageName string) {
	alg := utils.PathGenerator("pngquant")

	cmd := exec.Command(alg, "--quality=80-90", "--output", imageName, "--force", imageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			println("Невозможно сжать")
		}
		fmt.Println(err.Error())
	}
}
