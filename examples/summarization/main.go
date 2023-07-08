//go:build example
// +build example

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"text_summarizer/internal/clients"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	apiURL := os.Getenv("OPENAI_API_URL")
	apiKey := os.Getenv("OPENAI_API_KEY")

	client := clients.NewOpenAIClient(apiURL, apiKey)
	content := "Today, as the world lay shrouded in a dense mist, something remarkable happened in Fisantekraal. " +
		"Despite the cold and fog, a group of ten dedicated runners rose from their warm beds and made a conscious " +
		"choice to embrace the challenge of a 10km training run. Among them was our dear friend Samuel, whose presence " +
		"reminded us of the incredible strength and resilience that resides within the human spirit. Life has a way of " +
		"throwing unexpected obstacles on our paths. It tests our resolve, challenges our determination, and sometimes " +
		"fills our hearts with doubt. There are mornings when our bodies ache, when the warmth of our beds is enticing, " +
		"and when the world outside seems cloaked in an impenetrable mist of uncertainty. It is precisely during these " +
		"moments that the power of showing up shines its brightest. When we choose to show up, we are making a profound " +
		"statement to ourselves and to the world. We declare that our dreams, our goals, and our aspirations matter. " +
		"We acknowledge that progress is not always easy or comfortable, but it is through the act of showing up that " +
		"we create momentum. We build resilience and pave the way for personal growth. Today, as Samuel joined us on " +
		"this misty morning, his presence carried an indomitable spirit"

	summary, err := client.Summarize(content, "short")
	if err != nil {
		fmt.Println("Error summarizing:", err)
		return
	}

	fmt.Println("Summary:", summary)
}
