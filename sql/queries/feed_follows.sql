-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
        VALUES (
            $1,
            $2,
            $3,
            $4,
            $5
        )
    RETURNING *
)

SELECT 
    iff.*,
    f.name AS feed_name,
    u.name AS user_name
FROM inserted_feed_follow iff
INNER JOIN users u on u.id = iff.user_id
INNER JOIN feeds f on f.id = iff.feed_id;


-- name: GetFeedFollowsForUser :many
SELECT * FROM feeds
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1; 

-- name: Unfollow :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;
