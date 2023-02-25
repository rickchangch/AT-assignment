%:
	@:
ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
PROJECT_NAME = $(shell echo $(shell basename $(PWD)) | tr '[A-Z]' '[a-z]')
GUID := $(USER)_$(PROJECT_NAME)

.PHONY: run kill rerun

run:
	@NETWORK_ADDR_TAIL=$(call ARGS,1) GUID=$(GUID) docker compose -p $(GUID) up -d
kill:
	@GUID=$(GUID) docker compose -p $(GUID) down -v
rerun:
	@GUID=$(GUID) docker compose -p $(GUID) down -v && \
	NETWORK_ADDR_TAIL=$(call ARGS,1) GUID=$(GUID) docker compose -p $(GUID) up -d
clean:
	@docker rm -f $(docker ps -aq)
logs:
	@docker logs -f $(GUID)-app
