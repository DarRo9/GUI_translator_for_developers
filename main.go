package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"sync"
)

var from_e_to_r_s sync.Map
var from_r_to_e_s sync.Map

func main() {
	from_e_to_r := map[string]string{
		"access mechanism": "механизм доступа",
		"addressbuffer":    "адресный буфер",
		"appliance":        "устройство",
		"assignment":       "присвоение",
		"bandwidth":        "пропускная способность",
		"cloud storage":    "облачное хранилище данных",
		"compile":          "компилировать",
		"compress":         "сжимать",
		"credentials":      "учетные данные",
		"database":         "база данных",
		"debug":            "отлаживать",
		"drawback":         "недостаток",
		"eject":            "извлекать устройство",
		"encrypt":          "зашифровывать",
		"folder":           "папка",
		"layout":           "макет, разметка",
		"maintenance":      "поддержка",
		"namespace":        "пространство имен",
		"negotiate":        "вести переговоры",
		"outsource":        "осуществлять аутсорсинг",
		"password":         "пароль",
		"plug in":          "подключать",
		"prohibit":         "запрещать",
		"replace":          "заменить",
		"uninstall":        "удалять",
		"upload":           "загрузить",
		"validation":       "проверка",
		"variable":         "переменная",
		"vendor":           "поставщик",
	}
	from_r_to_e := map[string]string{
		"механизм доступа":          "access mechanism",
		"адресный буфер":            "addressbuffer",
		"устройство":                "appliance",
		"присвоение":                "assignment",
		"пропускная способность":    "bandwidth",
		"облачное хранилище данных": "cloud storage",
		"компилировать":             "compile",
		"сжимать":                   "compress",
		"учетные данные":            "credentials",
		"база данных":               "database",
		"отлаживать":                "debug",
		"недостаток":                "drawback",
		"извлекать устройство":      "eject",
		"зашифровывать":             "encrypt",
		"папка":                     "folder",
		"макет, разметка":           "layout",
		"поддержка":                 "maintenance",
		"пространство имен":         "namespace",
		"вести переговоры":          "negotiate",
		"осуществлять аутсорсинг":   "outsource",
		"пароль":                    "password",
		"подключать":                "plug in",
		"запрещать":                 "prohibit",
		"заменить":                  "replace",
		"удалять":                   "uninstall",
		"загрузить":                 "upload",
		"проверка":                  "validation",
		"переменная":                "variable",
		"поставщик":                 "vendor",
	}
	// Инициализируем GTK.
	gtk.Init(nil)

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("wdgts.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("window_main")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Отображаем все виджеты в окне
	win.ShowAll()

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	go func() {
		for k, v := range from_e_to_r {
			from_e_to_r_s.Store(k, v)
		}
	}()
	go func() {
		for k, v := range from_r_to_e {
			from_r_to_e_s.Store(k, v)
		}
	}()

	obj, _ = b.GetObject("entry_1")
	entry1 := obj.(*gtk.Entry)

	obj, _ = b.GetObject("entry_2")
	entry2 := obj.(*gtk.Entry)

	obj, _ = b.GetObject("entry_3")
	entry3 := obj.(*gtk.Entry)

	obj, _ = b.GetObject("entry_4")
	entry4 := obj.(*gtk.Entry)

	// Получаем кнопку
	obj, _ = b.GetObject("button_1")
	button1 := obj.(*gtk.Button)

	obj, _ = b.GetObject("button_2")
	button2 := obj.(*gtk.Button)

	obj, _ = b.GetObject("button_3")
	button3 := obj.(*gtk.Button)

	obj, _ = b.GetObject("label_3")
	label1 := obj.(*gtk.Label)

	// После нажатия на кнопку меняем картинку с помощью получения псевдорандомного числа
	button1.Connect("clicked", func() {
		text, err := entry4.GetText()
		if err == nil {
			if val, ok := from_e_to_r_s.Load(text); ok {
				// Устанавливаем текст из поля ввода метке
				str := fmt.Sprintf("%v", val)
				label1.SetText(str)
			} else {
				label1.SetText("в словаре нет такого слова, вы можете его добавить")
			}

		}
	})

	button2.Connect("clicked", func() {
		text, err := entry1.GetText()
		if err == nil {
			if val, ok := from_r_to_e_s.Load(text); ok {
				// Устанавливаем текст из поля ввода метке
				str := fmt.Sprintf("%v", val)
				label1.SetText(str)
			} else {
				label1.SetText("в словаре нет такого слова, вы можете его добавить")
			}

		}
	})

	button3.Connect("clicked", func() {
		text, err := entry2.GetText()
		text2, err2 := entry3.GetText()
		if err == nil && err2 == nil {
			from_e_to_r_s.Store(text, text2)
			from_r_to_e_s.Store(text2, text)
		}
	})

	gtk.Main()
}
