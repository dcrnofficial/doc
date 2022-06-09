# Insight API

The Insight API is accessible via HTTP via REST or WebSocket. 

To call the REST API, use the `/insight/api` path prefix. To call the Websocket API, use the `/insight/socket.io` path prefix.

POST methods require parameters to be passed in JSON objects.

# Endpoints

Below are the implemented Insight API endpoints and associated methods.

  * [Blocks](#blocks)
    + [/block/](#block)
    + [/block-index/](#block-index)
    + [/rawblock/ (hash)](#rawblock-hash)
    + [/rawblock/ (height)](#rawblock-height)
    + [/blocks/](#blocks)
  * [Transactions](#transactions)
    + [/tx/](#tx)
    + [/rawtx/](#rawtx)
    + [/txs/ (block)](#txs-block)
    + [/txs/ (address)](#txs-address)
    + [/tx/send/ (POST)](#txsend-post)
  * [Addresses](#addresses)
    + [/addr/](#addr)
    + [/addr/ (balance)](#addr-balance)
    + [/addr/ (totalSent)](#addr-totalsent)
    + [/addr/ (totalReceived)](#addr-totalreceived)
    + [/addr/ (unconfirmedBalance)](#addr-unconfirmedbalance)
    + [/addr/utxo/](#addrutxo)
    + [/addrs/utxo/](#addrsutxo)
    + [/addrs/utxo/ (POST)](#addrsutxo-post)
    + [/addrs/txs/](#addrstxs)
    + [/addrs/txs/ (POST)](#addrstxs-post)
  * [Status and Utility](#status-and-utility)
    + [/sync](#sync)
    + [/peer](#peer)
    + [/status](#status)
    + [/estimatefee](#estimatefee)


## Blocks 

Methods that work with blocks. 

### /block/

**URL:**  ```GET /block/{hash}```

**Description:** Retrieves summary of block by hash. 

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| hash              | `string`      |   Block hash       |  

**Request Example:**

```GET /block/00000000000000021afd116007a8aa68388992a76694adcd8f0bd49860b2f26b```

**Request Response:**

```
[
    {
        "hash": "00000000000000021afd116007a8aa68388992a76694adcd8f0bd49860b2f26b",
        "confirmations": 22,
        "size": 360,
        "height": 255,
        "version": 7,
        "merkleroot": "35135bdd693031f52f256c7ff8a44fb636762d7384cb37895e4e555626b55904",
        "tx": [
            "f2c0a8cec33ff6b7a0e39a56894a354d0f4b18cefec1a28eb0500b35edea153a"
        ],
        "time": 1653467706,
        "nonce": 407694534,
        "bits": "1a0a0000",
        "difficulty": 1677696,
        "previousblockhash": "000000000000016ced035399585356a20d40dff0e11a1b272f846947a97516d3",
        "nextblockhash": "00000000000006138b579b50d57c3b540139c0b9232b4dcc215473b6b01a1789",
        "reward": 7,
        "isMainChain": true
    }
]
```
<br/>

### /block-index/

**URL:**  ```GET /block-index/{height}```

**Description:** Retrieves hash of a block by block height.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| hash              | `"int64"`      |    Block hash       |  

**Request Example:**

```GET /block-index/255```

**Request Response:**

```
{
    "blockHash": "00000000000000021afd116007a8aa68388992a76694adcd8f0bd49860b2f26b"
}
```


<br/>

### /rawblock/ (hash)

**URL:**  ```GET /rawblock/{hash}```

**Description:** Retrieves raw block data by hash. 

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| hash              | `string`      |  Block hash         |  

**Request Example:**

```GET /rawblock/00000000000000021afd116007a8aa68388992a76694adcd8f0bd49860b2f26b ```

**Request Response:**

```
{
    "rawblock": "07000000d31675a94769842f271b1ae1f0df400da2565358995303ed6c010000000000000459b52656554e5e8937cb84732d7636b64fa4f87f6c252ff5313069dd5b1335643a4b2dd2c7e3d28458b7d20885b0a41916e1311f6e846f5c0f87711c050a350100000000000000000000000000000000000a1a00c2eb0b00000000ff000000680100003aea8d62c6ec4c18497a1c008857555b000000000000000000000000000000000000000000000000000000000101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff00ffffffff0300e1f5050000000000001976a91429dbcdffde4bf456b9df612b0916cc331ad4720288ac000000000000000000000e6a0cff0000001358cbece75428b60046c3230000000000001976a91462427fdc6c1afe612a391c9efaeacfcbfbb3b9db88ac0000000000000000010027b9290000000000000000ffffffff0800002f646372642f00"
}
```

<br/>

### /rawblock/ (height)

**URL:**  ```GET /rawblock/{height}```

**Description:** Retrieves raw block data by block height.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| height              | `"int64"`      |   Block height       |  

**Request Example:**

```GET /rawblock/255 ```

**Request Response:**

```
{
    "rawblock": "07000000d31675a94769842f271b1ae1f0df400da2565358995303ed6c010000000000000459b52656554e5e8937cb84732d7636b64fa4f87f6c252ff5313069dd5b1335643a4b2dd2c7e3d28458b7d20885b0a41916e1311f6e846f5c0f87711c050a350100000000000000000000000000000000000a1a00c2eb0b00000000ff000000680100003aea8d62c6ec4c18497a1c008857555b000000000000000000000000000000000000000000000000000000000101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff00ffffffff0300e1f5050000000000001976a91429dbcdffde4bf456b9df612b0916cc331ad4720288ac000000000000000000000e6a0cff0000001358cbece75428b60046c3230000000000001976a91462427fdc6c1afe612a391c9efaeacfcbfbb3b9db88ac0000000000000000010027b9290000000000000000ffffffff0800002f646372642f00"
}
```

<br/>

### /blocks/

**URL:**  ```GET /blocks```

**Description:** Retrieves summaries of blocks by time. 

**Parameters:**

| Parameter           | Type                   | Description                   | 
| -------------------- | ---------------------- | ---------------------- |
| limit              | `"int64"`  |  (optional) Maximum number of blocks to return. If `limit` is not specified (or set to '0'), blocks for the  24 hours after `blockDate` will be returned.   |  
| blockDate              | `"int64"`      |  (optional) Date to start searching for blocks (YYYY-MM-DD). If `blockDate` is not specified, blockDate defaults to current date.      |  

**Request Example:**

```GET /blocks?limit=5&blockDate=2022-05-25```

**Request Response:**

```
{
    "blocks": [
        {
            "height": 288,
            "size": 10310,
            "hash": "00000000000000e8ce8059f99111025f72c4b8300639649ec747d04a45da55b4",
            "time": "2022-05-25T09:14:49Z",
            "txlength": 23
        },
        {
            "height": 287,
            "size": 10310,
            "hash": "00000000000001aebd1008ab070498abea7f44c0d37ede1b40a28fcd99a6a594",
            "time": "2022-05-25T09:13:21Z",
            "txlength": 23
        },
        {
            "height": 286,
            "size": 10310,
            "hash": "0000000000000715fbbb0dc2fea7b39f912a2296be13ef0df0d66b7ecb4398f6",
            "time": "2022-05-25T09:09:33Z",
            "txlength": 23
        },
        {
            "height": 285,
            "size": 6280,
            "hash": "00000000000009a0dc69c854efa1660ab3d83d0775b35e6578ea69fc23b37431",
            "time": "2022-05-25T09:08:32Z",
            "txlength": 21
        },
        {
            "height": 284,
            "size": 10311,
            "hash": "00000000000002df5757d9657532ea50add14c9361497cc60497464cda5b5a44",
            "time": "2022-05-25T09:08:09Z",
            "txlength": 23
        }
    ],
    "length": 5,
    "pagination": {
        "next": "2022-05-26",
        "prev": "2022-05-24",
        "currentTs": 1653523199,
        "current": "2022-05-25",
        "isToday": true,
        "more": true
    }
}
```

<br/>

## Transactions 

Methods that work with transactions.

### /tx/

**URL:**  ```GET /tx/{hash}```

**Description:** Retrieves transaction data by transaction hash (txid). 

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| hash              | `string`      |  Transaction hash (txid) of transaction        |  

**Request Example:**

```GET /tx/ff56ce33b326c7233c0517d52c94ba226c038772512c0282c66a488583bb6b23 ```

**Request Response:**

```
{
    "txid": "ff56ce33b326c7233c0517d52c94ba226c038772512c0282c66a488583bb6b23",
    "version": 1,
    "locktime": 0,
    "vin": [
        {
            "txid": "564e3c546078365f4045bc61e0519cc63f53e8e73685e84d0cf406469fb07f71",
            "vout": 31,
            "sequence": 4294967295,
            "n": 0,
            "scriptSig": {
                "hex": "46304302205497e630309cec08ccdeaafacefbc508aadcafd177694e5647b9f04dcf5efe57021f0a06f61893b30c878e5520753741b9d89b3571c5e437df08813854f96d0f710121030480c2768c70cb0879cc820f854c284b111b6b150fa35cc62a5a2050129fb565",
                "asm": "304302205497e630309cec08ccdeaafacefbc508aadcafd177694e5647b9f04dcf5efe57021f0a06f61893b30c878e5520753741b9d89b3571c5e437df08813854f96d0f7101 030480c2768c70cb0879cc820f854c284b111b6b150fa35cc62a5a2050129fb565"
            },
            "addr": "Dsnd2jV4AYnNfdgUYEhix4UMMK95fQp7ASF",
            "valueSat": 200002980,
            "value": 2.0000298
        }
    ],
    "vout": [
        {
            "value": 2,
            "n": 0,
            "scriptPubKey": {
                "hex": "ba76a914fdba1349ce643e8ac4b97f6bdce68d9fabe5223b88ac",
                "asm": "OP_SSTX OP_DUP OP_HASH160 fdba1349ce643e8ac4b97f6bdce68d9fabe5223b OP_EQUALVERIFY OP_CHECKSIG",
                "addresses": [
                    "Dsp6VMhrb25D4kDaq9pKDyswmTepNoWyQhH"
                ],
                "type": "stakesubmission"
            },
            "spentTxId": null,
            "spentIndex": null,
            "spentHeight": null
        },
        {
            "value": 0,
            "n": 1,
            "scriptPubKey": {
                "hex": "6a1e9b0522325b5497548641f134834c962b0333f8a3a4cdeb0b000000000058",
                "asm": "OP_RETURN 9b0522325b5497548641f134834c962b0333f8a3a4cdeb0b000000000058",
                "addresses": [
                    "Dsf6aNrMEz9ro5SQHBYBKtRohXJVdVeycwp"
                ],
                "type": "sstxcommitment"
            },
            "spentTxId": null,
            "spentIndex": null,
            "spentHeight": null
        },
        {
            "value": 0,
            "n": 2,
            "scriptPubKey": {
                "hex": "bd76a914000000000000000000000000000000000000000088ac",
                "asm": "OP_SSTXCHANGE OP_DUP OP_HASH160 0000000000000000000000000000000000000000 OP_EQUALVERIFY OP_CHECKSIG",
                "addresses": [
                    "DsQxuVRvS4eaJ42dhQEsCXauMWjvopWgrVg"
                ],
                "type": "sstxchange"
            },
            "spentTxId": null,
            "spentIndex": null,
            "spentHeight": null
        }
    ],
    "blockhash": "00000000000002a4d504b08244d63caea003faf617b916706b38cca6bf23e52f",
    "blockheight": 290,
    "confirmations": 1,
    "time": 1653470417,
    "blocktime": 1653470417,
    "valueOut": 2,
    "size": 295,
    "valueIn": 2.0000298,
    "fees": 0.0000298
}
```

<br/>

### /rawtx/

**URL:**  ```GET /rawtx/{hash}```

**Description:** Retrieves raw transaction data by transaction hash (txid). 

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| hash              | `string`      |   Transaction hash (txid) of transaction        |  

**Request Example:**

```GET rawtx/ff56ce33b326c7233c0517d52c94ba226c038772512c0282c66a488583bb6b23 ```

**Request Response:**

```
{
    "rawtx": "0100000001717fb09f4606f40c4de88536e7e8533fc69c51e061bc45405f367860543c4e561f00000000ffffffff0300c2eb0b0000000000001aba76a914fdba1349ce643e8ac4b97f6bdce68d9fabe5223b88ac00000000000000000000206a1e9b0522325b5497548641f134834c962b0333f8a3a4cdeb0b000000000058000000000000000000001abd76a914000000000000000000000000000000000000000088ac000000000000000001a4cdeb0b0000000021010000020000006946304302205497e630309cec08ccdeaafacefbc508aadcafd177694e5647b9f04dcf5efe57021f0a06f61893b30c878e5520753741b9d89b3571c5e437df08813854f96d0f710121030480c2768c70cb0879cc820f854c284b111b6b150fa35cc62a5a2050129fb565"
}
```
<br/>

### /txs/ (block) 

**URL:**  ```GET /txs``` 

**Description:** Retrieves all transactions in a block by block hash.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| block              | `string`      |   Block hash        |  

**Request Example:**

```GET /txs?block=00000000000000021afd116007a8aa68388992a76694adcd8f0bd49860b2f26b ```

**Request Response:**

```
{
    "pagesTotal": 1,
    "txs": [
        {
            "txid": "f2c0a8cec33ff6b7a0e39a56894a354d0f4b18cefec1a28eb0500b35edea153a",
            "version": 1,
            "locktime": 0,
            "isCoinBase": true,
            "vin": [
                {
                    "vout": 0,
                    "sequence": 4294967295,
                    "n": 0,
                    "scriptSig": {},
                    "valueSat": 700000000,
                    "value": 7,
                    "coinbase": "00002f646372642f"
                }
            ],
            "vout": [
                {
                    "value": 1,
                    "n": 0,
                    "scriptPubKey": {
                        "hex": "76a91429dbcdffde4bf456b9df612b0916cc331ad4720288ac",
                        "asm": "OP_DUP OP_HASH160 29dbcdffde4bf456b9df612b0916cc331ad47202 OP_EQUALVERIFY OP_CHECKSIG",
                        "addresses": [
                            "DsUnEWbLXrqMWTEMkEGsUC9R1mP6p3FDCRN"
                        ],
                        "type": "pubkeyhash"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                },
                {
                    "value": 0,
                    "n": 1,
                    "scriptPubKey": {
                        "hex": "6a0cff0000001358cbece75428b6",
                        "asm": "OP_RETURN ff0000001358cbece75428b6",
                        "type": "nulldata"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                },
                {
                    "value": 6,
                    "n": 2,
                    "scriptPubKey": {
                        "hex": "76a91462427fdc6c1afe612a391c9efaeacfcbfbb3b9db88ac",
                        "asm": "OP_DUP OP_HASH160 62427fdc6c1afe612a391c9efaeacfcbfbb3b9db OP_EQUALVERIFY OP_CHECKSIG",
                        "addresses": [
                            "DsZvTNhA34Dx7R9mz8W5J5MwzGZa2MDmine"
                        ],
                        "type": "pubkeyhash"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                }
            ],
            "blockhash": "00000000000000021afd116007a8aa68388992a76694adcd8f0bd49860b2f26b",
            "blockheight": 255,
            "confirmations": 38,
            "time": 1653467706,
            "blocktime": 1653467706,
            "valueOut": 7,
            "size": 178,
            "valueIn": 7
        }
    ]
}
```

<br/>

### /txs/ (address)

**URL:**  ```GET /txs```

**Description:** Retrieves all transactions by address. Transactions are returned by time in descending order.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address    |  

**Request Example:**

```GET /txs?address=DseRvgwcGeCJJC4AUWimHSj8YzRMJ4ryvBM```

**Request Response:**

```
{
    "pagesTotal": 0,
    "txs": []
}
```

<br/>

### /tx/send/ (POST)

**URL:**  ```POST /tx/send ``` 

**Description:** Broadcasts transaction to network.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| rawtx              | `string`      |   Signed transaction as hex string       |  


**Request Example:**

```
curl -X POST \
 https://data.dcrn.xyz/insight/api/tx/send \
 -H 'Cache-Control: no-cache' \
 -H 'Content-Type: application/json' \
 -d '{"rawtx":"010000000184de71690b97b4cbac6e723570a25a56295b30aa26345bfc40c0609c87e23f1d0100000000ffffffff02809698000000000000001976a9143eb656115197956125365348c542e37b6d3d259988ac00811b2c0000000000001976a914e4c9ada86ba67b2c082fac1aa09fb17fccc6833688ac000000000000000001ffffffffffffffff00000000ffffffff6b483045022100b5020004e60a4d26c99a00316ed4f51131b99838fd8fab936d55f9475719bb3b022063b86596adba004367a138f62c73d59f8f3b2cded4b10b5fc0ebe92a4cf41c840121029ec6a82a9646c090decf20806029f332cc11a5c6fa17c943ee72c31b9707a433"}'
```
**Request Response:**
```
{
    "txid": "18a4eeed058c2266512863d03f79651cd38d94d6d682ae3f1e4aad0178c6998f"
}
```
<br/>

## Addresses 

Methods that work with addresses.

### /addr/

**URL:**  ```GET /addr/{address}```

**Description:** Retrieves transactions by addresses. Can optionally return transactions from `N` to `M`.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address    |  
| from              | `int64`      |   (optional) Starting transaction index    |  
| to              | `int64`      |   (optional) Ending transaction index    |  
| noTxList              | `boolean`      |   (optional) If `noTxList` = '1', response will not include a list of txids     |  

**Request Example:**

```GET /addr/DshRLMBrqyDmMpp9gcz3bLgCnjMa75mrHsS?from=100&to=200?noTxList=1 ```

**Request Response:**

```
{
    "addrStr": "DshRLMBrqyDmMpp9gcz3bLgCnjMa75mrHsS",
    "balance": 0,
    "balanceSat": 0,
    "totalReceived": 0,
    "totalReceivedSat": 0,
    "totalSent": 0,
    "totalSentSat": 0,
    "unconfirmedBalance": 0,
    "unconfirmedBalanceSat": 0,
    "unconfirmedTxApperances": 0,
    "txApperances": 0
}
```

<br/>

### /addr/ (balance)

**URL:**  ```GET /addr/{address}/balance```

**Description:** Retrieves address balance in atoms (the smallest unit of Decred; 1 DCR = 100,000,000 atoms).

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address    |  

**Request Example:**

```GET /addr/Dcur2mcGjmENx4DhNqDctW5wJCVyT3Qeqkx/balance ```

**Request Response:**

```
0
```

<br/>

### /addr/ (totalSent)

**URL:**  ```GET /addr/{address}/totalSent```

**Description:** Retrieves total amount sent from an address in atoms (the smallest unit of Decred; 1 DCR = 100,000,000 atoms)

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address  |  

**Request Example:**

```GET addr/Dcur2mcGjmENx4DhNqDctW5wJCVyT3Qeqkx/totalSent ```

**Request Response:**

```
0
```

<br/>

### /addr/ (totalReceived)

**URL:**  ```GET /addr/{address}/totalReceived```

**Description:** Retrieves total amount received by an address.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address  |  

**Request Example:**

```GET /addr/Dcur2mcGjmENx4DhNqDctW5wJCVyT3Qeqkx/totalReceived ```

**Request Response:**

```
0
```

<br/>

### /addr/ (unconfirmedBalance)

**URL:**  ```GET /addr/{address}/unconfirmedBalance```

**Description:** Retrieves unconfirmed balance for an address.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address |  

**Request Example:**

```GET /addr/DseRvgwcGeCJJC4AUWimHSj8YzRMJ4ryvBM/unconfirmedBalance```

**Request Response:**

```
0
```

<br/>

### /addr/utxo

**URL:**  ```GET /addr/{address}/utxo```

**Description:** Retrieves Unspent Transaction Outputs (UTXO) for an address.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address |  


**Request Example:**

```GET /addr/TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9/utxo```

**Request Response:**

```
[
    {
        "address": "TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9",
        "txid": "8dda237797440e747b3451ed586dbb0a6ea8a6fc16638315d67332246c7fbbb6",
        "vout": 0,
        "ts": 1653535363,
        "scriptPubKey": "76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac",
        "height": 569,
        "amount": 0.8,
        "satoshis": 80000000,
        "confirmations": 1
    },
    {
        "address": "TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9",
        "txid": "2db865d62c9fdee15a995d7da6f820d3cfc08b2b595a23dabaf6f9f627b2d9f2",
        "vout": 2,
        "ts": 1653535263,
        "scriptPubKey": "76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac",
        "height": 568,
        "amount": 6.00431682,
        "satoshis": 600431682,
        "confirmations": 2
    }
]
```

<br/>

### /addrs/utxo

**URL:**  ```GET /addrs/{addr0, addr1, ...}/utxo```

**Description:** Retrieves Unspent Transaction Outputs (UTXO) for multiple addresses.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address |  

**Request Example:**

```GET /addrs/TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9,Tsd7xGcyjcMaTG1yujW7vk4BaVwkco4DYn7/utxo ```

**Request Response:**

```
[
    {
        "address": "TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9",
        "txid": "13acbb6fa4ea76cc63c2e4c41589f32c56fa32499d03c1fcb8c55380c99fd6e3",
        "vout": 0,
        "ts": 1653460090,
        "scriptPubKey": "76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac",
        "height": 92,
        "amount": 1,
        "satoshis": 100000000,
        "confirmations": 480
    },
    {
        "address": "Tsd7xGcyjcMaTG1yujW7vk4BaVwkco4DYn7",
        "txid": "870524f9891e1f0836e2569e1375aa2a65617cc2f1a05cd29494e5ca594d58a3",
        "vout": 2,
        "ts": 1653449640,
        "scriptPubKey": "76a91484bb7c5b590e5d31dc11ece3159d28864d66401688ac",
        "height": 1,
        "amount": 6,
        "satoshis": 600000000,
        "confirmations": 571
    }
]
```

<br/>

### /addrs/utxo (POST)

**URL:**  ```POST /addrs/utxo``` 

**Description:** Retrieves Unspent Transaction Outputs (UTXO) for multiple addresses.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address |  


**Request Example:**

```
curl -X POST \
https://alpha.dcrdata.org/insight/api/addrs/utxo \
 -H 'Cache-Control: no-cache' \
 -H 'Content-Type: application/json' \
 -d '{"addrs":"TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9,Tsd7xGcyjcMaTG1yujW7vk4BaVwkco4DYn7"}'
```

**Request Response:**

```
[
      {
        "address": "TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9",
        "txid": "13acbb6fa4ea76cc63c2e4c41589f32c56fa32499d03c1fcb8c55380c99fd6e3",
        "vout": 0,
        "ts": 1653460090,
        "scriptPubKey": "76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac",
        "height": 92,
        "amount": 1,
        "satoshis": 100000000,
        "confirmations": 480
    },
    {
        "address": "Tsd7xGcyjcMaTG1yujW7vk4BaVwkco4DYn7",
        "txid": "870524f9891e1f0836e2569e1375aa2a65617cc2f1a05cd29494e5ca594d58a3",
        "vout": 2,
        "ts": 1653449640,
        "scriptPubKey": "76a91484bb7c5b590e5d31dc11ece3159d28864d66401688ac",
        "height": 1,
        "amount": 6,
        "satoshis": 600000000,
        "confirmations": 571
    }
]
```

<br/>


### /addrs/txs 

**URL:**  ```GET /addrs/{addr0, addr1, ...}/txs ``` 

**Description:** Retrieves transactions for multiple addresses. Transactions are sorted in descending order, from the most recent to the oldest. For example, setting `from` = '2' and `to` = '5', will return the second most recent transaction to the fifth most recent transaction for each address provided. 

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address |  
| from              | `int64`      |   (optional) Starting transaction index  |  
| to              | `int64`      |   (optional) Ending transaction index |  

**Request Example:**

```GET /addrs/TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9,Tsd7xGcyjcMaTG1yujW7vk4BaVwkco4DYn7/txs?from=2&to=3```

**Request Response:**

```
{
    "totalItems": 881,
    "from": 2,
    "to": 3,
    "items": [
        {
            "txid": "2e1547224e347390d78e5135789f1efca15e49aeddd44fb7aa794deb19679809",
            "version": 1,
            "locktime": 0,
            "isCoinBase": true,
            "vin": [
                {
                    "vout": 0,
                    "sequence": 4294967295,
                    "n": 0,
                    "scriptSig": {},
                    "valueSat": 700000000,
                    "value": 7,
                    "coinbase": "00002f646372642f"
                }
            ],
            "vout": [
                {
                    "value": 1,
                    "n": 0,
                    "scriptPubKey": {
                        "hex": "76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac",
                        "asm": "OP_DUP OP_HASH160 f4f3e728ea60e5ce3fa4a32958925a67cf6c221c OP_EQUALVERIFY OP_CHECKSIG",
                        "addresses": [
                            "TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9"
                        ],
                        "type": "pubkeyhash"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                },
                {
                    "value": 0,
                    "n": 1,
                    "scriptPubKey": {
                        "hex": "6a0c400200002ad8c519aada086d",
                        "asm": "OP_RETURN 400200002ad8c519aada086d",
                        "type": "nulldata"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                },
                {
                    "value": 6.00122331,
                    "n": 2,
                    "scriptPubKey": {
                        "hex": "76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac",
                        "asm": "OP_DUP OP_HASH160 f4f3e728ea60e5ce3fa4a32958925a67cf6c221c OP_EQUALVERIFY OP_CHECKSIG",
                        "addresses": [
                            "TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9"
                        ],
                        "type": "pubkeyhash"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                }
            ],
            "blockhash": "0000000a45e4971b071134b4d2d7cbc8b9283c60548cc63b2496abd457f0dd03",
            "blockheight": 576,
            "confirmations": 2,
            "blocktime": 1653536107,
            "valueOut": 7.00122331,
            "size": 178,
            "valueIn": 7
        }
    ]
}
```
<br/>


### /addrs/txs/ (POST)

**URL:**  ```POST /addrs/txs```

**Description:** Retrieves transactions for multiple addresses.

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| address              | `string`      |   Address |  
| from              | `int64`      |   (optional) Starting transaction index|  
| to              | `int64`      |   (optional) Ending transaction index|  
| noScriptSig              | `boolean`      |   (optional) If `noScriptSig` = '1', omits ScriptSig from all inputs |  
| noSpent              | `boolean`      |   (optional) If `noSpent` = '1', omits spend information per output |  
| noAsm              | `boolean`      |   (optional) If `noAsm` = '1', omits script asm from results |  


**Request Example:**

```
curl -X POST \
https://data.dcrn.xyz/insight/api/addrs/txs\
 -H 'Cache-Control: no-cache' \
 -H 'Content-Type: application/json' \
 -d '{"addrs":"TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9,Tsd7xGcyjcMaTG1yujW7vk4BaVwkco4DYn7","from":"1","to":"2","noSpent":"1","noScriptSig":"1"}'
 ```

**Request Response:**

```
{
    "totalItems": 951,
    "from": 1,
    "to": 2,
    "items": [
        {
            "txid": "fc0acaf5a9b9abbf01daa52ea181c2ecc68ea4f47b72e31e89b7991c98b53d49",
            "version": 1,
            "locktime": 0,
            "isCoinBase": true,
            "vin": [
                {
                    "vout": 0,
                    "sequence": 4294967295,
                    "n": 0,
                    "valueSat": 700000000,
                    "value": 7,
                    "coinbase": "00002f646372642f"
                }
            ],
            "vout": [
                {
                    "value": 1,
                    "n": 0,
                    "scriptPubKey": {
                        "hex": "76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac",
                        "addresses": [
                            "TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9"
                        ],
                        "type": "pubkeyhash"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                },
                {
                    "value": 0,
                    "n": 1,
                    "scriptPubKey": {
                        "hex": "6a0c750200008a59729f3d44b6b7",
                        "type": "nulldata"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                },
                {
                    "value": 6.00010612,
                    "n": 2,
                    "scriptPubKey": {
                        "hex": "76a914f4f3e728ea60e5ce3fa4a32958925a67cf6c221c88ac",
                        "addresses": [
                            "TsoMKY3P4yXVg7a7Q5em119BFrhCCtr1HT9"
                        ],
                        "type": "pubkeyhash"
                    },
                    "spentTxId": null,
                    "spentIndex": null,
                    "spentHeight": null
                }
            ],
            "blockhash": "00000025e785f5edd652809e11584ec56359557ac7467f5588bd733c2db2d933",
            "blockheight": 629,
            "confirmations": 1,
            "time": 1653545852,
            "blocktime": 1653545852,
            "valueOut": 7.00010612,
            "size": 178,
            "valueIn": 7
        }
    ]
}
```

<br/>

## Status and Utility

Methods that provide utilities or relay network status.

### /sync/

**URL:**  ```GET /sync```

**Description:** Retrieves status of dcrndata's synchronization with the connected node (dcrnd).


**Request Example:**

```GET /sync```

**Request Response:**

```
{
    "status": "finished",
    "blockChainHeight": 631,
    "syncPercentage": 100,
    "height": 631,
    "error": null,
    "type": "from RPC calls"
}
```

<br/>

### /peer/

**URL:**  ```GET /peer```

**Description:** Retrieves Peer-to-Peer (P2P) data sync status.  

**Request Example:**

```GET /peer/ ```

**Request Response:**

```
{
    "connected": true,
    "host": "127.0.0.1",
    "port": null
}
```

<br/>

### /status/

**URL:**  ```GET /status```

**Description:** Retrieves status of Decred-Next network. If `q` is set to a parameter from the table below, only that parameter is returned. If `q` is not specified, all status parameters are returned. 


| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| q              | `string`      |   (optional) getInfo </br> (optional) getDifficulty </br> (optional) getBestBlockHash </br> (optional) getLastBlockHash </br>|

**Request Example:**

```
GET /status?q=getBestBlockHash
```

**Request Response:**

```
{
    "bestblockhash": "000000006a9d766817f9fc6eecf4d18cd00ba7a0dc2fcbacf3bc6807c586d44e"
}
```

<br/>

### /estimatefee/

**URL:**  ```GET /estimatefee```

**Description:** Retrieves an estimate of the fee required for a transaction to be included within a certain number of blocks (`nbBlocks`). If `nbBlocks` is not specified, it defaults to 2. 

**Parameters:**

| Parameter           | Type                   |  Description                   | 
| -------------------- | ---------------------- | ---------------------- | 
| nbBlocks             | `int64`      |   (optional) Number of blocks within which the transaction should be mined |  

**Request Example:**

```GET /utils/estimatefee?nbBlocks=5 ```

**Request Response:**

```
{
    "5": 0.0001
}
```
