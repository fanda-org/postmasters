package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	api "github.com/fanda-org/postmasters/api"
	"github.com/fanda-org/postmasters/config"
	"github.com/fanda-org/postmasters/database"
	"github.com/fanda-org/postmasters/database/models/system"
	"github.com/fanda-org/postmasters/web"
	"github.com/gorilla/mux"
	"github.com/xorcare/pointer"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", 15*time.Second, "The duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	var dir string
	flag.StringVar(&dir, "dir", "./public", "The directory to serve files from. Defaults to the public dir")
	flag.Parse()

	var useHTTP2 bool
	flag.BoolVar(&useHTTP2, "http2", true, "Enable HTTP/2")
	flag.Parse()

	cfg := config.New()
	database.Migrate(cfg.DB)
	// db, err := database.New(cfg.DB)
	// if err != nil {
	// 	//log.Fatal("Could not connect database")
	// 	panic(err)
	// }
	// defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()

	//views := template.Must(template.ParseGlob("./views/*.hbs"))
	//usersViews := template.Must(template.ParseGlob("./views/users/*.hbs"))
	//tmpl, _ := getTemplates()
	env := &config.Env{
		Config: cfg,
		//DB:          db,
		Logger:    &log.Logger{},
		WebRouter: router,
		APIRouter: apiRouter,
		//Views:      views,
		//UsersViews: tmpl,
	}

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/",
		// 	http.FileServer(http.Dir(dir))))
		cacheControlWrapper(http.FileServer(http.Dir(dir))))) //http.FileServer(http.Dir(dir))))
	//router.NotFoundHandler = http.HandlerFunc(homeHandler)

	//tmpl := template.Must(template.ParseFiles("./views/shared/layout.hbs"))

	// apiURL := fmt.Sprintf("%s://%s:%d/", env.Config.Web.Scheme,
	// 	env.Config.Web.Host, env.Config.Web.Port)

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	users := []system.User{}
	// 	response, err := http.Get(apiURL + "api/users")
	// 	if err != nil {
	// 		fmt.Printf("The HTTP request failed with error %s\n", err)
	// 		return
	// 	}

	// 	data, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		fmt.Printf("Read response body failed with error %s\n", err)
	// 		return
	// 	}

	// 	json.Unmarshal(data, &users)

	// 	tmpl := sharedViews.Lookup("layout.hbs")
	// 	tmpl.Execute(w, users)
	// })

	// API
	api := &api.App{}
	api.Initialize(env, apiRouter)

	// WEB
	web := web.Web{}
	web.Initialize(env, router)

	// HTTP/2
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS11,
	}
	log.Printf("HTTP/2 enabled: %t", useHTTP2)
	if useHTTP2 {
		tlsConfig.NextProtos = []string{"h2"}
	} else {
		tlsConfig.NextProtos = []string{"http/1.1"}
	}

	srvAddr := fmt.Sprintf("%s:%d", env.Config.Web.Host, env.Config.Web.Port)
	fmt.Println("Server address: ", srvAddr)
	// HTTP SERVER with both web and api routers configured
	srv := &http.Server{
		Handler:   router,
		Addr:      srvAddr,
		TLSConfig: tlsConfig,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: env.Config.Web.WriteTimeout,
		ReadTimeout:  env.Config.Web.ReadTimeout,
		IdleTimeout:  env.Config.Web.IdleTimeout,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			//"../ssl/certificate.crt",
			//"../ssl/privateKey.key")
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Shutting down")
	os.Exit(0)
}

func cacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var modTime time.Time
		// fi, err := fh.Stat()
		// if err != nil {
		// 	modTime = fi.ModTime()
		// } else {
		// 	modTime = time.Now()
		// }
		// etag := "\"" + file + modTime.String() + "\""
		w.Header().Set("Etag", r.RequestURI)
		w.Header().Set("Cache-Control", "max-age=2592000") // 30 days

		h.ServeHTTP(w, r)
	})
}

func testDatabase(dbConfig *config.DBConfig) {
	database.Migrate(dbConfig)

	db, err := database.New(dbConfig)
	if err != nil {
		panic("Failed to connect database. Error: " + err.Error())
	}
	// defer database.Close(db)
	defer db.Close()
	fmt.Printf("Database created\n")

	var sysorg = system.Organization{}
	db.Find(&sysorg, "org_code = ?", "SYS")
	//db.Where().First(&sysorg)
	if sysorg.OrgName == "" {
		billAddr := &system.Address{
			Attention: pointer.String("Balamurugan T"), AddrLine1: pointer.String("No.8/9, V.V. Koil street"),
			AddrLine2: pointer.String("Saidapet"), City: pointer.String("CHENNAI"),
			State: pointer.String("Tamilnadu"), Country: pointer.String("INDIA"),
			PostalCode: pointer.String("600015"), Phone: pointer.String("9940180875"), Fax: pointer.String("FAX"),
			AddrType: "BILL"}

		shipAddr := &system.Address{
			Attention: pointer.String("Balamurugan T"), AddrLine1: pointer.String("No.53/15A, Narayana palayam street"),
			City: pointer.String("KANCHIPURAM"), State: pointer.String("Tamilnadu"),
			Country: pointer.String("INDIA"), PostalCode: pointer.String("631501"),
			Phone: pointer.String("9940180875"), Fax: pointer.String("FAX"),
			AddrType: "SHIP"}

		priContact := &system.Contact{
			Salutation: pointer.String("Mr"), FirstName: pointer.String("Balamurugan"),
			LastName: pointer.String("Thanikachalam"), Email: pointer.String("software.balu@gmail.com"),
			WorkPhone: pointer.String("044 33504292"), Mobile: pointer.String("9940180875"),
			Designation: pointer.String("Sr.Manager - Technology"),
			Department:  pointer.String("SSW"), IsPrimary: true}

		secContact := &system.Contact{
			Salutation: pointer.String("Mrs"), FirstName: pointer.String("Seethalakshmi"),
			LastName: pointer.String("Balamurugan"), Email: pointer.String("seethabala2002@gmail.com"),
			Mobile: pointer.String("9789813931")}

		sysorg := system.Organization{
			OrgCode: "SYS", OrgName: "System", Description: pointer.String("System organization")}

		org := system.Organization{
			OrgCode: "BALA", OrgName: "Bala Personal", Description: pointer.String("Bala's personal account"),
			Addresses: []*system.Address{billAddr, shipAddr},
			Contacts:  []*system.Contact{priContact, secContact}}

		usr := &system.User{
			UserName: "sysuser", Email: "sysuser@system.com", FirstName: pointer.String("SysUser")}
		usr.SetPassword("Welcome@123")
		sysorg.Users = append(sysorg.Users, usr)

		// db.Create(&usr)
		db.Create(&sysorg)
		db.Create(&org)

		fmt.Printf("SystemOrg created\n")
	} else {
		fmt.Printf("SystemOrg already exists. Name: " + sysorg.OrgName)
	}
}
