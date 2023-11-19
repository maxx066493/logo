curl -L https://github.com/NethermindEth/eigenlayer/releases/download/<VERSION>/eigenlayer-linux-<ARCH> --output eigenlayer
chmod +x eigenlayer
sudo apt update && sudo apt upgrade -y
sudo apt install make clang pkg-config libssl-dev build-essential git gcc chrony curl jq ncdu bsdmainutils htop net-tools lsof fail2ban wget -y
cd ~
wget https://golang.org/dl/go1.21.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile
go version
curl -L https://github.com/NethermindEth/eigenlayer/releases/download/<VERSION>/eigenlayer-linux-<ARCH> --output eigenlayer
chmod +x eigenlayer
curl -L https://github.com/NethermindEth/eigenlayer/releases/download/0.4.2/eigenlayer-linux-amd64 --output eigenlayer
chmod +x eigenlayer
curl -L https://github.com/NethermindEth/eigenlayer/releases/download/v0.4.3/eigenlayer-linux-amd64 --output eigenlayer
chmod +x eigenlayer
curl -L https://github.com/NethermindEth/eigenlayer/releases/download/v0.4.3/eigenlayer-linux-arm64 --output eigenlayer
chmod +x eigenlayer
curl -L https://github.com/NethermindEth/eigenlayer/releases/download/v0.4.3/eigenlayer-linux-amd64-ubuntu-20-04 --output eigenlayer
chmod +x eigenlayer
curl -L https://github.com/NethermindEth/eigenlayer/releases/download/v0.4.3/eigenlayer-linux-arm64-ubuntu-20-04 --output eigenlayer
chmod +x eigenlayer
source ~/.profile
echo $GOBIN
export GOBIN=$GOPATH/bin
export PATH=$GOBIN:$PATH
source ~/.profile
echo $GOBIN
go install github.com/NethermindEth/eigenlayer/cmd/eigenlayer@latest
git clone https://github.com/NethermindEth/eigenlayer.git
cd eigenlayer
ls
sudo find / -type d -name "eigenlayer"
cd /root/go/pkg/mod/github.com/\!nethermind\!eth/eigenlayer@v0.4.3/cmd/eigenlayer
go build -o eigenlayer main.go
./eigenlayer
eigenlayer
eigenlayer operator
cd ~/eigenlayer  
eigenlayer operator keys create --key-type ecdsa TEST1
eigenlayer operator keys create --key-type bls test1
eigenlayer operator keys list
eigenlayer operator config create
eigenlayer operator status operator.yaml
eigenlayer operator status operator-config.yaml
eigenlayer operator update operator-config.yaml
