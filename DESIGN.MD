##############################

##GOAL

Develop a framework that allows the user to bridge the gap between recording UI tests in Selenium IDE and having a functional testing suite.

Web application, Vue.js 
##############################

## COMPONENTS

### Backend

Developed in Go. Process recordings, structure them based on a coherent order, fix possible errors, and generate a runnable test suite.

getSequences:
    Returns the sequences in persistence

addSequence:
    Obtains a recording, adds it to the recording pool, and generates an entry in a map. Save the map serialized to a json.

removeSequence:
    Removes the recording file and deletes the entry from the map

renameSequence:
    Renames the sequence.



Model:

Sequence: FileName, ClassName, id, Precedents: {}

SequenceMap: Map<String, Sequence>


### Frontend

Developed with Vue. Provide a simple UI that allows the user to drag & drop recordings, configure premises, choose pre and post test actions (screenshots, etc...), define tolerable errors, and run the tests in the same page. SPA structure.



### USE CASE

User opens App.
User links recordings within the Order Screen view of recordings, to set pre-conditions and order. 
Order Screen can be viewed as nodes or as json with pre-condition field.

### FEATURES
Add button to view content of file (See vue)


### PENDING

DELETE ID FROM PRECEDENT ARRAYS WHEN DELETING AN ELEMENT