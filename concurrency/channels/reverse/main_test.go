package main

import (
	"strings"
	"testing"
)

func Benchmark_walker(b *testing.B) {

	for i := 0; i < b.N; i++ {
		benchmarkWalker()
	}

}

var str = `
JSON and JSONB Mapping ¶
pgx includes built-in support to marshal and unmarshal between Go types and the PostgreSQL JSON and JSONB.

Inet and CIDR Mapping ¶
pgx encodes from net.IPNet to and from inet and cidr PostgreSQL types. In addition, as a convenience pgx will encode from a net.IP; it will assume a /32 netmask for IPv4 and a /128 for IPv6.

Custom Type Support ¶
pgx includes support for the common data types like integers, floats, strings, dates, and times that have direct mappings between Go and SQL. In addition, pgx uses the github.com/jackc/pgx/pgtype library to support more types. See documention for that library for instructions on how to implement custom types.

See example_custom_type_test.go for an example of a custom type for the PostgreSQL point type.

pgx also includes support for custom types implementing the database/sql.Scanner and database/sql/driver.Valuer interfaces.

If pgx does cannot natively encode a type and that type is a renamed type (e.g. type MyTime time.Time) pgx will attempt to encode the underlying type. While this is usually desired behavior it can produce surprising behavior if one the underlying type and the renamed type each implement database/sql interfaces and the other implements pgx interfaces. It is recommended that this situation be avoided by implementing pgx interfaces on the renamed type.

Raw Bytes Mapping ¶
[]byte passed as arguments to Query, QueryRow, and Exec are passed unmodified to PostgreSQL.

Transactions ¶
Transactions are started by calling Begin.

tx, err := conn.Begin(context.Background())
if err != nil {
    return err
}
// Rollback is safe to call even if the tx is already closed, so if
// the tx commits successfully, this is a no-op
defer tx.Rollback(context.Background())

_, err = tx.Exec(context.Background(), "insert into foo(id) values (1)")
if err != nil {
    return err
}

err = tx.Commit(context.Background())
if err != nil {
    return err
}
The Tx returned from Begin also implements the Begin method. This can be used to implement pseudo nested transactions. These are internally implemented with savepoints.

Use BeginTx to control the transaction mode.

Prepared Statements ¶
Prepared statements can be manually created with the Prepare method. However, this is rarely necessary because pgx includes an automatic statement cache by default. Queries run through the normal Query, QueryRow, and Exec functions are automatically prepared
on first execution and the prepared statement is reused on subsequent executions. See ParseConfig for information on how to customize or disable the statement cache.
`

func Benchmark_walker2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkWalker2()
	}

}
func Benchmark_walker3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkWalker3()
	}

}
func Benchmark_walker4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkWalker4()
	}

}

func benchmarkWalker() {
	s := strings.Fields(str)
	reverseStrStream := make(chan string) // Q. how can you do this with buffer channels
	newr := []string{}

	walker(s, reverseStrStream)
	for i := 0; i < len(s); i++ {
		newr = append(newr, <-reverseStrStream)
	}
}

func benchmarkWalker2() {
	s := strings.Fields(str)
	newr := []string{}
	for i := 0; i < len(s); i++ {
		newr = append(newr, walker2(s)[i])
	}
}

func benchmarkWalker3() {
	s := strings.Fields(str)
	newr := []string{}
	for r := range walker3(s) {
		newr = append(newr, r)
	}
}

func benchmarkWalker4() {
	strSlice := strings.Fields(str)
	reverseStrStream := make(chan string)
	strStream := make(chan string, len(strSlice))
	newr := []string{}

	// create a walker
	for i := 0; i < (len(strSlice) / 2); i++ {
		go walker4(strStream, reverseStrStream)
	}

	go func() {
		for _, s := range strSlice {
			strStream <- s
		}
	}()

	for i := 0; i < len(strSlice); i++ {
		newr = append(newr, <-reverseStrStream)
	}
	close(reverseStrStream)
	close(strStream)
}
