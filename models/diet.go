package models

type (
	GENDER                int
	ACTIVITY_FACTOR       int
	CALCULATION_OPERATION int
)

const (
	FEMALE GENDER = iota
	MALE
)

const (
	SEDENTARY ACTIVITY_FACTOR = iota
	LIGHTLY_ACTIVE
	MODERATELY_ACTIVE
	VERY_ACTIVE
	EXTRA_ACTIVE
)

const (
	PLUS CALCULATION_OPERATION = iota
	MINUS
)

type (
	Diet struct {
		Name string `json:"name" bson:"name"`
	}

	DietType struct {
		Name        string                `json:"name" bson:"name"`
		Description string                `json:"description" bson:"description"`
		Operation   CALCULATION_OPERATION `json:"operation" bson:"operation"`
		Amount      float32               `json:"amount" bson:"amount"`
	}

	DietUsecase interface {
		CountREE(gender GENDER, weight float32, height float32, age int) float32
		CountCA(ree float32, af ACTIVITY_FACTOR) float32
		CountDCR(ca float32, dietType DietType) float32
	}

	DietRepository interface{}
)
