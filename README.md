# blnote

*This project is in its early stages of development as of 7/27/24 and is functional, but unstable.*

A note taking application which displays your notes as interconnected nodes.

This tool is meant to hold and connect a bunch of small, single-thought notes that link together to form a larger idea. It is *not* meant to hold and organize several full-sized notes, like a notebook does.

## Usage

Currently, it is easiest to run the project with Go's toolchain (go run .) and follow the prompts as they appear.

```bash
go run .
```
```txt
Welcome to blnote!
Please enter a command:
```
If you are unsure of where to start, you can type 'help' to see the help menu:
```txt
Please enter a command:
help

Welcome to the help menu!
The available commands are:
new <note name>
link <node1> <node2>
check <note_name>
remove <note_name>
read
show
help
exit
```

Upon using the `exit` command, the program will write any saved nodes with their links to a file called `./NodeStorage/nodes.json`. If the directory does not exist, it should create it. To exit this program without saving your nodes to a file, you must use an interrupt signal, like Ctrl-C.

If you have used the tool previously, using the `read` command will load any saved nodes in the path from above. If no JSON file is found at the path, a message will be printed saying so, but the program will continue running. 

**Note**: `link`ing nodes is purely visual. Right now, the only functionality it provides is that the note names will be displayed as being connected when `show` is used. This may help you to keep your ideas organized, but it does not provide any inherent functionality. 
  - Additionally: there is a bug right now where if one of a pair of linked nodes is deleted, they will both still be displayed as if they exist. In the future, I plan to have the deleted node's name still be visible, but with some kind of denotation that is has been removed. 

  
## Commands

**new** `<note name>`
 - Creates a new Node with the given name. Will prompt for text after entering name.
   - Note: names cannot include whitespace yet. Anything past the first word will be omitted.

**link** `<node1>` `<node2>`
 - Links `<node2>` to `<node1>`. For example:
  - link node_one node_two
  - node_one <- node_two (when they're displayed.)

**check** `<note_name>`
 - Displays the saved text from inside `<note_name>`. 
 
**remove** `<note_name>`
 - Removes `<note_name>` from your list.
  - Note: if `<note_name>` contains a link, it will still be displayed when `show` is used.

**read**
 - Attempts to read nodes from previous sessions. If no nodes exist, a message is printed saying so.
 
**show**
 - Displays a list of any node names from your current session. Previous nodes must be `read` before they can be `show`n. 
  - This command does not display any of the text from the notes. 

**help**
 - Prints a list of commands which can be used. 

**exit**
 - Saves any nodes from your current session to './NodeStorage/nodes.json' then exits the program.
 - **To exit without saving your current session, please use Ctrl-C. If you start a new session, then immediately `exit`, it will save a blank file.** 

