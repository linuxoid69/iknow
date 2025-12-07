package menu

import (
	"fmt"

	"github.com/linuxoid69/iknow/internal/fs"
)


const (
	IconMain string = "icons/main.png"
)

func GetIcon(path string) (icon []byte, err error) {
	icon, err = fs.FS.ReadFile(path)
	if err != nil {
		return icon, fmt.Errorf("can't read an icon file %w", err)
	}
	return icon, nil
}
