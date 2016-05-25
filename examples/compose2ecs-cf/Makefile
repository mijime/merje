aws-cf-template.json: cluster.yml docker-compose.yml compose2ecs-cf.tmpl
	merje docker-compose.yml --format compose2ecs-cf.tmpl | \
		merje cluster.yml - -iyaml -o aws-cf-template.json

test: docker-compose.yml
	docker-compose up -d
