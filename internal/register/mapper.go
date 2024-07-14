package register

import "time"

func mapUserRequestToModelUser(userReq UserRequest) UserDataComplete {
	return UserDataComplete{
		User: User{
			ID:                   userReq.ID,
			IdentificationTypeID: userReq.IdentificationTypeID,
			Name:                 userReq.Name,
			Email:                userReq.Email,
			Phone:                userReq.Phone,
			Active:               userReq.Active,
			Password:             userReq.Password,
		},
		Motorcycle: Motorcycle{
			UserID:                   userReq.ID,
			Plate:                    userReq.Motorcycle.Plate,
			Brand:                    userReq.Motorcycle.Brand,
			Model:                    userReq.Motorcycle.Model,
			Year:                     userReq.Motorcycle.Year,
			SOATFile:                 userReq.Motorcycle.SOATFile,
			PhotoFile:                userReq.Motorcycle.PhotoFile,
			IdentificationFile:       userReq.Motorcycle.IdentificationFile,
			MechanicalTechnicianFile: userReq.Motorcycle.MechanicalTechnicianFile,
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
		},
	}
}
