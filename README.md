# Go Ember Todo List
Implements a back end for https://github.com/ottogiron/todo-list using  <a href="https://github.com/ottogiron/chapi">Chapi</a>


##Drone build

Install drone
```
curl http://downloads.drone.io/drone-cli/drone_darwin_amd64.tar.gz | tar zx
sudo cp drone /usr/local/bin
```

```
drone exec
```

## Development

### Prerrequisites
* Install go
  * https://golang.org/doc/install

### Fork The repository

### Run the server

```bash
  mkidr -p $GOPATH/github.com/<my_github_user_name>
  cd $GOPATH/src/github.com/<my_github_user_name>
  git clone https://github.com/<my_github_user_name>/gotodo.git
  cd gotodo
  go get

  go run server.go
```
