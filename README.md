### Monitor HTTP

###### Uptime monitor, latency and requests

<br />

Features:
- Latency requests (Provided automatically from websocket)

<br />

Go extra packages:
- gorilla/mux
- gorilla/websocket
- mgo.v2/bson
BurntSushi/toml

<br />

System requirements:
    `Golang`, `MongoDB`
    
<br />
    
How to run:
  1. Clone project,
  2. Enter in project folder,
  2. Download list of dependencies,
  4. Run or configure `go build`: ```go run main.go```,


To do:
- [ ] List all dependencies (Working...)
- [x] Websocket connection handshake.
- [ ] Websocket list of clients 
- [ ] Websocket emit message to clients
- [ ] Websocket emit message to specific client
- [ ] Find solution to resume connection handshake.
- [ ] Automatically close connections when over
- [x] CRUD with database
- [ ] Attach ClientSide 
- [ ] API for ClientSide communication
- [ ] Login/Dashboard Page
- [ ] Dynamically URL inserts and render components
- [ ] Link reactivity between components and database
- [ ] Mount more graphics in painel
- [ ] Export data report