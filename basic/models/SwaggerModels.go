package models

type SwaggerCompanyRegister struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	WorkField *string `json:"workField,omitempty"`
}

type SwaggerCustomerRegister struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Birth    *string `json:"birth,omitempty"`
}

type SwaggerWorkCreate struct {
	Name        string  `json:"name"`
	WorkField   string  `json:"workField"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type SwaggerWorkAction struct {
	WorkID int `json:"workID"`
}

type SwaggerWorkFields struct {
	WorkFileds []string `json:"workFields"`
}

type SwaggerGetGeneral struct {
	ID        int     `json:"ID"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	WorkField *string `json:"workField,omitempty"`
	Birth     *string `json:"birth,omitempty"`
	Amount    float64 `json:"amount"`
}

type SwaggerWork struct {
	ID          int     `json:"ID"`
	CompanyID   int     `json:"companyID"`
	Name        string  `json:"name"`
	WorkField   string  `json:"workField"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Date        string  `json:"date"`
	CompanyName string  `json:"companyName"`
}

type SwaggerLogin struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
}
