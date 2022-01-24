package services

type ItunesRssModel struct {
	Feed FeedModel `json:"feed"`
}

type FeedModel struct {
	Entry []ItunesReviewModel `json:"entry"`
}

type ItunesReviewModel struct {
	Author    AuthorModel `json:"author"`
	UpdatedAt ItunesLabel `json:"updated"`
	Title     ItunesLabel `json:"title"`
	Rating    ItunesLabel `json:"im:rating"`
	ReviewId  ItunesLabel `json:"id"`
	Content   ItunesLabel `json:"content"`
}

type AuthorModel struct {
	Name ItunesLabel `json:"name"`
}

type ItunesLabel struct {
	Label string `json:"label"`
}

// {
// 	author: {
// 	uri: {
// 	label: "https://itunes.apple.com/us/reviews/id107909396"
// 	},
// 	name: {
// 	label: "Socratic Method"
// 	},
// 	label: ""
// 	},
// 	updated: {
// 	label: "2013-03-27T14:56:53-07:00"
// 	},
// 	im:rating: {
// 	label: "5"
// 	},
// 	im:version: {
// 	label: "1.03"
// 	},
// 	id: {
// 	label: "777118441"
// 	},
// 	title: {
// 	label: "Perfection!"
// 	},
// 	content: {
// 	label: "Downloaded this app this morning and used it for a conference dinner for my job. I was uber impressed that it worked so well with so many people at the table! Great job",
// 	attributes: {
// 	type: "text"
// 	}
// 	},
// 	link: {
// 	attributes: {
// 	rel: "related",
// 	href: "https://itunes.apple.com/us/review?id=595068606&type=Purple%20Software"
// 	}
// 	},
// 	im:voteSum: {
// 	label: "1"
// 	},
// 	im:contentType: {
// 	attributes: {
// 	term: "Application",
// 	label: "Application"
// 	}
// 	},
// 	im:voteCount: {
// 	label: "1"
// 	}
// 	},
