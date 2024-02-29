package house

import (
	"database/sql"
	//"github.com/ClickHouse/clickhouse-go"
	//"github.com/ClickHouse/clickhouse-go/v2"
	//"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	_ "github.com/ClickHouse/clickhouse-go"
	"go_service/config"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("clickhouse", "tcp://localhost:9000?database=default&username=root&debug=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

	//connect, err := sql.Open("clickhouse", "tcp://localhost:9000?&database=go_service&username=default&debug=true")
	//if err != nil {
	//	return nil, err
	//}
	//if err := connect.Ping(); err != nil {
	//	if exception, ok := err.(*clickhouse.Exception); ok {
	//		fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
	//	} else {
	//		fmt.Println(err)
	//	}
	//	return nil, err
	//}
	//
	//return connect, nil

	//
	//var (
	//	ctx       = context.Background()
	//	conn, err = clickhouse.Open(&clickhouse.Options{
	//		Addr: []string{"<CLICKHOUSE_SECURE_NATIVE_HOSTNAME>:9440"},
	//		Auth: clickhouse.Auth{
	//			Database: cfg.ClickHouse.Database,
	//			Username: cfg.ClickHouse.Username,
	//			Password: cfg.ClickHouse.Password,
	//		},
	//		ClientInfo: clickhouse.ClientInfo{
	//			Products: []struct {
	//				Name    string
	//				Version string
	//			}{
	//				{Name: "an-example-go-client", Version: "0.1"},
	//			},
	//		},
	//
	//		Debugf: func(format string, v ...interface{}) {
	//			fmt.Printf(format, v)
	//		},
	//		TLS: &tls.Config{
	//			InsecureSkipVerify: true,
	//		},
	//	})
	//)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//if err := conn.Ping(ctx); err != nil {
	//	if exception, ok := err.(*clickhouse.Exception); ok {
	//		fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
	//	}
	//	return nil, err
	//}
	//return conn, nil
}
