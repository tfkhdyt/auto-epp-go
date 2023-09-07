PROGRAM_NAME=auto-epp-go

OUTPUT_DIR=./bin
BINARY_DIR=/usr/bin
SERVICE_DIR=/usr/lib/systemd/system

install: build
	install -Dm755 $(OUTPUT_DIR)/$(PROGRAM_NAME) $(BINARY_DIR)/$(PROGRAM_NAME)
	install -Dm644 ./$(PROGRAM_NAME).service -t $(SERVICE_DIR)
	systemctl daemon-reload
	systemctl enable --now $(PROGRAM_NAME).service

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(OUTPUT_DIR)/$(PROGRAM_NAME) .

clean:
	rm -f $(OUTPUT_DIR)/*

uninstall:
	systemctl disable --now $(PROGRAM_NAME).service
	rm -f $(SERVICE_DIR)/$(PROGRAM_NAME).service
	rm -f $(BINARY_DIR)/$(PROGRAM_NAME)

reinstall: uninstall install

.PHONY: install build clean uninstall reinstall
