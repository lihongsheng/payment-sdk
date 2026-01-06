package params

//	{
//	   "type": "object",
//	   "properties": {
//	       "param_type": {
//	           "type": "string"
//	       },
//	       "param_object": {
//	           "type": "object",
//	           "properties": {
//	               "param_type": {
//	                   "type": "string"
//	               },
//	               "validate": {
//	                   "type": "string"
//	               }
//	           },
//	           "x-apifox-orders": [
//	               "param_type",
//	               "validate"
//	           ],
//	           "required": [
//	               "param_type",
//	               "validate"
//	           ]
//	       }
//	   },
//	   "x-apifox-orders": [
//	       "param_type",
//	       "param_object"
//	   ],
//	   "required": [
//	       "param_type",
//	       "param_object"
//	   ]
//	}
type Param struct {
	// 变量在html 展示的标签名称
	Label string `json:"label"`
	// name:变量名称
	Name string `json:"name"`
	// 变量类型 String|Int|Array|Object|Bool
	Type string `json:"type"`
	// 验证格式
	ValidateReg string `json:"validate_reg"`
	// String|Int|Email|Phone|Domain|Url|Reg
	ValidateType string `json:"validate_type"`
	// object 的子属性
	Properties []Param `json:"object"`
	// 默认值变量默认值：String|Int|Bool
	Default string `json:"default"`
	// Values 针对 String|Int|Bool 提供的可选性
	Values []Value `json:"values"`
	// InputType : select | radio | checkbox | text | number | password
	InputType string `json:"input_type"`
}
type Value struct {
	// 值
	Value string `json:"value"`
	// 描述
	Label string `json:"label"`
}
