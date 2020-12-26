package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Food struct {
	ingredients []string
	allergens   []string
}

func commonIngredientsByAllergen(foods []Food) map[string]map[string]bool {
	common := make(map[string]map[string]bool)
	for _, f := range foods {
		for _, a := range f.allergens {
			_, ok := common[a]
			if !ok {
				common[a] = make(map[string]bool)
				for _, i := range f.ingredients {
					common[a][i] = true
				}
			}
		}
	}

	for _, f := range foods {
		foodIngr := make(map[string]bool)
		for _, i := range f.ingredients {
			foodIngr[i] = true
		}

		for _, allergen := range f.allergens {
			for ingredient := range common[allergen] {
				_, ok := foodIngr[ingredient]
				if !ok {
					delete(common[allergen], ingredient)
				}
			}
		}
	}
	return common
}

func getAllergens(foods []Food) []string {
	allergens := make(map[string]bool)
	for _, f := range foods {
		for _, a := range f.allergens {
			allergens[a] = true
		}
	}
	result := make([]string, 0)
	for allergen := range allergens {
		result = append(result, allergen)
	}
	return result
}

func getInput(r io.Reader) []Food {
	b, _ := ioutil.ReadAll(r)
	input := string(b)

	foodRe := regexp.MustCompile(`^(.*) \(contains (.*)\)$`)

	foods := make([]Food, 0)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		matches := foodRe.FindStringSubmatch(line)

		ingredients := strings.Split(matches[1], " ")
		allergens := strings.Split(matches[2], ", ")

		foods = append(foods, Food{ingredients, allergens})
	}

	return foods
}

func assertAllergen(allergen, ingredient string, common map[string]map[string]bool) map[string]map[string]bool {
	for a, is := range common {
		if a != allergen {
			for i := range is {
				if i == ingredient {
					delete(common[a], i)
				}
			}
		}
	}
	return common
}

func findAllergensByIngredient(common map[string]map[string]bool) map[string]string {
	found := make(map[string]string)
	done := false
	for !done {
		done = true
		for allergen, ingredients := range common {
			if len(ingredients) == 1 {
				for ingredient := range ingredients {
					_, ok := found[ingredient]
					if !ok {
						done = false
						found[ingredient] = allergen
						common = assertAllergen(allergen, ingredient, common)
					}
				}
			}
		}
	}
	return found
}

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	foods := getInput(stdinReader)

	common := commonIngredientsByAllergen(foods)
	found := findAllergensByIngredient(common)

	count := 0
	for _, f := range foods {
		for _, i := range f.ingredients {
			_, ok := found[i]
			if !ok {
				count++
			}
		}
	}
	fmt.Println("The number of safe foods:", count)

	pairs := make([][2]string, 0)
	for ingredient, allergen := range found {
		pairs = append(pairs, [2]string{allergen, ingredient})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
	ingredients := make([]string, 0)
	for _, pair := range pairs {
		ingredients = append(ingredients, pair[1])
	}
	fmt.Println("The dangerous ingredients:", strings.Join(ingredients, ","))
}
