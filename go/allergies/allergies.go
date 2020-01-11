package allergies

// type allergicResult struct {
// 	substance string
// 	result    bool
// }

// Allergies returns an array of all allergic substances
func Allergies(n uint) []string {

	substances := []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
	result := make([]string, 0)
	for i := 0; i < len(substances); i++ {
		t := uint(1 << i)

		if (t & n) == t {
			result = append(result, substances[i])
		}
	}
	return result
}
