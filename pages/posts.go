package pages

import (
	"time"
)

var postsReverseChronological []*Page

func AllPosts() []*Page {
	return postsReverseChronological
}

func init() {
	postsReverseChronological = []*Page{
		&Page{
			Path: "/posts/happy-tenth-birthday-pull-requests",

			Title:         "Happy 10th Birthday Pull Requests!",
			ExpandedTitle: "Happy 10th Birthday Pull Requests!",
			Description:   "Github launched Pull Requests on the 23rd of February 2008. 10 years later we have a quick look back.",

			PublishDate: publishDate("2018-02-24"),

			ImagePath:    "/post-images/moon-boot.png",
			ImageCaption: "Post delayed by accident",

			ContentPath: "posts/happy-tenth-birthday-pull-requests.md",
		},
		&Page{
			Path: "/posts/developers-on-call",

			Title:         "Developers On Call",
			ExpandedTitle: "Developers On Call",
			Description:   "Developers not on call benefit from a status quo that is both morally questionable and ultimately unhelpful for software quality.",

			PublishDate: publishDate("2018-02-11"),

			ImagePath:    "/post-images/tools.png",
			ImageCaption: "My toolbox as allegory",

			ContentPath: "posts/developers-on-call.md",
		},
		&Page{
			Path: "/posts/book-review-one-strategy",

			Title:         "Book Review: One Strategy",
			ExpandedTitle: "Book Review: One Strategy by Steven Sinofsky",
			Description:   "I spent a year reading this book and if you manage products or engineers you should too.",

			PublishDate: publishDate("2018-01-17"),

			ImagePath:    "/post-images/one-strategy.png",
			ImageCaption: "It was cheaper to buy a physical copy",

			ContentPath: "posts/book-review-one-strategy.md",
		},
		&Page{
			Path: "/posts/three-flavours-of-iteration",

			Title:         "Three flavours of iteration",
			ExpandedTitle: "Three flavours of iteration",
			Description:   "In the project kickoff meeting your team has all agreed to iterate towards a solution. Sounds great. Everyone's agreed. (Narrator's voice) They hadn't",

			PublishDate: publishDate("2018-01-11"),

			ImagePath:    "/post-images/catlike-agility.jpg",
			ImageCaption: "This kind of agility?",

			ContentPath: "posts/three-flavours-of-iteration.md",
		},
		&Page{
			Path: "/posts/2017-in-review",

			Title:         "2017 In Review",
			ExpandedTitle: "2017 In Review",
			Description:   "A quick braindump of the year primarily to test out the new blogging feature I built for my site.",

			PublishDate: publishDate("2017-12-31"),

			ImagePath:    "/post-images/top-nine-2017.png",
			ImageCaption: "My Instagram top nine of 2017",

			ContentPath: "posts/2017-in-review.md",
		},
	}
}

func publishDate(date string) *time.Time {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil
	}
	return &time
}
