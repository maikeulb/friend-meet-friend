echo waiting for db
./wait-for-it.sh db:5432

flyway -url=jdbc:postgresql://db/postgres -user=postgres -password=P@ssw0rd! migrate
