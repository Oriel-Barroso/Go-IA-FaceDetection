package Imagen

import (
	"bytes"
	"fmt"
	iaFuncs "github.com/Oriel-Barroso/golangBackend/IA"
	"github.com/gofiber/fiber/v2"
	"image/jpeg"
	"os"
	"strconv"
)

func PostImage(c *fiber.Ctx) error {
	req, _ := c.Request().MultipartForm()
	for _, espacioMemoria := range req.File {
		for idx, valor := range espacioMemoria {
			file, _ := valor.Open()
			imagen, _ := jpeg.Decode(file) //Convertimos el request en imagen
			buf := new(bytes.Buffer)
			_ = jpeg.Encode(buf, imagen, &jpeg.Options{100}) //Convertimos la imagen en io.Writer para poder ser convertida en bytesImg, se guardan en buf
			bytesImg := buf.Bytes()                          //Escribimos los bytesImg
			err := os.WriteFile("./Imagenes/imgNew"+strconv.Itoa(idx)+".jpeg", bytesImg, 0644)
			if err != nil {
				return err
			}
		}
	}
	vals := iaFuncs.StartIA()
	fmt.Println(vals)
	return nil
}
