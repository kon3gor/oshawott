package ydbs

import (
	"fmt"
	"os"

	"github.com/kon3gor/oshawott"
	"github.com/yandex-cloud/ydb-go-sdk/v2"
	"github.com/yandex-cloud/ydb-go-sdk/v2/connect"
	"github.com/yandex-cloud/ydb-go-sdk/v2/table"
)

const ydbUrl = "grpcs://ydb.serverless.yandexcloud.net:2135/?database=/ru-central1/b1g1mak5k5l52mp3rc74/etn3un20r669i07bkn3g"

var txc = table.TxControl(
	table.BeginTx(table.WithSerializableReadWrite()),
	table.CommitTx(),
)

func (ys YdbStorage) saveUrl(key oshawott.Key, url string) {
	session, err := ys.conn.Table().CreateSession(ys.ctx.Ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return
	}
	defer session.Close(ys.ctx.Ctx)

	_, _, err = session.Execute(ys.ctx.Ctx, txc,
		`--!syntax_v1
		declare $key as Utf8;
		declare $value as Utf8;
		upsert into Links (key, value) values ($key, $value)
		`,
		table.NewQueryParameters(
			table.ValueParam("$key", ydb.UTF8Value(string(key))),
			table.ValueParam("$value", ydb.UTF8Value(url)),
		),
	)

	if err != nil {
		//todo: handle this
	}
}

func (ys YdbStorage) getKey(url string) (oshawott.Key, bool) {
	session, err := ys.conn.Table().CreateSession(ys.ctx.Ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return oshawott.NoKey, false
	}
	defer session.Close(ys.ctx.Ctx)

	_, res, err := session.Execute(ys.ctx.Ctx, txc,
		`--!syntax_v1
		declare $value as Utf8;
		select key from Links where value = $value
		`,
		table.NewQueryParameters(
			table.ValueParam("$value", ydb.UTF8Value(url)),
		),
	)

	if err != nil {
		//todo: handle this
		fmt.Println(err)
		return "", false

	}

	var key string
	for res.NextResultSet(ys.ctx.Ctx, "key") {
		for res.NextRow() {
			err := res.Scan(&key)

			// Error handling.
			if err != nil {
				fmt.Println(err)
				return "", false
			}

			return oshawott.Key(key), true
		}
	}

	if res.Err() != nil {
		fmt.Println(res.Err())
	}

	return oshawott.NoKey, false
}

func (ys YdbStorage) getValue(k oshawott.Key) (string, bool) {
	session, err := ys.conn.Table().CreateSession(ys.ctx.Ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return "", false
	}
	defer session.Close(ys.ctx.Ctx)

	_, res, err := session.Execute(ys.ctx.Ctx, txc,
		`--!syntax_v1
		declare $key as Utf8;
		select value from Links where key = $key
		`,
		table.NewQueryParameters(
			table.ValueParam("$key", ydb.UTF8Value(string(k))),
		),
	)

	if err != nil {
		//todo: handle this
		fmt.Println(err)
		return "", false

	}

	var value string
	for res.NextResultSet(ys.ctx.Ctx, "value") {
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

func (ys YdbStorage) getUsedKeys() ([]string, error) {
	session, err := ys.conn.Table().CreateSession(ys.ctx.Ctx)
	if err != nil {
		//todo: handle this shit
		fmt.Println(err)
		return nil, err
	}
	defer session.Close(ys.ctx.Ctx)

	_, res, err := session.Execute(ys.ctx.Ctx, txc,
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
	for res.NextResultSet(ys.ctx.Ctx, "key") {
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

func conn(ctx oshawott.AppContext) (*connect.Connection, error) {
	sf, found := os.LookupEnv("SA_FILE")
	if !found {
		//todo: send error here
		fmt.Println("dang")
		return nil, nil
	}

	return connect.New(
		ctx.Ctx,
		connect.MustConnectionString(ydbUrl),
		connect.WithServiceAccountKeyFileCredentials(sf),
	)
}
