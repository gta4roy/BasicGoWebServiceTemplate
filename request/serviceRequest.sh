#!/bin/sh 

url="127.0.0.1:9098"
basepath="${url}/api/v1/directory"

healthCheck(){
	curl -X "GET" http://${url}/health
}

addPhone(){
	curl -X POST "http://${basepath}/phone" -H "Content-Type: application/json" -d '{"name": "Abhijit Roy", "phone": "7829712286" }'
}
getPhone(){
	curl -X GET "http://${basepath}/phone/7829712287" -H "Content-Type: application/json"
}
healthCheck
addPhone
getPhone
