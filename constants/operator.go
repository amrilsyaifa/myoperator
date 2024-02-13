package constant

type ValidationConfig struct {
	MaxLength int `json:"max_length,omitempty" binding:"required"`
}

type Operator struct {
	Name             string            `json:"name,omitempty" binding:"required"`
	Operator         string            `json:"operator,omitempty" binding:"required"`
	Code             []int             `json:"code,omitempty" binding:"required"`
	Key              string            `json:"key,omitempty" binding:"required"`
	ValidationConfig *ValidationConfig `json:"validation_config,omitempty"`
}

type Message string

const (
	VALID     Message = "valid"
	INVALID   Message = "invalid"
	BELOW_MIN Message = "bellow minimum length"
	ABOVE_MAX Message = "above maximum length"
	NOT_FOUND Message = "not found"
)

var Operators = []Operator{
	{Name: "Kartu Halo", Operator: "Telkomsel", Key: "telkomsel", Code: []int{11}},
	{Name: "Simpati", Operator: "Telkomsel", Key: "telkomsel", Code: []int{12, 13, 21}},
	{Name: "Loop", Operator: "Telkomsel", Key: "telkomsel", Code: []int{22}},
	{Name: "Kartu As", Operator: "Telkomsel", Key: "telkomsel", Code: []int{21, 23, 52, 53}},
	{Name: "by.U/Kartu As", Operator: "Telkomsel", Key: "by.u", Code: []int{51}},
	{Name: "Indosat M2", Operator: "Indosat Ooredoo", Key: "indosat", Code: []int{14}},
	{Name: "Matrix", Operator: "Indosat Ooredoo", Key: "indosat", Code: []int{55}},
	{Name: "Mentari", Operator: "Indosat Ooredoo", Key: "indosat", Code: []int{58}},
	{Name: "Mentari/Matrix", Operator: "Indosat Ooredoo", Key: "indosat", Code: []int{15, 16}},
	{Name: "IM3", Operator: "Indosat Ooredoo", Key: "indosat", Code: []int{56, 57}},
	{Name: "XL", Operator: "XL Axiata", Key: "xl", Code: []int{17, 18, 19, 59, 77, 78, 79}},
	{Name: "Axis", Operator: "XL Axiata", Key: "axis", Code: []int{31, 32, 33, 38}},
	{Name: "3", Operator: "3", Key: "tri", Code: []int{95, 96, 97, 98, 99}, ValidationConfig: &ValidationConfig{MaxLength: 13}},
	{Name: "Smartfren", Operator: "Smartfren", Key: "smartfren", Code: []int{81, 82, 83, 84, 85, 86, 87, 88, 89}},
	{Name: "Net1", Operator: "Net1", Key: "net1", Code: []int{27, 28}},
	{Name: "ByRU", Operator: "ByRU", Key: "byru", Code: []int{68}},
}
