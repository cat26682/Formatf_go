package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/disintegration/imaging"
)

//форматы

var formatfoto = []string{"jpg", "jpeg", "png", "gif", "bmp", "tiff", "webp", "heif", "heic", "raw", "dng", "psd", "avif", "ico", "cur", "pcx", "ppm", "pgm", "pbm", "pgf", "xcf", "svg"}
var th = "."

func main() {

	// Генерируем случайное число от 0 до 99
	randomInt := rand.Intn(99999999)
	randomIntStr := fmt.Sprintf("%d", randomInt)
	fmt.Println(randomIntStr)

	//проверяем формат файла
	var q = []string{"png"}
	// Проверяем, есть ли общие элементы
	fmt.Println(hasCommonElement(q, formatfoto)) // Вывод: true

	// Открываем изображение
	img, err := imaging.Open("ф.png")
	if err != nil {
		log.Fatalf("Ошибка при открытии изображения: %v", err)
	}
	fmt.Println("введите расширение файла")
	var ext string
	fmt.Scan(&ext)
	// Сохраняем изображение
	err = imaging.Save(img, "conwert"+randomIntStr+th+ext)
	if err != nil {
		log.Fatalf("Ошибка при сохранении изображения: %v", err)
	} else {
		os.Remove("inp.jpg")

	}
}

// Функция для проверки совпадений
func hasCommonElement(slice1, slice2 []string) bool {
	for _, item1 := range slice1 {
		for _, item2 := range slice2 {
			if item1 == item2 {
				return true
			}
		}
	}
	return false
}
