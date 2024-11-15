package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// CallGeminiAPI generates a description for the given item name
func CallGeminiAPI(itemName string) (string, error) {
	ctx := context.Background()

	// Create a new Gemini AI client
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer client.Close()

	// Use the "gemini-1.5-flash" model for content generation
	model := client.GenerativeModel("gemini-1.5-flash")

	// Prepare the prompt to generate description for the item
	prompt := fmt.Sprintf("deskripsikan secara singkat dan jelaskan manfaatnya untuk camping untuk item: %s", itemName)

	// Generate content based on the item name
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Extract and return the generated content
	return extractDescription(resp), nil
}

// Extracts the description from the API response
func extractDescription(resp *genai.GenerateContentResponse) string {
	// Print the full structure of the response to inspect it
	fmt.Printf("Full response structure: %+v\n", resp)

	// Check if there are candidates in the response
	if len(resp.Candidates) == 0 {
		return "No description generated."
	}

	// Loop through candidates to find the first valid description
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			// Loop through the parts in the content
			for _, part := range cand.Content.Parts {
				// Print the part structure to inspect its fields
				fmt.Printf("Part content: %+v\n", part)

				// At this point, look at the fields of 'part' after inspecting the printed structure.
				// Access the correct field based on what you find in the part structure.
				if part != nil {
					// Example: Assuming 'part' has a field named 'Text' or similar
					// Replace 'Text' with the actual field based on the inspection.
					// If there is a 'Content' field or another text-related field, access that.
					// Use the correct field name based on the inspection here.
					return fmt.Sprintf("%v", part) // Adjust this to match the correct field
				}
			}
		}
	}

	return "No description generated."
}
