utxo

https://medium.com/%E4%B8%80%E5%80%8B%E5%AE%B9%E6%98%93%E5%81%A5%E5%BF%98%E7%9A%84%E5%A4%A7%E5%AD%B8%E7%94%9F/%E4%BB%80%E9%BA%BC%E6%98%AFutxo-40b62e73c84d
https://ithelp.ithome.com.tw/articles/10216976
https://kknews.cc/zh-tw/tech/g92mrem.html
https://ethereum.stackexchange.com/questions/326/what-are-the-pros-and-cons-of-ethereum-balances-vs-utxos


https://eth.wiki/en/fundamentals/design-rationale


比特幣 區塊鏈 幾種交易標準詳解 P2PKH、P2PK、MS、P2SH加密方式
https://codertw.com/%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80/474756/

可替代性（fungibility）
https://www.blocktempo.com/litecoin-creator-charlie-lee-to-make-coin-more-fungible-and-private-1/


什么是Plasma？Plasma Cash？
https://www.bishijie.com/shendu/45752.html
https://ethfans.org/posts/simplest-introduction-to-ethereum-plasma-mvp-and-plasma-cash
https://www.chainnews.com/zh-hant/articles/095190738084.htm

區塊鏈 Blockchain – 創世區塊、區塊、Merkle Tree、Hash
https://www.samsonhoi.com/274/blockchain_genesis_block_merkle_tree

區塊鏈如何運用merkle tree驗證交易真實性
https://www.itread01.com/content/1548294679.html

https://blog.csdn.net/chunlongyu/article/details/80417356
https://kknews.cc/zh-tw/code/yp2ya4n.html
https://medium.com/@twedusuck/%E6%AF%94%E7%89%B9%E5%B9%A3-%E4%BB%A5%E5%A4%AA%E5%9D%8A%E7%9A%84%E4%B8%80%E4%BA%9B%E5%95%8F%E9%A1%8C%E4%BB%8B%E7%B4%B9-%E4%BA%8C-bc06a5e7f8fc

cosmos
https://cypherpunks-core.github.io/news/2020/02/13/Cosmos-%E5%8D%80%E5%A1%8A%E9%8F%88%E7%9A%84%E5%B7%A5%E4%BD%9C%E5%8E%9F%E7%90%86-Part-1-%E6%AF%94%E8%BC%83Cosmos-%E8%88%87%E6%AF%94%E7%89%B9%E5%B9%A3-%E4%BB%A5%E5%A4%AA%E5%9D%8A/
https://zhuanlan.zhihu.com/p/38252058


https://www.freecodecamp.org/news/build-a-blockchain-in-golang-from-scratch/

https://www.guru99.com/blockchain-tutorial.html

https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc

https://medium.com/@lhartikk/a-blockchain-in-200-lines-of-code-963cc1cc0e54 
-> too simple, only make you understand simple function what you need to wirite blockchain  
https://applicature.com/blog/blockchain-technology/blockchain-code-examples

https://github.com/yeasy/blockchain_guide
https://github.com/izqui/blockchain
https://github.com/Jeiwan/blockchain_go
https://github.com/tendermint/tendermint
https://github.com/web3coach/the-blockchain-bar


blockchain database 
https://medium.com/@linpoan/%E5%8D%80%E5%A1%8A%E9%8F%88%E7%9A%84%E5%95%8F%E9%A1%8C%E8%88%87%E6%8C%91%E6%88%B0-1be828019f89

https://arxiv.org/pdf/2003.14271.pdf

https://github.com/ZtesoftCS/go-ethereum-code-analysis/blob/master/go-ethereum%E6%BA%90%E7%A0%81%E9%98%85%E8%AF%BB%E7%8E%AF%E5%A2%83%E6%90%AD%E5%BB%BA.md

coinbase
https://www.chainnews.com/zh-hant/articles/157318076757.htm

https://github.com/miguelmota/ethereum-development-with-go-book/tree/master/code

https://cypherpunks-core.github.io/ethereumbook_zh/07.html

https://fullstacks.org/materials/ethereumbook/07_transactions.html

https://ithelp.ithome.com.tw/articles/10216297

1. It is an account based blockchain, which means you don’t have to handle UTXO.
2. Each newly observed account has 100 units of initial coins. For example, when your
blockchain sees a transaction with a new sender or a new recipient, it credits the new
sender or/and the new recipient 100 coins.
3. The blockchain generates a new block every 10 seconds.
4. Each block header has the following information:
    1. Block hash: Define your own hashing, can be something like SHA256(Block
    header)
    2. Parent hash: Previous block hash.
    3. Block height.
    4. Transactions: Array of transaction hashes belong to this block.
5. Each transaction has the following information:
    1. Transaction hash: hash of this transaction.
    2. From: from address.
    3. To: to address.
    4. Value: coins to transact.
6. Input: your program needs to take transactions as input. It can be in a file format,
    interactive console, command line, grpc API, JSON RPC API, or anything you prefer.
    Each transaction input entry contains the following information: from address, to
address, and value.
7. Your program needs to be able to dump blockchain logs somewhere, can be console or
file.
8. In the README file, please include the following message:
1. How to run your program?
2. Design details about your program.
3. All the references you have gone through while working on this assignment.

https://medium.facilelogin.com/the-mystery-behind-block-time-63351e35603a