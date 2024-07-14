package register

import "time"

type UserDataComplete struct {
	User       User       `json:"user"`
	Motorcycle Motorcycle `json:"motorcycle"`
}

type User struct {
	ID                   string    `json:"id"`
	IdentificationTypeID int       `json:"identification_type_id"`
	Name                 string    `json:"name"`
	Email                string    `json:"email"`
	Phone                string    `json:"phone"`
	Active               bool      `json:"active"`
	Password             string    `json:"password"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type Motorcycle struct {
	ID                       int       `json:"id"`
	UserID                   string    `json:"user_id"`
	Plate                    string    `json:"plate"`
	Brand                    string    `json:"brand"`
	Model                    string    `json:"model"`
	Year                     int       `json:"year"`
	SOATFile                 string    `json:"soat_file"`
	PhotoFile                string    `json:"photo_file"`
	IdentificationFile       string    `json:"identification_file"`
	MechanicalTechnicianFile string    `json:"mechanical_technician_file"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

type UserRequest struct {
	ID                   string     `json:"id"`
	IdentificationTypeID int        `json:"identification_type_id"`
	Name                 string     `json:"name"`
	Email                string     `json:"email"`
	Phone                string     `json:"phone"`
	Active               bool       `json:"active"`
	Password             string     `json:"password"`
	Motorcycle           Motorcycle `json:"motorcycle"`
}

type UserMotorcycles struct {
	User       User         `json:"user"`
	Motorcycle []Motorcycle `json:"motorcycle"`
}
