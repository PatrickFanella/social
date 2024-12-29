package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/patrickfanella/social/internal/store"
)

var usernames = []string{
	"alex_jones", "bella_smith", "charlie_brown", "daniel_williams", "emma_johnson",
	"fiona_miller", "george_clark", "hannah_davis", "isaac_evans", "julia_wilson",
	"kyle_thomas", "lily_moore", "mason_taylor", "natalie_anderson", "oliver_jackson",
	"peter_white", "quinn_harris", "rachel_martin", "samuel_thompson", "tina_roberts",
	"ursula_lewis", "victor_hall", "wendy_allen", "xavier_young", "yasmine_king",
	"zoe_scott", "aaron_adams", "brenda_walker", "chris_baker", "diana_nelson",
	"eric_carter", "faith_cooper", "gavin_hill", "harper_green", "ian_wood",
	"jackson_bryant", "kelly_bell", "logan_morris", "mia_reed", "nathan_griffin",
	"olivia_adams", "paul_brooks", "quincy_butler", "rebecca_barnes", "steven_foster",
	"tracy_morgan", "ursula_ross", "vincent_perez", "willow_wright", "xander_hamilton",
}

var titles = []string{
	"10 Tips for Effective Time Management",
	"How to Start a Successful Blog in 2024",
	"The Future of Artificial Intelligence: Trends to Watch",
	"Mastering Remote Work: Strategies for Productivity",
	"SEO Best Practices for Beginners",
	"Top 5 Programming Languages to Learn This Year",
	"Building Healthy Habits: A Step-by-Step Guide",
	"Financial Planning Tips for Young Professionals",
	"The Art of Storytelling in Marketing",
	"How to Build Your Personal Brand Online",
	"The Ultimate Guide to Fitness for Beginners",
	"Travel Hacks to Save Money on Your Next Trip",
	"Mindfulness Techniques for Stress Relief",
	"How to Create Engaging Social Media Content",
	"Breaking Down the Basics of Cryptocurrency",
	"Improving Your Writing Skills: A Practical Approach",
	"The Power of Networking: Building Meaningful Connections",
	"Top 10 Tools for Graphic Designers",
	"Productivity Apps That Will Change Your Workflow",
	"Beginner's Guide to Investing in Stocks",
	"How to Monetize Your Blog and Earn Passive Income",
	"Effective Leadership Skills for Team Success",
	"DIY Projects to Refresh Your Home Decor",
	"Exploring the World of Virtual Reality",
	"The Psychology Behind Successful Marketing Campaigns",
	"How to Start a Podcast: A Beginner's Guide",
	"Developing a Morning Routine That Works for You",
	"Building Confidence Through Public Speaking",
	"Simple Recipes for Healthy Meal Prep",
	"Understanding Data Privacy and Cybersecurity",
	"The Best Books to Read for Personal Growth",
	"Tips for Creating Stunning Photography",
	"How to Balance Work and Life Effectively",
	"Breaking Down Machine Learning Concepts",
	"Steps to Launch Your First Online Business",
	"Exploring Renewable Energy Solutions",
	"How to Write a Winning Resume and Cover Letter",
	"The Impact of Social Media on Mental Health",
	"Organizing Your Workspace for Maximum Productivity",
	"Why Emotional Intelligence Matters in Leadership",
	"Steps to Create a Successful Marketing Plan",
	"How to Learn a New Language Quickly and Easily",
	"Travel Destinations You Should Visit This Year",
	"Fitness Trends to Watch in 2024",
	"How to Stay Motivated When Working from Home",
	"Turning Your Hobby Into a Profitable Business",
	"Essential Skills for a Career in Tech",
	"How to Build a Portfolio That Gets You Hired",
	"The Science Behind Habit Formation",
	"Tips for Managing Anxiety and Stress",
	"The Ultimate Guide to Freelancing Success",
}

var contents = []string{
	"Discover practical strategies to manage your time effectively and boost productivity.",
	"Step-by-step guide to launching a successful blog and monetizing it in 2024.",
	"AI is evolving rapidly. Explore its potential and prepare for the future.",
	"Remote work is here to stay. Learn how to master productivity at home.",
	"SEO basics made simple. Optimize your website and attract more visitors.",
	"Which programming languages should you learn in 2024? Find out here.",
	"Good habits lead to a healthier life. Learn how to build them effectively.",
	"Young professionals often overlook financial planning. Start now and secure your future.",
	"Storytelling isn’t just for books—it's key to great marketing strategies.",
	"Building a personal brand online requires consistency and creativity. Here's how.",
	"Need beginner fitness tips? Start with these simple yet effective routines.",
	"Travel smarter, not harder. Discover hacks to save money on your next trip.",
	"Mindfulness reduces stress. Try these techniques for a calmer mind.",
	"Social media engagement is key. Learn how to create shareable content.",
	"Cryptocurrency can be confusing, but this beginner's guide simplifies it.",
	"Better writing opens doors. Enhance your skills with these actionable tips.",
	"Networking builds connections. Discover ways to expand your professional circle.",
	"Graphic designers need the right tools. Explore must-haves for creatives.",
	"Boost productivity with these apps. Make every minute count!",
	"Stocks are a great investment. Here’s how to get started.",
	"Blog monetization is easier than you think. Start earning passive income.",
	"Effective leaders inspire others. Learn skills to lead teams successfully.",
	"DIY home decor projects can refresh your space on a budget. Try these ideas.",
	"Virtual reality is transforming industries. Discover the latest innovations.",
	"Marketing psychology impacts buying decisions. Learn how to use it.",
	"Podcasting is booming. Here's how to create and grow your own show.",
	"Morning routines set the tone for your day. Build one that works for you.",
	"Public speaking can be nerve-wracking. Build confidence with these tips.",
	"Meal prep saves time and keeps you healthy. Start with these recipes.",
	"Data privacy matters. Protect yourself with these essential tips.",
	"Books can change your perspective. Here are must-reads for personal growth.",
	"Photography isn’t just a skill—it’s an art. Learn how to capture stunning shots.",
	"Work-life balance is crucial. Avoid burnout with these strategies.",
	"Machine learning demystified. Understand its basics and applications.",
	"Launch your online business successfully with this step-by-step plan.",
	"Renewable energy is the future. Explore sustainable options today.",
	"Resumes and cover letters make first impressions. Craft ones that stand out.",
	"Social media impacts mental health. Understand its effects and set boundaries.",
	"A tidy workspace boosts efficiency. Organize yours with these tips.",
	"Emotional intelligence is vital for leaders. Learn why and how to improve it.",
	"Marketing plans drive success. Follow these steps to create one.",
	"Learning a new language is rewarding. Start speaking faster with these tips.",
	"Dreaming of travel? Check out these must-visit destinations.",
	"Fitness trends keep changing. Stay updated and stay fit.",
	"Working from home can be tough. Stay motivated with these strategies.",
	"Turn hobbies into income. Learn how to start your side hustle.",
	"Tech careers demand key skills. Find out what you need to thrive.",
	"Portfolios showcase your skills. Build one that gets you hired.",
	"Habits shape success. Learn how to form good ones that stick.",
	"Stress is inevitable. Manage it effectively with these techniques.",
	"Freelancing offers freedom. Succeed with these tips and tricks.",
}

var tags = []string{
	"Productivity", "Time Management", "Blogging Tips", "AI Trends", "Remote Work",
	"SEO Basics", "Programming Languages", "Healthy Habits", "Financial Planning", "Storytelling",
	"Personal Branding", "Fitness Tips", "Travel Hacks", "Mindfulness", "Social Media Content",
	"Cryptocurrency", "Writing Skills", "Networking", "Graphic Design", "Productivity Tools",
	"Investing", "Passive Income", "Leadership Skills", "DIY Projects", "Virtual Reality",
	"Marketing Psychology", "Podcasting", "Morning Routines", "Public Speaking", "Meal Prep",
	"Data Privacy", "Books", "Photography", "Work-Life Balance", "Machine Learning",
	"Online Business", "Renewable Energy", "Resumes and Cover Letters", "Mental Health",
	"Workspace Organization", "Emotional Intelligence", "Marketing Plans", "Language Learning",
	"Travel Destinations", "Fitness Trends", "Working from Home", "Side Hustles", "Tech Careers",
	"Portfolios", "Habit Formation", "Stress Management", "Freelancing",
}

var comments = []string{
	"Great tips! I’ll definitely try these out.",
	"This was really insightful. Thank you!",
	"I never thought about it this way before.",
	"Very helpful and easy to understand.",
	"This article changed my perspective completely.",
	"I’m bookmarking this for future reference.",
	"Can you share more examples on this topic?",
	"Loved this! Keep up the great work.",
	"Super informative and well-written.",
	"This is exactly what I needed to read today.",
	"Thanks for breaking it down so clearly.",
	"I’ll be applying these tips immediately!",
	"I shared this with my team, and they loved it.",
	"Interesting perspective. I’d like to hear more.",
	"Great advice! I’m inspired to start now.",
	"This was very detailed and actionable.",
	"Can you recommend more resources on this?",
	"I feel more confident about trying this now.",
	"Your tips are always so practical.",
	"Such a relatable post. Thank you!",
	"I learned something new today—thank you!",
	"I can’t wait to implement this advice.",
	"This made me think differently about the topic.",
	"Excellent write-up. Keep it coming!",
	"You’ve simplified a complex topic. Impressive!",
	"I appreciate the clarity in your explanations.",
	"This will be super useful for my project.",
	"Looking forward to more posts like this.",
	"Brilliant ideas—thank you so much!",
	"Definitely trying this method tomorrow.",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment", err)
			return
		}
	}
log.Println("Seeding Complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		
		posts[i] = &store.Post{
			UserID: user.ID,
			Title: titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID: posts[rand.Intn(len(posts))].ID,
			UserID: users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}