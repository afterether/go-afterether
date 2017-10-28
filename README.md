# AfterEther cryptocurrency project

AfterEther is a project to create a copy of Ether (ETH), the currency of Ethereum.
Our currency is named AET and it runs under networkid=233

http://afterether.org

## Building instructions



To build 

1. Download these 2 files:
	Ethereum geth project : https://github.com/ethereum/go-ethereum/archive/v1.7.2.tar.gz
	AfterEther patch : https://github.com/afterether/go-afterether-patch/releases/download/v1.7.2/go-afterether-v1.7.2.diff

2. Uncompress Ethereum:

	tar zxf go-afterether-v1.7.2-bin.tar.gz
	cd go-ethereum

3. Apply AfterEther patch:

	patch -p1 < go-afterether-v1.7.2.diff

4. Build 

	make all

5. Initialize the blockchain

	build/bin/geth init afterether-genesis.json

6. Create an account

	build/bin/geth account new

7. Start `geth`
	
	build/bin/geth console

8. Mine

	\>eth.coinbase=[YOUR ACCOUNT ADDRESS]

	\>miner.start(1)




	
