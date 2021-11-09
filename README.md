# The Two Pine Games
This project handle the bowling scores rules described in the specs and [here](https://www.youtube.com/watch?v=aBe71sD8o8c&ab_channel=StephenAnderson). 

It's a command-line application to score a game of ten-pin bowling. You can find the rules [here](https://en.wikipedia.org/wiki/Ten-pin_bowling#Rules_of_play).

Two Pines Application receives a file path containing a list of scores for each player as we you can see in the next example:

```bash
Jeff	10
John	3
John	7
Jeff	7
Jeff	3
John	6
John	3
Jeff	9
Jeff	0
John	10
Jeff	10
John	8
John	1
Jeff	0
Jeff	8
John	10
Jeff	8
Jeff	2
John	10
Jeff	F
Jeff	6
John	9
John	0
Jeff	10
John	7
John	3
Jeff	10
John	4
John	4
Jeff	10
Jeff	8
Jeff	1
John	10
John	9
John	0
``` 

... and should generate a file with with the result of the match:

```bash
Frame		1		2		3		4		5		6		7		8		9		10
Jeff
Pinfalls		X	7	/	9	0		X	0	8	8	/	F	6		X		X	X	8	1	
Score		20		39		48		66		74		84		90		120		148		167		
John
Pinfalls	3	/	6	3		X	8	1		X		X	9	0	7	/	4	4	X	9	0	
Score		16		25		44		53		82		101		110		124		132		151		
```

### Mandatory Features
* Each line represents a player and a chance with the subsequent number of pins knocked down.
* An 'F' indicates a foul on that chance and no pins knocked down (identical for scoring to a roll of 0).
* The columns in each row are tab-separated.


## About the solution
The soolution has be coded defining the following important interfaces:

```
                     +-----------+  +----------------------------------------+
                 +---+  Reader   |  | model                                  |
+------------+   |   +-----------+  |                                        |
|            |   |                  |   +------+   +--------+   +--------+   |
| Controller +---+------------------+---+ Game +---+ Player +---+ Frame  |   |
|            |   |                  |   +------+   +--------+   +--------+   |
+------------+   |   +-----------+  |                                        |
                 +---+  Printer  +--+                                        |
                     +-----------+  +----------------------------------------+
```

* Controller: Orchestrates the overall flow. It only needs a reader and a printer.
* Reader: Defines the interface that should implemented to read data to calculate scores.
* Printer: Defines the interface that should implemented to print the result.
* Game: Defines the interface to interact with the model.
* Player: Defines the "api" to interact with the player.
* Frame: Defines the common api for all kinds of frames. All of them implements Frame's methods.
* NormalFrame: Defines the methods for a normal frame (try 1 + try 2 < 10).
* SpeareFrame: Defines the methods for a frame with two tries and score == 10. 
* StrikeFrame: Defines the methods for a frame with one try with score == 10. 
* LastFrame: Defines the methods for the last frame (special case). 

### Important Note 1: Calculating score of each frame
In order to calculate the score of each player, I coded a **double dispatch** strategy between the player and 
frames. So, when the player iterates the list of frames, calls a frame method sending itself as a parameter. 
Some frames "knows" which method the player must call to be able to calculate your own score.

### Important Note 2: Printing the results
To print the entire model (player, frames, and tries) I coded the solution by relaying on the **Visitor Pattern**. 
I believe that is an elegant solution to resolve this task in this context. 
Maybe are more lines of code, but the solution is more legible and extensible.


### Build and Running

```bash
$ cd cmd/twopines
$ go build
$ ./twopines path_to_file.txt (you can try samples in assets folder)

## Sample:
$ ./twopines ../../assets/scores.txt
```  