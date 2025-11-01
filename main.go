package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	data, err := os.ReadFile("/tmp/test.png")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	systray.SetIcon(data)
	systray.SetTooltip("Простое трей-приложение")

	// Добавляем элементы меню
	mStatus := systray.AddMenuItem("Статус: Работает", "Статус приложения")
	mClick := systray.AddMenuItem("Нажми меня!", "Тестовая кнопка")
	mNotify := systray.AddMenuItem("Notify", "Notify")
	systray.AddSeparator()
    mTest := systray.AddMenuItemCheckbox("test","test", false)
    mTest.SetIcon(data)
    mTest.AddSubMenuItem("test", "test")
	mQuit := systray.AddMenuItem("Выход", "Закрыть приложение")

	// Счетчик кликов
	counter := 0

	// Обработка событий
	go func() {
		for {
			select {
			case <-mClick.ClickedCh:
				counter++
				mStatus.SetTitle(fmt.Sprintf("Кликов: %d", counter))
				systray.SetTooltip(fmt.Sprintf("Приложение (кликов: %d)", counter))
			case <-mNotify.ClickedCh:
				Toast(data)
			case <-mTest.ClickedCh:
				mTest.Check()
			case <-mQuit.ClickedCh:
				fmt.Println("Завершение работы...")
				systray.Quit()
				return
			}
		}
	}()

	// Обновление времени в тултипе
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				systray.SetTooltip(fmt.Sprintf("Работает с: %s",
					time.Now().Format("15:04")))
			}
		}
	}()
}

func onExit() {
	fmt.Println("Приложение закрыто")
	os.Exit(0)
}

func Toast(icon []byte) {

	// err := beeep.Notify("Заголовок", "Сообщение", "/tmp/test.png")
	// if err != nil {
	// 	panic(err)
	// }

	beeep.Alert("Внимание!", "Что-то случилось", "assets/warning.png")

	// Или просто
	// beeep.Notify("Title", "Message body", "")
}
