DYNAMO_TABLE_NAME=ecatrom2000

list-tables:
	aws dynamodb list-tables --endpoint-url http://localhost:4566

create-table:
	aws dynamodb create-table-localstack \
		--table-name ecatrom2000 \
		--attribute-definitions \
			AttributeName=EntryID,AttributeType=S \
		--key-schema \
			AttributeName=EntryID,KeyType=HASH \
		--provisioned-throughput \
			ReadCapacityUnits=5,WriteCapacityUnits=5 \
		--table-class STANDARD \
		--provisioned-throughput \
			ReadCapacityUnits=5,WriteCapacityUnits=5 \
		--endpoint-url \
			http://localhost:4566

create-table-aws:
	aws dynamodb create-table \
		--table-name ecatrom2000 \
		--attribute-definitions \
			AttributeName=key,AttributeType=S \
		--key-schema \
			AttributeName=key,KeyType=HASH \
		--provisioned-throughput \
			ReadCapacityUnits=5,WriteCapacityUnits=5 \
		--table-class STANDARD \
		--provisioned-throughput \
			ReadCapacityUnits=5,WriteCapacityUnits=5 \
		--endpoint-url \
			https://dynamodb.us-east-1.amazonaws.com/