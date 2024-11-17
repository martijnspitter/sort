# Build Your Own Sort Tool

This challenge is to build your own version of the Unix command line tool sort!

It sounds like a simple utility, but if you build everything in this challenge you’ll be using five different sorting algorithms and at least three different data structures.

The Unix command line tools are a great metaphor for good software engineering and they follow the Unix Philosophies of:

- Design for simplicity; add complexity only where you must.
- Design programs to be connected to other programs - each tool can be easily connected to other tools to create incredibly powerful compositions.
- Design programs to be connected to other programs - each tool can be easily connected to other tools, via files and streams, to create incredibly powerful compositions.

Following these philosophies has made the simple unix command line tools some of the most widely used software engineering tools which can be chained together to create far more complex and powerful set of tools that you’d expect.

You can read more about the Unix Philosophy on the blog.

## The Challenge - Building sort

The functional requirements for sort are concisely described by it’s man page - give it a go in your local terminal now:

`man sort`

The TL/DR version is: the sort utility sorts text files by lines. A line is a record separated from the subsequent record by a newline.

We’re going to ignore binary files for this challenge, though I’d encourage you to consider extending your implementation to support them once you’ve finished the challenge.

### Step Zero

Like all good software engineering we’re zero indexed! In this step you’re going to set your environment up ready to begin developing and testing your solution.

I’ll leave you to setup your IDE / editor of choice and programming language of choice. After that here’s what I’d like you to do to be ready to test your solution.

Download the this text from Project Gutenberg and save it as test.txt.

Then we will create a list of all the words in the file and save it as words.txt.

`% tr -s '[[:punct:][:space:]]' '\n' < test.txt |sed '/^[0-9]/d' > words.txt`

Please note how powerful the unix command line tools are. If you’re not familiar with them you’re missing out on some incredibly power tools that can do a lot data munging. It’s a great topic for self study.

### Step 1

In this step your goal is to implement the essence of the sort program, we want to be able to run sort and have it open a file and output the lines in the file sorted lexicographically.

Your goal is to be able to run the following command and see the same output:

```bash
% sort words.txt | uniq | head -n5
A
ACTUAL
AGREE
AGREEMENT
AND
```

### Step 2

In this step your goal is to implement the same functionality as the above but with the unique functionality built into sort, i.e. the -u option we saw when we did man sort.

```bash
% sort -u words.txt | head -n5
A
ACTUAL
AGREE
AGREEMENT
AND
```

You can achieve this whilst sorting, before sorting or after sorting the list. Think about the algorithms and data structures involved and the trade-offs you’ll be making in terms of computational complexity, space complexity and extensibility of the software.

Hint: just like writing code professionally, sometimes it pays to check you have all the known requirements before you start coding.

### Step 3

In this step your goal is to allow the user of your sort program to select which sort algorithm to use. Checking man sort, we see that the standard Unix version offers radix sort, merge sort, quicksort and heapsort.

Implement as many of the different sort algorithms as you like and think carefully about how they interact with the -u flag.

There are thousands of words in our test text, so you might like to implement some unit testing / using the different sorting algorithms as a way to learn or practice test-drive development using your stacks testing library.

### Step 4

In this step your goal is to implement random sort. Random sort is defined by the man page of sort as:

```
-R, -random-sort, -sort=random
Sort by a random order. This is a random permutation of the inputs except that
the equal keys sort together. It is implemented by hashing the input keys and
sorting the hash values. The hash function is chosen randomly. The hash function
is randomized by /dev/random content.
```

It’s random so your results won’t match mine, but it should look something like this:

```
% sort --random-sort -u words.txt | head -n5
gales
represent
satisfied
Wealth
awful
```
