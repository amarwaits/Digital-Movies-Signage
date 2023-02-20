# simple-go-grpc-wallet

#### Objective
As a test you are required to design and create a user wallet service from scratch, which is to record customer balance, and all transactions occurred through the system.

#### Requirements
1. Design a data structure for storing customer ledgers, balance details etc.
2. By making use of the created data structure, create a gRPC service to support below features
	a. Create user wallet – assuming user profile is created already in other systems and a unique user id (uuid) is created
	b. Record keeping for user's transaction (credit to or debit from user's wallet) in ledgers
	c. Retrieve user wallet summary 
	d. Get user transaction history (pagination required)
3. The service should also support
	 a. Multiple currency (e.g ETC, BTC, USD etc.)
	 b. Multiple users
4. The service should have proper unit testing coverage
5. Proper logging in the service
6. Documentation
7. Anything you think it can make the service better

\* Please try not to over complicate the task.
\* You can use any third-party framework or libraries that you think it can help on you task.	
\* Please provide a repository URL (e.g Github) to send it after completion.

#### Tech Design
We have used in-memory data structures to implement this assignment.
* `proto/wallet.proto` defines the gRPC wallet service interface.
* `src/wallet/wallet.go` implements the above interface.
* `src/wallet/wallet_test.go` contains some unit tests.
* `cmd/server.go` contains the code to start the gRPC server.
* `cmd/client/client.go` is included that contains the client code for gRPC call(s) to server.

##### Dependencies:
1. Go v1.18 or above (Installation guide: https://go.dev/doc/install)
2. Protobuf v3.x (Installation guide: https://grpc.io/docs/protoc-installation/)
3. Go plugins:
	1. Install the protocol compiler plugins for Go using the following commands:
		>`$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`

		>`$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`
	2. Update your  `PATH`  so that the  `protoc`  compiler can find the plugins:
		>`$ export PATH="$PATH:$(go env GOPATH)/bin"`

##### Running the code
- Run server after moving the root directory of repository
	>`go run cmd/server.go`
- Run client after moving the root directory of repository
	>`go run cmd/client/client.go`

##### Test cases
- Run unit tests after moving the root directory of repository
	>`go test -v ./...`
