package forms

import "github.com/go-playground/validator/v10"

// Profile service's input value (from body)
type ProfileForm struct {
	Nickname 		string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Age	  	 		int    `form:"age" json:"age" binding:"required,min=0,max=150"`
	Introduce 		string `form:"introduce" json:"introduce" binding:"required"`
}

// Custom validation error messages for ProfileForm 
func (f ProfileForm) Validate() string {
    validate := validator.New()
    err := validate.Struct(f)

    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            switch err.Field() {
            case "Nickname":
                return f.NicknameError(err.Tag())
            case "Age":
                return f.AgeError(err.Tag())
            case "Introduce":
                return f.IntroduceError(err.Tag())
            }
        }
    }
    return ""
}

// Custom validation error messages for Nickname field
func (f ProfileForm) NicknameError(tag string) string {
    switch tag {
    case "required":
        return "Nickname is required"
    case "min":
        return "Nickname is too short"
    case "max":
        return "Nickname is too long"
    default:
        return "Invalid nickname"
    }
}

// Custom validation error messages for Age field
func (f ProfileForm) AgeError(tag string) string {
    switch tag {
    case "required":
        return "Age is required"
    case "min":
        return "Age cannot be negative"
    case "max":
        return "Age is not realistic"
    default:
        return "Invalid age"
    }
}

// Custom validation error messages for Introduce field
func (f ProfileForm) IntroduceError(tag string) string {
    if tag == "required" {
        return "Introduce is required"
    }
    return "Invalid introduction"
}