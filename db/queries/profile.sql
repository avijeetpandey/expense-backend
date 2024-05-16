-- name: CreateProfile :one
INSERT into profile (
    name,
    email,
    bio
) values (
    $1, $2, $3
    ) RETURNING *;


-- name: GetProfile :one
SELECT * from profile where id = $1 LIMIT 1;

-- name: UpdateProfileName :exec
UPDATE profile set name = $2 WHERE id = $1;

-- name: UpdateProfileBio :exec
UPDATE profile set bio = $2 WHERE id = $1;

-- name: UpdateProfileEmail :exec
UPDATE profile set email = $2 WHERE id = $1;

-- name: DeleteProfile :exec
DELETE from profile where id = $1;