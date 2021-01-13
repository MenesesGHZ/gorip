# Gorip v1.0
#### It is a tool for login with multi-accounts into Facebook and commit basic interactions:
 - Make multi-reactions given a postURL and a reactionID.
 - Get basic information from the user logged.
 - Multi-Scrap given facebookURLs.
 #### Without the need of Web Browser !!!

## Example Usage
### Import Gorip
```go
import "github.com/menesesghz/gorip/fbrip"
```
### Login into Facebook
```go
//Create usersRip
users := []*fbrip.UserRip{
    fbrip.CreateUser("mockuser@domainname.top","super_secret_pass"),
    fbrip.CreateUser("mockuser2@domainname.top","super_secret_pass")
}

// Main loop
for i,user := range users{
  //Making GET Request to facebook.com, and saving Cookies need it for login.
  user.Sense()
  
  //Login into Fb. Gathering extra body parameters and cookies need it for be logged in.
  user.Rip()
  
  /*... Actions ...*/
}
```
### Get Basic Info
```go
//Action configuration for getting basic info
actionConfig := fbrip.ActionConfig{GetBasicInfo:true}

// Main loop
for i,user := range users{
    /*...*/
    user.Do(actionConfig)
    fmt.Printf("\n> User: %s | Gender:[ %s ] Birthday:[ %s ]\n",user.Info.Name,user.Info.Gender,user.Info.Birthday)
 }

```
### Make Reactions
```go

actionConfig := fbrip.ActionConfig{
  React: FALTAAA
}
user := gorip

```
### Scrap Facebook Url
```go
//creating userRip and Login
user := fbrip.CreateUser("mockuser@domainname.top","super_secret_pass"),
user.Sense()
isLogged := user.Rip()

if isLogged {
  //Defining urls to scraps.
  urlSlice := []string{
    "https://www.facebook.com/GolangSociety",
    "https://www.facebook.com/googlemexico/"
  }

  //Define path where the HTMLs are going to be saved.
  path := "./scraps"

  //Action configuration for scrap
  scrap := fbrip.CreateScrap(path, urlSlice)
  actionConfig := fbrip.ActionConfig{Scrap:scrap}

  //User Do action
  user.Do(actionConfig)
}
```
### Doing the same above but with rip.json
```go
// Reading users and action config from JSON
users,actionConfig := fbrip.ReadRip("./rip.json")

// Main Loop 
for _,user := range users{
  // Login sequence
  user.Sense()
  isLogged := user.Rip()
  if isLogged {
    fmt.Println("* Logged Successfully.")
    u.Do(actionConfig)
    fmt.Printf("\n> User: %s | Gender:[ %s ] Birthday:[ %s ]\n",u.Info.Name,u.Info.Gender,u.Info.Birthday)
    fmt.Printf("* Actions Completed for -> %s\n\n",u.Parameters["email"])
   }
}
```
#### rip.json template
```json
{
	"Users":[
		{
			"Parameters":{
				"email":"someFacebookEmail_1@domain.com",
				"pass":"super_secret_password"
			}	
		},
		{
			"Parameters":{
				"email":"someFacebookEmail@domain.com",
				"pass":"super_secret_password"
			}
		}
	],
	"ActionConfig":{
		"GetBasicInfo":true,
		"React":[
      {
        "Url":"",
        "Id":""
      }
    ],
		"Scrap":{
			"Urls":[
			],
			"FolderPath":""
		}	
	}
}
```


## Future work to implement
### Actions:
- Add feature to allow m
- Make post(s) in own Profile Page or given Url(s) 
- Send Random Friend Requests
- Make Comment(s) given a postUrl
### Features:
- Create Facebook user without the need of web browser. (maybe)
### New packages:
- instarip (Instagram)
- twtrip (Twitter)
- gmailrip (Gmail)
