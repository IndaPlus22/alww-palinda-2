The program [many2many.go](src/many2many.go) contains four
producers that together send 32 strings over a channel. At the
other end there are two consumers that receive the strings.
Describe what happens, and explain why it happens, if you make the
following changes in the program. Try first to reason your way
through, and then test your hypothesis by changing and running the
program.

- What happens if you switch the order of the statements
  `wgp.Wait()` and `close(ch)` in the end of the `main` function?

  - **Answer**: The program will crash as the channel is closed before all producers have sent their message

- What happens if you move the `close(ch)` from the `main` function
  and instead close the channel in the end of the function
  `Produce`?

  - **Answer**: The program will crash as the channel is closed before all producers have sent their message as it will close after one producer is done.

- What happens if you remove the statement `close(ch)` completely?
  - **Answer**: Nothing will happen in this case as the wg.Wait() assures that all operations are completed prior to the closing of the channel. However, as the lifetime of the channel is tied to the main function it is unnecessary to close the channel.
- What happens if you increase the number of consumers from 2 to 4?
  - **Answer**: You will double the number of concurrent operations or gorutines which depending on cpu architecture could double to execution time of the program
- Can you be sure that all strings are printed before the program
  stops?
  - **Answer**: No, the program only waits for all producers to be done so it is possible that some strings might not be printed.

Finally, modify the code by adding a new WaitGroup that waits for
all consumers to finish.
