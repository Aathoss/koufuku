package framework

import (
	"database/sql"
	"os"

	"github.com/fsnotify/fsnotify"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gitlab.com/koufuku/logger"
)

var (
	//Variable de version
	err       error
	DBKoufuku *sql.DB
)

//LoadConfiguration charge les paramètres / variables
func init() {
	logger.InfoLogger.Println("----- [Koufuku] Démarrage du bot")
	logger.InfoLogger.Println("----- [Config] en préparation")

	//Configuration de l'heure sûr le serveur
	logger.InfoLogger.Println("----- [Config] Initialisation l'heure")
	os.Setenv("TZ", "Europe/Paris")

	//Chargement de la configuration du serveur
	logger.InfoLogger.Println("----- [Config] Initialisation du fichier de config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		logger.ErrorLogger.Println(err)
		os.Exit(10)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.InfoLogger.Println("Modification de la config :", e.Name)
	})

	//Connexion à la base de données lyana
	logger.InfoLogger.Println("----- [Config] Initialisation de la base de données [-Koufuku-]")
	dbUser := viper.GetString("MySql.Koufuku.dbuser")
	dbPass := viper.GetString("MySql.Koufuku.dbmdp")
	dbName := viper.GetString("MySql.Koufuku.dbname")
	dbIP := viper.GetString("MySql.Koufuku.dbip")
	dbPort := viper.GetString("MySql.Koufuku.dbport")

	DBKoufuku, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIP+":"+dbPort+")/"+dbName)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	if err = DBKoufuku.Ping(); err != nil {
		logger.ErrorLogger.Println(err)
	}

	if err != nil {
		os.Exit(10)
	}
	logger.InfoLogger.Println("----- [Config] Configuration charger")
}
