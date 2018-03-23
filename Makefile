BINARY_SVR = gagt_server
BINARY_CLI = gagt_client

# Build the project
all: clean server client

server:
	cd server; \
	go build -o ../${BINARY_SVR} .; \
	cd - >/dev/null

client:
	cd client; \
	go build -o ../${BINARY_CLI} .; \
	cd - >/dev/null


clean:
	-rm -f ${BINARY_SVR}
	-rm -f ${BINARY_CLI}

.PHONY: server client clean
