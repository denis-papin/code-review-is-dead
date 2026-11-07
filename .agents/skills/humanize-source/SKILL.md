---
name: humanize-source
description: >-
  Rewrite a Markdown (.md) document so it reads clearly and gets straight to the
  point — short sentences, plain words, no filler. Use this whenever the user
  hands over a .md file (README, spec, design doc, notes, changelog, blog draft)
  and asks to "humanize", "simplify", "tighten", "clean up", "make it clearer",
  "cut the fluff", "make it more readable", or otherwise improve the writing.
  Trigger even if they don't say "humanize" — any request to make an existing
  Markdown file shorter, plainer, or less corporate/AI-sounding should use this.
  Not for generating new content from scratch and not for non-Markdown files.
---

# Humanize Source

Take a Markdown file written in dense, wordy, or robotic prose and rewrite it so a
busy human can read it fast. The goal is clarity and brevity, not decoration.

## What good output looks like

The reader should get the point in the first sentence of any section. Specifics:

- **Short sentences.** One idea each. Break long sentences at the "and", "which", "in order to".
- **Plain words.** "use" not "utilize", "help" not "facilitate", "about" not "in regards to", "now" not "at this point in time".
- **Active voice.** "The server caches the result" beats "The result is cached by the server".
- **Lead with the point.** Put the conclusion or instruction first, then the reason. Cut throat-clearing intros like "It is important to note that".
- **No filler.** Drop "in order to", "it should be noted", "as previously mentioned", "basically", "essentially", "very", "really", hedging chains, and restated headings.
- **Concrete over abstract.** Name the actual thing instead of "the relevant component" or "this functionality".

## How to work

1. **Read the whole file first.** Understand what it's actually about before changing a word. Note the audience (devs? end users? teammates?) and the document type — a changelog is trimmed differently than a tutorial.

2. **Rewrite section by section, trimming freely.** You have permission to cut redundant sentences, merge overlapping sections, and drop low-value detail to get to the point. Tighten wording aggressively. But never invent facts, and never drop a fact, step, warning, or code/command the reader needs to act. Brevity must not cost correctness.

3. **Preserve Markdown structure and meaning.** Keep code blocks, commands, links, tables, and inline code exactly as-is unless the surrounding prose is what's wordy. Keep heading hierarchy sensible. Don't change technical claims, version numbers, or API names.

4. **Ask only when genuinely unclear.** If simplifying a sentence risks changing a substantive claim, leave it intact and continue improving the unambiguous passages. Ask a focused question only when that uncertainty prevents the requested result. Example: "This paragraph says the cache 'self-heals' — does that mean it retries failed writes, or rebuilds on restart? I want to simplify it without changing what it claims." Don't ask about style or tone; just simplify those. One or two sharp questions beat a long interview.

5. **Edit the file in place.** Update the original Markdown file unless the user requests a separate copy.

6. **Report what you changed.** After editing, give a 2-4 line summary: roughly how much shorter it is, what kinds of changes you made (e.g. "cut the intro boilerplate, merged the two install sections, shortened sentences"), and flag ambiguous passages left unchanged. Preserve substantive information.

## Before / after examples

**Wordy → tight:**

Before:
> It is important to note that, in order to successfully complete the installation process, the user will first need to ensure that they have the necessary prerequisites installed on their system.

After:
> Before installing, make sure the prerequisites are installed.

**Buried point → leading point:**

Before:
> There are a number of different factors that can affect performance, and after extensive testing we have come to the conclusion that the most impactful one is the size of the cache.

After:
> Cache size affects performance the most.

**Robotic → human:**

Before:
> This feature provides users with the capability to leverage automated workflows in order to optimize their productivity.

After:
> This feature automates repetitive work so you get more done.

## What not to do

- Don't add a "Summary" or "Conclusion" the original didn't have.
- Don't make it sound chatty or add personality the document doesn't call for — plain is the target, not casual.
- Don't strip technical precision to make a sentence shorter.
- Don't touch files that aren't Markdown. If the input isn't a `.md` file, say so and stop.
