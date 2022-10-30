# Nimbus

[![PkgGoDev](https://pkg.go.dev/badge/github.com/albertollamaso/nimbus)](https://pkg.go.dev/github.com/albertollamaso/nimbus)

Nimbus is a library for simplify the usage of [Badger](https://github.com/dgraph-io/badger) key-value (KV) database.

## Goals of this project

- Have fun and learn a lot
- Implement an abstraction of Badger capabilities with little improvements

## Usage

Using Nimbus is easy. First, use go get to install the latest version of the library.

```
go get -u github.com/albertollamaso/nimbus@latest
```

Next, include Nimbus in your application:

```
import "github.com/albertollamaso/nimbus"
```

## Examples

- Add a key value pair

```
package main

import (
	"fmt"
	"github.com/albertollamaso/nimbus"
)

func main() {

	db, err := nimbus.InitDB("db")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = nimbus.AddKV(db, "me@example.com", "MyWeakPassword")
	if err != nil {
		fmt.Println(err)
	}
}
```

- Remove a key value pair

```
package main

import (
	"fmt"
	"github.com/albertollamaso/nimbus"
)

func main() {

	db, err := nimbus.InitDB("db")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = nimbus.RemoveKV(db, "me@example.com")
	if err != nil {
		fmt.Println(err)
		fmt.Println("couldn't delete key me@example.com that is not in the database")
		fmt.Println("-----------------------------")
	} else {
		fmt.Println("key: me@example.com has been removed")
		fmt.Println("-----------------------------")
	}
}
```

- Read the value for a given key

```
package main

import (
	"fmt"
	"github.com/albertollamaso/nimbus"
)

func main() {

	db, err := nimbus.InitDB("db")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	value, err := nimbus.ReadKV(db, "me@example.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("value of me@example.com is: ", value)
	fmt.Println("-----------------------------")
}
```

- List all key value pairs from the database

```
package main

import (
	"fmt"
	"github.com/albertollamaso/nimbus"
)

func main() {

	db, err := nimbus.InitDB("db")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	valuesAfter, err := nimbus.ListKV(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all KV pairs in database are:")
	for k, v := range valuesAfter {
		fmt.Println(k, v)
	}
}
```

## TODO

- Remove prefix (bulk)
- Prometheus exporter
- pub/sub mechanism
- Message broker implementation like AMQP
- In-Memory Mode/Diskless Mode
- Encryption Mode
- Merge Operations
- Setting Time To Live(TTL) 
- Support User Metadata on Keys
- Prefix scans
- Key-only iteration
- Stream

## Licence

Nimbus is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/albertollamaso/nimbus/blob/master/LICENSE.txt)
