package _355_design_twitter

type Tweet struct {
	id        int
	timeStamp int
}

type Twitter struct {
	tweets    map[int][]Tweet
	follows   map[int]map[int]bool
	timeStamp int
}

func Constructor() Twitter {
	return Twitter{
		tweets:    make(map[int][]Tweet),
		follows:   make(map[int]map[int]bool),
		timeStamp: 0,
	}
}

func (tw *Twitter) PostTweet(userId int, tweetId int) {
	tw.timeStamp++

	newTweet := Tweet{
		id:        tweetId,
		timeStamp: tw.timeStamp,
	}

	tw.tweets[userId] = append(tw.tweets[userId], newTweet)
}

func (tw *Twitter) GetNewsFeed(userId int) []int {
	var candidates []Tweet

	candidates = append(candidates, tw.tweets[userId]...)

	for followeeId := range tw.follows[userId] {
		if followeeId != userId {
			candidates = append(candidates, tw.tweets[followeeId]...)
		}
	}

	for i := 0; i < len(candidates); i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[j].timeStamp > candidates[i].timeStamp {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}

	feedSize := 10

	if len(candidates) < 10 {
		feedSize = len(candidates)
	}

	res := make([]int, feedSize)

	for i := 0; i < feedSize; i++ {
		res[i] = candidates[i].id
	}

	return res
}

func (tw *Twitter) Follow(followerId int, followeeId int) {
	if tw.follows[followerId] == nil {
		tw.follows[followerId] = make(map[int]bool)
	}

	tw.follows[followerId][followeeId] = true
}

func (tw *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := tw.follows[followerId]; ok {
		delete(tw.follows[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
