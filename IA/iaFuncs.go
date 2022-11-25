package IA

import (
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	resized "golang.org/x/image/draw"
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func StartIA() []float64 {
	r := [][]float64{}
	r = OpenImage()
	var sd = []float64{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}
	var neuronas = 80
	var iteraciones = 100
	var iteracion = 0
	var iteracion_r = 0
	var iteracion_sd = 0
	var listaP = []float64{}
	//var errores = [][]float64{}
	var weights = genWeights(neuronas-1, neuronas)
	var iter = 0
	errores := make([][]opts.LineData, 10)
	for iter != iteraciones {
		var listaInicial = capaInicial(r[iteracion_r], weights, neuronas-1)
		weights = modifyWeight(listaInicial, weights, r[iteracion_r], sd[iteracion_sd], true)
		var listaFinal = capaFinal(listaInicial, weights)
		weights = modifyWeight(listaFinal, weights, listaInicial, sd[iteracion_sd], false)
		var listaNew = append(listaInicial[1:], listaFinal[1:]...)
		listaNew = modificarSlice(listaNew)
		weights = modificarMatriz(weights)
		listaP = append(listaP, listaNew[0])
		for i := 0; i < len(weights); i++ {
			if i == 0 {
				var errorVal = sd[iteracion_sd] - listaNew[i]
				var deltaMinuscula = listaNew[i] * (1 - listaNew[i]) * errorVal
				errores[iteracion_r] = append(errores[iteracion_r], opts.LineData{Value: errorVal})
				weights[i] = backpropagation(weights[i], listaInicial, deltaMinuscula)
			} else {
				var errorVal = sd[iteracion_sd] - listaNew[i]
				var deltaMinuscula = listaNew[i] * (1 - listaNew[i]) * errorVal
				listaInicial = r[iteracion_r]
				var deltaMayuscula = listaNew[i] * (1 - listaNew[i]) * deltaMinuscula
				weights[i] = backpropagation(weights[i], listaInicial, deltaMayuscula)
			}
		}
		iter += 1
		iteracion_r += 1
		iteracion += 1
		iteracion_sd += 1
		weights = insertToMatrix(weights, weights[0], len(weights))
		weights = weights[1:]
		if iteracion_r == 10 {
			iteracion_r = 0
			iteracion_sd = 0
		}
	}
	listValues := []float64{listaP[len(listaP)-10], listaP[len(listaP)-9], listaP[len(listaP)-8],
		listaP[len(listaP)-7], listaP[len(listaP)-6], listaP[len(listaP)-5], listaP[len(listaP)-4],
		listaP[len(listaP)-3], listaP[len(listaP)-2], listaP[len(listaP)-1],
	}
	genLines(errores)
	return listValues
}

func genLines(errorSlice [][]opts.LineData) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.InitOpts{Theme: types.ThemeWesteros},
		charts.TitleOpts{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		})

	// Put data into instance
	line.AddXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"})
	colors := []string{"yellow", "blue", "red", "black",
		"orange", "green", "purple", "pink",
		"grey", "brown"}
	for idx, list := range errorSlice {
		line.AddYAxis("Category "+colors[idx], list)
		line.SetSeriesOptions(charts.LineStyleOpts{Color: colors[idx]})
	}
	f, _ := os.Create("linea.html")
	_ = line.Render(f)
}

func capaInicial(r []float64, weights [][]float64, neuronas int) []float64 {
	var lista = []float64{1}
	for i := 0; i < neuronas; i++ {
		lista = append(lista, perceptron(r, weights[i]))
	}
	return lista
}

func capaFinal(lista []float64, weights [][]float64) []float64 {
	var lista1 = []float64{1}
	for i := len(weights) - 1; i < len(weights); i++ {
		lista1 = append(lista1, perceptron(lista, weights[i]))
	}
	return lista1
}

func perceptron(listaE []float64, weights []float64) float64 {
	fmt.Println(len(listaE), "len listaE")
	fmt.Println(len(weights), "len weights")
	var i = len(listaE) - 1
	var z = 0
	var sum float64 = 0
	for j := i - (len(listaE) - 1); j < i+1; j++ {
		sum += listaE[j] * weights[z]
		z += 1
	}
	var y = 1 / (1 + (math.Pow(math.E, -sum)))
	return y
}

func modifyWeight(lista []float64, weights [][]float64, r []float64, sd float64, booleano bool) [][]float64 {
	if booleano == true {
		for i := 1; i < len(lista); i++ {
			weights[i-1] = calculateWeights(lista[i], weights[i-1], r, sd)
		}
	} else if booleano == false {
		weights[len(weights)-1] = calculateWeights(lista[1], weights[len(weights)-1], r, sd)
	}
	return weights
}

func calculateWeights(num float64, weights []float64, listaE []float64, sd float64) []float64 {
	var lR = 0.6
	var calcError = sd - num
	var delta = num * (1 - num) * calcError
	for val := 0; val < len(weights); val++ {
		weights[val] = weights[val] + lR*listaE[val]*delta
	}
	return weights
}

func modificarSlice(lista []float64) []float64 {
	lista = insertToSlice(lista, lista[len(lista)-1], 0)
	lista = lista[:len(lista)-1]
	return lista
}

func modificarMatriz(lista [][]float64) [][]float64 {
	lista = insertToMatrix(lista, lista[len(lista)-1], 0)
	lista = lista[:len(lista)-1]
	return lista
}

func insertToSlice(a []float64, element float64, i int) []float64 {
	a = append(a[:i+1], a[i:]...)
	a[i] = float64(element)
	return a
}

func insertToMatrix(a [][]float64, element []float64, i int) [][]float64 {
	a = append(a[:i+1], a[i:]...)
	a[i] = element
	return a
}

func backpropagation(weights []float64, ent []float64, val float64) []float64 {
	var num = 0.6
	for i := 0; i < len(weights); i++ {
		weights[i] = weights[i] + num*ent[i]*val
	}
	return weights
}

func genWeights(ocultas int, neuronas int) [][]float64 {
	var pesos2 = make([][]float64, neuronas)
	var j = 0
	var k = 0
	var rango = ocultas*8001 + neuronas
	var z = 1
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < rango; i++ {
		if z == 1 {
			pesos2[j] = append(pesos2[j], r1.Float64()*-1.0)
			z *= -1
		} else {
			pesos2[j] = append(pesos2[j], r1.Float64()*1.0)
			z *= -1
		}
		k += 1
		if k == 8001 {
			j += 1
			k = 0
		}
		if j == len(pesos2) {
			j -= 1
		}
	}
	return pesos2
}

func lenImages(directory string) int {
	d, e := os.ReadDir(directory)
	if e != nil {
		panic(e)
	}
	return len(d)
}

func OpenImage() [][]float64 {
	numImagenes := lenImages("./Imagenes/")
	fmt.Println(numImagenes)
	slicePixels := make([][]float64, numImagenes)
	for i := 0; i < numImagenes; i++ {
		file, err := os.Open("./Imagenes/imgNew" + strconv.Itoa(i) + ".jpeg")
		if err != nil {
			fmt.Println("Error: File could not be opened")
			os.Exit(1)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
			}
		}(file)

		img := decodeImagen(file)
		imgResized := resizeImage(img)
		imgGray := convertToGray(imgResized)
		pixels, err := getPixels(imgGray)
		if err != nil {
			fmt.Println("Error: Image could not be decoded")
			fmt.Println(pixels)
			os.Exit(1)
		}
		slicePixels[i] = append(slicePixels[i], 1)
		for _, val := range pixels {
			slicePixels[i] = append(slicePixels[i], float64(val.R))
		}
		fmt.Println(len(slicePixels[0]), "len de pixels")
	}
	return slicePixels
}

func resizeImage(img image.Image) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X/15, img.Bounds().Max.Y/16))
	resized.CatmullRom.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)
	return dst
}

func decodeImagen(file io.Reader) image.Image {
	img, _ := jpeg.Decode(file)
	return img
}

func convertToGray(img image.Image) *image.Gray {
	result := image.NewGray(img.Bounds())
	draw.Draw(result, result.Bounds(), img, img.Bounds().Min, draw.Src)
	return result
}

func getPixels(img image.Image) ([]Pixel, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixel []Pixel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel = append(pixel, rgbaToPixel(img.At(x, y).RGBA()))
		}
	}

	return pixel, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	var pix = Pixel{int(r / 255)}
	return pix
}

type Pixel struct {
	R int
}
