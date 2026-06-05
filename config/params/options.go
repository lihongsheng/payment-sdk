package params

type Type string

const (
	String Type = "String"
	Int    Type = "Int"
	Array  Type = "Array"
	Bool   Type = "Bool"
)

type ValueType string

const (
	ValueString ValueType = "String"
	ValueInt    ValueType = "Int"
	ValueBool   ValueType = "Bool"
)

type InputType string

const (
	InputSelect   InputType = "select"
	InputRadio    InputType = "radio"
	InputCheckbox InputType = "checkbox"
	InputText     InputType = "text"
	InputNumber   InputType = "number"
	InputPassword InputType = "password"
	InputTextarea InputType = "textarea"
)

type ValidateType string

// String|Int|Email|Phone|Domain|Url|Reg
const (
	ValidateString     ValidateType = "String"
	ValidateInt        ValidateType = "Int"
	ValidateEmail      ValidateType = "Email"
	ValidatePhone      ValidateType = "Phone"
	ValidateDomain     ValidateType = "Domain"
	ValidateUrl        ValidateType = "Url"
	ValidateReg        ValidateType = "Reg"
	ValidateRsaPublic  ValidateType = "RsaPublic"
	ValidateRsaPrivate ValidateType = "RsaPrivate"
	ValidateRsaCert    ValidateType = "RsaCert"
)

type Option struct {
	// 变量在html 展示的标签名称
	Label string `json:"label"`
	// name:变量名称
	Name string `json:"name"`
	// 变量类型 String|Int|Array|Object|Bool
	Type Type `json:"type"`
	// 验证格式
	ValidateReg string `json:"validate_reg"`
	// String|Int|Email|Phone|Domain|Url|Reg
	ValidateType ValidateType `json:"validate_type"`
	// 默认值变量默认值：String|Int|Bool
	Default string `json:"default"`
	// Values 针对 String|Int|Bool 提供的可选性
	Values []Value `json:"values"`
	// InputType : select | radio | checkbox | text | number | password
	InputType InputType `json:"input_type"`
	// 是否必须
	Require bool `json:"require"`
}
type Value struct {
	// 值
	Value string `json:"value"`
	// 描述
	Label string `json:"label"`
	// 变量类型 String|Int|Bool
	Type ValueType `json:"type"`
}
