# goLangToAerospike
attempt to spin up an aerospike server, load with test data, and provide access

NOTE: you will have to change the ip address the webserver will use to reach aerospike untill i find a better discovery method, that ip is in the main func of src/webserver/webserver.go

to run this test, use the shell scripts to clean.sh and build.sh

you may have to install the aerospike client module

once the webserver is built

docker-compose up --build -d

will launch the aerospike CE and webserver with client

some sample curl requests:

this will return a record

curl http://localhost:8000/user/11?api_key=42

any other api key will result in 401 response for user 11, other users will return not found

the tests currently only are working while the aero spike container is up

