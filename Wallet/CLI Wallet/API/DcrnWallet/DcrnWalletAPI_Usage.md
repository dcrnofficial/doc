# DcrnWalletAPI Usage

Clients use RPC to interact with the wallet.  A client may be implemented in any language directly supported by [gRPC](https://www.grpc.io/), languages capable of performing [FFI](https://en.wikipedia.org/wiki/Foreign_function_interface) with these, and languages that share a common runtime (e.g. Scala, Kotlin, and Ceylon for the JVM, F# for the CLR, etc.).

The rest of this document provides short examples of how to quickly get started by implementing a basic client that fetches the balance of the default account (account 0) from a testnet wallet listening on `localhost:19111`(mainnet:9111) in several different languages:

- [DcrnWalletAPI Usage](#dcrnwalletapi-usage)
	- [Go](#go)
		- [VersionService](#versionservice)
			- [Version](#version)
		- [WalletService](#walletservice)
			- [Balance](#balance)
		- [WalletLoaderService](#walletloaderservice)
			- [WalletExists](#walletexists)
		- [SeedService](#seedservice)
			- [GenerateRandomSeed](#generaterandomseed)


Unless otherwise stated under the language example, it is assumed that
gRPC is already already installed.  The gRPC installation procedure
can vary greatly depending on the operating system being used and
whether a gRPC source install is required.  Follow the [gRPC install
instructions](https://github.com/grpc/grpc/blob/master/INSTALL) if
gRPC is not already installed.  A full gRPC install also includes
[Protocol Buffers](https://github.com/google/protobuf) (compiled with
support for the proto3 language version), which contains the protoc
tool and language plugins used to compile this project's `.proto`
files to language-specific bindings.

## Go

The native gRPC library (gRPC Core) is not required for Go clients (a pure Go implementation is used instead) and no additional setup is required to generate Go bindings.


### VersionService
#### Version
```Go
package main

import (
	"fmt"
	"path/filepath"

	pb "github.com/decred/dcrwallet/rpc/walletrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/decred/dcrd/dcrutil"
)

var dcrnwalletHomeDir = dcrutil.AppDataDir("dcrwallet", false)
var certificateFile = filepath.Join(dcrnwalletHomeDir, "rpc.cert")

func main() {
	creds, err := credentials.NewClientTLSFromFile(certificateFile, "localhost")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := grpc.Dial("localhost:19111", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := pb.NewVersionServiceClient(conn)

	versionRequest := &pb.VersionRequest{}
	VersionResponse, err := c.Version(context.Background(), versionRequest)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("VersionResponse:%v\n", VersionResponse)
}
```


### WalletService
#### Balance
```Go
package main

import (
	"fmt"
	"path/filepath"

	pb "github.com/decred/dcrwallet/rpc/walletrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/decred/dcrd/dcrutil"
)

var certificateFile = filepath.Join(dcrutil.AppDataDir("dcrwallet", false), "rpc.cert")

func main() {
	creds, err := credentials.NewClientTLSFromFile(certificateFile, "localhost")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := grpc.Dial("localhost:19111", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	c := pb.NewWalletServiceClient(conn)

	balanceRequest := &pb.BalanceRequest{
		AccountNumber:         0,
		RequiredConfirmations: 1,
	}
	balanceResponse, err := c.Balance(context.Background(), balanceRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("balanceResponse: %v\n", balanceResponse)

}
```


### WalletLoaderService
#### WalletExists
```Go
package main

import (
	"fmt"
	"path/filepath"

	pb "github.com/decred/dcrwallet/rpc/walletrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/decred/dcrd/dcrutil"
)

var certificateFile = filepath.Join(dcrutil.AppDataDir("dcrwallet", false), "rpc.cert")

func main() {
	creds, err := credentials.NewClientTLSFromFile(certificateFile, "localhost")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := grpc.Dial("localhost:19111", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := pb.NewWalletLoaderServiceClient(conn)
	walletExistsRequest := &pb.WalletExistsRequest{}
	walletExistsResponse, err := c.WalletExists(context.Background(), walletExistsRequest)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("walletExistsResponse: %v\n", walletExistsResponse)

}

```


### SeedService
#### GenerateRandomSeed
```Go
package main

import (
	"fmt"
	"path/filepath"

	pb "github.com/decred/dcrwallet/rpc/walletrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/decred/dcrd/dcrutil"
)

var certificateFile = filepath.Join(dcrutil.AppDataDir("dcrwallet", false), "rpc.cert")

func main() {
	creds, err := credentials.NewClientTLSFromFile(certificateFile, "localhost")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := grpc.Dial("localhost:19111", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := pb.NewSeedServiceClient(conn)
	generateRandomSeedRequest := &pb.GenerateRandomSeedRequest{}
	generateRandomSeedResponse, err := c.GenerateRandomSeed(context.Background(), generateRandomSeedRequest)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("SeedBytes: %v\n", generateRandomSeedResponse.SeedBytes)
	fmt.Printf("SeedHex: %v\n", generateRandomSeedResponse.SeedHex)
	fmt.Printf("SeedMnemonic: %v\n", generateRandomSeedResponse.SeedMnemonic)

}

```
