package db

import{
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "gopkg.in/yaml.v2"
    "io/ioutil"
}

type Config struct{
	Database struct{
        Host     string `yaml:"host"`
        Port     int    `yaml:"port"`
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        DBName   string `yaml:"dbname"`
        SSLMode  string `yaml:"sslmode"`		
	}`yaml:"database"`
}

func InitDB() (*sql.DB, error) {
	// Read config file 
	configFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading config file : %v", err)
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
    if err != nil {
        return nil, fmt.Errorf("error parsing config file: %v", err)
    }

	// create connection string
	connStr := fmt.Sprintf(        
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMode,
	)

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
        return nil, fmt.Errorf("error connecting to database: %v", err)
    }

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

    return db, nil
}
