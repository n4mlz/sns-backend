-- Keep relationship state idempotent under concurrent requests.
ALTER TABLE `follows`
  ADD UNIQUE INDEX IF NOT EXISTS `follows_unique_pair` (`follower_user_id`, `following_user_id`);

ALTER TABLE `likes`
  ADD UNIQUE INDEX IF NOT EXISTS `likes_unique_pair` (`post_id`, `user_id`);
