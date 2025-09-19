package minifiers

import (
	"fmt"
	"github.com/NikitaVi/image_minifier/internal/utils"
	"os/exec"
)

func MinifierWEBP(imageName string) {
	metaDataMin := utils.PathGenerator("webpmux")
	recodeMin := utils.PathGenerator("cwebp")
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
