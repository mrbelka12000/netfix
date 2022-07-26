start_in_detach:
	@docker-compose up -d --build

start_in_normal:
	@docker-compose up --build

cleanVolumes:
	@docker-compose down
	@docker volume rm netfix_users netfix_basic netfix_postgresUsers netfix_postgresBasic netfix_postgresBilling