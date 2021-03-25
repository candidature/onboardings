package config

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)
var lock = &sync.Mutex{}

type Single struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InfoLog *log.Logger
	InProduction bool
	Scs *scs.SessionManager
}
var SingleInstance *Single

//AppConfig holds Application Config.


func SetConfig()  {

	if SingleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if SingleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			SingleInstance = &Single{UseCache: true}
			tc, err:= CreateTemplateCache()
			if err!= nil {
				log.Fatal("can not create template cache")
			}
			SingleInstance.TemplateCache = tc
			SingleInstance.Scs = scs.New()
			SingleInstance.Scs.Lifetime = 24 * time.Hour

			//Persist the session even if browser window is closed.
			SingleInstance.Scs.Cookie.Persist = true
			SingleInstance.Scs.Cookie.SameSite = http.SameSiteLaxMode
			SingleInstance.Scs.Cookie.Secure = SingleInstance.InProduction
		} else {
			fmt.Println("Single Instance already created-1")
		}
		fmt.Printf("Cache is %t", SingleInstance.UseCache)
	} else if SingleInstance.UseCache == false {

		fmt.Println("Not using CACHEEEE")
		tc, err:= CreateTemplateCache()
		if err!= nil {
			log.Fatal("can not create template cache")
		}
		SingleInstance.TemplateCache = tc

	} else {
		fmt.Println("Single Instance already created-2")
	}
	fmt.Println("Single instanec is ", SingleInstance.TemplateCache)

}

func SetScsSession(scsSession *scs.SessionManager) {
	SingleInstance.Scs = scsSession
}

func SetInProduction(inProduction bool) {
	SingleInstance.InProduction = inProduction
}
func  SetUseCache(cacheUse bool) {
	fmt.Println("Resetting the cache...to ", cacheUse)
	SingleInstance.UseCache = cacheUse
}

func  GetConfig()  *Single{
	if SingleInstance.UseCache == false {
		tc, err:= CreateTemplateCache()
		if err!= nil {
			log.Fatal("can not create template cache")
		}
		SingleInstance.TemplateCache = tc
	}
	fmt.Printf("Returning Single instance!!")
	return SingleInstance
}


var functions = template.FuncMap{}

func CreateTemplateCache() (map[string]*template.Template, error) {

	log.Println("Creating tempplate cache start...")
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./src/github.gwd.broadcom.net/pg942665/go-code-1/templates/*.html.tmpl")

	if err!=nil {
		fmt.Println("Error readng templates")
		return myCache,err
	}

	for _ , page := range pages {
		fmt.Println("Parsing " , page)
		baseName := filepath.Base(page)
		fmt.Println("Page is currently ", page)
		fmt.Println("BaseName " , baseName)
		ts,err:= template.New(baseName).Funcs(functions).ParseFiles(page)
		if err!=nil {
			return myCache,err
		}

		matches, err:= filepath.Glob("./src/github.gwd.broadcom.net/pg942665/go-code-1/templates/*.layout.tmpl")
		if err!=nil {
			fmt.Println("Error from filepath.Glob")
			return myCache,err
		}

		if len(matches) > 0 {
			ts,err = ts.ParseGlob("./src/github.gwd.broadcom.net/pg942665/go-code-1/templates/*.layout.tmpl")
			if err!=nil {
				fmt.Println("Error from ts.ParseGlob")
				return myCache,err
			}
		}
		fmt.Println("I am here.")
		myCache[baseName] = ts
	}

	return myCache,nil
}
