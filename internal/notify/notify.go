package notify

import "github.com/gen2brain/beeep"

func Notifying() error {
	if err := beeep.Alert("Изменения", "У вас есть изменения в", "assets/warning.png"); err != nil {
		return err
	}

	return nil
}
