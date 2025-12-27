package config

// {
//    "type": "object",
//    "properties": {
//        "param_type": {
//            "type": "string"
//        },
//        "param_object": {
//            "type": "object",
//            "properties": {
//                "param_type": {
//                    "type": "string"
//                },
//                "validate": {
//                    "type": "string"
//                }
//            },
//            "x-apifox-orders": [
//                "param_type",
//                "validate"
//            ],
//            "required": [
//                "param_type",
//                "validate"
//            ]
//        }
//    },
//    "x-apifox-orders": [
//        "param_type",
//        "param_object"
//    ],
//    "required": [
//        "param_type",
//        "param_object"
//    ]
//}

type Param struct {
	// 变量在html 展示的标签名称
	Label string `json:"label"`
	// 变量类型 String|Int|Array|Object|Bool
	Type string `json:"type"`
	// 验证格式
	ValidateReg string `json:"validate_reg"`
	// String|Int|Email|Phone|Domain|Url|Reg
	ValidateType string `json:"validate_type"`
	Object       *Param `json:"object"`
}
