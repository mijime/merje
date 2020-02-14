aws-cf-template.json: compose2ecs-cf.tmpl docker-compose.yml cluster.yml
	merje -encode compose2ecs-cf.tmpl docker-compose.yml | \
		merje -decode yaml - cluster.yml | jq . > $@

test: docker-compose.yml
	docker-compose up -d
