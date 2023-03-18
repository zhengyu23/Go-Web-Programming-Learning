package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type DB struct {
	mutex *sync.Mutex
	store map[string][3]float64
}

var directory = "chapter9_GoConcurrency/9.4_Mosaic/"
var pictureDirecotry = "chapter9_GoConcurrency/9.4_Mosaic/"

func main() {
	//cpuNum := runtime.NumCPU()
	//fmt.Println("cpuNum=", cpuNum)
	//runtime.GOMAXPROCS(1)
	//fmt.Println("cpuNum=", runtime.GOMAXPROCS(1))
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(directory + "public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(directory + "upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	// ① 获取用户上传的目标图片，以及瓷砖图片的尺寸
	// ② 对用户上传的目标图片进行解码
	// ③ 腐植瓷砖图数据库
	// ④ 为每张瓷砖图片设置起始点
	// ⑤ 对目标图片分割出的每张子图进行迭代
	// ⑥ 将图片编码为JPEG格式看然后同故宫base64字符串将其传输至浏览器

	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image") // 从web输入
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	original, _, _ := image.Decode(file)
	bounds := original.Bounds()

	// 原图片副本，在副本上进行mosaic操作
	newImage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))

	db := cloneTilesDB()

	sp := image.Point{} // source point
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}

			nearest := nearest(color, &db)
			file, err := os.Open(nearest)
			if err == nil {
				tileImg, _, err := image.Decode(file)
				if err == nil {
					t := resize(tileImg, tileSize) // 将tilePicture缩小到指定宽度

					tile := t.SubImage(t.Bounds())
					// SubImage returns an image representing the portion of the image p visible through r
					// 自己看自己?

					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
					draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("error :", err, nearest)
				}
			} else {
				fmt.Println("error: ", nearest)
			}
			file.Close()
		}
	}

	// 解码成base64
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newImage, nil)

	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())
	t1 := time.Now()
	image := map[string]string{
		"original": originalStr,
		"mosaic":   mosaic,
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles(directory + "results.html")
	t.Execute(w, image)

}

var TILESDB map[string][3]float64

// averageColor 计算一张图的平均颜色
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{
		r / totalPixels,
		g / totalPixels,
		b / totalPixels,
	}
}

// resize 缩放图片至指定宽度
func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWidth // ratio 比例
	// 等比例缩放
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.X/ratio,
		bounds.Max.X/ratio, bounds.Max.Y/ratio))
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(a >> 8)})
		}
	}
	return *out
}

// tilesDB 扫描瓷砖图片所在目录，创建瓷砖图片数据库
func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db...")
	db := make(map[string][3]float64)
	//files, _ := os.ReadDir(directory + "tiles_test")
	files, _ := os.ReadDir(pictureDirecotry + "tiles")
	for _, f := range files {
		//name := directory + "tiles_test/" + f.Name()
		name := pictureDirecotry + "tiles/" + f.Name()
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB:", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	return db
}

// neatest 获取与目标图片平均颜色最接近的瓷砖图片
func nearest(target [3]float64, db *map[string][3]float64) string {
	var filename string
	smallest := 1000000.0
	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	//delete(*db, filename)
	return filename
}

// distance 计算两个三元组之间的欧几里得距离
func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}
func sq(n float64) float64 {
	return n * n
}

// cloneTilesDB 复制自传图片数据库副本
func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}
