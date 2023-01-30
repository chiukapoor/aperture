// Code generated by "enumer -type=ArithmeticOperator -output=arithmetic-operator-string.go"; DO NOT EDIT.

package components

import (
	"fmt"
	"strings"
)

const _ArithmeticOperatorName = "UnknownArithmeticAddSubMulDivXorLShiftRShift"

var _ArithmeticOperatorIndex = [...]uint8{0, 17, 20, 23, 26, 29, 32, 38, 44}

const _ArithmeticOperatorLowerName = "unknownarithmeticaddsubmuldivxorlshiftrshift"

func (i ArithmeticOperator) String() string {
	if i < 0 || i >= ArithmeticOperator(len(_ArithmeticOperatorIndex)-1) {
		return fmt.Sprintf("ArithmeticOperator(%d)", i)
	}
	return _ArithmeticOperatorName[_ArithmeticOperatorIndex[i]:_ArithmeticOperatorIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ArithmeticOperatorNoOp() {
	var x [1]struct{}
	_ = x[UnknownArithmetic-(0)]
	_ = x[Add-(1)]
	_ = x[Sub-(2)]
	_ = x[Mul-(3)]
	_ = x[Div-(4)]
	_ = x[Xor-(5)]
	_ = x[LShift-(6)]
	_ = x[RShift-(7)]
}

var _ArithmeticOperatorValues = []ArithmeticOperator{UnknownArithmetic, Add, Sub, Mul, Div, Xor, LShift, RShift}

var _ArithmeticOperatorNameToValueMap = map[string]ArithmeticOperator{
	_ArithmeticOperatorName[0:17]:       UnknownArithmetic,
	_ArithmeticOperatorLowerName[0:17]:  UnknownArithmetic,
	_ArithmeticOperatorName[17:20]:      Add,
	_ArithmeticOperatorLowerName[17:20]: Add,
	_ArithmeticOperatorName[20:23]:      Sub,
	_ArithmeticOperatorLowerName[20:23]: Sub,
	_ArithmeticOperatorName[23:26]:      Mul,
	_ArithmeticOperatorLowerName[23:26]: Mul,
	_ArithmeticOperatorName[26:29]:      Div,
	_ArithmeticOperatorLowerName[26:29]: Div,
	_ArithmeticOperatorName[29:32]:      Xor,
	_ArithmeticOperatorLowerName[29:32]: Xor,
	_ArithmeticOperatorName[32:38]:      LShift,
	_ArithmeticOperatorLowerName[32:38]: LShift,
	_ArithmeticOperatorName[38:44]:      RShift,
	_ArithmeticOperatorLowerName[38:44]: RShift,
}

var _ArithmeticOperatorNames = []string{
	_ArithmeticOperatorName[0:17],
	_ArithmeticOperatorName[17:20],
	_ArithmeticOperatorName[20:23],
	_ArithmeticOperatorName[23:26],
	_ArithmeticOperatorName[26:29],
	_ArithmeticOperatorName[29:32],
	_ArithmeticOperatorName[32:38],
	_ArithmeticOperatorName[38:44],
}

// ArithmeticOperatorString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ArithmeticOperatorString(s string) (ArithmeticOperator, error) {
	if val, ok := _ArithmeticOperatorNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ArithmeticOperatorNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ArithmeticOperator values", s)
}

// ArithmeticOperatorValues returns all values of the enum
func ArithmeticOperatorValues() []ArithmeticOperator {
	return _ArithmeticOperatorValues
}

// ArithmeticOperatorStrings returns a slice of all String values of the enum
func ArithmeticOperatorStrings() []string {
	strs := make([]string, len(_ArithmeticOperatorNames))
	copy(strs, _ArithmeticOperatorNames)
	return strs
}

// IsAArithmeticOperator returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ArithmeticOperator) IsAArithmeticOperator() bool {
	for _, v := range _ArithmeticOperatorValues {
		if i == v {
			return true
		}
	}
	return false
}
