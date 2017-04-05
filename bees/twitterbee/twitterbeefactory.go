/*
 *    Copyright (C) 2014-2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 *      Johannes Fürmann <johannes@weltraumpflege.org>
 */

package twitterbee

import (
	"github.com/muesli/beehive/bees"
)

// TwitterBeeFactory is a factory for TwitterBees.
type TwitterBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *TwitterBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := TwitterBee{
		Bee: bees.NewBee(name, factory.ID(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// ID returns the ID of this Bee.
func (factory *TwitterBeeFactory) ID() string {
	return "twitterbee"
}

// Name returns the name of this Bee.
func (factory *TwitterBeeFactory) Name() string {
	return "Twitter"
}

// Description returns the description of this Bee.
func (factory *TwitterBeeFactory) Description() string {
	return "Tweets and reacts to events in your Twitter timeline"
}

// Image returns the filename of an image for this Bee.
func (factory *TwitterBeeFactory) Image() string {
	return factory.ID() + ".png"
}

// LogoColor returns the preferred logo background color (used by the admin interface).
func (factory *TwitterBeeFactory) LogoColor() string {
	return "#00abec"
}

// Options returns the options available to configure this Bee.
func (factory *TwitterBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "consumer_key",
			Description: "Consumer key for Twitter API",
			Type:        "string",
		},
		{
			Name:        "consumer_secret",
			Description: "Consumer secret for Twitter API",
			Type:        "string",
		},
		{
			Name:        "access_token",
			Description: "Access token for Twitter API",
			Type:        "string",
		},
		{
			Name:        "access_token_secret",
			Description: "API secret for Twitter API",
			Type:        "string",
		},
	}
	return opts
}

// Actions describes the available actions provided by this Bee.
func (factory *TwitterBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "tweet",
			Description: "Posts a new tweet",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "status",
					Description: "Text of the Status to tweet, may be no longer than 140 characters",
					Type:        "string",
					Mandatory:   true,
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "follow",
			Description: "Follows another user",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the user you want to follow",
					Type:        "string",
					Mandatory:   true,
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "unfollow",
			Description: "Unfollows another user",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the user you want to unfollow",
					Type:        "string",
					Mandatory:   true,
				},
			},
		},
	}
	return actions
}

// Events describes the available events provided by this Bee.
func (factory *TwitterBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "direct_message",
			Description: "is triggered when someone sends you a direct message",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the DM's author",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "text content of the DM",
					Type:        "string",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "tweeted",
			Description: "is triggered when you tweeted something",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "text",
					Description: "text content of the tweet",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the tweet",
					Type:        "url",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "tweet",
			Description: "is triggered whenever someone you follow tweets",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the tweet's author",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "text content of the tweet",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the tweet",
					Type:        "url",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "mention",
			Description: "is triggered whenever someone mentions you on Twitter",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the mention's author",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "text content of the mention",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the mention",
					Type:        "url",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "retweeted",
			Description: "is triggered when you retweeted something",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the user that you retweeted",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "Text of the retweeted tweet",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the retweeted tweet",
					Type:        "url",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "retweet",
			Description: "is triggered when someone retweets one of your tweets",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the user that retweeted your tweet",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "Text of the retweeted tweet",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the retweeted tweet",
					Type:        "url",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "liked",
			Description: "is triggered when you liked something",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the user that originally wrote the liked tweet",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "Text of the liked tweet",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the liked tweet",
					Type:        "url",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "like",
			Description: "is triggered when someone likes one of your tweets",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the user that liked your tweet",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "Text of the liked tweet",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the liked tweet",
					Type:        "url",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "unliked",
			Description: "is triggered when you unliked something",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the user that originally wrote the unliked tweet",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "Text of the unliked tweet",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the unliked tweet",
					Type:        "url",
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "unlike",
			Description: "is triggered when someone un-likes one of your tweets",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "username",
					Description: "Twitter handle of the user that liked your tweet",
					Type:        "string",
				},
				{
					Name:        "text",
					Description: "Text of the liked tweet",
					Type:        "string",
				},
				{
					Name:        "url",
					Description: "URL of the liked tweet",
					Type:        "url",
				},
			},
		},
	}
	return events
}

func init() {
	f := TwitterBeeFactory{}
	bees.RegisterFactory(&f)
}
