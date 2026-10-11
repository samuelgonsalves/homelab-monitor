curl -fsS -m 10 -X POST "http://192.168.1.105:8089/backup-status"  -H "Content-Type: application/json"

curl -fsS -m 10 -X POST "http://192.168.1.105:8089/backup-status-update"  -H "Content-Type: application/json"   -d '{"status":true}'
