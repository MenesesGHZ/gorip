
<div align="center" width="100%">
  <img src="./assets/reaper.png" align="center" width="auto" height="auto">
   <h1 align="center">Gorip v1.0</h1>
</div>
<p align="center">
	<a><img src="https://img.shields.io/badge/Version-1.0.0-green.svg" alt="Version"></a>
	<a><img src="https://img.shields.io/badge/Made%20with-Go-blue.svg" alt="Go"></a>
</p>

#### It is a tool for login with multi-accounts into Facebook and commit basic interactions:
 - Make multi-reactions given Facebook Post URLs and Reaction IDs.
 - Get basic information from the user logged.
 - Multi-Scrap given Facebook URLs.
 #### Without the need of Web Browser !!!

## Example Usage
### Import Gorip
```go
import "github.com/menesesghz/gorip/fbrip"
```
### Login into Facebook
```go
//Create a UsersRip slice
users := []fbrip.UserRip{
    fbrip.CreateUser("mock@email","mockpass"),
    fbrip.CreateUser("mock@email","mockpass"),
}
// Main loop
for _,user := range users{
  //Making GET Request to facebook.com, and saving Cookies need it for login.
  user.Sense()
  
  //Login into Fb. Gathering extra body parameters and cookies need it for be logged in.
  isLogged := user.Rip()
 
   if isLogged{
  	/*... Actions ...*/
  }
}
```
### Get Basic Info
```go
//Action configuration for getting basic info
actionConfig := fbrip.ActionConfig{GetBasicInfo:true}

// Main loop
for _,user := range users{
    /*...*/
    user.Do(&actionConfig)
    fmt.Printf("\n> User: %s | Gender:[ %s ] Birthday:[ %s ]\n",user.Info.Name,user.Info.Gender,user.Info.Birthday)
 }

```
### Make multiple Reactions in multiple accounts
The reactions Id are the same that Facebook provides and are listed below:

<br>
<div align="center" width="100%">
  <img src="./assets/reactions.png" align="center" width="310px" height="auto">
</div>
<br>

- "1" -> Like
- "2" -> Love
- "3" -> Care
- "4" -> Haha
- "5" -> Wow
- "6" -> Sad 
- "7" -> Angry

```go
//Create usersRip
users := []fbrip.UserRip{
    fbrip.CreateUser("mock@email","password"),
}

//Urls Slice & reactionIds Slice (urls -> reactionsIds)
urls := []string{
	"https://www.facebook.com/googlemexico/videos/756896714867588/",
	"https://www.facebook.com/GolangSociety/photos/p.1832113056840497/1832113056840497/",
}
reactionIds := []string{"1","7"}

// Making action configuration
actionConfig := fbrip.ActionConfig{React: fbrip.CreateReact(reactionIds,urls),}

//Main Loop
for _,user := range users{
	//Login	
	user.Sense()
	isLogged := user.Rip()
	if isLogged {
		//Do	
		user.Do(&actionConfig)
	}
}
```
### Saving Facebook Scraps given Urls
```go
//creating userRip and Login
user := fbrip.CreateUser("mock@email","password")
user.Sense()
isLogged := user.Rip()

if isLogged {
      //Defining urls to scrap.
      urlSlice := []string{
              "https://www.facebook.com/GolangSociety",
              "https://www.facebook.com/googlemexico/",
      }

      //Define path where the HTMLs are going to be saved.
      path := "./scraps"

      //Action configuration for scrap
      scrap := fbrip.CreateScrap(path, urlSlice)
      actionConfig := fbrip.ActionConfig{Scrap:*scrap}

      //User Do action
      user.Do(&actionConfig)
}
```
### Doing the same thing as before, but with rip.json
```go
// Reading Users and Action configuration from JSON
users,actionConfig := fbrip.ReadRip("./rip.json")

// Main Loop 
for _,user := range users{
  // Login sequence
  user.Sense()
  isLogged := user.Rip()
  if isLogged {    
    user.Do(actionConfig)
    
    fmt.Printf("\n> User: %s | Gender:[ %s ] Birthday:[ %s ]\n",user.Info.Name,user.Info.Gender,user.Info.Birthday)
    fmt.Printf("* Actions Completed for -> %s\n\n",user.Parameters["email"])
   }
}
```
#### rip.json template
<a href="https://github.com/MenesesGHZ/gorip/blob/main/rip.json">rip.json</a>

## Future work to implement
### Actions:
- Make post(s) in own Profile Page or given Url(s) 
- Send Random Friend Requests
- Make Comment(s) for given Facebook post Url(s)
- And more...
### Features:
- Create Facebook user without the need of web browser. (maybe)
### New packages:
- instarip (Instagram)
- twtrip (Twitter)
- gmailrip (Gmail) (maybe)

## About The Project
This project is the evidence for my first serious try to Go language. I wanted to make something with it that challenges me. As you can see the code is not perfect, and it can be optimized a lot. If you are interest in improving the project, to add new cool functions or whatever. I invite you to send me a message and see what we can do. I am widely open to add colaborators to make this project better. Also I if you have any suggestion, proposal, or something that you would like to use in this package, please let me know.

## About Me
A Computer Science Student at CETYS.
If you want to contact me here is my email: <a href="mail:gerardo.meneses.hz@gmail.com">gerardo.meneses.hz@gmail.com</a>


## Thanks to
- My secondary account that has been banned for multiple testing requests. Rest in peace. For now...

<div align="center" width="100%">
  <img src="./assets/result.png" align="center" width="450px" height="auto">
</div>
