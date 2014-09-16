package main

type NutritionalValue interface {
	convertTo(interface{}) NutritionQuantity
}

type NutritionQuantity struct {
	Value int
	Units string
}

type NutritionFact struct {
	Name         string
	Quantites    []NutritionQuantity
	ServingSize  int
	ServingUnits string
}

type Ingredient struct {
	Name     string
	Quantity float32
	Units    string
}

type Meal struct {
	Name           string
	NutritionFacts []NutritionFact
	Ingredients    []Ingredient
}

type MealPlan struct {
}

/**

{
    serving_size: {
        value: 1,
        units: "cup"
    },
    total_fat: {
        value: 13,
        units: "g"
    },
    cholesterol: {
        value: 30,
        units: "mg"
    }
}









*/
