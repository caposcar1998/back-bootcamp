migrateup:
	migrate -path migrations -database "mysql://${dbuser}:${dbpass}@tcp(${dbserver}:${dbport})/${dbname}" -verbose up
migratedown:
	migrate -path migrations -database "mysql://${dbuser}:${dbpass}@tcp(${dbserver}:${dbport})/${dbname}" -verbose down