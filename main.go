/*
Copyright © 2021 Tejas Wanjari

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import "github.com/tejas-w/portal/cmd"

func main() {
	cmd.Execute()
}

//var text = `
//PREFACE
//
//The following Tales are meant to be submitted to the young reader as an
//introduction to the study of Shakespeare, for which purpose his words
//are used whenever it seemed possible to bring them in; and in whatever
//has been added to give them the regular form of a connected story,
//diligent are has been taken to select such words as might least
//interrupt the effect of the beautiful English tongue in which he wrote:
//therefore, words introduced into our language since his time have been
//as far as possible avoided.
//
//In those tales which have been taken from the Tragedies, the young
//readers will perceive, when they come to see the source from which
//these stories are derived, that Shakespeare's own words, with little
//alteration, recur very frequently in the narrative as well as in the
//dialogue; but in those made from the Comedies the writers found
//themselves scarcely ever able to turn his words into the narrative
//form: therefore it is feared that, in them, dialogue has been made use
//of too frequently for young people not accustomed to the dramatic form
//of writing. But this fault, if it be a fault, has been caused by an
//earnest wish to give as much of Shakespeare's own words as possible:
//and if the 'He said,' and 'She said,' the question and the reply,
//should sometimes seem tedious to their young ears, they must pardon it,
//because it was the only way in which could be given to them a few hints
//and little foretastes of the great pleasure which awaits them in their
//elder years, when they come to the rich treasures from which these
//small and valueless coins are extracted; pretending to no other merit
//than as faint and imperfect stamps of Shakespeare's matchless image.
//Faint and imperfect images they must be called, because the beauty of
//his language is too frequently destroyed by the necessity of changing
//many of his excellent words into words far less expressive of his true
//sense, to make it read something like prose; and even in some few
//places, where his blank verse is given unaltered, as hoping from its
//simple plainness to cheat the young reader into the belief that they
//are reading prose, yet still his language being transplanted from its
//own natural soil and wild poetic garden, it must want much of its
//native beauty.
//
//It has been wished to make these Tales easy reading for very young
//children. To the utmost of their ability the writers have constantly
//kept this in mind; but the subjects of most of them made this a very
//difficult task. It was no easy matter to give the histories of men and
//women in terms familiar to the apprehension of a very young mind. For
//young ladies too, it has been the intention chiefly to write; because
//boys being generally permitted the use of their fathers' libraries at a
//much earlier age than girls are, they frequently have the best scenes
//of Shakespeare by heart, before their sisters are permitted to look
//into this manly book; and, therefore, instead of recommending these
//Tales to the perusal of young gentlemen who can read them so much
//better in the originals, their kind assistance is rather requested in
//explaining to their sisters such parts as are hardest for them to
//understand: and when they have helped them to get over the
//difficulties, then perhaps they will read to them (carefully selecting
//what is proper for a young sister's ear) some passage which has pleased
//them in one of these stories, in the very words of the scene from which
//it is taken; and it is hoped they will find that the beautiful
//extracts, the select passages, they may choose to give their sisters in
//this way will be much better relished and understood from their having
//some notion of the general story from one of these imperfect
//abridgments; which if they be fortunately so done as to prove
//delightful to any of the young readers, it is hoped that no worse
//effect will result than to make them wish themselves a little older,
//that they may be allowed to read the Plays at full length (such a wish
//will be neither peevish nor irrational). When time and leave of
//judicious friends shall put them into their hands, they will discover
//in such of them as are here abridged (not to mention almost as many
//more, which are left untouched) many surprising events and turns of
//fortune, which for their infinite variety could not be contained in
//this little book, besides a world of sprightly and cheerful characters,
//both men and women, the humour of which it was feared would be lost if
//it were attempted to reduce the length of them.
//
//What these Tales shall have been to the young readers, that and much more it is the writers' wish that the true Plays of Shakespeare may prove to them in older years enriches of the fancy, strengtheners of virtue, a withdrawing from all selfish and mercenary thoughts, a lesson of all sweet and honourable thoughts and actions, to teach courtesy, benignity, generosity, humanity: for of examples, teaching these virtues, his pages are full.
//`
//
//func main() {
//	var lines []string
//	for i, line := range strings.Split(text, ".") {
//		lines = append(lines, fmt.Sprint(i, "> ", strings.Trim(line, "\n")))
//	}
//
//	p := portal.New(&portal.Options{
//		Height: 8,
//		Width:  80,
//	})
//	defer p.Close()
//	in := p.Open()
//	for _, l := range lines {
//		in <- l
//		time.Sleep(200 * time.Millisecond)
//	}
//}
