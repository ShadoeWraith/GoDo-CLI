# A simple Go CLI To-Do application.

I made this to work alongside the [GoDo Web App and API](https://github.com/ShadoeWraith/GoDo), and locally offline

## Commands

#### When built in the future will replace `go run main.go {command}` to `godo {command}`

- godo list
- godo create -title? {title} -description? {description}
- godo update -id? {id} -title? {title} -description? {description}
- godo delete -id? {id}
- godo -help
- godo -version
- godo -setuser {userID}


## If you wanna run it
`git clone {project}`

`cd {path/to/project/dir}`

`touch tasks.json`

`go run main.go {command}`


## Future Plans
- Finish current commands
- Add Push and Pull commands to get data from Web API
- Make Set User allow interaction with Web API
