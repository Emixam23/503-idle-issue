.PHONY: all

all: install-proxies start-envoy install-oha run-api test
	@echo "✅ All done."


run-test: stop run-api start-envoy test

install-proxies: install-envoy

install-envoy:
	wget -O envoy https://github.com/envoyproxy/envoy/releases/download/v1.34.1/envoy-1.34.1-linux-x86_64 && chmod +x envoy
	@echo "✅ Proxies installed"

start-envoy:
	./envoy -c ./proxies/envoy.yaml
	@echo "✅ Envoy started"

install-oha:
	wget -O oha https://github.com/hatoo/oha/releases/download/v1.9.0/oha-macos-arm64 && chmod +x oha
	@echo "✅ Oha installed"

stop:
	$(shell zsh kill_port_process 3000)
	$(shell zsh kill_port_process 8080)

run-api:
	go run main.go

test:
	./oha -c 100 -z 5m --burst-delay 10s --burst-rate 100 http://localhost:8080/
