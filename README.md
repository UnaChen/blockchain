
### Outline
* [Usage](#usage)
* [Basic terms](#basic-terms)
* [FAQ](#faq)
* [Reference](#reference)
---
### Usage
* Run node server
```sh
# build source code
cd cmd/node
go build

# run node grpc server (default port: 5566)
./node run -p ${PORT}
```

* Run node client
```sh
# run node grpc client related cmd (default server: "127.0.0.1:5566")

# add tx (default from/to: random string)
./node addtx -h
./node addtx -s ${SERVER} -from ${FROM} -to ${TO} -v ${VALUE}


# list related status 
./node list -h
./node status -h

# more info
./node -h
```

* (optional) Run node grpc gateway
```sh
# run node grpc gateway
cd cmd/node_gateway
go build
./node_gateway -grpc ${SERVER} -p ${PORT}

# add tx
curl --location --request POST 'http://localhost:7788/tx/add' \
--header 'Content-Type: application/json' \
--data-raw '{
        "from": "0x71a0DF94Dd9c9390Cf5e0BD87190F111cd55C325",
        "to": "0x0F86Cd82Ff4BA6FbB1655F44EA26b93d150E11E0",
        "value": 100
}'
# list balances
curl http://127.0.0.1:7788/balance/list
# list blocks
curl http://127.0.0.1:7788/block/list
# get node status
curl http://127.0.0.1:7788/node/status

```

---
### Basic terms 
1. 什麼是公私鑰? 公鑰私鑰運作場景是什麼? 什麼是簽名?
    * **公鑰(Public key)**: 能被廣泛的發佈與流傳
    * **私鑰(Private key)**: 不公開, 被妥善保管
    * **運作場景**: </br>
    運作原理是傳送方與接收方在傳送之前，先把彼此的公鑰傳給對方，當傳送方要傳送時，就用接收方的公鑰將訊息加密，接收方收到加密訊息後，再用自己的密鑰解開，這樣即使有心人拿到公鑰，只要沒拿到接收方的私鑰，也還是無法解密訊息。
    * **數位簽章(Digital Signature)**:
        * 由來: 假設有人在網路上找到了接收方的公鑰，假造一封有病毒的訊息，再用接收方公鑰加密傳過去，這樣接收方一打開不就傻傻中計了嗎?!
        * 用途: **Double Confirm**了，也就是傳送方除了使用接收方的公鑰加密外，也使用自己的私鑰對該封加密訊息的Hash簽名

2. 對稱式與非對稱式加密有何不同?
    * **對稱式加密**- 傳送方與接收方的加解密皆使用**同**一把密鑰
    * **非對稱式加密**- 傳送方與接收方加解密用**不同**密鑰, 雙方皆存在一鑰匙 (公鑰＋私鑰）

3. 什麼是 Challenge-Response Authentication?
    * 其中一方提出問題，另一方必須提供有效答案才能進行身份驗證 
    
4. 什麼是多簽 (Multisignature)? 比特幣的多簽地址是什麼?
    * **多簽 Multisignature** </br>
        一種特定類型的數位簽章，而此類型的簽名將允許兩個以上用戶作為一組來簽署文檔。因此，多重簽名則通過多個單一簽名的組合來產生。
    * **Creating a Multisignature Address with Bitcoin-Qt** </br>
        1. Gather (or generate) 3 bitcoin addresses 
            *(RPC cmd: getnewaddress or getaccountaddress)*
        2. Get their public keys.
            *(RPC cmd:  validateaddress)*
        3. Then create a 2-of-3 multisig address using addmultisigaddress; e.g.
        
        ```shell
        # addmultisigaddress returns the multisignature address. 
        > bitcoind addmultisigaddress 2 '["1st pubkey","2nd pubkey","3rd pubkey."]'
        ```
        
5. 什麼是 Hash 函數? 有什麼特性? 什麼是 Proof of Work (PoW)?
    * **Hash Function** </br>
    主要是將不定長度訊息的輸入，演算成固定長度雜湊值的輸出，且所計算出來的值必須符合兩個主要條件：
        * 由雜湊值是無法反推出原來的訊息
        * 雜湊值必須隨明文改變而改變
    * **Proof of Work (PoW)** </br>
    一種對應服務與資源濫用、或是阻斷服務攻擊的經濟對策。一般是要求使用者進行一些耗時適當的複雜運算，並且答案能被服務方快速驗算，以此耗用的時間、裝置與能源做為擔保成本，以確保服務與資源是被真正的需求所使用
* Reference:
    * [基礎密碼學](https://medium.com/@RiverChan/%E5%9F%BA%E7%A4%8E%E5%AF%86%E7%A2%BC%E5%AD%B8-%E5%B0%8D%E7%A8%B1%E5%BC%8F%E8%88%87%E9%9D%9E%E5%B0%8D%E7%A8%B1%E5%BC%8F%E5%8A%A0%E5%AF%86%E6%8A%80%E8%A1%93-de25fd5fa537)
    * [Challenge–Response Authentication - Wikipedia](https://en.wikipedia.org/wiki/Challenge%E2%80%93response_authentication)
    * [authentication](http://systw.net/note/af/sblog/more.php?id=152)
    * [什麼是多重簽名錢包](https://academy.binance.com/zt/security/what-is-a-multisig-wallet)
    * [區塊鏈入門系列 | 比特幣的多重簽名技術 Multisignature](https://www.itread01.com/hkpqcif.html)
    * [Bitcoin Multisignature - Wikipedia](https://en.bitcoin.it/wiki/Multisignature)
    * [雜湊 (Hash)](https://ithelp.ithome.com.tw/articles/10208884)
    * [proof-of-work (PoW) - Wikipedia](https://en.wikipedia.org/wiki/Proof_of_work)
---
### FAQ
1. Compare the difference between account/UTXO based models.</br>

| | Account| UTXO (Unspent Transaction Output)|
|---|---|---|
|原理|||
|Examples|Ethereum|Bitcoin|

2. How to ensure transaction order in an account based model?
3. What is transaction/block?
4. Why is setting block generation time necessary?
5. When to update the account balance?

---
### Reference 
* UTXO 
    * https://medium.com/%E4%B8%80%E5%80%8B%E5%AE%B9%E6%98%93%E5%81%A5%E5%BF%98%E7%9A%84%E5%A4%A7%E5%AD%B8%E7%94%9F/%E4%BB%80%E9%BA%BC%E6%98%AFutxo-40b62e73c84d
    * https://ithelp.ithome.com.tw/articles/10216976
    * https://kknews.cc/zh-tw/tech/g92mrem.html
    * https://ethereum.stackexchange.com/questions/326/what-are-the-pros-and-cons-of-ethereum-balances-vs-utxos
    * https://eth.wiki/en/fundamentals/design-rationale

* 比特幣 區塊鏈 幾種交易標準詳解 P2PKH、P2PK、MS、P2SH加密方式
    * https://codertw.com/%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80/474756/

* 可替代性（fungibility）
    * https://www.blocktempo.com/litecoin-creator-charlie-lee-to-make-coin-more-fungible-and-private-1/


* 什么是Plasma？Plasma Cash？
    * https://www.bishijie.com/shendu/45752.html
    * https://ethfans.org/posts/simplest-introduction-to-ethereum-plasma-mvp-and-plasma-cash
    * https://www.chainnews.com/zh-hant/articles/095190738084.htm

    * [區塊鏈 Blockchain – 創世區塊、區塊、Merkle Tree、Hash](https://www.samsonhoi.com/274/blockchain_genesis_block_merkle_tree)
    * [區塊鏈如何運用merkle tree驗證交易真實性](https://www.itread01.com/content/1548294679.html)
    * https://blog.csdn.net/chunlongyu/article/details/80417356
    * https://kknews.cc/zh-tw/code/yp2ya4n.html
    * https://medium.com/@twedusuck/%E6%AF%94%E7%89%B9%E5%B9%A3-%E4%BB%A5%E5%A4%AA%E5%9D%8A%E7%9A%84%E4%B8%80%E4%BA%9B%E5%95%8F%E9%A1%8C%E4%BB%8B%E7%B4%B9-%E4%BA%8C-bc06a5e7f8fc
* coinbase
    * https://www.chainnews.com/zh-hant/articles/157318076757.htm


* cosmos
    * https://cypherpunks-core.github.io/news/2020/02/13/Cosmos-%E5%8D%80%E5%A1%8A%E9%8F%88%E7%9A%84%E5%B7%A5%E4%BD%9C%E5%8E%9F%E7%90%86-Part-1-%E6%AF%94%E8%BC%83Cosmos-%E8%88%87%E6%AF%94%E7%89%B9%E5%B9%A3-%E4%BB%A5%E5%A4%AA%E5%9D%8A/
https://zhuanlan.zhihu.com/p/38252058

* blockchain database 
    * https://medium.com/@linpoan/%E5%8D%80%E5%A1%8A%E9%8F%88%E7%9A%84%E5%95%8F%E9%A1%8C%E8%88%87%E6%8C%91%E6%88%B0-1be828019f89

* github
    * https://github.com/yeasy/blockchain_guide
    * https://github.com/izqui/blockchain
    * https://github.com/Jeiwan/blockchain_go
    * https://github.com/tendermint/tendermint
    * https://github.com/web3coach/the-blockchain-bar
    * https://github.com/ethereum/go-ethereum.git
    * https://github.com/ZtesoftCS/go-ethereum-code-analysis/blob/master/go-ethereum%E6%BA%90%E7%A0%81%E9%98%85%E8%AF%BB%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA.md
    * https://github.com/miguelmota/ethereum-development-with-go-book/tree/master/code

* others 
    * https://www.freecodecamp.org/news/build-a-blockchain-in-golang-from-scratch/
    * https://www.guru99.com/blockchain-tutorial.html
    * https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc
    *  https://applicature.com/blog/blockchain-technology/blockchain-code-examples
    * https://arxiv.org/pdf/2003.14271.pdf
    * https://cypherpunks-core.github.io/ethereumbook_zh/07.html
    * https://fullstacks.org/materials/ethereumbook/07_transactions.html
    * https://ithelp.ithome.com.tw/articles/10216297
    * https://learnmeabitcoin.com/technical/block-header
    * https://medium.facilelogin.com/the-mystery-behind-block-time-63351e35603a
