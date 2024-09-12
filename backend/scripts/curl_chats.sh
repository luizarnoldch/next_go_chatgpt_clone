#!/bin/bash

# Base URL for the API
BASE_URL="http://localhost:4000"

# Test GET / (root endpoint)
echo "Testing GET / (root)"
curl -X GET "$BASE_URL/"
echo -e "\n"

# Test GET /chats (Get all chats)
echo "Testing GET /chats"
curl -X GET "$BASE_URL/chats"
echo -e "\n"

# # Test POST /chats (Create a new chat)
# echo "Testing POST /chats (Create Chat)"
# curl -X POST "$BASE_URL/chats" \
#     -H "Content-Type: application/json" \
#     -d '{"title": "New Chat", "description": "This is a new chat"}'
# echo -e "\n"

# # Test GET /chats/:id (Get chat by ID)
# echo "Testing GET /chats/1 (Get Chat by ID)"
# curl -X GET "$BASE_URL/chats/1"
# echo -e "\n"

# # Test PUT /chats/:id (Update chat)
# echo "Testing PUT /chats/1 (Update Chat)"
# curl -X PUT "$BASE_URL/chats/1" \
#     -H "Content-Type: application/json" \
#     -d '{"title": "Updated Chat Title", "description": "Updated description"}'
# echo -e "\n"

# # Test DELETE /chats/:id (Delete chat)
# echo "Testing DELETE /chats/1 (Delete Chat)"
# curl -X DELETE "$BASE_URL/chats/1"
# echo -e "\n"

# # Test POST /upload (Upload file)
# echo "Testing POST /upload (Upload File)"
# curl -X POST "$BASE_URL/upload" \
#     -F "file=@/path/to/your/file.txt"
# echo -e "\n"

# # Test POST /chat/completion (Chat Completion)
# echo "Testing POST /chat/completion"
# curl -X POST "$BASE_URL/chat/completion" \
#     -H "Content-Type: application/json" \
#     -d '{"prompt": "Hello, how are you?", "model": "gpt-3"}'
# echo -e "\n"
