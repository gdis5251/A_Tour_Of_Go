package main
import (
    "image"
    "image/color"
    "image/png"
    "log"
    "math"
    "os"
)

func main() {
	// 设置图片大小
	const size = 400

	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))

	// 遍历图片的每个像素点，将其设置为白色
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	// 在图片上绘制正弦函数
	for x := 0; x < size; x++ {
		// 让 sin 的值在 0 ~ 2PI 之间
		s := float64(x) * 2 * math.Pi / size

		// 对 sin 的幅度为一半的像素，向下偏移一半的像素并翻转
		y := size / 2 - math.Sin(s) * size / 2

		// 绘制图形
		pic.SetGray(x, int(y), color.Gray{0})
	}

	// 创建文件
	file, err := os.Create("./sin.png")

	if err != nil {
		log.Fatal(err)
	}

	// 使用 PNG 格式将数据写入文件
	png.Encode(file, pic)

	file.Close()
}