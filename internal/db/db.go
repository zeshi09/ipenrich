// package db
//
// import (
//
//	"context"
//	"fmt"
//
//	"github.com/jackc/pgx/v5"
//	"github.com/zeshi09/ipenrich/model"
//
// )
//
//	func SaveToPostgres(results []model.EnrichedIP) error {
//		connStr := "postgres://superset:superset@localhost:5432/postgres"
//		// 	conn, err := pgx.Connect(context.Background(), connStr)
//		// 	if err != nil {
//		// 		return fmt.Errorf("failed to connect to postgres: %w", err)
//		// 	}
//		// 	defer conn.Close(context.Background())
//		//
//		// 	_, err = conn.Exec(context.Background(), `
//		// 		CREATE TABLE IF NOT EXISTS enriched_ips (
//		// 			log_file TEXT,
//		// 			ip TEXT,
//		// 			country TEXT,
//		// 			city TEXT,
//		// 			org TEXT,
//		// 			abuse_score INTEGER,
//		// 			vt_score TEXT
//		// 		);
//		// 	`)
//		// 	if err != nil {
//		// 		return fmt.Errorf("failed to create table: %w", err)
//		// 	}
//		//
//		// 	for _, r := range results {
//		// 		_, err = conn.Exec(context.Background(), `
//		// 			INSERT INTO enriched_ips (log_file, ip, country, city, org, abuse_score, vt_score)
//		// 			VALUES ($1, $2, $3, $4, $5, $6, $7)
//		// 		`, r.LogFile, r.Ip, r.Geo["country"], r.Geo["city"], r.Geo["org"], r.Abuse, r.Vt["score"])
//		// 		if err != nil {
//		// 			return fmt.Errorf("insert failed: %w", err)
//		// 		}
//		// 	}
//		//
//		// 	return nil
//		// }
//
//		conn, err := pgx.Connect(context.Background(), connStr)
//		if err != nil {
//			return fmt.Errorf("failed to connect to postgres: %w", err)
//		}
//		defer conn.Close(context.Background())
//
//		_, err = conn.Exec(context.Background(), `
//			CREATE TABLE IF NOT EXISTS enriched_ips (
//				log_file TEXT,
//				ip TEXT,
//				country TEXT,
//				city TEXT,
//				org TEXT,
//				abuse_score INTEGER,
//				vt_malicious INTEGER,
//				vt_suspicious INTEGER,
//				vt_harmless INTEGER,
//				vt_undetected INTEGER
//			);
//		`)
//		if err != nil {
//			return fmt.Errorf("failed to create table: %w", err)
//		}
//
//		for _, r := range results {
//			_, err = conn.Exec(context.Background(), `
//				INSERT INTO enriched_ips (
//					log_file, ip, country, city, org,
//					abuse_score,
//					vt_malicious, vt_suspicious, vt_harmless, vt_undetected
//				) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
//			`,
//				r.LogFile,
//				r.Ip,
//				r.Country,
//				r.City,
//				r.Org,
//				r.AbuseScore,
//				r.VTMalicious,
//				r.VTSuspicious,
//				r.VTHarmless,
//				r.VTUndetected,
//			)
//			if err != nil {
//				return fmt.Errorf("insert failed: %w", err)
//			}
//		}
//
//		return nil
//	}
package db

import (
	"context"
	"fmt"
	// "os"

	"github.com/jackc/pgx/v5"
	"github.com/zeshi09/ipenrich/model"
)

func SaveToPostgres(results []model.EnrichedIP) error {
	// Подключение по DSN
	// connStr := os.Getenv("POSTGRES_DSN")
	connStr := "postgres://casaos:casaos@192.168.2.103:5434/ipenrich"
	// if connStr == "" {
	// 	return fmt.Errorf("POSTGRES_DSN env var is not set")
	// }

	ctx := context.Background()

	// Подключаемся
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to Postgres: %w", err)
	}
	defer conn.Close(ctx)

	// Создание таблицы (если нет)
	_, err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS enriched_ips (
			log_file TEXT,
			ip TEXT,
			dns TEXT,
			country TEXT,
			city TEXT,
			org TEXT,
			abuse_score INT,
			vt_malicious INT,
			vt_suspicious INT,
			vt_harmless INT,
			vt_undetected INT
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	// Вставка записей
	for _, r := range results {
		_, err = conn.Exec(ctx, `
			INSERT INTO enriched_ips (
				log_file, ip, dns, country, city, org,
				abuse_score, vt_malicious, vt_suspicious,
				vt_harmless, vt_undetected
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		`,
			r.LogFile,
			r.Ip,
			r.Dns,
			r.Country,
			r.City,
			r.Org,
			r.AbuseScore,
			r.VTMalicious,
			r.VTSuspicious,
			r.VTHarmless,
			r.VTUndetected,
		)
		if err != nil {
			return fmt.Errorf("insert failed for %s: %w", r.Ip, err)
		}
	}

	return nil
}

