create-gun: 
	curl http://localhost:8080/guns \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","name": "AWP","price": 4500}'