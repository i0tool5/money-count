CREATE_USER="CREATE USER tempuser PASSWORD 'temppass'"
DROP_USER="DROP USER tempuser"

database-create:
	docker exec postgre psql -h 127.0.0.1 -U postgres -c $(CREATE_USER)
	docker exec postgre psql -h 127.0.0.1 -U postgres -c "CREATE DATABASE tempdb"
	docker exec postgre psql -h 127.0.0.1 -U tempuser -d tempdb -c "CREATE TABLE payments (id bigserial PRIMARY KEY, user_id bigint not null, date date DEFAULT current_date,type text,description text,amount bigint)"
	docker exec postgre psql -h 127.0.0.1 -U tempuser -d tempdb -c "CREATE TABLE users (id bigserial PRIMARY KEY,username  text NOT NULL UNIQUE,firstname text,lastname  text,password  text)"

database-drop:
	docker exec postgre psql -h 127.0.0.1 -U postgres -c "DROP DATABASE tempdb"
	docker exec postgre psql -h 127.0.0.1 -U postgres -c $(DROP_USER)

