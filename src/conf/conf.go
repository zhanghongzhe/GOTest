package conf

func init()  {
	Conf.DB.Name="OnlineDB"
	Conf.DB.UserName="sa"
	Conf.DB.Pwd="colipu"
	Conf.DB.Host="10.10.4.201"
	Conf.DB.Port=1433
}

var (
	Conf config
)

type config struct {
	DB database
}

type database struct {
	Name     string `toml:"name"`
	UserName string `toml:"user_name"`
	Pwd      string `toml:"pwd"`
	Host     string `toml:"host"`
	Port     int `toml:"port"`
}