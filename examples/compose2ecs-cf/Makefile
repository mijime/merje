aws-cf-template.json: cluster.yml docker-compose.yml compose2ecs-cf.tmpl
	merje docker-compose.yml --format compose2ecs-cf.tmpl | \
		merje cluster.yml - -iyaml -o aws-cf-template.json

prod: aws-cf-template.json
	aws cloudformation create-stack \
		--stack-name wordpress --template-body file://$$(pwd)/$<

test: docker-compose.yml
	docker-compose up -d
