# Go Todo List
Implements a todo list using <a href="http://emberjs.com/">Ember</a> and Chapi <a href="https://github.com/ottogiron/chapi">Chapi</a>


## Development

### Prerrequisites
* Install go
  * https://golang.org/doc/install
* <a href="https://nodejs.org/en/">Node 4.x.x</a>
* <a href="http://ember-cli.com/">Ember CLI</a>

### First Fork The repository

### Run the server

```bash
  mkidr -p $GOPATH/github.com/<my_github_user_name>
  cd $GOPATH/src/github.com/<my_github_user_name>
  git clone https://github.com/ottogiron/gotodo.git
  cd gotodo
  go get
  cd client && npm install && bower install  && ember build && cd ..  //ember stuff
  go run server.go
```
