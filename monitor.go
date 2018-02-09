package main

import (
	"fmt"
	"log"
	"strconv"
)

var fl *FollowersList
var err error

func indexOf(array []int64, e int64) int {
	for i, ele := range array {
		if ele == e {
			return i
		}
	}
	return -1
}

func Monitor(c *Client, userID string) error {
	if fl == nil {
		log.Printf("Initialising followers...")

		fl, err = c.GetFollowersList(userID)

		if err != nil {
			return err
		}

		return nil
	}

	cf, err := c.GetFollowersList(userID)

	if err != nil {
		return err
	}

	for _, f := range fl.Ids {
		if indexOf(cf.Ids, f) == -1 {
			// Not found -> Unfollow
			u, err := c.GetUserInfo(strconv.FormatInt(f, 10))
			if err != nil {
				return err
			}

			msg := fmt.Sprintf("👎 @%s stopped following you.", u.ScreenName())
			c.Notify(userID, msg)
		}
	}

	fl = cf

	return nil
}

// func monitor(client Client, username string) {
// 	previousFollowers, err := getFollowers(username)
// 	if err != nil {
// 		log.Printf("Error getting initial followers!!: %s", err)
// 		return
// 	}

// 	log.Printf("Initial followers updated!")
// 	sendDM(username, "🤖 Started!")

// 	for {
// 		// sleep
// 		time.Sleep(time.Duration(refreshTime) * time.Minute)

// 		actualFollowers, err := getFollowers(username)
// 		if err != nil {
// 			log.Printf("Error getting followers: %s", err)
// 			sendDM(username, fmt.Sprintf("🤖 Error getting followers: %s", err))
// 		} else {
// 			// Check unfollows
// 			for _, f := range previousFollowers.Ids {
// 				if IndexOf(actualFollowers.Ids, f) == -1 {
// 					// Not found -> Unfollow
// 					u, err := getUserInfo(f)
// 					if err != nil {
// 						log.Printf("Error getting user info [%d]: %s", f, err)
// 					} else {
// 						msg := fmt.Sprintf("👎 @%s stopped following you.", u.ScreenName())
// 						sendDM(username, msg)
// 					}
// 				}
// 			}

// 			// Check follows
// 			for _, f := range actualFollowers.Ids {
// 				if IndexOf(previousFollowers.Ids, f) == -1 {
// 					// Not found -> Unfollow
// 					u, err := getUserInfo(f)
// 					if err != nil {
// 						log.Printf("Error getting user info [%d]: %s", f, err)
// 					} else {
// 						msg := fmt.Sprintf("👍 @%s started following you!", u.ScreenName())
// 						sendDM(username, msg)
// 					}
// 				}
// 			}

// 			// update followers
// 			previousFollowers = actualFollowers
// 			log.Printf("Previous followers updated!")
// 		}
// 	}
// }
