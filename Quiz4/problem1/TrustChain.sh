# Create root certificate 
./cert-go-linux-amd64-v2-3-2 create private-key -o root/root.key.pem
./cert-go-linux-amd64-v2-3-2 create cert -t root -y Cfg.yml

# Create intermediate certificate 
./cert-go-linux-amd64-v2-3-2 create private-key -o intermediate/intermediate.key.pem
./cert-go-linux-amd64-v2-3-2 create csr -t intermediate -y Cfg.yml
./cert-go-linux-amd64-v2-3-2 create cert -t intermediate -y Cfg.yml

# Create server certificate 
./cert-go-linux-amd64-v2-3-2 create private-key -o server/server.key.pem
./cert-go-linux-amd64-v2-3-2 create csr -t server -y Cfg.yml
./cert-go-linux-amd64-v2-3-2 create cert -t server -y Cfg.yml

# Create client certificate
./cert-go-linux-amd64-v2-3-2 create private-key -o client/client.key.pem
./cert-go-linux-amd64-v2-3-2 create csr -t client -y Cfg.yml
./cert-go-linux-amd64-v2-3-2 create cert -t client -y Cfg.yml

