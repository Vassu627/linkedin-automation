package search

import "fmt"

// SearchQuery represents user targeting criteria
type SearchQuery struct {
	JobTitle string
	Company  string
	Keywords string
}

// Result represents a found profile
type Result struct {
	ProfileURL string
}

// SearchProfiles simulates LinkedIn profile search with pagination
func SearchProfiles(query SearchQuery, pages int) []Result {
	fmt.Println("Starting search with criteria:")
	fmt.Println("Job Title:", query.JobTitle)
	fmt.Println("Company:", query.Company)
	fmt.Println("Keywords:", query.Keywords)

	results := []Result{}
	seen := make(map[string]bool)

	for page := 1; page <= pages; page++ {
		fmt.Println("Processing search page:", page)

		// simulate 5 results per page
		for i := 1; i <= 5; i++ {
			url := fmt.Sprintf(
				"https://linkedin.com/in/mock-profile-%d-%d",
				page, i,
			)

			// duplicate detection
			if seen[url] {
				continue
			}

			seen[url] = true
			results = append(results, Result{
				ProfileURL: url,
			})
		}
	}

	fmt.Println("Search completed. Profiles found:", len(results))
	return results
}
