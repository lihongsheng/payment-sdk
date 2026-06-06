package params

// InputType 表单控件类型，由前端按此渲染输入组件。
type InputType string

const (
	InputSelect   InputType = "select"
	InputRadio    InputType = "radio"
	InputCheckbox InputType = "checkbox"
	InputText     InputType = "text"
	InputNumber   InputType = "number"
	InputPassword InputType = "password"
	InputTextarea InputType = "textarea"
	InputSwitch   InputType = "switch"
)

// ValidateType 字段校验类型，前后端共用。
type ValidateType string

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

// SchemaType JSON Schema 类型，符合 JSON Schema 2020-12 子集。
type SchemaType string

const (
	SchemaString  SchemaType = "string"
	SchemaInteger SchemaType = "integer"
	SchemaNumber  SchemaType = "number"
	SchemaBoolean SchemaType = "boolean"
	SchemaArray   SchemaType = "array"
	SchemaObject  SchemaType = "object"
)

// Schema 描述支付通道配置的字段树（JSON Schema 2020-12 子集）。
// 前端按 Schema 动态渲染表单并对提交内容做校验。
type Schema struct {
	Type        SchemaType         `json:"type,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	Properties  map[string]*Schema `json:"properties,omitempty"`
	Required    []string           `json:"required,omitempty"`
	Items       *Schema            `json:"items,omitempty"`
	Enum        []any              `json:"enum,omitempty"`
	OneOf       []SchemaEnum       `json:"oneOf,omitempty"`
	Default     any                `json:"default,omitempty"`
	Pattern     string             `json:"pattern,omitempty"`
	Format      string             `json:"format,omitempty"`
	// Examples 标准 JSON Schema 示例值数组；前端取第一项作为 placeholder。
	Examples []any     `json:"examples,omitempty"`
	UI       *UISchema `json:"x-ui,omitempty"`
	// Order 控制对象内字段的展示顺序（properties 是 map，无序）。
	Order []string `json:"x-order,omitempty"`
}

// SchemaEnum oneOf 单项；Const 是值，Title 是展示标签。
type SchemaEnum struct {
	Const any    `json:"const"`
	Title string `json:"title,omitempty"`
}

// UISchema 渲染层附加信息（不属于 JSON Schema 标准）。
type UISchema struct {
	InputType    InputType    `json:"input_type,omitempty"`
	ValidateType ValidateType `json:"validate_type,omitempty"`
	Rows         int          `json:"rows,omitempty"`
	Placeholder  string       `json:"placeholder,omitempty"`
}