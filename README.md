frequency-english
=========

An example program for https://github.com/solvip/frequency

corpus.txt contains the combined text of Jane Eyre, Moby Dick, Paradise Lost, Tarzan of the Apes, The origin of the species, The return of Tarzan, The Voyage of the Beagle and Wuthering Heights - as found on Project Gutenberg

First, the analyzer is trained with corpus.txt and the analyzer state saved to disk.

Then this README, the accompanying LICENSE and the original corpus.txt are frequency analyzed and their score printed to stdout.
For reference, a block of random data is also frequency analyzed

Example output:
```
$ ./frequency-english 
Score of README.md: 0.538007
Score of LICENSE: 0.460335
Score of corpus.txt: 1.000000
Score of 500 random bytes: 0.061490
```

License
=======
Copyright 2014 Sölvi Páll Ásgeirsson

The MIT License