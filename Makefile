run:
	run go main.go
update:
	sh update_books.sh
checkout:
	curl "localhost:8080/checkout?id=1" --request "PATCH"
return:
	curl "localhost:8080/return?id=1" --request "PATCH"