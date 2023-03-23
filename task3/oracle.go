// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Socrates"
	venue  = "Athens"
	prompt = "~ "
)

func printSlow(text string, delay int) {}

func main() {
	grumble("Hail, stranger, and welcome to Athens, the home of philosophy. If you have come seeking answers, then you have come to the right place. For I, Socrates, am here to engage with you in dialogue and to help dispel any doubts that trouble your mind. So speak freely, my friend, and let us together seek the truth.", 30)
	fmt.Println()
	grumble("Pray, do not keep your heart's desire concealed any longer. Speak, and let us hear what is truly in your heart.", 30)

	questions := Gadfly()
	pupil := bufio.NewReader(os.Stdin)
	for {
		line, _ := pupil.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			grumble("My dear friend, I must confess that I am puzzled by your line of questioning. For it seems to me that you are asking about matters that are of little consequence, and that do not lead us towards a deeper understanding of truth and wisdom.\n\nI do not mean to be dismissive or unkind, but rather to remind us both of the importance of asking meaningful and thoughtful questions, ones that have the potential to shed light on the mysteries of life and the human condition. So let us turn our attention towards more worthy inquiries, and see where our dialogue takes us.", 50)
			continue
		} else {
			fmt.Printf("%s heard: %s\n", star, line)
			think := []rune("Wait a moment, my friend. I need to reflect and consider before I can offer you an answer.")
			for _, x := range think {
				time.Sleep(30 * time.Millisecond)
				fmt.Printf("%c", x)
			}
			fmt.Printf("\n")
			time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)
			fmt.Printf(".")
			time.Sleep(time.Duration(1+rand.Intn(4)) * time.Second)
			fmt.Printf(".")
			time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
			fmt.Printf(".\n")
		}
		questions <- line // The channel doesn't block.
		fmt.Print(prompt)
	}
}

func Gadfly() chan<- string {
	questions := make(chan string)
	response := make(chan string)

	go getQuestion(questions, response)
	go getResponse(response)

	return questions
}

func getQuestion(q <-chan string, r chan<- string) {
	for {
		question := <-q
		go enlightenment(question, r)
	}
}
func getResponse(r <-chan string) {
	for {
		response := <-r
		go grumble(response, 500)
	}
}

func randomWisdom(r chan<- string) {
	// Cook up some pointless nonsense.
	nonsense := []string{
		"The only true wisdom is in knowing you know nothing.",
		"I cannot teach anybody anything, I can only make them think.",
		"The unexamined life is not worth living.",
		"To be yourself in a world that is constantly trying to make you something else is the greatest accomplishment.",
		"An honest man is always a child.",
		"I know that I am intelligent, because I know that I know nothing.",
		"He who is not a good servant will not be a good master.",
		"The only good is knowledge, and the only evil is ignorance.",
		"The greatest way to live with honor in this world is to be what we pretend to be.",
		"As for me, all I know is that I know nothing.",
		"True wisdom comes to each of us when we realize how little we understand about life, ourselves, and the world around us.",
		"The shortest and surest way to live with honor in the world, is to be in reality what we would appear to be; and if we observe, we shall find, that all human virtues increase and strengthen themselves by the practice of them.",
		"Beware the barrenness of a busy life.",
		"No man has the right to be an amateur in the matter of physical training. It is a shame for a man to grow old without seeing the beauty and strength of which his body is capable.",
		"Envy is the ulcer of the soul.",
		"It is not living that matters, but living rightly.",
		"I am not an Athenian or a Greek, but a citizen of the world.",
		"Death may be the greatest of all human blessings.",
		"Beauty is the bait which with delight allures man to enlarge his kind.",
		"The end of life is to be like God, and the soul following God will be like Him.",
	}
	r <- nonsense[rand.Intn(len(nonsense))]
}

func enlightenment(q string, r chan<- string) {
	question := strings.ToLower(q)
	if strings.Contains(question, "what is the meaning of life") {
		r <- "Ah, the eternal question of the meaning of life! My friend, this is a question that has been asked by philosophers and thinkers for countless generations. In my own view, the meaning of life is something that each individual must discover for themselves, through their own inquiry and reflection.\n\n For me, the purpose of life is to seek truth and wisdom, to cultivate virtue and goodness, and to live in accordance with the principles of justice and morality. I believe that the pursuit of these goals is what gives life its highest meaning and value, and that each person must strive to find their own unique path towards these ends.\n\n Ultimately, my friend, the meaning of life is a question that cannot be answered definitively or universally. Each of us must find our own way, guided by our own reason, intuition, and experience."
	} else if strings.Contains(question, "virtue") {
		r <- "Virtue, my dear friend, is the quality of excellence or moral goodness that distinguishes a person who is living in accordance with the highest standards of human conduct. In my view, virtue is not something that can be attained through external rewards or accolades, but rather is the result of an internal transformation of the soul.\n\nThe virtues that I hold in highest regard include wisdom, courage, justice, temperance, and piety. These qualities are not simply desirable in themselves, but are essential for living a good and fulfilling life.\n\nFor example, wisdom is the virtue of knowledge and understanding, and is necessary for making sound judgments and decisions. Courage is the virtue of facing danger or adversity with steadfastness, and is essential for living a life of integrity and honor. Justice is the virtue of fairness and equity, and is necessary for creating a just and harmonious society. Temperance is the virtue of self-control and moderation, and is essential for avoiding excess and living a balanced life. Piety is the virtue of reverence and respect for the divine, and is essential for living a life of purpose and meaning.\n\nThese virtues, my friend, are not the exclusive domain of the elite or the privileged, but are accessible to all who seek to live a good and fulfilling life."
	} else if strings.Contains(question, "piety") {
		r <- "Piety, my friend, is the quality of showing reverence, respect, and devotion towards the divine. It is the recognition that there is something greater than ourselves, a power or force that transcends the material world and gives meaning and purpose to our lives.\n\nFor me, piety is not simply a matter of performing religious rituals or ceremonies, but is a state of mind and heart that permeates all aspects of our lives. It is the recognition that everything we do, every thought we think, and every action we take is ultimately connected to a higher power or purpose.\n\nPiety involves acknowledging our own limitations and recognizing that we are not the masters of our own fate, but are instead part of a larger cosmic order that is beyond our understanding. It requires humility, reverence, and a sense of awe in the face of the mystery and grandeur of the universe.\n\nUltimately, my friend, piety is a personal and individual matter, and each person must find their own way towards the divine. Whether through prayer, meditation, or other spiritual practices, the pursuit of piety is a lifelong journey of seeking and striving towards something greater than ourselves."
	} else if strings.Contains(question, "absolute morality") {
		r <- "Ah, the question of absolute morality! This is a question that has been debated by philosophers and thinkers for centuries. In my own view, there are certain moral principles that are universal and apply to all people, regardless of their culture, society, or personal beliefs.\n\nFor example, I believe that it is always wrong to intentionally harm others or to act unjustly towards them. This principle, known as the Golden Rule, is a fundamental moral principle that is found in many different cultures and traditions. Similarly, I believe that it is always virtuous to seek knowledge and wisdom, to cultivate self-control and moderation, and to act with honesty and integrity in all our dealings with others.\n\nHowever, while there may be certain moral principles that are universally applicable, the application of these principles can vary depending on the particular circumstances of a situation. Ethics is a complex and multifaceted field, and moral judgments must take into account a variety of factors, including the intentions of the actor, the consequences of their actions, and the social and cultural context in which those actions occur.\n\nTherefore, while I believe that there are certain moral principles that are absolute and universal, the application of those principles is always subject to interpretation and debate. Ultimately, my friend, the pursuit of ethical truth and the cultivation of virtue is an ongoing and never-ending process, one that requires constant reflection, dialogue, and inquiry."
	} else if strings.Contains(question, "plan") || strings.Contains(question, "propose") || strings.Contains(question, "arrange") || strings.Contains(question, "plot") {
		r <- "The phrase \"man proposes, God disposes\" suggests that while human beings may make plans and set goals, the ultimate outcome of those plans is ultimately determined by a higher power or force, such as God.\n\nIn my own view, this phrase reflects a deep understanding of the limits of human knowledge and control. While we may strive to achieve certain ends, we cannot always predict or control the outcomes of our actions. There are many factors that are beyond our control, such as the actions of other people, the forces of nature, or the workings of fate.\n\nAt the same time, I believe that human beings have a responsibility to act with wisdom, virtue, and prudence, even in the face of uncertainty and the unknown. We must do our best to make good decisions, to act with integrity and honor, and to seek truth and knowledge in all our endeavors.\n\nIn the end, my friend, the phrase \"man proposes, God disposes\" reminds us of the importance of humility, respect, and awe in the face of the mysteries of life. We must recognize our own limitations and the limitations of our knowledge and control, while also striving to live in accordance with the highest principles of wisdom and virtue."
	} else {
		randomWisdom(r)
	}

}

func grumble(r string, delay int) {
	response_words := strings.Fields(r)
	time.Sleep(time.Duration(5+rand.Intn(delay)) * time.Millisecond)
	for _, x := range response_words {
		chars := []rune(x)
		for _, y := range chars {
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("%c", y)
		}
		fmt.Printf(" ")
	}
	fmt.Printf("\n")
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
