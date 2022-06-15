
<div align="center" width="100%">
  <img src="./assets/reaper.png" align="center" width="auto" height="auto">
   <h1 align="center">Gorip v1.5</h1>
</div>
<p align="center">
	<a><img src="https://img.shields.io/badge/Version-1.5-green.svg" alt="Version"></a>
	<a><img src="https://img.shields.io/badge/Made%20with-Go-blue.svg" alt="Go"></a>
</p>

#### A tool for login into Facebook and commit basic interactions:
 - Make facebook reaction given an facebook URL.
 - Get basic information from the user logged.
 - Scrap www.facebook.com and mbasic.facebook.com

## Example Usage
### Import Gorip
```go
import "github.com/menesesghz/gorip/fbrip"
```
### Login into Facebook
```go
user := fbrip.NewUserRip("mock@email","mockpass"),
isLogged := user.Rip()
if isLogged{
    /*... do actions ...*/
}
```
### Get Basic Info
```go
user := fbrip.NewUserRip("mock@email","mockpass"),
isLogged := user.Rip()
if isLogged{
    user.GetBasicInfo()
    fmt.Printf("\n> User: %s | Gender:[ %s ] Birthday:[ %s ]\n", user.Info.Name, user.Info.Gender, user.Info.Birthday)
 }
```
### Make a Reaction
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
fbrip.NewUser("mock@email","password"),
isLogged := user.Rip()
if isLogged {
  react := fbrip.NewReaction("2", "https://www.facebook.com/RickandMorty/photos/pcb.5282285888534857/5282285578534888/")
  success := fbrip.Do(user, react)
  if success {
    fmt.Println("You have reacted '<3' to a Rick and Morty image !")
  } else {
    fmt.Println("You haven't reacted ;(")
  }
}

```
### Saving Facebook Scraps given Urls
```go
user := fbrip.CreateUser("mock@email","password")
isLogged := user.Rip()
if isLogged {
    scrap_facebook := fbrip.NewScrap(
      "https://www.facebook.com/profile.php?id=100008137277101",
      "./scraps/",
      "some-name-facebook-original",
      false, // scrap from www.facebook.com
    )
    scrap_mbasic := fbrip.NewScrap(
      "https://www.facebook.com/profile.php?id=100008137277101",
      "./scraps/",
      "some-name-facebook-mbasic", // if empty string 
      true, // scrap from mbasic.facebook.com
    )
    fbrip.Do(client, scrap_facebook)
    fbrip.Do(client, scrap_mbasic)
}
```
#### rip.json template (TODO)
<a href="https://github.com/MenesesGHZ/gorip/blob/main/rip.json">rip.json</a>

## Future work to implement
### Actions:
- Make post(s) in own Profile Page or given Url(s) 
- Send Random Friend Requests
- Make Comment(s) for given Facebook post Url(s)
- And more...
### Features:
- Create Facebook user without the need of web browser. (maybe)
### Things to Improve:
- Improve code for reading json file. (I was pretty dummie in golang at that time for reading json files lol)
- Improve fbrip/ code structure in general.
### New packages (maybe):
- instarip (Instagram)
- twtrip (Twitter)
- gmailrip (Gmail) (maybe)
- tikrip (TikTok) (maybe)


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
