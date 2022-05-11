package celeritas

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/okeefem2/celeritas/render"
	"github.com/okeefem2/celeritas/session"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	JetViews *jet.Set
	Session  *scs.SessionManager
	DB       Database

	// config config
}

// type config struct {
// 	port     string
// 	renderer string
// 	// cookie cookieConfig
// 	sessionType string

// }

// Config pattern
// Log env data into private config data
// Init package config data
// Pass env data to package config data
// Call Init function to fully init config data?
// pattern is somewhat confusing. I would just use functions for alot of this

func (c *Celeritas) New(rootPath string) error {
	// Initialize

	pathConfig := initPaths{
		rootPath: rootPath,
		// I think using templates for this would be better
		folderNames: []string{
			"handles", "migrations", "views", "data", "client", "tmp", "logs", "middleware",
		},
	}

	err := c.Init(pathConfig)

	if err != nil {
		return err
	}

	err = c.checkDotEnv(rootPath)

	if err != nil {
		return err
	}

	// Read the .env file and load into the app for usage
	err = godotenv.Load(rootPath + "/.env")

	if err != nil {
		return err
	}

	// create loggers
	infoLog, errorLog := c.startLoggers()

	c.InfoLog = infoLog
	c.ErrorLog = errorLog

	c.Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		c.ErrorLog.Fatal("DEBUG env variable must be parseable into a bool")
	}

	dbType := strings.ToLower(os.Getenv("DATABASE_TYPE"))
	// Connect to database
	if dbType != "" {
		db, err := c.OpenDB(dbType, c.buildDSN())

		if err != nil {
			errorLog.Println(err)
			os.Exit(1)
		}
		c.DB = Database{
			DataType: dbType,
			Pool: db,
		}
	}

	// I'd like to think about this a little more.
	c.Version = version
	c.RootPath = rootPath

	// // Seems a bit odd having duplicated configs across packages... want to revisit this maybe?
	// // Whi not just directly add the ENV to the Session struct?
	// Basically he is creating a "junk drawer" config struct to hold all this info that is used to configure other things
	// c.config = config{
	// 	port:     os.Getenv("PORT"),
	// 	renderer: os.Getenv("RENDERER"),
	// 	cookie: cookieConfig{
	// 		name: os.Getenv("COOKIE_NAME"),
	// 		lifetime: os.Getenv("COOKIE_LIFETIME"),
	// 		persist: os.Getenv("COOKIE_PERSISTS"),
	// 		secure: os.Getenv("COOKIE_SECURE"),
	// 		domain: os.Getenv("COOKIE_DOMAIN"),
	// 	},
	// 	sessionType: os.Getenv("SESSION_TYPE"),
	// }

	// Personally I think I would just pass this stuff to a function maybe rather than a receiver?
	// Idk using the receiver method is more OOP in my eyes. I'd prefer to pass the struct to a function and get
	// back the thing I need, but either way works fine I suppose
	s := session.Session{
		CookieLifetime: os.Getenv("COOKIE_LIFETIME"),
		CookiePersist:  os.Getenv("COOKIE_PERSISTS"),
		CookieSecure:   os.Getenv("COOKIE_SECURE"),
		CookieDomain:   os.Getenv("COOKIE_DOMAIN"),
		CookieName:     os.Getenv("COOKIE_NAME"),
		SessionType:    os.Getenv("SESSION_TYPE"),
	}

	c.Session = s.InitSession()

	c.InfoLog.Println("Session initialized", c.Session)

	// Init default routes middleware
	c.Routes = c.routes().(*chi.Mux)

	// This maybe doesn't need to live on the Celeritas struct
	c.JetViews = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	c.Render = &render.Render{
		Renderer: os.Getenv("RENDERER"),
		RootPath: c.RootPath,
		Port:     os.Getenv("PORT"),
		JetViews: c.JetViews,
	}

	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath

	// Init all folder paths specified by the initPaths config
	for _, path := range p.folderNames {
		// create folder if it doesn't exist
		err := c.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListenAndServe starts tge web server
func (c *Celeritas) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     c.ErrorLog,
		Handler:      c.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	defer c.DB.Pool.Close()

	c.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	c.ErrorLog.Fatal(err)
}

func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}

	return nil
}

// This is useful because we can add
func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	// | being the bitwise or operator
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// log shortfile gives info on where the error occurred
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}

func (c *Celeritas) buildDSN() string {
	var dsn string

	switch strings.ToLower(os.Getenv("DATABASE_TYPE")) {
		case "postgres", "postgresql":
			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)

			if os.Getenv("DATABASE_PASS") != "" {
				dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
			}
		default:
	}

	return dsn
}
