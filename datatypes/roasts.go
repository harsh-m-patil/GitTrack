package datatypes

var languageRoasts = map[string]string{
	"HTML":       "Oh, you write HTML? Congrats on being the guy who sets up the chairs while the real devs do the work.",
	"CSS":        "CSS? Oh, you mean Can't Style Sht*—you must really enjoy spending hours trying to center a div like it’s a form of self-harm!",
	"JavaScript": "So you know JavaScript explain what is a promise !!",
	"Python":     "Oh, Python? The language that tricks you into thinking you're a programmer while you drown in indentation errors like a confused toddler learning to walk.",
	"Java":       "Java? I guess you love typing out novels just to print 'Hello, World!' It’s almost like you want to scare off new developers.",
	"Go":         "Go? More like 'Stop.' You must really enjoy debugging goroutines that disappear faster than your motivation to write tests.",
	"C++":        "C++? Oof, someone’s a masochist. You really must enjoy wrestling with memory leaks and cryptic error messages while everyone else moves on with their lives.",
	"Ruby":       "Ruby? Oh, you like your code elegant and unreadable. Just admit it—you’re just trying to impress your friends with your ‘artistic’ coding skills!",
	"PHP":        "PHP? Still holding onto that '90s nostalgia, huh? Let me guess, your idea of modern web development is using an old, rusty bicycle to win the Tour de France.",
	"TypeScript": "TypeScript—because JavaScript wasn’t frustrating enough for you. Gotta love adding types just to make your bugs dress up nicely.",
	"C#":         "C#? Ah, Microsoft’s flavor of Java. You must enjoy pretending you're coding something cutting-edge while living in the shadow of real developers.",
	"Rust":       "Rust? Oh, you’re the one who shouts 'Rewrite it in Rust!' in every PR, thinking it’ll magically solve all your problems. Good luck with that!",
	"Swift":      "Swift? I guess you love writing iOS apps that look great until the next Apple update crashes them harder than your last relationship.",
	"Perl":       "Perl?! Wow, you must love pain and cryptic one-liners. Your code looks more like ancient hieroglyphics than anything useful!",
	"Kotlin":     "Kotlin? Let me guess, you wanted to write modern code but decided that overcomplicating it was a better use of your time. Android developers are rolling their eyes.",
	"Haskell":    "Haskell? Oh, you think you're too cool for normal programming, huh? Enjoy writing math proofs while the rest of us build real applications.",
	"Scala":      "Scala? Wow, look at you—too smart for Java but still managing to create confusion. Your peers must love deciphering your ‘intelligent’ code.",
	"Shell":      "Shell scripting—because why write clear code when you can confuse everyone, including yourself, with cryptic one-liners?",
	"Lua":        "Lua? Oh, so you’re one of those who think coding should be 'lightweight' and 'simple.' Enjoy debugging while your game engine crashes like your dreams.",
	"SQL":        "SQL? Ah, you're the magician of databases, until you forget one `JOIN` and suddenly your entire query turns into a tragicomedy.",
	"MATLAB":     "MATLAB? Really? A language that can't even start counting from zero? You're just trying to impress people while hiding from real programming!",
	"Zig":        "Zig? So you’re one of those hipsters who wants to feel special in systems programming while explaining why you’re not just using Rust.",
	"C":          "C? Wow, the granddaddy of programming languages! Nothing says 'I'm a programmer' quite like questioning every pointer you touch while drinking three cups of coffee.",
	"Solidity":   "Solidity? Ah, so you’re all in on the blockchain hype! I hope your smart contracts are as solid as your fantasies of becoming a crypto millionaire.",
}

func RoastMe(language string) string {
	if roast, ok := languageRoasts[language]; ok {
		return roast
	}
	return "Language not found"
}
