simple golang server that has three routes ,the home route / the add note route and the getallnotes route,the home route
displays a html form that does a post request to the add note route which creates a folder called notes on the server directory
and stores a file which is appended with a note stored as json ,if the folder does not exist the server creates it if the folder
and file exist the file is appended with the new note.The getallnotes route reads the file and displays the data to the frontend.
Used files instead of a database to store the user data
