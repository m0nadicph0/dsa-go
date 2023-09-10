package floodfill

type Image [][]int

func FloodFill(image Image, startRow int, startCol int, color int) Image {
	if len(image) == 0 || image[startRow][startCol] == color {
		return image
	}

	originalColor := image[startRow][startCol]

	floodFillHelper(image, startRow, startCol, originalColor, color)

	return image
}

func floodFillHelper(image Image, row int, col int, originalColor int, newColor int) {
	if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) || image[row][col] != originalColor {
		return
	}

	image[row][col] = newColor

	floodFillHelper(image, row+1, col, originalColor, newColor)
	floodFillHelper(image, row-1, col, originalColor, newColor)
	floodFillHelper(image, row, col+1, originalColor, newColor)
	floodFillHelper(image, row, col-1, originalColor, newColor)
}
