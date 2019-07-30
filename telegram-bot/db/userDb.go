package db

func returnUsersSlice() []botUser {

	botUserSlice := []botUser{
		{
			Username:    "vovavovinko",
			Password:    "provalionok_pidor",
			AccessLevel: superAdminLevel,
		},
		{
			Username:    "vladoleg",
			Password:    "sohonevich_pidor",
			AccessLevel: adminLevel,
		},
	}

	return botUserSlice
}
