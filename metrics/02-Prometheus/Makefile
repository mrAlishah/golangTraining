DEVICES_URL := localhost:8080
PROMETHRUS_URL := localhost:8081/metrics
compose-up:
		docker-compose up --build

run:  show-devices show-prometheuse create-devices update-devices login


breakln-%:
	@printf "\n\n -------------------------------------------------- $*\n"
.SECONDEXPANSION:

show-devices: breakln-$$@
	$(info "---> get devices")
	curl ${DEVICES_URL}/devices

show-prometheuse: breakln-$$@
	@echo "--->get prometheuse data"
	curl ${PROMETHRUS_URL}

create-devices: breakln-$$@
	@echo "\n---> create new devices"
	curl -d '{"id": 3, "mac": "96-40-D1-32-D7-1A", "firmware": "3.03.00"}' ${DEVICES_URL}/devices

update-devices: breakln-$$@
	@echo "\n---> update device devices"
	curl -X PUT -d '{"firmware": "2.3.0"}' localhost:8080/devices/1

login: breakln-$$@
	@echo "\n---> login"
	curl localhost:8080/login