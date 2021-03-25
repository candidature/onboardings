package handlers

import (
	"fmt"
	"github.com/candidature/onboardings/config"
	"github.com/candidature/onboardings/models"
	"github.com/candidature/onboardings/render"
	"log"
	"net/http"
)



func Home(w http.ResponseWriter, r *http.Request) {
	remoteIP:= r.RemoteAddr
	config.GetConfig().Scs.Put(r.Context(),"remoteIP", remoteIP)

	render.RenderTemplate(w,"home.html.tmpl", models.TD)

}


func About(w http.ResponseWriter, r *http.Request) {
	log.Print("Webcome to about..")
	remoteIP:= config.GetConfig().Scs.Get(r.Context(),"remoteIP")
	models.SetTeamplateDataStringMap("remoteIP", fmt.Sprintf("%v", remoteIP))

	models.SetTeamplateDataStringMap("test", "Hello again")
	render.RenderTemplate(w,"about.html.tmpl",models.TD)

}
