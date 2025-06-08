# Acknowledgements
This repo is a newly initialized tree-sitter repo, but it uses the grammar.js and scanner.c from https://github.com/DerekStride/tree-sitter-sql
The reason for not just forking and making PR's is that these changes (mostly involving making ; optional in most places) may conflict with other dialects of SQL, and I somehow couldn't make my NVIM config read and compile from the forked repo I made, so out of ignorance this is a new tree-sitter initialized from scratch.
I did add the LICENSE.md from the original repo and I think I did what I have to to ensure I am not breaking any license. If I am incorrect here, let me know and I will take this down.

# Differences so far from the original repo
- [] as quotes is permitted, and that allows an identifier to contain any characters
- go splits programs into several programs
- ; is optional everywhere, and even when they may be required in TSQL (CTE's for instance etc.. ), but my goal is to make a parser that is lenient without it breaking easily
- Support for åäöÅÄÖ in identifiers
