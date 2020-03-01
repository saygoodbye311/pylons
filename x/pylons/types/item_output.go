package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ItemOutput models the continuum of valid outcomes for item generation in recipes
type ItemOutput struct {
	Doubles DoubleParamList
	Longs   LongParamList
	Strings StringParamList
	Weight  int
}

func (io ItemOutput) String() string {
	return fmt.Sprintf(`ItemOutput{
		Doubles: %+v,
		Longs:   %+v,
		Strings: %+v,
		Weight:  %d,
	}`, io.Doubles, io.Longs, io.Strings, io.Weight)
}

func (io ItemOutput) GetWeight() int {
	return io.Weight
}

func (io ItemOutput) Item(cookbook string, sender sdk.AccAddress) (*Item, error) {
	// This function is used on ExecuteRecipe's AddExecutedResult, and it's
	// not acceptable to provide predefined GUID
	dblActualize, err := io.Doubles.Actualize(nil, nil)
	if err != nil {
		return nil, err
	}
	longActualize, err := io.Longs.Actualize(nil, nil)
	if err != nil {
		return nil, err
	}
	stringActualize, err := io.Strings.Actualize(nil, nil)
	if err != nil {
		return nil, err
	}

	return NewItem(cookbook, dblActualize, longActualize, stringActualize, sender), nil
}
