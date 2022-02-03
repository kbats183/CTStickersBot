package storage

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kbats183/CTStickersBot/pkg/core"
)

func (st *Storage) PingDB(ctx context.Context) (stickerCount int, userCount int, requestCount int, chosenStickerCount int, adminCount int, err error) {
	conn, err := st.clientPull.Acquire(ctx)
	if conn == nil || err != nil {
		return 0, 0, 0, 0, 0, err
	}
	defer conn.Release()

	sqlQuery := `SELECT 
(SELECT count(*) FROM sticker) AS sticker_count, 
(SELECT count(*) FROM users) as user_count, 
(SELECT count(*) FROM request) as request_count,
(SELECT count(*) FROM request WHERE chosen_sticker_id IS NOT NULL) as chosen_sticker_count,
(SELECT count(*) FROM admins) as admin_count;`
	row := conn.QueryRow(ctx, sqlQuery)
	err = row.Scan(&stickerCount, &userCount, &requestCount, &chosenStickerCount, &adminCount)
	if err != nil {
		return 0, 0, 0, 0, 0, err
	}
	return
}

func (st *Storage) CheckAdminTelegram(ctx context.Context, telegramID int64, telegramLogin string) (adminID int, err error) {
	conn, err := st.clientPull.Acquire(ctx)
	if conn == nil || err != nil {
		return 0, err
	}
	defer conn.Release()

	sqlQuery := `SELECT id FROM admins WHERE tg_id = $1 AND tg_login = $2 AND active IS TRUE;`
	rows, err := conn.Query(ctx, sqlQuery, telegramID, telegramLogin)
	if err != nil {
		return
	}
	for rows.Next() {
		err := rows.Scan(&adminID)
		if err != nil {
			return 0, err
		}
		return adminID, nil
	}
	return
}

func (st *Storage) CreateSticker(ctx context.Context, sticker *tgbotapi.Sticker, stickerText string) (int, error) {
	conn, err := st.clientPull.Acquire(ctx)
	if conn == nil || err != nil {
		return 0, err
	}
	defer conn.Release()

	var newStickerID int
	sqlQuery := `
INSERT INTO sticker(tg_file_id, tg_set_name, text_content) VALUES ($1, $2, $3)
ON CONFLICT (tg_file_id) DO UPDATE SET text_content = $3
RETURNING id;`
	row := conn.QueryRow(ctx, sqlQuery, sticker.FileID, sticker.SetName, stickerText)
	err = row.Scan(&newStickerID)
	if err != nil {
		return 0, err
	}
	return newStickerID, err
}

func (st *Storage) SearchStickers(ctx context.Context, userQuery []string, limit int) ([]core.StickerAnswer, error) {
	conn, err := st.clientPull.Acquire(ctx)
	if conn == nil || err != nil {
		return nil, err
	}
	defer conn.Release()

	var stickers []core.StickerAnswer
	sqlQuery := `
WITH
query_words AS (
    SELECT word FROM UNNEST($1::VARCHAR[]) AS query(word)
),
matches_stickers AS (
    SELECT st.id, (SELECT count(*) FROM query_words WHERE LOWER(st.text_content) LIKE '%' || query_words.word || '%') as match_count FROM sticker st ORDER BY match_count DESC LIMIT $2::INT
)
SELECT st.id, st.tg_file_id, st.text_content FROM sticker st INNER JOIN matches_stickers ms ON st.id = ms.id ORDER BY match_count DESC, st.addition_time DESC LIMIT $2::INT;`
	rows, err := conn.Query(ctx, sqlQuery, userQuery, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sticker core.StickerAnswer
		err := rows.Scan(
			&sticker.ID,
			&sticker.FileID,
			&sticker.StickerTitle,
		)
		if err != nil {
			return nil, err
		}
		stickers = append(stickers, sticker)
	}
	return stickers, nil
}

func (st *Storage) SaveUserRequest(ctx context.Context, userTgID int64, userLogin string, tgID int, request string) error {
	conn, err := st.clientPull.Acquire(ctx)
	if conn == nil || err != nil {
		return err
	}
	defer conn.Release()

	sqlQuery := `
WITH user_ids AS (
	INSERT INTO users (tg_id, login) VALUES ($1, $2)
	ON CONFLICT (tg_id) DO UPDATE SET login = $2
	RETURNING id
)
INSERT INTO request(user_id, tg_id, text)
(SELECT ids.id, $3, $4 FROM (SELECT id FROM user_ids) ids);`
	_, err = conn.Query(ctx, sqlQuery, userTgID, userLogin, tgID, request)
	return err
}

func (st *Storage) SaveUserRequestChosenSticker(ctx context.Context, tgRequestID int, stickerID int) error {
	conn, err := st.clientPull.Acquire(ctx)
	if conn == nil || err != nil {
		return err
	}
	defer conn.Release()

	sqlQuery := `
UPDATE request SET chosen_sticker_id = $1 WHERE tg_id = $2;`
	_, err = conn.Query(ctx, sqlQuery, stickerID, tgRequestID)
	return err
}
