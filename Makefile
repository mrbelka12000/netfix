start:
	@docker-compose down
	@docker-compose up -d --build

cleanVolumes:
	@docker volume rm netfix_auth netfix_basic netfix_postgres