# goLangToAerospike
attempt to spin up an aerospike server, load with test data, and provide access

NOTE: although the webserver is now using docker dns, the tests that run locally still have to change the ip address the webserver will use to reach aerospike untill i find a better discovery method, that ip is in the body of src/webserver/webserver_test.go

to run this test, use the shell scripts to clean.sh and build.sh

you may have to install the aerospike client module

once the webserver is built

    docker-compose up --build -d

will launch the aerospike CE and webserver with client

### some sample curl requests:

this will return the single record written to aerospike on startup

    curl http://localhost:8000/user/11?api_key=42

response should be 200 OK with content:

    {"api_key":"42","company":"mindbodyengineer","first_name":"jonathan","last_name":"wilson"}

untill other data is loaded any other api key will result in 401 response for user 11, other users will return not found

the tests currently only are working while the aero spike container is up

### regarding loading additional test data

i played a little with the aerospike tools docker, the aql command is scripted in the aql-console.sh file, from the aql console adding test records is a matter of:

    insert into test.users(PK,api_key,first_name,last_name,company) values (1,"551","Mick","Jagger","The Rolling Stones")
    insert into test.users(PK,api_key,first_name,last_name,company) values (2,"551","Charlie","Watts","The Rolling Stones")
    insert into test.users(PK,api_key,first_name,last_name,company) values (3,"551","Ronnie","Wood","The Rolling Stones")
    insert into test.users(PK,api_key,first_name,last_name,company) values (4,"551","Keith","Richards","The Rolling Stones")

allowing arbitrary test data load
