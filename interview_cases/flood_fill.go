package interview_cases

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	start := image[sr][sc]
	image[sr][sc] = color
	// down
	if len(image) > sr+1 && image[sr+1][sc] == start {
		floodFill(image, sr+1, sc, color)
	}
	// up
	if sr-1 >= 0 && image[sr-1][sc] == start {
		floodFill(image, sr-1, sc, color)
	}
	// left
	if sc-1 >= 0 && image[sr][sc-1] == start {
		floodFill(image, sr, sc-1, color)
	}
	// right
	if len(image[sr]) > sc+1 && image[sr][sc+1] == start {
		floodFill(image, sr, sc+1, color)
	}
	return image
}
