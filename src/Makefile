
servers = gateway account game chat
.PHONY: clear
clear:
	for server in $(servers); do \
  		rm -rf ../bin/$$server; \
		echo $$server; \
	done
	rm -rf ../bin/logs

.PHONY: build
build:
	go build -o ../bin/gateway ./cmd/gateway/main.go
	go build -o ../bin/account ./cmd/account/main.go
	go build -o ../bin/game ./cmd/game/main.go
	go build -o ../bin/chat ./cmd/chat/main.go

.PHONY: logs
logs:
	mkdir -p ../bin/logs/gateway
	mkdir -p ../bin/logs/account
	mkdir -p ../bin/logs/game
	mkdir -p ../bin/logs/chat

.PHONY: startup
startup:logs
	./build/linux/startup.sh

.PHONY: shutdown
shutdown:
	./build/linux/shutdown.sh
