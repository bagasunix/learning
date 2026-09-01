package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type NutritionalInfo struct {
	Kcal float64 `json:"kcal"`
}

type Chocolate struct {
	Brand           string          `json:"brand"`
	Type            string          `json:"type"`
	Weights         []float64       `json:"weights"`
	NutritionalInfo NutritionalInfo `json:"nutritionalInformation"`
}

type ApiResponse struct {
	Page       int         `json:"page"`
	TotalPages int         `json:"total_pages"`
	Data       []Chocolate `json:"data"`
}

func energyBar(country string) string {
	page := 1
	bestEnergyRaw := -1.0
	bestResult := ""

	for {
		apiURL := fmt.Sprintf(
			"https://jsonmock.hackerrank.com/api/chocolates?countryOfOrigin=%s&page=%d",
			url.QueryEscape(country), page,
		)

		resp, err := http.Get(apiURL)
		if err != nil {
			fmt.Println("HTTP error:", err)
			break
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("Read body error:", err)
			break
		}

		var result ApiResponse
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("JSON error:", err)
			break
		}

		for _, choc := range result.Data {
			if len(choc.Weights) == 0 {
				continue
			}

			total := 0.0
			for _, w := range choc.Weights {
				total += w
			}
			avgWeight := total / float64(len(choc.Weights))

			energyRaw := choc.NutritionalInfo.Kcal * 0.01 * avgWeight
			if energyRaw > bestEnergyRaw {
				bestEnergyRaw = energyRaw
				bestResult = choc.Brand + " : " + choc.Type
			}
		}

		if result.TotalPages == 0 || page >= result.TotalPages {
			break
		}
		page++
	}

	return bestResult
}

func collision(speed []int32, pos int32) int32 {
	var count int32 = 0
	targetSpeed := speed[pos]

	for i := int32(0); i < int32(len(speed)); i++ {
		if i == pos {
			continue
		}
		if i < pos && speed[i] > targetSpeed {
			count++
		} else if i > pos && speed[i] < targetSpeed {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(energyBar("France"))
	fmt.Println(collision([]int32{6, 6, 1, 6, 3, 4, 6, 8}, 2))
}
