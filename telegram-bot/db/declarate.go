package db

const (
	userLevel       = 0
	adminLevel      = 1
	superAdminLevel = 2
)

type botUser struct {
	Username    string
	Password    string
	AccessLevel int
}