package core

import (
	"fmt"
	"os"

	"github.com/yandex-cloud/ydb-go-sdk/v2"
	"github.com/yandex-cloud/ydb-go-sdk/v2/connect"
	"github.com/yandex-cloud/ydb-go-sdk/v2/table"
)

const ydbUrl = "grpcs://ydb.serverless.yandexcloud.net:2135/?database=/ru-central1/b1g1mak5k5l52mp3rc74/etn3un20r669i07bkn3g"

var txc = table.TxControl(
	table.BeginTx(table.WithSerializableReadWrite()),
	table.CommitTx(),
)

func SaveUrl(ctx OshawottContext, key string, url string) {
	db, err := conn(ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return
	}
	defer db.Close()

	session, err := db.Table().CreateSession(ctx.ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return
	}
	defer session.Close(ctx.ctx)

	_, _, err = session.Execute(ctx.ctx, txc,
		`--!syntax_v1
		declare $key as Utf8;
		declare $value as Utf8;
		upsert into Links (key, value) values ($key, $value)
		`,
		table.NewQueryParameters(
			table.ValueParam("$key", ydb.UTF8Value(key)),
			table.ValueParam("$value", ydb.UTF8Value(url)),
		),
	)

	if err != nil {
		//todo: handle this
	}
}

func GetKey(ctx OshawottContext, v string) (string, bool) {
	db, err := conn(ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return "", false
	}
	defer db.Close()

	session, err := db.Table().CreateSession(ctx.ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return "", false
	}
	defer session.Close(ctx.ctx)

	_, res, err := session.Execute(ctx.ctx, txc,
		`--!syntax_v1
		declare $value as Utf8;
		select key from Links where value = $value
		`,
		table.NewQueryParameters(
			table.ValueParam("$value", ydb.UTF8Value(v)),
		),
	)

	if err != nil {
		//todo: handle this
		fmt.Println(err)
		return "", false

	}

	var key string
	for res.NextResultSet(ctx.ctx, "key") {
		for res.NextRow() {
			err := res.Scan(&key)

			// Error handling.
			if err != nil {
				fmt.Println(err)
				return "", false
			}

			return key, true
		}
	}

	if res.Err() != nil {
		fmt.Println(res.Err())
	}

	return "", false
}

func GetValue(ctx OshawottContext, k string) (string, bool) {
	db, err := conn(ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return "", false
	}
	defer db.Close()

	session, err := db.Table().CreateSession(ctx.ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return "", false
	}
	defer session.Close(ctx.ctx)

	_, res, err := session.Execute(ctx.ctx, txc,
		`--!syntax_v1
		declare $key as Utf8;
		select value from Links where key = $key
		`,
		table.NewQueryParameters(
			table.ValueParam("$key", ydb.UTF8Value(k)),
		),
	)

	if err != nil {
		//todo: handle this
		fmt.Println(err)
		return "", false

	}

	var value string
	for res.NextResultSet(ctx.ctx, "value") {
		for res.NextRow() {
			err := res.Scan(&value)

			// Error handling.
			if err != nil {
				fmt.Println(err)
				return "", false
			}

			return value, true
		}
	}

	if res.Err() != nil {
		fmt.Println(res.Err())
	}

	return "", false
}

func GetUsedKeys(ctx OshawottContext) ([]string, error) {
	db, err := conn(ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return nil, err
	}
	defer db.Close()

	session, err := db.Table().CreateSession(ctx.ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return nil, err
	}
	defer session.Close(ctx.ctx)

	_, res, err := session.Execute(ctx.ctx, txc,
		`--!syntax_v1
		select key from Links;
		`,
		table.NewQueryParameters(),
	)

	if err != nil {
		//todo: handle this
		fmt.Println(err)
		return nil, err

	}

	keys := make([]string, 0)
	var value string
	for res.NextResultSet(ctx.ctx, "key") {
		for res.NextRow() {
			err := res.Scan(&value)

			// Error handling.
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			keys = append(keys, value)

		}
	}

	if res.Err() != nil {
		fmt.Println(res.Err())
		return nil, err
	}

	return keys, nil
}

func conn(ctx OshawottContext) (*connect.Connection, error) {
	sf, found := os.LookupEnv("SA_FILE")
	if !found {
		//todo: send error here
		fmt.Println("dang")
		return nil, nil
	}

	return connect.New(
		ctx.ctx,
		connect.MustConnectionString(ydbUrl),
		connect.WithServiceAccountKeyFileCredentials(sf),
	)
}
