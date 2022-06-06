# ExchangeGuide

Last updated for CLI release v1.5.2.

This guide offers signposts to some useful resources, including how to set up exchange's dcrnd, dcrnwallet and dcrnctl. 

## Deploy
### Minimum Recommended Specifications
* 50 GB disk space (as of June 2022, increases over time)
* 4GB memory (RAM)
* Windows 7/8.x/10, macOS, Linux
* High uptime
  
**WARNING: If you decide to build from source and you plan to trade on mainnet, use the release branch instead of master. Furthermore, if you build dcrnd and dcrnwallet from source, you must use their release branches, not master.**


### Manual Installation
The newest binary releases can be found [here](https://github.com/Decred-Next/binary-release).Although most of this will be extract and go, instructions are provided for Linux below.

1. Download the correct file for your computer:
2. Navigate to download location and extract the .tar.gz file:
Ubuntu File Browser: simply right click on the .tar.gz file and select “Extract Here”.
Terminal: use the tar -xvzf filename.tar.gz command.
Both of these should extract the .tar.gz file into a folder that shares the same name. 


### Configuration
#### Configuration File Locations
All of the Decred-Next software, when started, reads from a configuration file to determine which settings it should enable/disable/set during that initial load. All of the command line startup flags (e.g. dcrnd --testnet) can be replaced by settings within the appropriate configuration file (e.g. dcrnd --testnet could be replaced by testnet=1 in dcrnd.conf).

These configuration files are located within the application home directory of the application. The location of these default home directories for Windows, macOS, and Linux are listed below:

|  OS  | dcrnd |dcrnwallet| dcrnctl |
|:----:|:-----:|:--------:|:-------:|
|Linux|~/.dcrnd/|~/.dcrnwallet/|~/.dcrnctl/|
|Windows|%LOCALAPPDATA%\Dcrnd\ |%LOCALAPPDATA%\Dcrnwallet\ |%LOCALAPPDATA%\Dcrnctl\ |
|MacOS|~/Library/Application Support/Dcrnd/|~/Library/Application Support/Dcrnwallet/|~/Library/Application Support/Dcrnctl/|

Each of these folders is allowed its own .conf file, named after the individual application (e.g. dcrnd uses dcrnd.conf). Please also note that the Dcrnd and Dcrnwallet home directories are automatically created when each application is first launched. You will have to manually create a Dcrnctl home directory to utilize a config file.

The Manual Installation method includes sample configuration files within the .zip/.tar.gz. It is recommended to copy these config files into the appropriate directory described above, and rename them to remove ‘sample-‘. These files have many settings commented out (comments are not read by the program during runtime) so all of these settings are effectively disabled. You can enable these pre-written settings by simply deleting the semi-colon before the line.


### Minimum Configuration
At the very minimum, for dcrnd, dcrnwallet, and dcrnctl to be able to communicate with each other, they need to be launched with the same rpcuser/rpcpass combination. For manual configuration please follow these steps:
1. If the OS-specific home directories listed in Configuration File Locations do not exist, please create them for dcrnd, dcrnwallet, and dcrnctl.
   
2. Choose an arbitrary username and password, these will only be used for each application to communicate via remote procedure call. The easiest configuration is to set them all equal.

3. Using your favorite text editor, create a new text file and add the following lines:
    ```
    [Application Options]
    rpcuser=<chosen-username>
    rpcpass=<chosen-password>
    ```
    Save it as dcrnd.conf in dcrnds home directory.
4. Create another new text file and add the following lines:
    ```
    [Application Options]
    rpcuser=<chosen-username>
    rpcpass=<chosen-password>
    ```
    Save it as dcrnwallet.conf in dcrnwallet's home directory.
5. Create a third text file and add the following lines:
    ```
    [Application Options]
    rpcuser=<chosen-username>
    rpcpass=<chosen-password>
    ```
    Save it as dcrnctl.conf in dcrnctl‘s home directory.
    

### Port
 The default ports dcrd listen on
|            | mainnet | testnet | simnet |
|:----------:|:-------:|:-------:|:------:|
|Peer to Peer|  9108   |  19108  | 18555  |
|RPC Server  |  9109   |  19109  | 18556  |

The default ports dcrnwallet listen on
|               | mainnet | testnet | simnet |
|:-------------:|:-------:|:-------:|:------:|
|JSON-RPC Server|   9110  |  19110  | 19557  |
|gRPC Server    |   9111  |  19111  | 19558  |


### Additional Configuration Options
All command line options can be put in the config file. The sample config files in the release package give additional options or you can run one of the programs with the -h flag to show a list and description of all options for the specified application.

Dcrnd has a number of [configuration](./../Wallet/CLI%20Wallet/CliWallet.md#dcrndCli) options, which can be viewed by running: `$ dcrnd --help`.

Dcrnwallet has a number of [configuration](./../Wallet/CLI%20Wallet/CliWallet.md#dcrnwalletCli) options, which can be viewed by running: `$ dcrnwallet --help`.

## Dcrnd launch(Running a Full Node)

The first time launching dcrnd, it will connect to the Decred-Next network and begin downloading the blockchain.
1. **Configure RPC Username and Password** -  To allow dcrnwallet and dcrnctl to communicate with dcrnd, the configuration files must have rpcuser and rpcpass settings enabled.

2. **Start dcrnd** - Navigate to the /decred directory and launch dcrnd.(`./dcrnd`) See the [Operating System Differences](./../Wallet/CLI%20Wallet/CliWallet.md##osd) page for OS-specific commands. - dcrd will boot up, begin connecting to peers and downloading the full Decred blockchain.

3. **Enable incoming connections** - dcrd will automatically begin downloading the blockchain and connecting to peers. To maximize your positive impact on the Decred network, it is important to allow inbound peers to connect to your node. Allowing inbound peers will allow new participants in the Decred network to connect to your node, and also enable your node to serve lightweight clients such as SPV wallets. This may require changing the settings of your router. If this is the case: 1. Find your Local IP address. 1. Edit the settings of your router to open port 9108 (port forwarding). This process depends on the type of router you have. - If running a firewall, you’ll also need to configure it to allow inbound connections on port 9108.
4. **Wait for dcrd to sync** - Check a Decred block explorer such as dcrdata and wait until the block height displayed matches the block height shown locally.
5. **Leave running** - Leave your node running, online, in a safe space, 24/7/365.

Congratulations, you are now running a Decred-Next full node to support the network!

Note that this connection will be used in the future. You must leave this dcrnd instance running in order to use dcrwallet.



## DcrnWallet

**Prerequisites:
Setup dcrnd and have it running in the background.**

In order to run dcrnwallet, a wallet.db must exist within dcrnwallet's home directory. In order for that file to exist, you must create a new wallet. If you delete your wallet.db or used another installation process, you’ll have to run the manual wallet creation command.

The wallet creation and launch steps can be found [here](./../Wallet/CLI%20Wallet/CliWallet.md#mwcc).


### Generate addresses offline
**Prerequisites:
Create batch of addresses for one account**

The project "dcrnaddrgen" was ready for address and account operations. You can find the func 
`generateAddresses(seedHex string, filename string, addrNum uint32)` in https://github.com/Decred-Next/dcrnaddrgen. 
The func is in file `main.go`.  

**Paramters** 
1. **seedHex:** Your wallet seed in Hexadecimal format. Saved when you create your wallet  
2. **filename:** Addresses will write in the file  
addrNum: How many addresses you want generate.  
And this is a sample example for invoke it.
~~~
func TestGenerateAddress(t *testing.T) {
	generateAddresses("a0b66fcdf463a26753d5d86c5022ef95220d6efb0cf7b133a855ff9e299ccaa9", "/home/ubuntu/address.txt", 100)
}
~~~

### Send transcation by code
**Prerequisites: Call wallet RPC interface with code to send transaction and get account info**  
In order to call wallet RPC interface with code we prepared ./httpclient.go. Its a util to call the wallet rpc
here is the guide to use it.  
1. **get account info**
~~~
func main() {
	marshalledStr := "{\"jsonrpc\":\"1.0\",\"method\":\"getbalance\",\"params\":[],\"id\":1}"
	response := sendPostRequest(marshalledStr, 1)
}
~~~
MarshalledStr:  
~~~
{
	"jsonrpc": "1.0",         //jsonrpc version
	"method": "getbalance",   //rpc method name (get account info)
	"params": [],             //params
	"id": 1                   //id
}
~~~
Response:
~~~
{
	"jsonrpc": "1.0",         
	"result": {
		"balances": [
			{
				"accountname": "default",                //account name
				"immaturecoinbaserewards":20.42105366,   //The total value of all immature coinbase outputs.
				"immaturestakegeneration":0,             //The total value of all immature stakebase outputs
				"lockedbytickets":0,                     //The total value of all tickets that are currently locked, and awaiting vote.
				"spendable": 190.17713242,               //The spendable balance, given some number of required confirmations, counted in Atoms. This equals the total balance when the required number of confirmations is zero and there are no immature coinbase outputs.
				"total": 210.59818608,                   //The total balance
				"unconfirmed" :0,                        //The total value of all unconfirmed transactions with with reference to the minimum number of confirmations for a transaction (minconf). If minconf is 0 unconfirmed will be 0, otherwise unconfirmed will be the total balance of transactions that do fulfill the requested minconf.
				"votingauthority": 0                     //The total value of all tickets that the account has voting authority over.
			},
			{
				"accountname": "imported",
				"immaturecoinbaserewards": 0,
				"immaturestakegeneration": 0,
				"lockedbytickets": 0,
				"spendable": 0,
				"total": 0,
				"unconfirmed": 0,
				"votingauthority": 0
			}
		],
		"blockhash": "0000000ea47d6816b673cbe250fdd68e5d9e82bcc4fe61521b268e5ef90e8127",
		"totalimmaturecoinbaserewards": 20.42105366,     //The total value of all account
		"totalspendable": 190.17713242,                  //The total spendable of all account
		"cumulativetotal" :210.59818608                  
	},
	"error ":null,
	"id ":1
}
~~~  

2. **Send transaction**
~~~
func main() {
	marshalledStr := "{\"jsonrpc\":\"1.0\",\"method\":\"sendfrom\",\"params\":[\"default\", \"TsmiBuWig2hUiV7TX4gsr91ZU9dhj6QJKdm\", 10],\"id\":1}
	response := sendPostRequest(marshalledStr, 1)
}
~~~
MarshalledStr:
~~~
{
    "jsonrpc":"1.0",       //jsonrpc version
    "method":"sendfrom",   //method name
    "params":[
        "default",         //send coin from account name
        "TsmiBuWig2hUiV7TX4gsr91ZU9dhj6QJKdm",     //send to address
        "10"               //coin num
    ],
    "id":1                 //id
}
~~~
Response
~~~
{
    "jsonrpc":"1.0",       //jsonrpc version
    "result":"0f4a2224330870e883ebc4ea819a082227567c07bcceb37e8bab8673ef3adf6a",    //transcation hash
    "error":null,          
    "id":1                 //id
}
~~~

### Sign transaction offline
**Prerequisites: Generate unsigned transaction; Sign the unsigned transaction and send it to a dcrnd node**
Generate unsigned transaction and sign the unsigned transaction part view the [offlinesigner](https://github.com/Decred-Next/offlinesigner).  
Through the offlinesigner you will get a hexadecimal signed transaction then you can send the transaction by calling dcrnd 
jsonRPC interface Through ./httpclient.go.  
1. send raw transaction
~~~
func main() {
	marshalledStr := "{\"jsonrpc\":\"1.0\",\"method\":\"sendrawtransaction\",\"params\":[\"01000000043aafc4fd6adcec85aa677df0f99808117602a9d1498caff845125ff53bdd1e7b0100000000ffffffff3c6daaaa92dd6362d3b41669a9b51648bce231d1847dc8af266a7bfa0fcbdd4e0200000000ffffffff3d9c486319fe05ad2a090762689c789780a682adbace055130b58d1c65fee9130100000000ffffffff44b0b296c187ff4563c45f653c3aa1852f8c03bff7bc38184242fdf99cac91040200000000ffffffff0200ca9a3b0000000000001976a914cd634b3621a0144824012ee4bcd3379ce866979e88ac7484810f0000000000001976a91427821402540fc3d55b090ab0224ed5aac202c30988ac000000000000000004e73414130000000000000000ffffffff6b483045022100db660f94e25eda2173e24f4e8682183a1f9d07136750e6e9e0237d39e53a758302205b620c938689cfd8513f9c24ab0a250264fe3215acaf419300869f2912fbf65c012103ebb7eab9893a42e7e4babcc08b905c11c9172a8f4c72877e3ef425045138b8a0412e76150000000000000000ffffffff6a4730440220352276ce0dfc0e95ab240536b996a5931af5ffc92f0b310e53123af0eeca86390220728a6d0a433c0fe74f59c1c86772cc6f883f395da6a792b5c55f22d74188bc8c012102a23f4b3ba56c92c9a64ecefce2d4da918b957a06601f65d4f29b9c3edb25654aa2de1c0d0000000000000000ffffffff6a47304402200d7186549f1a651feea0ea53aaafbaf0401c935f59969ebb8d59f970644f6e2802207d014427cf5e853356646454874892c4ba028dbfcb48cc45dbc3641c0cf52932012103231d30147cffb693fb4f3c93fd6b0ba9eaae53253c9fd61a6f839e8c0654914f002a75150000000000000000ffffffff6a47304402206ac5cd441dd32e377ac8a7745762c436e1651fed96fe6889c558c2d6c6e7527902201d36541a312c0604bd7144e08451e1164e5ff0e78c0bb07d4023f656541055cb012102a23f4b3ba56c92c9a64ecefce2d4da918b957a06601f65d4f29b9c3edb25654a\"],\"id\":1}
	//marshalledStr is jsonRPC request body
	//2 means request dcrnd jsonRPC server
	response := sendPostRequest(marshalledStr, 2)
}
~~~
MarshalledStr:
~~~
{
    "jsonrpc":"1.0",                 //jsonrpc version
    "method":"sendrawtransaction",   //method name
    "params":[                       //hexadecimal signed transaction
        "01000000043aafc4fd6adcec85aa677df0f99808117602a9d1498caff845125ff53bdd
        9fe05ad2a090762689c789780a682adbace055130b58d1c65fee9130100000000fffffff
        f44b0b296c187ff4563c45f653c3aa1852f8c03bff7bc38184242fdf99cac91040200000"
    ],
    "id":1                 //id
}
~~~
Response:
~~~
{
    "jsonrpc":"1.0",       //jsonrpc version
    "result":"0f4a2224330870e883ebc4ea819a082227567c07bcceb37e8bab8673ef3adf6a",    //transcation hash
    "error":null,          
    "id":1                 //id
}
~~~

## Dcrnctl using

**Prerequisites:
Setup dcrnwallet and have it running in the background.**

Dcrnctl is the client that controls dcrd and dcrnwallet via remote procedure call (RPC). You can use dcrnctl for many things such as checking your balance, buying tickets, creating transactions, and viewing network information.

Dcrnctl is not a daemon - it does not run permanently in the background - it calls the requested RPC method, prints the response, and then terminates immediately.

The dcrnctl usage and some common commands can be found [here](./../Wallet/CLI%20Wallet/CliWallet.md#ctlUsage).

More common RPC commands for dcrd and dcrwallet can be found [here](./../Wallet/CLI%20Wallet/CliWallet.md#ctlCommand).



## Security

### whitelist
