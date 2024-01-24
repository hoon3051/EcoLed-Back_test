package forms

import "github.com/go-playground/validator/v10"

//Post Service's input value
type PostForm struct {
    Title string `json:"title" validate:"required,min=2,max=100"`
    Body  string `json:"body" validate:"required,min=1,max=1000"`
}

// Custom validation error messages for PostForm 
func (f PostForm) Validate() string {
    validate := validator.New()
    err := validate.Struct(f)

    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            switch err.Field() {
            case "Title":
                return f.TitleError(err.Tag())
            case "Body":
                return f.BodyError(err.Tag())
            }
        }
    }
    return ""
}

// Custom validation error messages for Title field
func (f PostForm) TitleError(tag string) string {
    switch tag {
    case "required":
        return "Title is required"
    case "min":
        return "Title is too short"
    case "max":
        return "Title is too long"
    default:
        return "Invalid title"
    }
}

// Custom validation error messages for Body field
func (f PostForm) BodyError(tag string) string {
    switch tag {
    case "required":
        return "Body is required"
    case "min":
        return "Body is too short"
    case "max":
        return "Body is too long"
    default:
        return "Invalid body"
    }
}