package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
    if prepTime == 0 {
        return len(layers) * 2
    }
    return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodle := 0
    sauce := 0.0
    for _, v := range layers {
        if v == "noodles" {
            noodle += 50
        } else if v == "sauce" {
            sauce += 0.2
        }
    }
    return noodle, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
    myList[len(myList) - 1] = friendList[len(friendList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    newReceip := make([]float64, len(quantities))
    factor := float64(portions) / 2
    for i, v := range(quantities) {
        newReceip[i] = v * factor
    }
    return newReceip
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
