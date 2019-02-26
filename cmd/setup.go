// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Start the web interface",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setup called")
		initWeb()
	},
}

// Setup is a struct we make to add values to and then create funcs in to pass
// as the handler so that it has both the context and the values we put in it.
type Setup struct {
	password string
	secret   string
}

func (s *Setup) login(c echo.Context) error {
	//username := c.FormValue("username")
	password := c.FormValue("password")

	if password == s.password {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "authed"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(s.secret))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// Check if steamcmd is installed and functional
func checkSteam() (path string, st string, err error) {
	path, err = exec.LookPath("steamcmd")
	if err != nil || path == "" {
		log.Printf("Error: %v", err)
		log.Printf("Path: %v", path)
		return path, "", err
	}
	return path, "", nil

}

// Check if docker is installed and if it is, get a version number.
func checkDocker() (path string, do DockerOutput, err error) {
	// docker version --format '{ "os": "{{.Server.Os}}" , "version":
	// "{{.Server.Version}}" ,"goversion": "{{.Server.GoVersion}}" }'
	//var do DockerOutput
	path, err = exec.LookPath("docker")
	if err != nil || path == "" {
		log.Printf("Error: %v", err)
		log.Printf("Path: %v", path)
		return path, do, err
	}
	log.Printf("Path: %v", path)
	format := []string{"version", "--format", `{ "os": "{{.Server.Os}}" , "version": "{{.Server.Version}}" ,"goversion": "{{.Server.GoVersion}}" }`}
	cmd := exec.Command("docker", format...)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Error: %v", err)
		log.Printf("Output: %v", string(out))
		return path, do, err
	}
	log.Printf("Output: %v", string(out))
	var stuff DockerOutput
	err = json.Unmarshal(out, &stuff)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Version: %v", stuff.Version)
	return path, stuff, nil

}

// DockerOutput to capture the results of checking our docker verson
type DockerOutput struct {
	Os        string `json:"os"`
	Version   string `json:"version"`
	Goversion string `json:"goversion"`
}

// TemplateRegistry struct
type TemplateRegistry struct {
	templates *template.Template
}

// Render Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func initWeb() {
	// We need Auth
	// https://auth0.com/blog/authentication-in-golang/

	pwd, err := pwGen(12)
	if err != nil {
		log.Panicf("pwGen failed generating pwd : %v", err)
	}
	log.Printf("Go to http://localhost:1323 and login with %v", pwd)
	sec, err := pwGen(12)
	if err != nil {
		log.Panicf("pwGen failed generating sec : %v", err)
	}

	s := &Setup{password: pwd, secret: sec}

	e := echo.New()
	e.Static("/css", "frontend/css")
	e.File("/favicon.ico", "favicon.ico")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("frontend/*.html")),
	}

	e.POST("/login", s.login)
	e.GET("/login", s.login)
	// e.POST("/refresh", jwt.RefreshTokenHandler())
	// r := e.Group("/restricted")
	// r.Use(jwt.AuthRequired())
	//r.GET("", restricted)
	r := e.Group("/")
	r.Use(middleware.JWT([]byte(s.secret)))
	r.GET("", startSetup)
	r.GET("step/:id", doStep)
	e.Logger.Fatal(e.Start(":1323"))
}

func startSetup(c echo.Context) error {

	// What do we need to check?
	// Ubuntu LTS
	//		- Current
	//		- Local
	// Docker
	//		- Current
	//		- Local
	// Docker-compose
	//		- Current
	//		- Local
	// Steamcmd
	//		- Current
	//		- Local
	// Atlas Binaries
	//		- Current
	//		- Local
	// Environment Variables
	//		- TZ (Default: Etc/UTC)
	//		- DOMAINNAME
	//			- existing domain
	//			- dyndns domain
	// 		- ATLASIP (Get this from the domain)
	// 		- GAMEPORT1 (Default: 27005)
	// 		- GAMEPORT2 (Default: 27006)
	// 		- ATLASQUERYPORT (Default: 27015)
	// 		- RCONPORT (Default: 27025)
	// 		- SEAMLESSPORT (Default: 27020)
	// 		- ADMINPASS (Default: create a random one)
	//
	// If those are all installed and at the right versions
	// Check container status
	// 		- Redis
	//		- Atlas Server (1x1)

	// check if we are authed.
	user := c.Get("user").(*jwt.Token)
	if user == "authed" {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	pagination := `<nav aria-label="Game Server Setup">
		<ul class="pagination justify-content-center">
		  <li class="page-item disabled">
			<a class="page-link" href="#" tabindex="-1">Start</a>
		  </li>
		  <li class="page-item"><a class="page-link" href="/step/1">Next</a></li>
		</ul>
	  </nav>`
	main := ``
	path, do, err := checkDocker()
	log.Printf("Output: %v", do)
	if err != nil || do.Version == "" {
		if err != nil {
			main += `<div class="alert alert-danger" role="alert">Error checking for Docker: `
			main += err.Error()
			main += `</div>`
		}

		if do.Version == "" {
			main += `<div class="alert alert-dark" role="alert">Docker Version information is blank</div>`
		}
	} else {
		main += `<div class="list-group list-group-accent">
					<div class="list-group-item list-group-item-accent-success list-group-item-success">Docker <span class="badge badge-success">Success</span></div>
					<div class="list-group-item list-group-item-accent-success">Path : `
		main += path
		main += `</div>
					<div class="list-group-item list-group-item-accent-success">Version : `
		main += do.Version
		main += `</div>
					<div class="list-group-item list-group-item-accent-success">OS : `
		main += do.Os
		main += `</div>
					<div class="list-group-item list-group-item-accent-success">Go Version : `
		main += do.Goversion
		main += `</div>	
				  </div>`
	}
	t, err := os.Stat("~/.steamcmd")
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("File info : %v", t.Name())
	}

	// path = ""
	// path, st, err := checkSteam()
	// log.Printf("Output: %v", path)
	if err != nil || path == "" {
		if err != nil {
			main += `<div class="alert alert-danger" role="alert">Error checking for Steamcmd: `
			main += err.Error()
			main += `</div>`
		}

		if path == "" {
			main += `<div class="alert alert-dark" role="alert">Steamcmd path not found</div>`
		}
	} else {
		main += `<div class="list-group list-group-accent">
					<div class="list-group-item list-group-item-accent-success list-group-item-success">Steamcmd <span class="badge badge-success">Success</span></div>
					<div class="list-group-item list-group-item-accent-success">Path : `
		main += path
		// main += `</div>
		// 			<div class="list-group-item list-group-item-accent-success">Version : `
		// main += st
		// main += `</div>
		// 			<div class="list-group-item list-group-item-accent-success">OS : `
		// main += st
		// main += `</div>
		// 			<div class="list-group-item list-group-item-accent-success">Go Version : `
		// main += st
		main += `</div>	
				  </div>
				  `
	}

	// Ok here we likely need to check it we have a config file
	// if we do, pull values from that to populate the fields we display, or
	// maybe just go straight to step/1
	// If we don't we should pop a form, get all the values, validate them and
	// save it out to some kind of yaml file or something.
	// Dropdown for picking game (only Atlas right now supported)

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name":       "Go Gaming Automated Server Manager",
		"msg":        "Hello, Yeeter",
		"author":     "Bobblehead",
		"desc":       "Installer for Go Gaming Automated Server Manager",
		"main":       template.HTML(main),       // Dont forget to go back and work out the security implications of this.  Doesnt appear to be escaped?
		"pagination": template.HTML(pagination), // Dont forget to go back and work out the security implications of this.  Doesnt appear to be escaped?
	})
}

func doStep(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World!")
	id, _ := strconv.Atoi(c.Param("id"))
	next := id + 1
	prev := id - 1

	if id == 0 {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	pagination := `<nav aria-label="Game Server Setup">
		<ul class="pagination justify-content-center">
		  <li class="page-item disabled">
			<a class="page-link" href="#" tabindex="-1">Previous</a>
		  </li>`
	if id > 1 {
		pagination += `<li class="page-item"><a class="page-link" href="/step/`
		pagination += strconv.Itoa(prev)
		pagination += `">`
		pagination += strconv.Itoa(prev)
		pagination += `</a></li>`
	}
	pagination += `<li class="page-item active">
			  <span class="page-link">`
	pagination += strconv.Itoa(id)
	pagination += `
			  <span class="sr-only">(current)</span>
			  </span>
		  	</li>
			  <li class="page-item"><a class="page-link" href="/step/`
	pagination += strconv.Itoa(next)
	pagination += `">`
	pagination += strconv.Itoa(next)
	pagination += `</a></li>
			<li class="page-item">
			<a class="page-link" href="/step/`
	pagination += strconv.Itoa(next)
	pagination += `">`
	pagination += `Next</a>
			</li>
		</ul>
	  </nav>`

	if id > 0 {
		main := `<div>Next is /step/`
		main += strconv.Itoa(next)
		main += ` and Previous is /step/`
		main += strconv.Itoa(prev)
		main += `</div>`
	} // Probably need to do some kinf of case/switch here to redefine what main with each step
	// Might even need to call out to a function and pass in the step and return
	// the main content based on validating a bunch of things.

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name":       "Go Gaming Automated Server Manager",
		"msg":        "Step" + strconv.Itoa(id),
		"author":     "Bobblehead",
		"desc":       "Manager for Atlas Server Grid on Docker",
		"main":       template.HTML("<div>Next is /step/" + strconv.Itoa(next) + " and Previous is /step/" + strconv.Itoa(prev) + "</div>"),
		"pagination": template.HTML(pagination),
	})
}

func init() {

	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func pwGen(n int) (string, error) {

	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes, err := bytesGen(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}
	return string(bytes), nil
}

func bytesGen(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
