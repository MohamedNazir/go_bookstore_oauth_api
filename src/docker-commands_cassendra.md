docker run --name cassandra -d -p 7199:7199 -p 7000:7000 -p 9042:9042 -p 9160:9160 -p7001:7001 cassandra:3.11.8
docker exec -it cassandra cqlsh -e "CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1}"
docker exec -it cassandra cqlsh
