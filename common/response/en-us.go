package response

var enUSText = map[int]string{
	Success:            "success",
	Fail:               "Fail",
	ParamError:         "Parameter error",
	AuthorizationError: "Authorization error",
	AccountError:       "address error",
	NotRegister:        "NotRegister",
	PasswordErr:        "PasswordErr",
	SmsCodeErr:         "SmsCodeErr",
	UsernameErr:        "UsernameErr",
	AccountRegistered:  "Account Registered",
	AccountFormatErr:   "AccountFormatErr",
	UpdateErr:          "update data error",
	FileFormatErr:      "file format or size too big",

	ServerNotExist:          "server pkg not exist",
	BenefitCodeErr:          "benefit code err",
	BenefitCodeUsed:         "benefit code has used",
	BenefitNotEnough:        "Insufficient equity frequency",
	PermissionErr:           "user permission err",
	AddressFormatErr:        "address format err",
	AccountBalanceNotEnough: "account balance not enough ",
	PayStatusErr:            "pay status err",
	FileAmountErr:           "file amount err",
	FileNotExist:            "file not exist",
	FileStatusErr:           "file status err",
}
