COMMON = $(wildcard  */docker-compose.*)

all: docker-compose-dev.yml docker-compose-prd.yml

docker-compose-%.yml: $(COMMON) env-%/*/docker-compose.*
	merje -o $@ $^
