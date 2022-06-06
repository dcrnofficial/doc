# Configuration

## What do you mean by configuration files 
Each application (dcrnd, dcrnwallet, dcrnctl) can have its own configuration files. Use `-h` and look at the path in parentheses of the configuration file option (-C, --configfile) to see the default path. Create a text file at the path and named according to that path you just looked up.

Then you can use the dcrnd sample config file and dcrnwallet sample config file to set whatever options you want. You can do the same thing for dcrnctl too. The format is the same. Every command line option listed by `-h` can be specified in the config files (just use the long option name).

Once those are created and in place, you do not have to add all of the options to the command line all the time.

## Port
 The default ports dcrnd listen on
|            | mainnet | testnet | simnet |
|:----------:|:-------:|:-------:|:------:|
|Peer to Peer|  9108   |  19108  | 18555  |
|RPC Server  |  9109   |  19109  | 18556  |

The default ports dcrnwallet listen on
|               | mainnet | testnet | simnet |
|:-------------:|:-------:|:-------:|:------:|
|JSON-RPC Server|   9110  |  19110  | 19557  |
|gRPC Server    |   9111  |  19111  | 19558  |