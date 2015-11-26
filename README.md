# DB_PROJECT

## Environment Setup:
## goLang: [GoLang](https://golang.org/dl/)
### Linux Ubuntu:
* Extract Tar File: 'tar -c /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz'
* Add PATH Environment Variable: 'export PATH=$PATH:/usr/local/go/bin'
* Set GOPATH Environment Variable To Point To Workspace: 'export GOPATH=$HOME/"dir of workspace"'
* Test: 'go version'
###	Windows:
* Install By Using MSI installer
* Set Environment Path:
	* Control Panel
	* System and Security
	* System
	* Advanced System Settings
	* Environment Variables
	* Add The Directory Of The Go/Bin Folder Into System Variables PATH 
* Test: 'go version'

## NPM/Node.js: [NodeJS](https://nodejs.org/en/download/)
###	Linux Ubuntu:
* sudo apt-get install curl
* curl -sL https://deb.nodesource.com/setup_4.x | sudo -E bash -
* sudo apt-get install -yes nodejs
* sudo apt-get install npm
###	Windows:
* Just Run Installer From Download Page

## Node Modules:
* Open platform_ui directory
* Remove Folder node_modules
* npm install gulp -g --save-dev
* npm install gulp --save-dev
* npm install jquery -g --save-dev
* npm install jquery --save-dev
* npm install lodash -g --save-dev
* npm install lodash --save-dev
* npm install react -g --save-dev
* npm install react --save-dev
* npm install browserify -g --save-dev
* npm install browserify --save-dev
* npm install reactify -g --save-dev
* npm install reactify --save-dev
* npm install vinyl-source-stream -g --save-dev
* npm install vinyl-source-stream --save-dev
* npm install react-tools -g --save-dev
* npm install react-tools --save-dev

## Setup BackEnd:
###	Linux Ubuntu:
* Open platform directory
* Compile Code: 'make'
* Run Server: './server'
###	Windows:
* Open platform directory
* Compile code: 'go build server.go'
* Run Server: 'server.exe'

##	Setup FrontEnd:
* Open platform directory
* 'gulp build'

##	How To Submit Work To GitHub
* git add --all
* git commit -m "What Work Are You Committing"
* git push origin master

## How To Update Local Repository
* git pull master origin
* make sure to update any environment changes
