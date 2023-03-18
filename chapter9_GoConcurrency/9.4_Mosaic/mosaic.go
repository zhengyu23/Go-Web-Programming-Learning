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
	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image") // 从web输入
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	db := cloneTilesDB()
	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X/2, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)
	c, t1 := combine(bounds, c1, c2, c3, c4) // 以扇形方式聚拢多个子图，并合成一个完整的图片

	// 解码成base64
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	image := map[string]string{
		"original": originalStr,
		"mosaic":   <-c,
		"duration": fmt.Sprintf("%v ", (<-t1).Sub(t0)),
	}
	t, _ := template.ParseFiles(directory + "results.html")
	t.Execute(w, image)
}

// cut 分割图片
func cut(original image.Image, db *DB, tileSize, x1, y1, x2, y2 int) <-chan image.Image {
	c := make(chan image.Image)
	sp := image.Point{}
	go func() {
		newImage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))
		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				r, g, b, _ := original.At(x, y).RGBA()
				color := [3]float64{float64(r), float64(g), float64(b)}

				nearest := db.nearest(color)
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
		c <- newImage.SubImage(newImage.Rect)
	}()
	return c
}

// combine 扇形聚拢
func combine(r image.Rectangle, c1, c2, c3, c4 <-chan image.Image) (<-chan string, <-chan time.Time) {
	c := make(chan string) // 这个函数将返回一个通道作为执行结果
	t1 := make(chan time.Time)

	// 创建一个匿名函数goroutine
	go func() {
		var wg sync.WaitGroup // 使用等待组去同步各个字图片的复制操作
		img := image.NewNRGBA(r)
		copy := func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
			draw.Draw(dst, r, src, sp, draw.Src)
			wg.Done() // 每复制完一张字图片,就对计数器执行一次减一操作
		}
		wg.Add(4) // 把等待组计数器的值设置为4
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool
		for { // 无限循环等待复制操作完成
			select { // 等待各个通道的返回值
			case s1, ok1 = <-c1:
				go copy(img, s1.Bounds(), s1, image.Point{X: r.Min.X, Y: r.Min.Y})
			case s2, ok2 = <-c2:
				go copy(img, s2.Bounds(), s2, image.Point{X: r.Max.X / 2, Y: r.Min.Y})
			case s3, ok3 = <-c3:
				go copy(img, s3.Bounds(), s3, image.Point{X: r.Min.X, Y: r.Max.Y / 2})
			case s4, ok4 = <-c4:
				go copy(img, s4.Bounds(), s4, image.Point{X: r.Max.X / 2, Y: r.Max.Y / 2})
			}
			if ok1 && ok2 && ok3 && ok4 {
				break
			}
		}
		wg.Wait()
		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2, img, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
		t1 <- time.Now()
	}()
	return c, t1
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
	files, _ := os.ReadDir(directory + "tiles")
	for _, f := range files {
		//name := directory + "tiles_test/" + f.Name()
		name := directory + "tiles/" + f.Name()
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
func (db *DB) nearest(target [3]float64) string {
	var filename string

	db.mutex.Lock()
	smallest := 1000000.0
	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	//delete(db.store, filename)
	db.mutex.Unlock()
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
func cloneTilesDB() DB {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	tiles := DB{
		store: db,
		mutex: &sync.Mutex{},
	}
	return tiles
}
