package main

import (
	"fmt"
	"log"
	"strconv"
)

var (
	fl  *FollowersList
	err error
)

// Check if element is in array
func indexOf(array []int64, e int64) int {
	for i, ele := range array {
		if ele == e {
			return i
		}
	}
	return -1
}

// Monitor user profile for changes
func Monitor(c *Client, userID string) error {
	if fl == nil {
		log.Printf("Getting current Followers...")

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
			u, err := c.GetUserInfo(strconv.FormatInt(f, 10))

			if err != nil {
				return err
			}

			msg := fmt.Sprintf("@%s Stopped following you.", u.ScreenName())
			c.Notify(userID, msg)
		}
	}

	fl = cf

	return nil
}
