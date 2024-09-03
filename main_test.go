package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestAddImage(t *testing.T) {
	baseImage := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

	err := addImage(baseImage, "image_example.png", image.Point{0, 0})
	if err != nil {
		t.Fatalf("addImage failed: %v", err)
	}

	tempFile, err := os.Create("temp_image_output.png")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer tempFile.Close()
	defer os.Remove("temp_image_output.png")

	if err := png.Encode(tempFile, baseImage); err != nil {
		t.Fatalf("Error encoding image: %v", err)
	}

	goldenFile, err := os.Open("testdata/image_only.golden")
	if err != nil {
		t.Fatalf("Error opening golden file: %v", err)
	}
	defer goldenFile.Close()

	goldenImage, err := png.Decode(goldenFile)
	if err != nil {
		t.Fatalf("Error decoding golden image: %v", err)
	}

	tempFile, err = os.Open("temp_image_output.png")
	if err != nil {
		t.Fatalf("Error opening temporary file: %v", err)
	}
	defer tempFile.Close()

	tempImage, err := png.Decode(tempFile)
	if err != nil {
		t.Fatalf("Error decoding temporary image: %v", err)
	}

	if !compareImages(goldenImage, tempImage) {
		t.Fatalf("Generated image does not match golden file")
	}
}

func TestAddText(t *testing.T) {
	baseImage := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

	err := addText(baseImage, "Hello, I am a dog!", image.Point{325, 950}, color.Black, 50)
	if err != nil {
		t.Fatalf("addText failed: %v", err)
	}

	tempFile, err := os.Create("temp_text_output.png")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer tempFile.Close()
	defer os.Remove("temp_text_output.png")

	if err := png.Encode(tempFile, baseImage); err != nil {
		t.Fatalf("Error encoding image: %v", err)
	}

	goldenFile, err := os.Open("testdata/text_only.golden")
	if err != nil {
		t.Fatalf("Error opening golden file: %v", err)
	}
	defer goldenFile.Close()

	goldenImage, err := png.Decode(goldenFile)
	if err != nil {
		t.Fatalf("Error decoding golden image: %v", err)
	}

	tempFile, err = os.Open("temp_text_output.png")
	if err != nil {
		t.Fatalf("Error opening temporary file: %v", err)
	}
	defer tempFile.Close()

	tempImage, err := png.Decode(tempFile)
	if err != nil {
		t.Fatalf("Error decoding temporary image: %v", err)
	}

	if !compareImages(goldenImage, tempImage) {
		t.Fatalf("Generated image does not match golden file")
	}
}

func TestMainFunction(t *testing.T) {
	main()

	generatedFile, err := os.Open("output.png")
	if err != nil {
		t.Fatalf("Error opening generated file: %v", err)
	}
	defer generatedFile.Close()
	defer os.Remove("output.png")

	generatedImage, err := png.Decode(generatedFile)
	if err != nil {
		t.Fatalf("Error decoding generated image: %v", err)
	}

	goldenFile, err := os.Open("testdata/image_generated.golden")
	if err != nil {
		t.Fatalf("Error opening golden file: %v", err)
	}
	defer goldenFile.Close()

	goldenImage, err := png.Decode(goldenFile)
	if err != nil {
		t.Fatalf("Error decoding golden image: %v", err)
	}

	if !compareImages(goldenImage, generatedImage) {
		t.Fatalf("Generated image does not match golden file")
	}
}

func compareImages(img1, img2 image.Image) bool {
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()

	if !bounds1.Eq(bounds2) {
		return false
	}

	for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
		for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
			if img1.At(x, y) != img2.At(x, y) {
				return false
			}
		}
	}

	return true
}
