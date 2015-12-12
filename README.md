# DB_PROJECT

## Environment Setup:
## goLang: [GoLang](https://golang.org/dl/)
### Linux Ubuntu:
* Extract Tar File: 'tar -c /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz'
* Add PATH Environment Variable: 'export PATH=$PATH:/usr/local/go/bin'
* Set GOPATH Environment Variable To Point To Workspace: 'export GOPATH=$HOME/"dir of workspace"'
* Test: 'go version'

## NPM/Node.js: [NodeJS](https://nodejs.org/en/download/)
###	Linux Ubuntu:
* sudo apt-get install curl
* curl -sL https://deb.nodesource.com/setup_4.x | sudo -E bash -
* sudo apt-get install -yes nodejs
* sudo apt-get install npm

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
* npm install request -g --save-dev
* npm install request --save-dev

## Start Backend Services
* sudo -s
* source start.sh

## Setup BackEnd:
###	Linux Ubuntu:
* Open platform directory
* Compile Code: 'make'
* Run Server: './server'

## Setup FrontEnd:
* Open platform directory
* 'gulp build'
