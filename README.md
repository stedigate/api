api
===

This is a simple API server that can be used to demonstrate how to build a RESTful API server in Go.


Avalanche Commands
```shell
go run . avalanche balance -c AVAX -w 0x323a6a107667d3c53a9681e9ab4e02d3c7b1E559
go run . avalanche balance -c USDT -w 0x323a6a107667d3c53a9681e9ab4e02d3c7b1E559
go run . avalanche balance -c EURC -w 0x323a6a107667d3c53a9681e9ab4e02d3c7b1E559
go run . avalanche balance -c USDC -w 0x323a6a107667d3c53a9681e9ab4e02d3c7b1E559
go run . avalanche simulate -a 0.001 -c AVAX -d 0x323a6a107667d3c53a9681e9ab4e02d3c7b1E559 -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
go run . avalanche simulate -a 0.01 -c USDC -d 0x323a6a107667d3c53a9681e9ab4e02d3c7b1E559 -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
go run . avalanche transfer -a 0.01 -c AVAX -d 0x323a6a107667d3c53a9681e9ab4e02d3c7b1E559 -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
go run . avalanche transfer -a 0.01 -c USDC -d 0x323a6a107667d3c53a9681e9ab4e02d3c7b1E559 -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
```

Ethereum Commands
```shell
go run . ethereum balance -c ETH -w 0xd3D3a295bE556Cf8cef2a7FF4cda23D22c4627E8
go run . ethereum balance -c USDT -w 0x7f58c6204867584F8B816d25c30007C108E57622
go run . ethereum balance -c USDC -w 0x7f58c6204867584F8B816d25c30007C108E57622
go run . ethereum simulate -a 0.001 -c AVAX -d 0x7f58c6204867584F8B816d25c30007C108E57622 -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
go run . ethereum simulate -a 0.01 -c USDC -d 0x7f58c6204867584F8B816d25c30007C108E57622 -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
go run . ethereum transfer -a 0.01 -c AVAX -d 0x7f58c6204867584F8B816d25c30007C108E57622 -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
go run . ethereum transfer -a 0.01 -c USDC -d 0x7f58c6204867584F8B816d25c30007C108E57622 -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
```

Solana Commands
```shell
go run . solana balance -c USDt -w 9VCJnELK742auACYWcJEEJTbNGebEwNSGE1Ck26YttEJ
go run . solana balance -c USDC -w 9VCJnELK742auACYWcJEEJTbNGebEwNSGE1Ck26YttEJ
go run . solana balance -c SOL -w 9VCJnELK742auACYWcJEEJTbNGebEwNSGE1Ck26YttEJ
go run . solana simulate -a 1 -c SOL -d 6jmPKEMVd3yJWCbpQf1iLgGBVKSKZaX4J6CBqXPTyRvQ -s x3R5X2LC6EW5NsgK8rEVKThjxLwS2FwfAVVD44CR1gN5BJ3GHmPiYKkQSahM9Yd88c4zMUa8GV8Zb4HWrdHTE94
go run . solana transfer -a 1 -c SOL -d 6jmPKEMVd3yJWCbpQf1iLgGBVKSKZaX4J6CBqXPTyRvQ -s x3R5X2LC6EW5NsgK8rEVKThjxLwS2FwfAVVD44CR1gN5BJ3GHmPiYKkQSahM9Yd88c4zMUa8GV8Zb4HWrdHTE94
go run . solana simulate -a 1 -c USDC -d 6jmPKEMVd3yJWCbpQf1iLgGBVKSKZaX4J6CBqXPTyRvQ -s x3R5X2LC6EW5NsgK8rEVKThjxLwS2FwfAVVD44CR1gN5BJ3GHmPiYKkQSahM9Yd88c4zMUa8GV8Zb4HWrdHTE94
go run . solana transfer -a 1 -c USDC -d 6jmPKEMVd3yJWCbpQf1iLgGBVKSKZaX4J6CBqXPTyRvQ -s x3R5X2LC6EW5NsgK8rEVKThjxLwS2FwfAVVD44CR1gN5BJ3GHmPiYKkQSahM9Yd88c4zMUa8GV8Zb4HWrdHTE94

```

Tron Commands
```shell
go run . tron balance -c TRX -w TRGWbaVcixGSQbWzztvQm58n4zti2BitgD
go run . tron balance -c USDT -w TRGWbaVcixGSQbWzztvQm58n4zti2BitgD
go run . tron transfer -a 10 -c TRX -d TRGWbaVcixGSQbWzztvQm58n4zti2BitgD -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
go run . tron transfer -a 10 -c USDT -d TRGWbaVcixGSQbWzztvQm58n4zti2BitgD -s d4c10efce84ed7023abd7193eda45aa300c9ce890746ea5d0fc281bc3f4f5d46
```



Wallet Solana
```
x3R5X2LC6EW5NsgK8rEVKThjxLwS2FwfAVVD44CR1gN5BJ3GHmPiYKkQSahM9Yd88c4zMUa8GV8Zb4HWrdHTE94
9VCJnELK742auACYWcJEEJTbNGebEwNSGE1Ck26YttEJ
```

```
3zxrjd2M9TjXnvHxMyG1fk7rrf2hGKYubYJ3scrRFZC2ihZmBccqwwivqqXgSFf81f9fAGNRikSavXvEGJCD6Axx
6jmPKEMVd3yJWCbpQf1iLgGBVKSKZaX4J6CBqXPTyRvQ
```