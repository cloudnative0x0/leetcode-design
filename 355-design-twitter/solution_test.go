package _355_design_twitter

import (
	"reflect"
	"testing"
)

func TestTwitterExample(t *testing.T) {
	twitter := Constructor()

	twitter.PostTweet(1, 5)
	assertFeed(t, &twitter, 1, []int{5})

	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	assertFeed(t, &twitter, 1, []int{6, 5})

	twitter.Unfollow(1, 2)
	assertFeed(t, &twitter, 1, []int{5})
}

func TestOwnTweetsAreReturnedNewestFirst(t *testing.T) {
	twitter := Constructor()

	twitter.PostTweet(7, 101)
	twitter.PostTweet(7, 102)
	twitter.PostTweet(7, 103)

	assertFeed(t, &twitter, 7, []int{103, 102, 101})
}

func TestFeedCombinesSeveralUsersByTimestamp(t *testing.T) {
	twitter := Constructor()

	twitter.Follow(1, 2)
	twitter.Follow(1, 3)

	twitter.PostTweet(2, 20)
	twitter.PostTweet(1, 10)
	twitter.PostTweet(3, 30)
	twitter.PostTweet(2, 21)
	twitter.PostTweet(1, 11)

	assertFeed(
		t,
		&twitter,
		1,
		[]int{11, 21, 30, 10, 20},
	)
}

func TestFeedContainsAtMostTenTweets(t *testing.T) {
	twitter := Constructor()

	for tweetID := 1; tweetID <= 15; tweetID++ {
		twitter.PostTweet(1, tweetID)
	}

	assertFeed(
		t,
		&twitter,
		1,
		[]int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6},
	)
}

func TestRepeatedFollowDoesNotDuplicateTweets(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(2, 200)

	twitter.Follow(1, 2)
	twitter.Follow(1, 2)

	assertFeed(t, &twitter, 1, []int{200})
}

func TestFollowingSelfDoesNotDuplicateOwnTweets(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 100)
	twitter.Follow(1, 1)

	assertFeed(t, &twitter, 1, []int{100})
}

func TestUnknownUserHasEmptyFeed(t *testing.T) {
	twitter := Constructor()

	assertFeed(t, &twitter, 999, []int{})
}

func TestUnfollowMissingRelationshipDoesNothing(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 10)

	twitter.Unfollow(1, 2)
	twitter.Unfollow(99, 100)

	assertFeed(t, &twitter, 1, []int{10})
}

func assertFeed(
	t *testing.T,
	twitter *Twitter,
	userID int,
	want []int,
) {
	t.Helper()

	got := twitter.GetNewsFeed(userID)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(
			"GetNewsFeed(%d) = %v; want %v",
			userID,
			got,
			want,
		)
	}
}
