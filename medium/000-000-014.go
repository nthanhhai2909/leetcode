package main

import "fmt"

// Link: https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/description/

/*
You have information about n different recipes. You are given a string array recipes and a 2D string array ingredients. The ith recipe has the name recipes[i], and you can create it if you have all the needed ingredients from ingredients[i]. Ingredients to a recipe may need to be created from other recipes, i.e., ingredients[i] may contain a string that is in recipes.

You are also given a string array supplies containing all the ingredients that you initially have, and you have an infinite supply of all of them.

Return a list of all the recipes that you can create. You may return the answer in any order.

Note that two recipes may contain each other in their ingredients.



Example 1:

Input: recipes = ["bread"], ingredients = [["yeast","flour"]], supplies = ["yeast","flour","corn"]
Output: ["bread"]
Explanation:
We can create "bread" since we have the ingredients "yeast" and "flour".
Example 2:

Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]
Output: ["bread","sandwich"]
Explanation:
We can create "bread" since we have the ingredients "yeast" and "flour".
We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
Example 3:

Input: recipes = ["bread","sandwich","burger"], ingredients = [["yeast","flour"],["bread","meat"],["sandwich","meat","bread"]], supplies = ["yeast","flour","meat"]
Output: ["bread","sandwich","burger"]
Explanation:
We can create "bread" since we have the ingredients "yeast" and "flour".
We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
We can create "burger" since we have the ingredient "meat" and can create the ingredients "bread" and "sandwich".


Constraints:

n == recipes.length == ingredients.length
1 <= n <= 100
1 <= ingredients[i].length, supplies.length <= 100
1 <= recipes[i].length, ingredients[i][j].length, supplies[k].length <= 10
recipes[i], ingredients[i][j], and supplies[k] consist only of lowercase English letters.
All the values of recipes and supplies combined are unique.
Each ingredients[i] does not contain any duplicate values.
*/

/*
	N - Len recipes or ingredients
	M - Max Len ingredients[ith]
	TC: O(N^2.M)
	SC: O(N)
*/

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	ans := make([]string, 0)
	size := len(recipes)
	memory := make(map[int]struct{}, size)
	ind := 0
	sets := make(map[string]struct{})
	for _, val := range supplies {
		sets[val] = struct{}{}
	}
	for ind < size {
		if _, ok := memory[ind]; ok {
			ind++
			continue
		}
		ingredient := ingredients[ind]
		ok := true
		for _, val := range ingredient {
			if _, cons := sets[val]; !cons {
				ok = false
				break
			}
		}

		if ok {
			sets[recipes[ind]] = struct{}{}
			ans = append(ans, recipes[ind])
			memory[ind] = struct{}{}
			ind = 0
		} else {
			ind++
		}
	}
	return ans
}

//----------------------------------------------------------------------------------------------------------------------

/*
	TC: O(M.N)
	Solution: topological-sort - https://www.interviewcake.com/concept/java/topological-sort
*/
func findAllRecipesV2(recipes []string, ingredients [][]string, supplies []string) []string {
	ans := make([]string, 0)
	graph := make(map[string][]string)
	indegrees := make(map[string]int)
	availables := make([]string, 0)

	for ind, recipe := range recipes {
		for _, ingred := range ingredients[ind] {
			graph[ingred] = append(graph[ingred], recipe)
			indegrees[recipe]++
		}
	}

	availables = append(availables, supplies...)

	for len(availables) > 0 {
		current := availables[0]
		availables = availables[1:]
		for _, dependentRecipe := range graph[current] {
			indegrees[dependentRecipe]--
			if indegrees[dependentRecipe] == 0 {
				availables = append(availables, dependentRecipe)
				ans = append(ans, dependentRecipe)
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(findAllRecipesV2([]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}))
}
