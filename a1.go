package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

//форматы

var formatfoto = []string{"jpg", "jpeg", "png", "gif", "bmp", "tiff", "webp", "heif", "heic", "raw", "dng", "psd", "avif", "ico", "cur", "pcx", "ppm", "pgm", "pbm", "pgf", "xcf", "svg"}
var th = "."

func formatF() {

	// Генерируем случайное число от 0 до 99
	randomInt := rand.Intn(99999999)
	randomIntStr := fmt.Sprintf("%d", randomInt)
	//fmt.Println(randomIntStr)

	fmt.Println("введите имя файла")

	var ext0 string
	fmt.Scan(&ext0)

	strings.Split(ext0, ".")
	if len(ext0) < 2 {
		log.Panic("ошибка")
	}
	//проверяем формат файла
	var q = []string{string(ext0[1])}
	// Проверяем, есть ли общие элементы
	var q1 = (hasCommonElement(q, formatfoto)) // Вывод: true
	if q1 == false {
		log.Panic("неверный формат файла")
	}

	// Открываем изображение
	img, err := imaging.Open(ext0)
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
	}
	os.Remove(ext0)
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
func ImF() {

}
func VVF() {

}
func ic() {

}
