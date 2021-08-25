package storage

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kbats183/CTStickersBot/pkg/core"
	"go.uber.org/zap"
)

func (st *Storage) CreateSticker(ctx context.Context, sticker *tgbotapi.Sticker, stickerText string) (int, error) {
	conn, err := st.clientPull.Acquire(ctx)
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

func (st *Storage) SearchStickers(ctx context.Context, userQuery []string, limit int, logger *zap.Logger) ([]core.StickerAnswer, error) {
	conn, err := st.clientPull.Acquire(ctx)
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
SELECT st.id, st.tg_file_id, st.text_content FROM sticker st INNER JOIN matches_stickers ms ON st.id = ms.id ORDER BY match_count DESC LIMIT $2::INT;`
	rows, err := conn.Query(ctx, sqlQuery, userQuery, limit)
	if err != nil {
		return nil, err
	}
	logger.Info("Have rows")
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

func (st *Storage) SaveUserRequest(ctx context.Context, userTgId int, userLogin string, request string) error {
	conn, err := st.clientPull.Acquire(ctx)
	defer conn.Release()

	sqlQuery := `
WITH user_ids AS (
	INSERT INTO users (tg_id, login) VALUES ($1, $2)
	ON CONFLICT (tg_id) DO UPDATE SET login = $2
	RETURNING id
)
INSERT INTO request(user_id, text)
(SELECT ids.id, $3 FROM (SELECT id FROM user_ids) ids);`
	_, err = conn.Query(ctx, sqlQuery, userTgId, userLogin, request)
	return err
}