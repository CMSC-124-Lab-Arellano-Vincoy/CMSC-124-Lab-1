# CMSC_124-Lab1

## Group Members
- Arellano, Junel
- Vincoy, Claire Dane

## Problem Analysis & Architecture
**Problem:** We need a lexical scanner to read source code without utilizing regular expressions, grouping raw characters into distinct tokens to feed into a parser.
**Architecture:** A character-by-character scanning loop written in Go, utilizing `start` and `current` pointers to isolate lexemes and the maximal munch principle for lookahead.
**Implementation Timeline:**
* Week 1: Single-character tokens, basic REPL, and CI setup.
* Week 2: Multi-character operators, keywords, strings, and testing.
* Week 3: Full token coverage, testing, and final debugging.

## Language Specification Draft
* **Host Language:** Go
* **File Extension:** `.mylang`
* **Token Output Format:** `Token(type=TYPE, lexeme=text, literal=val, line=N)`
* **Whitespace:** Ignored but counted for line numbers.
* **Token Vocabulary (Initial):** `LEFT_PAREN`, `RIGHT_PAREN`, `PLUS`, `MINUS`, `STAR`, `SLASH`, `EOF`.
