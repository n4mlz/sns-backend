# db design

## users
- id
    - PK, Firebase
- user_name
    - UK
- display_name
- biography
    - text
    - null ok
- icon_url
- bgimage_url
- created_at
    - timestamp

## posts
- id
    - PK, SK
- user_id
    - FK
- content
    - text
- created_at
    - timestamp

## comments
- id
    - PK, SK
- post_id
    - FK
- user_id
    - FK
- content
    - text
- created_at
    - timestamp

## replies
- id
    - PK, SK
- comment_id
    - FK
- user_id
    - FK
- content
    - text
- created_at
    - timestamp

## likes
- id
    - PK, SK
- post_id
    - FK
- user_id
    - FK
- created_at
    - timestamp

## follows
- id
    - PK, SK
- follower_user_id
    - FK
- following_user_id
    - FK

## post_notifications
- id
    - PK, SK
- user_id
    - FK
- comment_id
    - FK
    - null ok
- reply_id
    - FK
    - null ok
