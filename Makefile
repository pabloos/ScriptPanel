include .env

#########
# modes #
#########
.PHONY: test
.SILENT: build clean-data unit-test

########
# vars #
########

## cmds
docker=docker-compose
remdir= yes|rm -R 

## dirs
ldapCertDir=ldapserver/certs/
mongoCertDir=mongoserver/certs/

inputldapCertDir=certs-maker/ldap-certs/
inputmongoCertDir=certs-maker/mongo-certs/

## constant files
composefile=docker-compose.yml

keyfile=key.pem
certfile=cert.pem

###################
# Primary Actions #
###################
build:
	mkdir -p webserver/bin
	$(docker) --file $(composefile) up --build builder
	mv arlequin-latest-linux webserver/bin/

run: docker-clean
	@$(docker) up --build  mongoserver ldapserver ftpserver dockerserver webserver #build just added

run-silent: docker-clean
	@$(docker) up --detach

########
# test #
########
unit-test:
	$(docker) up --build tester
	yes | $(docker) rm tester
	# go tool cover -html=scriptpanel/coverage.out
	rm scriptpanel/coverage.out

	##local:
	# go test -v ./... -coverprofile=coverage.out
	# go tool cover -func=coverage.out
	# #go tool cover -html=coverage.out
	# rm coverage.out

functional-test: run-silent unit-test certs
	@$(docker) --file test/functional/Chrome/$(composefile) up --abort-on-container-exit --force-recreate --build
	@$(MAKE) docker-clean
	@$(docker) --file test/functional/Chrome/$(composefile) down

test: functional-test

clean-data:
	-rm -R fileserver/admin/ ldapserver/config/ ldapserver/database/ mongoserver/data/

docker-stop:
	@$(docker) stop

docker-clean: docker-stop
	@yes | $(docker) rm

#########
# certs #
#########
certs: ldap-certs mongo-certs

clean-ldap-certs:
	$(remdir) $(ldapCertDir)

ldap-certs: clean-ldap-certs
	mkdir -p $(ldapCertDir)
	$(docker) -f certs-maker/$(composefile) up ldap-certs-maker
	mv $(inputldapCertDir)* $(ldapCertDir)
	mv $(ldapCertDir)$(LDAP_HOST)/$(certfile) $(ldapCertDir)
	mv $(ldapCertDir)$(LDAP_HOST)/$(keyfile) $(ldapCertDir)
	$(remdir) $(ldapCertDir)$(LDAP_HOST)
	$(docker) -f certs-maker/$(composefile) up openssl
	mv certs-maker/openssl/dhparam.pem $(ldapCertDir)

# ldap-certs-localy: clean-ldap-certs
# 	mkdir $(ldapCertDir)
# 	cd $(ldapCertDir) && minica -domains scriptpanel.com 
# 	cd $(ldapCertDir)scriptpanel.com && mv $(certfile) ../$(certfile) 
# 	cd $(ldapCertDir)scriptpanel.com && mv $(keyfile) ../$(keyfile)
# 	$(remdir) $(ldapCertDir)scriptpanel.com
# 	openssl dhparam -dsaparam -out $(ldapCertDir)dhparam.pem 4096 

clean-mongo-certs:
	-$(remdir) $(inputmongoCertDir)* $(mongoCertDir)

mongo-certs: clean-mongo-certs
	mkdir -p $(mongoCertDir)
	$(docker) -f certs-maker/$(composefile) up mongo-certs-maker
	mv $(inputmongoCertDir)$(MONGO_HOST)/* $(mongoCertDir)
	mv $(inputmongoCertDir)minica.pem $(mongoCertDir)
	mv $(inputmongoCertDir)minica-key.pem $(mongoCertDir)
	cat $(mongoCertDir)$(keyfile) $(mongoCertDir)$(certfile) > $(mongoCertDir)mongo.pem #mongo needs to get both of them concatenated in the same file
	$(remdir) $(inputmongoCertDir)$(MONGO_HOST)

# mongo-certs-localy: clean-mongo-certs
# 	mkdir -p $(mongoCertDir)
# 	cd $(mongoCertDir) && minica -domains $(MONGO_HOST)
# 	mv $(mongoCertDir)$(MONGO_HOST)/$(certfile) $(mongoCertDir)
# 	mv $(mongoCertDir)$(MONGO_HOST)/$(keyfile) $(mongoCertDir)
# 	cat $(mongoCertDir)$(keyfile) $(mongoCertDir)$(certfile) > $(mongoCertDir)mongo.pem #mongo needs to get both of them concatenated
# 	$(remdir) $(mongoCertDir)$(MONGO_HOST)

############
# API-REST #
############
rest-validate:
	swagger validate ./scriptpanel/cmd/REST/swagger.yml

rest-generate: rest-validate
	swagger generate server \
		--target=./scriptpanel/cmd/REST/ \
		--spec=./scriptpanel/cmd/REST/swagger.yml \
		--exclude-main \
		--name=rest