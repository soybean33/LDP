package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	var epsilon float64 = 8
	var repeat int = 10
	aggregationOriginal := make(map[int][]float64, 0)
	aggregationDE := make(map[int][]float64, 0)
	aggregationSHE := make(map[int][]float64, 0)
	aggregationSUE := make(map[int][]float64, 0)
	aggregationOUE := make(map[int][]float64, 0)
	MSEDE := make(map[int]float64, 0)
	MSESHE := make(map[int]float64, 0)
	MSESUE := make(map[int]float64, 0)
	MSEOUE := make(map[int]float64, 0)
	AVGMSEDE := make(map[int]float64, 0)
	AVGMSESHE := make(map[int]float64, 0)
	AVGMSESUE := make(map[int]float64, 0)
	AVGMSEOUE := make(map[int]float64, 0)
	// summary
	var min, max, avg float64
	var minMovieId, maxMovieId int

	fmt.Println("main start")
	fmt.Println("Epsilon is ", epsilon)
	// Original
	aggregationOriginal = original()

	var sel int = 2

	switch sel {
	case 1:
		// DE
		for i := 0; i < repeat; i++ {
			fmt.Print(i+1, " > ")
			aggregationDE = directEncoding(epsilon)
			for key, _ := range aggregationOriginal {
				var tempSUM float64 = 0
				for j := 0; j < len(aggregationOriginal[key]); j++ {
					tempSUM = tempSUM + math.Pow((aggregationDE[key][j]-aggregationOriginal[key][j]), 2)
				}
				MSEDE[key] = tempSUM / float64(len(aggregationOriginal[key]))
				AVGMSEDE[key] = AVGMSEDE[key] + MSEDE[key]
			}
		}
		for key, _ := range aggregationOriginal {
			AVGMSEDE[key] = AVGMSEDE[key] / float64(repeat)
		}

		for key, _ := range AVGMSEDE {
			if key == 1 {
				min = AVGMSEDE[key]
				max = AVGMSEDE[key]
				minMovieId = key
				maxMovieId = key
			}
			if AVGMSEDE[key] > max {
				maxMovieId = key
				max = AVGMSEDE[key]
			}
			if AVGMSEDE[key] < min {
				minMovieId = key
				min = AVGMSEDE[key]
			}
			avg = avg + AVGMSEDE[key]
		}
		fmt.Println("")
		fmt.Println("avg MSE of DE is", avg/float64(len(aggregationOriginal)))
		fmt.Println("max MSE of DE is ", max, " movie ID is ", maxMovieId, aggregationOriginal[maxMovieId], aggregationDE[maxMovieId])
		fmt.Println("min MSE of DE is ", min, " movie ID is ", minMovieId, aggregationOriginal[minMovieId], aggregationDE[minMovieId])

	case 2:
		// SHE
		for i := 0; i < repeat; i++ {
			fmt.Print(i+1, " > ")
			aggregationSHE = summationHistogramEncoding(epsilon)
			for key, _ := range aggregationOriginal {
				var tempSUM float64 = 0
				for j := 0; j < len(aggregationOriginal[key]); j++ {
					tempSUM = tempSUM + math.Pow((aggregationSHE[key][j]-aggregationOriginal[key][j]), 2)
				}
				MSESHE[key] = tempSUM / float64(len(aggregationOriginal[key]))
				AVGMSESHE[key] = AVGMSESHE[key] + MSESHE[key]
			}
			for key, _ := range aggregationOriginal {
				AVGMSESHE[key] = AVGMSESHE[key] / float64(repeat)
			}
		}

		for key, _ := range AVGMSESHE {
			if key == 1 {
				min = AVGMSESHE[key]
				max = AVGMSESHE[key]
				minMovieId = key
				maxMovieId = key
			}
			if AVGMSESHE[key] > max {
				maxMovieId = key
				max = AVGMSESHE[key]
			}
			if AVGMSESHE[key] < min {
				minMovieId = key
				min = AVGMSESHE[key]
			}
			avg = avg + AVGMSESHE[key]
		}
		fmt.Println("")
		fmt.Println("avg MSE of SHE is", avg/float64(len(aggregationOriginal)))
		fmt.Println("max MSE of SHE is ", max, " movie ID is ", maxMovieId, aggregationOriginal[maxMovieId], aggregationSHE[maxMovieId])
		fmt.Println("min MSE of SHE is ", min, " movie ID is ", minMovieId, aggregationOriginal[minMovieId], aggregationSHE[minMovieId])

	case 3:
		// SUE
		for i := 0; i < repeat; i++ {
			fmt.Print(i+1, " > ")
			aggregationSUE = symmetricUnaryEncoding(epsilon)
			for key, _ := range aggregationOriginal {
				var tempSUM float64 = 0
				for j := 0; j < len(aggregationOriginal[key]); j++ {
					tempSUM = tempSUM + math.Pow(aggregationSUE[key][j]-aggregationOriginal[key][j], 2)
				}
				MSESUE[key] = tempSUM / float64(len(aggregationOriginal[key]))
				AVGMSESUE[key] = AVGMSESUE[key] + MSESUE[key]
			}
		}
		for key, _ := range aggregationOriginal {
			AVGMSESUE[key] = AVGMSESUE[key] / float64(repeat)
		}

		for key, _ := range AVGMSESUE {
			if key == 1 {
				min = AVGMSESUE[key]
				max = AVGMSESUE[key]
				minMovieId = key
				maxMovieId = key
			}
			if AVGMSESUE[key] > max {
				maxMovieId = key
				max = AVGMSESUE[key]
			}
			if AVGMSESUE[key] < min {
				minMovieId = key
				min = AVGMSESUE[key]
			}
			avg = avg + AVGMSESUE[key]
		}
		fmt.Println("")
		fmt.Println("avg MSE of SUE is", avg/float64(len(aggregationOriginal)))
		fmt.Println("max MSE of SUE is ", max, " movie ID is ", maxMovieId, aggregationOriginal[maxMovieId], aggregationSUE[maxMovieId])
		fmt.Println("min MSE of SUE is ", min, " movie ID is ", minMovieId, aggregationOriginal[minMovieId], aggregationSUE[minMovieId])

	case 4:
		// OUE
		for i := 0; i < repeat; i++ {
			fmt.Print(i+1, " > ")
			aggregationOUE = optimizedUnaryEncoding(epsilon)
			for key, _ := range aggregationOriginal {
				var tempSUM float64 = 0
				for j := 0; j < len(aggregationOriginal[key]); j++ {
					tempSUM = tempSUM + math.Pow(aggregationOUE[key][j]-aggregationOriginal[key][j], 2)
				}
				MSEOUE[key] = tempSUM / float64(len(aggregationOriginal[key]))
				AVGMSEOUE[key] = AVGMSEOUE[key] + MSEOUE[key]
			}
		}
		for key, _ := range aggregationOriginal {
			AVGMSEOUE[key] = AVGMSEOUE[key] / float64(repeat)
		}

		for key, _ := range AVGMSEOUE {
			if key == 1 {
				min = AVGMSEOUE[key]
				max = AVGMSEOUE[key]
				minMovieId = key
				maxMovieId = key
			}
			if AVGMSEOUE[key] > max {
				maxMovieId = key
				max = AVGMSEOUE[key]
			}
			if AVGMSEOUE[key] < min {
				minMovieId = key
				min = AVGMSEOUE[key]
			}
			avg = avg + AVGMSEOUE[key]
		}
		fmt.Println("")
		fmt.Println("avg MSE of OUE is", avg/float64(len(aggregationOriginal)))
		fmt.Println("max MSE of OUE is ", max, " movie ID is ", maxMovieId, aggregationOriginal[maxMovieId], aggregationOUE[maxMovieId])
		fmt.Println("min MSE of OUE is ", min, " movie ID is ", minMovieId, aggregationOriginal[minMovieId], aggregationOUE[minMovieId])
	}
	fmt.Println("main end")
}

func original() map[int][]float64 {
	var userId, movieId, timestamp int
	var rating float64

	fmt.Println("Original start")

	fileRatings, _ := os.Open("ratings.csv")
	defer fileRatings.Close()

	// Aggregation
	aggregation := make(map[int][]float64, 0)
	numberOfMember := make(map[int]int, 0)
	tempRating := make([]float64, 11)
	for {
		_, err := fmt.Fscanf(fileRatings, "%d,%d,%f,%d\n", &userId, &movieId, &rating, &timestamp)
		if err == io.EOF {
			break
		}
		for i := 0; i < len(tempRating); i++ {
			if i == int((rating * 2)) {
				tempRating[i] = 1
			} else {
				tempRating[i] = 0
			}
		}
		if aggregation[movieId] == nil {
			aggregation[movieId] = append(aggregation[movieId], tempRating...)
		} else {
			for i := 0; i < len(tempRating); i++ {
				if tempRating[i] != 0 {
					aggregation[movieId][i]++
				}
			}
		}
		numberOfMember[movieId]++
	}

	fmt.Println("original end")
	return aggregation
}

func optimizedUnaryEncoding(epsilon float64) map[int][]float64 { //OUE
	var userId, movieId, timestamp int
	var rating float64

	fmt.Println("OUE start")
	fileRatings, _ := os.Open("ratings.csv")
	fileoptimizedUnaryEncoding, _ := os.Create("OptimizedUnaryEncoding.csv")
	defer fileRatings.Close()
	defer fileoptimizedUnaryEncoding.Close()

	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)

	for {

		_, err := fmt.Fscanf(fileRatings, "%d,%d,%f,%d\n", &userId, &movieId, &rating, &timestamp)
		if err == io.EOF {
			break
		}
		//Encoding
		ratingEncoding := make([]int, 11)
		for i := 0; i < len(ratingEncoding); i++ {
			if i == int(rating*2) {
				ratingEncoding[i] = 1
			} else {
				ratingEncoding[i] = 0
			}
		}
		//perturbing
		ratingPerturbing := make([]int, 11)
		for i := 0; i < len(ratingEncoding); i++ {
			var probability float64 = random.Float64()

			if ratingEncoding[i] == 1 {
				if probability < 1/2 {
					ratingPerturbing[i] = 1
				}
			}
			if ratingEncoding[i] == 0 {
				if probability < 1/(math.Exp(epsilon/2)+1) {
					ratingPerturbing[i] = 1
				}
			}
		}
		fmt.Fprint(fileoptimizedUnaryEncoding, userId, movieId, ratingPerturbing, "\n")
	}
	//aggregation
	optimizedUnaryEncoding, _ := os.Open("optimizedUnaryEncoding.csv")
	defer optimizedUnaryEncoding.Close()
	aggregation := make(map[int][]float64, 0)
	numberOfMember := make(map[int]int, 0)
	tempRating := make([]float64, 11)

	for {
		_, err := fmt.Fscanf(optimizedUnaryEncoding, "%d %d [%f %f %f %f %f %f %f %f %f %f %f]\n", &userId, &movieId, &tempRating[0], &tempRating[1], &tempRating[2], &tempRating[3], &tempRating[4], &tempRating[5], &tempRating[6], &tempRating[7], &tempRating[8], &tempRating[9], &tempRating[10])
		if err == io.EOF {
			break
		}
		if aggregation[movieId] == nil {
			aggregation[movieId] = append(aggregation[movieId], tempRating...)
		} else {
			for i := 0; i < len(tempRating); i++ {
				if tempRating[i] != 0 {
					aggregation[movieId][i]++
				}
			}
		}
		numberOfMember[movieId]++
	}

	for key, _ := range aggregation {
		for i := 0; i < len(tempRating); i++ {
			aggregation[key][i] = (aggregation[key][i] - (float64(numberOfMember[key]) * (1 / (math.Exp(epsilon/2) + 1)))) / ((1 / 2) - (1 / (math.Exp(epsilon/2) + 1)))
		}
	}
	fmt.Println("OUE end")
	return aggregation
}

func symmetricUnaryEncoding(epsilon float64) map[int][]float64 {
	var userId, movieId, timestamp int
	var rating float64

	fmt.Println("SUE start")
	fileRatings, _ := os.Open("ratings.csv")
	fileSymmetricUnaryEncoding, _ := os.Create("SymmetricUnaryEncoding.csv")
	defer fileRatings.Close()
	defer fileSymmetricUnaryEncoding.Close()

	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)

	for {

		_, err := fmt.Fscanf(fileRatings, "%d,%d,%f,%d\n", &userId, &movieId, &rating, &timestamp)
		if err == io.EOF {
			break
		}
		//Encoding
		ratingEncoding := make([]int, 11)
		for i := 0; i < len(ratingEncoding); i++ {
			if i == int(rating*2) {
				ratingEncoding[i] = 1
			} else {
				ratingEncoding[i] = 0
			}
		}
		//perturbing
		ratingPerturbing := make([]int, 11)
		for i := 0; i < len(ratingEncoding); i++ {
			var probability float64 = random.Float64()

			if ratingEncoding[i] == 1 {
				if probability < math.Exp(epsilon/2)/(math.Exp(epsilon/2)+1) {
					ratingPerturbing[i] = 1
				}
			}
			if ratingEncoding[i] == 0 {
				if probability < 1/(math.Exp(epsilon/2)+1) {
					ratingPerturbing[i] = 1
				}
			}
		}
		fmt.Fprint(fileSymmetricUnaryEncoding, userId, movieId, ratingPerturbing, "\n")
	}
	//aggregation
	SymmetricUnaryEncoding, _ := os.Open("SymmetricUnaryEncoding.csv")
	defer SymmetricUnaryEncoding.Close()
	aggregation := make(map[int][]float64, 0)
	numberOfMember := make(map[int]int, 0)
	tempRating := make([]float64, 11)

	for {
		_, err := fmt.Fscanf(SymmetricUnaryEncoding, "%d %d [%f %f %f %f %f %f %f %f %f %f %f]\n", &userId, &movieId, &tempRating[0], &tempRating[1], &tempRating[2], &tempRating[3], &tempRating[4], &tempRating[5], &tempRating[6], &tempRating[7], &tempRating[8], &tempRating[9], &tempRating[10])
		if err == io.EOF {
			break
		}
		if aggregation[movieId] == nil {
			aggregation[movieId] = append(aggregation[movieId], tempRating...)
		} else {
			for i := 0; i < len(tempRating); i++ {
				if tempRating[i] != 0 {
					aggregation[movieId][i]++
				}
			}
		}
		numberOfMember[movieId]++

	}

	for key, _ := range aggregation {
		for i := 0; i < len(tempRating); i++ {
			aggregation[key][i] = (aggregation[key][i] - (float64(numberOfMember[key]) * 1 / (math.Exp(epsilon/2) + 1))) / ((math.Exp(epsilon/2) / (math.Exp(epsilon/2) + 1)) - (1 / (math.Exp(epsilon/2) + 1)))
		}
	}
	fmt.Println("SUE end")
	return aggregation
}

func summationHistogramEncoding(epsilon float64) map[int][]float64 {
	var userId, movieId, timestamp int
	var rating float64

	fmt.Println("SHE start")

	fileRatings, _ := os.Open("ratings.csv")
	fileHistogramEncoding, _ := os.Create("HistogramEncoding.csv")
	defer fileRatings.Close()
	defer fileHistogramEncoding.Close()

	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)

	for {

		_, err := fmt.Fscanf(fileRatings, "%d,%d,%f,%d\n", &userId, &movieId, &rating, &timestamp)
		if err == io.EOF {
			break
		}
		//Encoding
		ratingEncoding := make([]float64, 11)
		for i := 0; i < len(ratingEncoding); i++ {
			if i == int(rating*2) {
				ratingEncoding[i] = 1.0
			} else {
				ratingEncoding[i] = 0
			}
		}

		//perturbing

		for i := 0; i < len(ratingEncoding); i++ {
			var probability float64 = random.Float64()
			if probability < 0.5 {
				ratingEncoding[i] = ratingEncoding[i] + (2/epsilon)*math.Log(2*probability)
			} else {
				ratingEncoding[i] = ratingEncoding[i] - (2/epsilon)*math.Log(2-(2*probability))
			}
		}
		fmt.Fprint(fileHistogramEncoding, userId, movieId, ratingEncoding, "\n")
	}
	//Aggregation
	summationHistogramEncoding, _ := os.Open("HistogramEncoding.csv")
	defer summationHistogramEncoding.Close()

	aggregation := make(map[int][]float64, 0)
	tempRating := make([]float64, 11)

	for {
		_, err := fmt.Fscanf(summationHistogramEncoding, "%d %d [%f %f %f %f %f %f %f %f %f %f %f]\n", &userId, &movieId, &tempRating[0], &tempRating[1], &tempRating[2], &tempRating[3], &tempRating[4], &tempRating[5], &tempRating[6], &tempRating[7], &tempRating[8], &tempRating[9], &tempRating[10])
		if err == io.EOF {
			break
		}
		if aggregation[movieId] == nil {
			aggregation[movieId] = append(aggregation[movieId], tempRating...)
		} else {
			for i := 0; i < len(tempRating); i++ {
				if tempRating[i] != 0 {
					aggregation[movieId][i] = aggregation[movieId][i] + tempRating[i]
				}
			}
		}

	}
	fmt.Println("SHE end")
	return aggregation
}

func directEncoding(epsilon float64) map[int][]float64 {
	var userId, movieId, timestamp int
	var rating float64

	fmt.Println("DE start")

	fileRatings, _ := os.Open("ratings.csv")
	fileDirectEncoding, _ := os.Create("DirectEncoding.csv")
	defer fileRatings.Close()
	defer fileDirectEncoding.Close()

	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	var d float64 = 11
	for {
		_, err := fmt.Fscanf(fileRatings, "%d,%d,%f,%d\n", &userId, &movieId, &rating, &timestamp)
		if err == io.EOF {
			break
		}

		var probability float64 = random.Float64()
		// Encoding
		//	// None
		// Perterbing
		// domain 을 0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5 로 보면 domainSize = 11

		if 0 < probability && probability <= 1/(math.Exp(epsilon)+d-1) {
			rating = 0
		} else if 1/(math.Exp(epsilon)+d-1) < probability && probability <= 2*(1/(math.Exp(epsilon)+d-1)) {
			rating = 0.5
		} else if 2*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 3*(1/(math.Exp(epsilon)+d-1)) {
			rating = 1
		} else if 3*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 4*(1/(math.Exp(epsilon)+d-1)) {
			rating = 1.5
		} else if 4*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 4*(1/(math.Exp(epsilon)+d-1)) {
			rating = 2
		} else if 5*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 6*(1/(math.Exp(epsilon)+d-1)) {
			rating = 2.5
		} else if 6*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 7*(1/(math.Exp(epsilon)+d-1)) {
			rating = 3
		} else if 7*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 8*(1/(math.Exp(epsilon)+d-1)) {
			rating = 3.5
		} else if 8*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 9*(1/(math.Exp(epsilon)+d-1)) {
			rating = 4
		} else if 9*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 10*(1/(math.Exp(epsilon)+d-1)) {
			rating = 4.5
		} else if 10*(1/(math.Exp(epsilon)+d-1)) < probability && probability <= 11*(1/(math.Exp(epsilon)+d-1)) {
			rating = 5
		} else {

		}
		fmt.Fprintf(fileDirectEncoding, "%d %d %.1f\n", userId, movieId, rating)
		// if domain 을 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5 로 보면 domainSize = 10
	}
	//Aggregation
	directEncoding, _ := os.Open("DirectEncoding.csv")
	defer directEncoding.Close()

	aggregation := make(map[int][]float64, 0)
	numberOfMember := make(map[int]int, 0)

	for {
		_, err := fmt.Fscanf(directEncoding, "%d %d %f\n", &userId, &movieId, &rating)
		if err == io.EOF {
			break
		}
		tempRating := make([]float64, 11)
		for i := 0; i < len(tempRating); i++ {
			if i == int(rating*2) {
				tempRating[i] = 1
			} else {
				tempRating[i] = 0
			}
		}
		if aggregation[movieId] == nil {
			aggregation[movieId] = append(aggregation[movieId], tempRating...)
		} else {
			for i := 0; i < len(tempRating); i++ {
				if tempRating[i] != 0 {
					aggregation[movieId][i]++
				}
			}
		}
		numberOfMember[movieId]++
	}
	for key, _ := range aggregation {
		for i := 0; i < 11; i++ {
			var q float64 = 1 // 1 or 10
			aggregation[key][i] = (aggregation[key][i] - (float64(numberOfMember[key]) * (q / (math.Exp(epsilon) + d - 1)))) / ((math.Exp(epsilon) / (math.Exp(epsilon) + d - 1)) - (q / (math.Exp(epsilon) + d - 1)))
		}
	}
	fmt.Println("DE end")

	return aggregation
}
