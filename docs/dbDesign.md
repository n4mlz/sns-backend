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
    - null ok
- bg_image_url
    - null ok
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
- sequence
    - int
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

## follows
- id
    - PK, SK
- following_user_id
    - FK
- follower_user_id
    - FK