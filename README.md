# How to test the Item Management APIs

```bash
### Get list of items
curl -X GET http://localhost:8000/items
### Post an item to the list
curl -X POST http://localhost:8000/items/create -H "Content-Type: application/json" -d "{\"name\":\"Monitor\"}"
```
