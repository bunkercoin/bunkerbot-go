NAME = bunkerbot-go
BIN_DIR = bin/
GOFLAGS = -v -race
TOKEN = example

all: options $(NAME)

options:
	@echo ${NAME} build options
	@echo GOFLAGS       = ${GOFLAGS}

$(NAME):
	go build $(GOFLAGS) -o $(BIN_DIR)$(NAME)

clean:
	rm -f $(BIN_DIR)$(NAME)

dep:
	go mod download

fmt:
	go fmt

vet:
	go vet

lint:
	revive

check: fmt vet lint

run: $(NAME)
	BOT_TOKEN=$(TOKEN) $(BIN_DIR)$(NAME)

re: clean all

rerun: re run

.PHONY: all options clean dep fmt vet lint check run re rerun
