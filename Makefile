start:
	@docker-compose up -d --build

start_with_logs:
	@docker-compose up --build

cleanVolumes:
	@docker-compose down
	@docker volume rm netfix_users netfix_billing netfix_basic netfix_postgresUsers netfix_postgresBasic netfix_postgresBilling netfix_kafka netfix_redis_data