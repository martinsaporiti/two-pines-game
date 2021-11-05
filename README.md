# The Two Pine Games

## The problem to solve



#### Considerations
##### Calculating Frame's Score.
In order to calculate the score of each player, I coded a **double dispatch** strategy between the player and 
frames. So, when the player iterates the list of frames, calls a frame method sending itself as a parameter. 
Some frames "knows" which method the player must call to be able to calculate your own score.

##### Printing the Results
To print the entire model (player, frames, and tries) I coded the solution by relaying on the **Visitor Pattern**. 
I believe that is an elegant solution to resolve this task in this context. 
Maybe are more lines of code, but the solution is more legible and extensible.


#### Build and Running

```bash
$ cd cmd
$ go build -o two_pines
$ ./two_pines path_to_file.txt
```  