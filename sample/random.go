package sample

import (
	"hello-grpc/pb"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY

	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
func randomCPUBrand() string {
	return randStringFromSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randStringFromSet("Nvidia", "AMD")
}
func randStringFromSet(set ...string) string {
	return set[rand.Intn(len(set))]
}

func RandomCPUName(name string) string {
	switch name {
	case "Intel":
		return randStringFromSet("Core i3", "Core i5", "core i7")
	default:
		return randStringFromSet("Ryzen 7 pro", "Ryzen 5 pro", "Ryzen 3 pro")
	}
}
func RandomGPUName(name string) string {
	switch name {
	case "Nvidia":
		return randStringFromSet("RTX 2060", "RTX 2070", "RTX 2080")
	default:
		return randStringFromSet("RX 590", "RX 580", "RX 570")
	}
}
func randomThreads(cores int) int {
	return randomInteger(cores, 12)
}

func randomCores() int {
	return randomInteger(2, 8)
}
func randomInteger(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}
func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInteger(1080, 4320)
	width := height * 16 / 9
	return &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}

}
func randomLaptopBrand() string {
	return randStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randStringFromSet("Inspiron", "Vostro")
	default:
		return randStringFromSet("Thinkpad", "IdeaPad")
	}
}

func randomID() string {
	return uuid.New().String()
}
