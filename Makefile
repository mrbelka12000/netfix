start:
	@docker-compose up -d --build

cleanVolumes:
	@docker-compose down
	@docker volume rm netfix_auth netfix_basic netfix_postgres