package minifiers

import (
	"github.com/NikitaVi/image_minifier/internal/utils"
	"os"
	"os/exec"
)

func MinifierJPG(imageName string) {
	alg := utils.PathGenerator("jpegoptim")

	cmd := exec.Command(alg, "--max=90", imageName)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		println(err.Error())
	}
}
