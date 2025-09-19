package file_init

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func EmbedInit(embeddedBinaries embed.FS) error {
	sys := runtime.GOOS
	algorithms, _ := embeddedBinaries.ReadDir("bin/" + sys)

	for _, a := range algorithms {
		tmpPath := filepath.Join(os.TempDir(), a.Name())

		data, err := embeddedBinaries.ReadFile("bin/" + sys + "/" + a.Name())
		if err != nil {
			fmt.Println("Error:", err)
		}

		if err = os.WriteFile(tmpPath, data, 0755); err != nil {
			return err
		}
	}

	return nil
}
