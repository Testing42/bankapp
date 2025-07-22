package fileops

/*
This is an example of creating another package
to do this you need to make a new directory with the same name as this file and then add this file to it
Then you will need to call the new package in your main package or any code/package that needs it
in this case that is bank.go

do that by going to the go.mod file and look at the base module you created
module example.com/bank
then add fileops to the end to import
example.com/bank/fileops



IT is important that you make each function used in this way start with a capital otherwise this won't work.
*/
import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}
	// the above if statment is how to handle potential errors in golang
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

//the underscore tell go to not worry about the error if one happens

func WriteFloatToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
